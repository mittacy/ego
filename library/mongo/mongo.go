package mongo

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"sync"
	"time"
)

var ErrNoInit = errors.New("mongo: please initialize with InitConnectConf() method")

var (
	isInit        bool                     // 是否初始化
	mongoPool     map[string]*mongo.Client // 单例池
	mongoPoolLock sync.RWMutex             // 单例池锁
)

func Get(name string) *mongo.Client {
	if !isInit {
		panic(ErrNoInit)
	}

	mongoPoolLock.RLock()
	if client, ok := mongoPool[name]; ok {
		mongoPoolLock.RUnlock()
		return client
	}
	mongoPoolLock.RUnlock()

	conf, isExist := connectConf[name]
	if !isExist {
		l.Errorf("mongo: 配置不存在, conf name: %s", name)
		return &mongo.Client{}
	}

	db, err := NewConnect(conf)
	if err != nil {
		l.Errorw("mongo: 连接失败", "conf", conf, "err", err)
		return &mongo.Client{}
	}

	mongoPoolLock.Lock()
	mongoPool[name] = db
	mongoPoolLock.Unlock()

	return db
}

func NewConnect(conf Conf) (*mongo.Client, error) {
	if !isInit {
		panic(ErrNoInit)
	}

	var uri string
	if conf.Password != "" {
		uri = fmt.Sprintf(UriPSWFormat, conf.User, conf.Password, conf.Host, conf.Port)
	} else {
		uri = fmt.Sprintf(UriFormat, conf.Host, conf.Port)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	// check connect
	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, err
	}

	return client, nil
}

type Mongo struct {
	MongoConfName string
}

func (ctl *Mongo) DB() *mongo.Client {
	return Get(ctl.MongoConfName)
}
