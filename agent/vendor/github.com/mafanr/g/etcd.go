package g

import (
	"bytes"
	"context"
	"fmt"
	"strconv"
	"sync"
	"time"

	"go.uber.org/zap"

	"github.com/sunface/talent"

	"go.etcd.io/etcd/clientv3"
	"go.etcd.io/etcd/mvcc/mvccpb"
)

const (
	// ApisRootPath 是通过Api查找服务器地址时，etcd中的根目录
	AppRootPath = "/mafanr/apps/"

	//更新API间隔(秒)
	ServiceQueryInterval = 60

	// 存储api和节点信息的过期时间
	AppStoreInterval = 15
	AppLeaseTime     = 120
)

type Etcd struct {
	*clientv3.Client
}

var ETCD = &Etcd{}

var apps *sync.Map

func (etcd *Etcd) Register(app string, addrs []string, port string) {
	etcd = InitEtcd(addrs)

	// 保存服务状态到etcd
	ip := talent.LocalIP()
	L.Info("register local ip to etcd", zap.String("ip", ip))

	host := ip + ":" + port
	go func() {
		for {
			err := etcd.storeServer(&ServerInfo{app, host, 0})
			if err != nil {
				L.Error("register to etcd error", zap.Error(err))
			}

			time.Sleep(time.Second * AppStoreInterval)
		}
	}()
}

func (etcd *Etcd) QueryAll(addrs []string) {
	etcd = InitEtcd(addrs)

	var err error
	apps, err = etcd.queryAll()
	if err != nil {
		L.Fatal("Get application node error", zap.Error(err))
	}

	go func() {
		ch := etcd.Watch(context.Background(), AppRootPath, clientv3.WithPrefix())
		for {
			update := <-ch
			for _, e := range update.Events {
				switch e.Type {
				case mvccpb.DELETE:
					s, err := etcd.parseK(e.Kv)
					if err != nil {
						L.Warn("parse etcd kv error", zap.Error(err))
						continue
					}

					appI, ok := apps.Load(s.Service)
					if ok {
						app := appI.([]*ServerInfo)
						// delete the application node
						for i, n := range app {
							if n.IP == s.IP {
								app = append(app[:i], app[i+1:]...)
							}
						}

						if len(app) == 0 {
							apps.Delete(s.Service)
						} else {
							apps.Store(s.Service, app)
						}
					}
				case mvccpb.PUT:
					s, err := etcd.parseKV(e.Kv)
					if err != nil {
						L.Warn("parse etcd kv error", zap.Error(err))
						continue
					}

					appI, ok := apps.Load(s.Service)
					if !ok {
						app := []*ServerInfo{s}
						apps.Store(s.Service, app)
					} else {
						app := appI.([]*ServerInfo)
						// check alread exist
						exist := false
						for _, n := range app {
							if n.IP == s.IP {
								exist = true
							}
						}
						if !exist {
							app = append(app, s)
						}
						apps.Store(s.Service, app)
					}
				}
			}

		}
	}()
}

func GetServer(target string) *ServerInfo {
	if apps == nil {
		return nil
	}
	appI, ok := apps.Load(target)
	if !ok || appI == nil {
		return nil
	}

	app := appI.([]*ServerInfo)
	if len(app) == 0 {
		return nil
	}

	return app[0]
}

func InitEtcd(addrs []string) *Etcd {
	cfg := clientv3.Config{
		Endpoints:   addrs,
		DialTimeout: 10 * time.Second,
	}

	cli, err := clientv3.New(cfg)
	if err != nil {
		panic(fmt.Errorf("Etcd init error: %v", err))
	}

	// 检测etcd是否链接成功
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	_, err = cli.Cluster.MemberList(ctx)
	if err != nil {
		L.Fatal("etcd init failed", zap.Error(err))
	}
	return &Etcd{cli}
}

func (etcd *Etcd) queryAll() (*sync.Map, error) {
	res, err := etcd.Get(context.Background(), AppRootPath, clientv3.WithPrefix())
	if err != nil {
		return nil, err
	}

	// services := make(map[string][]*ServerInfo)
	apps := &sync.Map{}
	for _, kv := range res.Kvs {
		s, err := etcd.parseKV(kv)
		if err != nil {
			L.Warn("parse etcd kv error", zap.Error(err))
			continue
		}
		appI, ok := apps.Load(s.Service)
		if !ok {
			apps.Store(s.Service, []*ServerInfo{s})
		} else {
			app := appI.([]*ServerInfo)
			apps.Store(s.Service, append(app, s))
		}
	}

	return apps, nil
}

func (etcd *Etcd) parseKV(kv *mvccpb.KeyValue) (*ServerInfo, error) {
	// /tfgo/service/TFE/10.50.38.63 0.0
	load, err := strconv.ParseFloat(talent.Bytes2String(kv.Value), 64)
	if err != nil {
		return nil, fmt.Errorf("load error:%v,key:%s, val: %s", err, string(kv.Key), string(kv.Value))
	}

	if !bytes.HasPrefix(kv.Key, talent.String2Bytes(AppRootPath)) {
		return nil, fmt.Errorf("node status invalid :%s", kv.Key)
	}

	k1 := kv.Key[len(AppRootPath):]
	ks := bytes.Split(k1, []byte("/"))
	if len(ks) != 2 {
		return nil, fmt.Errorf("node status invalid:%s", kv.Key)
	}

	srv := talent.Bytes2String(ks[0])
	ip := talent.Bytes2String(ks[1])

	return &ServerInfo{srv, ip, load}, nil
}

func (etcd *Etcd) parseK(kv *mvccpb.KeyValue) (*ServerInfo, error) {
	if !bytes.HasPrefix(kv.Key, talent.String2Bytes(AppRootPath)) {
		return nil, fmt.Errorf("node status invalid :%s", kv.Key)
	}

	k1 := kv.Key[len(AppRootPath):]
	ks := bytes.Split(k1, []byte("/"))
	if len(ks) != 2 {
		return nil, fmt.Errorf("node status invalid:%s", kv.Key)
	}

	srv := talent.Bytes2String(ks[0])
	ip := talent.Bytes2String(ks[1])

	return &ServerInfo{srv, ip, 0}, nil
}

//  key='/APIsRootPath/ApiName/ip:port' val='load--path'
func (etcd *Etcd) storeServer(s *ServerInfo) error {
	key := AppRootPath + s.Service + "/" + s.IP
	val := strconv.FormatFloat(s.Load, 'f', 1, 64)

	grant, err := etcd.Grant(context.TODO(), AppLeaseTime)
	if err != nil {
		return err
	}

	_, err = etcd.Put(context.TODO(), key, val, clientv3.WithLease(grant.ID))
	if err != nil {
		return err
	}

	return nil
}
