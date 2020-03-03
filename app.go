package go_demo

import (
	"go-demo/apis/gorpc"
	_ "go-demo/apis/web"
	_ "go-demo/core/accounts"
	_ "go-demo/core/envelopes"
	"go-demo/infra"
	"go-demo/infra/base"
)

func init() {
	infra.Register(&base.PropsStarter{})
	infra.Register(&base.DbxDatabaseStarter{})
	infra.Register(&base.ValidatorStarter{})
	infra.Register(&base.GoRPCStarter{})
	infra.Register(&gorpc.GoRpcApiStarter{})
	infra.Register(&base.IrisStarter{})
	infra.Register(&infra.WebApiStarter{})
}
