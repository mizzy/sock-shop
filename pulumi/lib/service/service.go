package service

import (
	"github.com/mizzy/sock-shop/pulumi/lib/service/carts"
	"github.com/mizzy/sock-shop/pulumi/lib/service/cartsdb"
	"github.com/mizzy/sock-shop/pulumi/lib/service/catalogue"
	"github.com/mizzy/sock-shop/pulumi/lib/service/frontend"
	"github.com/mizzy/sock-shop/pulumi/lib/service/orders"
	"github.com/mizzy/sock-shop/pulumi/lib/service/payment"
	"github.com/mizzy/sock-shop/pulumi/lib/service/queuemaster"
	"github.com/mizzy/sock-shop/pulumi/lib/service/rabbitmq"
	"github.com/mizzy/sock-shop/pulumi/lib/service/sessiondb"
	"github.com/mizzy/sock-shop/pulumi/lib/service/shipping"
	"github.com/mizzy/sock-shop/pulumi/lib/service/user"
	"github.com/mizzy/sock-shop/pulumi/lib/service/userdb"
	"github.com/mizzy/sock-shop/pulumi/lib/service/zipkin"
	"github.com/mizzy/sock-shop/pulumi/lib/service/zipkincron"
	"github.com/mizzy/sock-shop/pulumi/lib/service/zipkindb"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func NewServices(ctx *pulumi.Context) error {
	services := []func(ctx2 *pulumi.Context) error{
		frontend.NewFrontEnd,
		carts.NewCarts,
		cartsdb.NewCartsDB,
		catalogue.NewCatalogue,
		orders.NewOrders,
		payment.NewPayment,
		queuemaster.NewQueueMaster,
		rabbitmq.NewRabbitMq,
		sessiondb.NewSessionDB,
		shipping.NewShipping,
		user.NewUser,
		userdb.NewUserDB,
		zipkin.NewZipkin,
		zipkincron.NewZipkinCron,
		zipkindb.NewZipkinDb,
	}

	for _, s := range services {
		err := s(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}
