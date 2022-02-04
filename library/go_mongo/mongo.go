package go_mongo

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

func NewDatabase(name string) *mongo.Database {
	if !isInit {
		panic(ErrNoInit)
	}

	// 检查配置名
	conf, ok := connectConf[name]
	if !ok {
		l.Errorf("mongo: 配置不存在, conf name: %s", name)
		return &mongo.Database{}
	}

	// 获取client
	mongoPoolLock.RLock()
	if client, ok := mongoPool[name]; ok {
		mongoPoolLock.RUnlock()
		return client.Database(conf.Database)
	}
	mongoPoolLock.RUnlock()

	client, err := NewClient(conf)
	if err != nil {
		l.Errorw("mongo: 连接失败", "conf", conf, "err", err)
		return &mongo.Database{}
	}

	mongoPoolLock.Lock()
	mongoPool[name] = client
	mongoPoolLock.Unlock()

	return client.Database(conf.Database)
}

func NewClient(conf Conf) (*mongo.Client, error) {
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

type GoMongo struct {
	MongoConfName string
	CollationName string
}

func (ctl *GoMongo) MDB() *mongo.Database {
	return NewDatabase(ctl.MongoConfName)
}

func (ctl *GoMongo) MCollection(opts ...*options.CollectionOptions) *mongo.Collection {
	return NewDatabase(ctl.MongoConfName).Collection(ctl.CollationName, opts...)
}
