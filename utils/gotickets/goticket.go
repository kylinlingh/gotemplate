package gotickets

import (
	"fmt"
	"github.com/pkg/errors"
)

// The interface of go routinue pool
type GoRoutinueTickets interface {

	// take one ticket
	Take()

	// return one ticket
	Return()

	// check the pool is active
	Active() bool

	// count of all tickets
	Total() uint32

	// remain count of tickets
	Remainder() uint32
}

// myGoTickets implements the interface GoRoutinueTickets
type myGoTickets struct {
	total    uint32        // total count of tickets
	ticketCh chan struct{} // channel of ticket
	active   bool
}

// new a ticket pool
func NewGoTickets(total uint32) (GoRoutinueTickets, error) {
	gt := myGoTickets{}
	if !gt.init(total) {
		errMsg :=
			fmt.Sprintf("The goroutine ticket pool can NOT be initialized! (total=%d)\n", total)
		return nil, errors.New(errMsg)
	}
	return &gt, nil
}

func (gt *myGoTickets) init(total uint32) bool {
	if gt.active {
		return false
	}
	if total == 0 {
		return false
	}
	ch := make(chan struct{}, total)
	n := int(total)
	for i := 0; i < n; i++ {
		ch <- struct{}{}
	}
	gt.ticketCh = ch
	gt.total = total
	gt.active = true
	return true
}

func (gt *myGoTickets) Take() {
	<-gt.ticketCh
}

func (gt *myGoTickets) Return() {
	gt.ticketCh <- struct{}{}
}

func (gt *myGoTickets) Active() bool {
	return gt.active
}

func (gt *myGoTickets) Total() uint32 {
	return gt.total
}

func (gt *myGoTickets) Remainder() uint32 {
	return uint32(len(gt.ticketCh))
}
