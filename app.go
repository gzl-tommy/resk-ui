package resk

import (
	"github.com/gzl-tommy/resk/apis/gorpc"
	_ "github.com/gzl-tommy/resk/apis/web"
	_ "github.com/gzl-tommy/account"
	_ "github.com/gzl-tommy/resk/core/envelopes"
	"github.com/gzl-tommy/infra"
	"github.com/gzl-tommy/infra/base"
	"github.com/gzl-tommy/resk/jobs"
	_ "github.com/gzl-tommy/resk/public/ui"
)

func init() {
	infra.Register(&base.PropsStarter{})
	infra.Register(&base.DbxDatabaseStarter{})
	infra.Register(&base.ValidatorStarter{})
	infra.Register(&base.GoRPCStarter{})
	infra.Register(&gorpc.GoRpcApiStarter{})
	infra.Register(&jobs.RefundExpiredJobStarter{})
	infra.Register(&base.IrisServerStarter{})
	infra.Register(&infra.WebApiStarter{})
	infra.Register(&base.EurekaStarter{})
	infra.Register(&base.HookStarter{})
}
