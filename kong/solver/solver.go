package solver

import (
	"apigw/kong/crud"
	"apigw/util/logs"
	"github.com/hbagdi/go-kong/kong"
)

// Stats holds the stats related to a Solve.
type Stats struct {
	CreateOps int
	UpdateOps int
	DeleteOps int
}

var stats Stats
var registry crud.Registry

func BuildRegistry(client *kong.Client) {
	registry.MustRegister("service", &serviceCRUD{client: client})
	registry.MustRegister("route", &routeCRUD{client: client})
	registry.MustRegister("upstream", &upstreamCRUD{client: client})
	registry.MustRegister("target", &targetCRUD{client: client})
	registry.MustRegister("certificate", &certificateCRUD{client: client})
	registry.MustRegister("sni", &sniCRUD{client: client})
	registry.MustRegister("ca_certificate", &caCertificateCRUD{client: client})
	registry.MustRegister("plugin", &pluginCRUD{client: client})
	registry.MustRegister("consumer", &consumerCRUD{client: client})
	registry.MustRegister("key-auth", &keyAuthCRUD{client: client})
	registry.MustRegister("hmac-auth", &hmacAuthCRUD{client: client})
	registry.MustRegister("jwt-auth", &jwtAuthCRUD{client: client})
	registry.MustRegister("basic-auth", &basicAuthCRUD{client: client})
	registry.MustRegister("acl-group", &aclGroupCRUD{client: client})
	registry.MustRegister("oauth2-cred", &oauth2CredCRUD{client: client})
}

func recordOp(op crud.Op) {
	switch op {
	case crud.Create:
		stats.CreateOps = stats.CreateOps + 1
	case crud.Update:
		stats.UpdateOps = stats.UpdateOps + 1
	case crud.Delete:
		stats.DeleteOps = stats.DeleteOps + 1
	}
}

func Solve(e crud.Event) (crud.Arg, error) {
	var err error
	var result crud.Arg

	switch e.Op {
	case crud.Create:
		logs.Info("creating", e.Kind)
	case crud.Update:
		diffString, err := getDiff(e.OldObj, e.Obj)
		if err != nil {
			return nil, err
		}
		logs.Info("updating", e.Kind, diffString)
	case crud.Delete:
		logs.Info("deleting", e.Kind)
	default:
		panic("unknown operation " + e.Op.String())
	}

	result, err = registry.Do(e.Kind, e.Op, e)
	if err != nil {
		return nil, err
	}

	// record operation in both: diff and sync commands
	recordOp(e.Op)

	return result, nil
}
