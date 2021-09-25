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
	}

	for _, s := range services {
		err := s(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}
