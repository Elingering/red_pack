package go_demo

import (
	"go-demo/infra"
	"go-demo/infra/base"
)

func init() {
	infra.Register(&base.PropsStarter{})
	infra.Register(&base.DbxDatabaseStarter{})
	infra.Register(&base.ValidatorStarter{})
	infra.Register(&base.IrisStarter{})
}
