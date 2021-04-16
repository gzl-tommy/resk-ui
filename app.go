package resk

import (
	_ "github.com/gzl-tommy/account"
	"github.com/gzl-tommy/infra"
	"github.com/gzl-tommy/infra/base"
	_ "github.com/gzl-tommy/resk/core/envelopes"
	_ "gzl-tommy/resk-ui/views"
)

func init() {
	infra.Register(&base.PropsStarter{})
	infra.Register(&base.DbxDatabaseStarter{})
	infra.Register(&base.ValidatorStarter{})
	infra.Register(&base.IrisServerStarter{})
	infra.Register(&infra.WebApiStarter{})
	infra.Register(&base.EurekaStarter{})
	infra.Register(&base.HookStarter{})
}
