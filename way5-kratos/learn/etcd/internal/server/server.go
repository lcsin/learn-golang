package server

import (
	"etcd/internal/conf"
	etcdregistry "github.com/go-kratos/etcd/registry"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/google/wire"
	etcd "go.etcd.io/etcd/client/v3"
	"log"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(NewHTTPServer, NewGRPCServer, NewRegistrar)

func NewRegistrar(conf *conf.Registry) registry.Registrar {
	client, err := etcd.New(etcd.Config{
		Endpoints: []string{conf.Etcd.Address},
	})
	if err != nil {
		log.Fatal(err)
	}
	r := etcdregistry.New(client)
	return r
}
