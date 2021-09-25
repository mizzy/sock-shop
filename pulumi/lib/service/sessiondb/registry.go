package sessiondb

import (
	"github.com/mizzy/sock-shop/pulumi/lib/ecs"
	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/servicediscovery"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var registry *servicediscovery.Service

func newRegistry(ctx *pulumi.Context) error {
	var err error
	registry, err = servicediscovery.NewService(ctx, "session_db", &servicediscovery.ServiceArgs{
		DnsConfig: &servicediscovery.ServiceDnsConfigArgs{
			DnsRecords: &servicediscovery.ServiceDnsConfigDnsRecordArray{
				&servicediscovery.ServiceDnsConfigDnsRecordArgs{
					Ttl:  pulumi.Int(10),
					Type: pulumi.String("A"),
				},
			},
			NamespaceId: ecs.PrivateDnsNamespace.ID(),
		},
		HealthCheckCustomConfig: &servicediscovery.ServiceHealthCheckCustomConfigArgs{
			FailureThreshold: pulumi.Int(1),
		},
		Name: pulumi.String("session-db"),
	})
	if err != nil {
		return err
	}

	return nil
}
