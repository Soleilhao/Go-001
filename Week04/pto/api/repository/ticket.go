package repository

import "time"

type TicketRepo struct {
}

func NewTicketRepo() TicketRepo {
	return TicketRepo{}
}
func (t *TicketRepo) Create() int {
	//save the ticket to DB and return the tickiet id
	//return dummy id
	return time.Now().Second()
}
