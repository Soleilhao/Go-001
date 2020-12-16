package controllers

import (
	"Go-001/Week04/pto/api/domain"
	"fmt"

	"github.com/gin-gonic/gin"
)

type TicketController struct {
	Ticket domain.Ticket
}

func NewTicketController(ticket domain.Ticket) TicketController {
	return TicketController{Ticket: ticket}
}

type TicketDto struct {
	Id       int    `form:"id"`
	Comments string `form:"comments"`
}

func (t *TicketController) RaiseTicket(c *gin.Context) string {
	var dto TicketDto
	if c.ShouldBind(&dto) == nil {
		fmt.Println(dto.Comments)
	}
	t.Ticket.RaisePTORequest(dto.Comments)
	return fmt.Sprintf("Ticket submitted --> id:%d  comments:%s", t.Ticket.Id, t.Ticket.Comments)
}
