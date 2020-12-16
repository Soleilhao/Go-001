//+build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"Go-001/Week04/pto/api/controllers"
	"Go-001/Week04/pto/api/domain"
	"Go-001/Week04/pto/api/repository"

	"github.com/google/wire"
)

func InitTicketController() controllers.TicketController {
	wire.Build(controllers.NewTicketController, domain.NewTicket, repository.NewTicketRepo)
	return controllers.TicketController{}
}
