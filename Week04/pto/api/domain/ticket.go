package domain

import "Go-001/Week04/pto/api/repository"

type Ticket struct {
	Id       int
	Comments string
	Repo     repository.TicketRepo
}

func NewTicket(repo repository.TicketRepo) Ticket {
	return Ticket{Repo: repo}
}

func (t *Ticket) RaisePTORequest(comments string) {
	t.Comments = comments
	t.Id = t.Repo.Create()
}
