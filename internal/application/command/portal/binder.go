package portal

import (
	"airplane/internal/application/server/portal"
	"airplane/internal/components/apis"
	"airplane/internal/components/binder"
	"airplane/internal/components/launcher"
	"airplane/internal/components/logger"
	"airplane/internal/config"
	bookingCtrl "airplane/internal/controller/portal/rest/bookingpkg"
	flightCtrl "airplane/internal/controller/portal/rest/flight"
	paymentCtrl "airplane/internal/controller/portal/rest/payment"
	"airplane/internal/core/business/payment"
	"airplane/internal/core/business/seat"
	"airplane/internal/core/business/user"
	"airplane/internal/core/repositories/rdb"
	"airplane/internal/core/repositories/redis"
	"airplane/internal/core/usecase/booking"
	"airplane/internal/core/usecase/flight"
	paymentUsecase "airplane/internal/core/usecase/payment"
	"airplane/internal/infrastructure/database"
	"airplane/internal/infrastructure/mqcli"
	"airplane/internal/infrastructure/rediscli"
	"airplane/internal/infrastructure/redlock"
	"airplane/internal/infrastructure/snowflake"
	"go.uber.org/dig"
)

// container 引用 binder.Container 是因為我懶得寫 if error panic !!
func newBinder() *binder.Container {
	return binder.NewContainer(dig.New(), &container{})
}

type container struct {
}

func (b *container) Provider() *binder.Provider {
	provider := binder.NewProvider()
	b.registerProvides(provider)

	return provider
}
func (b *container) Invoker() *binder.Invoker {
	invoker := binder.NewInvoker()
	return invoker
}

func (b *container) registerProvides(provider *binder.Provider) {
	// base tools
	provider.Provide(config.NewConfig)
	provider.Provide(logger.New)
	provider.Provide(launcher.NewLauncher)
	provider.Provide(apis.New)
	// infrastructure
	provider.Provide(database.NewDatabaseClient)
	provider.Provide(rediscli.NewCacheClient)
	provider.Provide(snowflake.New)
	provider.Provide(redlock.New)
	provider.Provide(mqcli.NewMQClient)
	// repository
	provider.Provide(rdb.New)
	provider.Provide(redis.New)
	// application
	provider.Provide(portal.NewServer)
	// business
	provider.Provide(user.New)
	provider.Provide(payment.New)
	provider.Provide(seat.New)
	// usecase
	provider.Provide(booking.New)
	provider.Provide(flight.New)
	provider.Provide(paymentUsecase.New)
	// controller
	provider.Provide(bookingCtrl.New)
	provider.Provide(paymentCtrl.New)
	provider.Provide(flightCtrl.New)
}
