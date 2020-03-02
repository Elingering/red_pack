package base

import (
	"fmt"
	"github.com/tietang/props/kvs"
	"go-demo/infra"
	"sync"
)

var props kvs.ConfigSource

func Props() kvs.ConfigSource {
	return props
}

type PropsStarter struct {
	infra.BaseStarter
}

func (p *PropsStarter) Init(ctx infra.StarterContext) {
	props = ctx.Props()
	fmt.Println("初始化配置.")
	GetSystemAccount()
}

type SystemAccount struct {
	AccountNo   string
	AccountName string
	UserId      string
	Username    string
}

var systemAccount *SystemAccount
var systemAccountOnce sync.Once

func GetSystemAccount() *SystemAccount {
	systemAccountOnce.Do(func() {
		systemAccount = new(SystemAccount)
		err := kvs.Unmarshal(Props(), systemAccount, "system.account")
		if err != nil {
			panic(err)
		}
	})
	return systemAccount
}

func GetEnvelopeActivityLink() string {
	link := Props().GetDefault("envelope.domain", "/v1/envelop/link")
	return link
}

func GetEnvelopeDomain() string {
	domain := Props().GetDefault("envelope.domain", "http://localhost")
	return domain
}
