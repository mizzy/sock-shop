package catalogue

import (
	"github.com/mizzy/sock-shop/pulumi/lib/ecs"
	"github.com/mizzy/sock-shop/pulumi/lib/vpc"
	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func NewRdsImporter(ctx *pulumi.Context) error {
	_, err := ec2.NewInstance(ctx, "rds_importer", &ec2.InstanceArgs{
		Ami:                      pulumi.String("ami-56d4ad31"),
		AssociatePublicIpAddress: pulumi.Bool(true),
		InstanceType:             pulumi.String("t2.micro"),
		SubnetId:                 vpc.PublicSubnet1.ID(),
		Tags: pulumi.StringMap{
			"Name": pulumi.String("RDS Importer - sock-shop"),
		},
		UserData:            pulumi.Sprintf("#!/bin/bash -xe\nyum -y install mysql\nwget https://raw.githubusercontent.com/microservices-demo/catalogue/master/docker/catalogue-db/data/dump.sql\nmysql -u catalogue_user --password=default_password -h %s \\\n  -f -D socksdb < dump.sql", dbInstance.Address),
		VpcSecurityGroupIds: pulumi.StringArray{ecs.EcsSecurityGroup.ID()},
	})
	if err != nil {
		return err
	}

	return nil
}
