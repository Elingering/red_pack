package go_demo

import (
	_ "go-demo/apis/web"
	_ "go-demo/core/accounts"
	"go-demo/infra"
	"go-demo/infra/base"
)

func init() {
	infra.Register(&base.PropsStarter{})
	infra.Register(&base.DbxDatabaseStarter{})
	infra.Register(&base.ValidatorStarter{})
	infra.Register(&base.IrisStarter{})
	infra.Register(&infra.WebApiStarter{})
}
