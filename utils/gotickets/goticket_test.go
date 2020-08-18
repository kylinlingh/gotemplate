package gotickets

import (
	"github.com/pkg/errors"
	"github.com/smartystreets/goconvey/convey"
	"gotemplate/utils/log"
	"os"
	"testing"
	"time"
)

type obj struct{
	tickets GoRoutinueTickets
}

var concurrency uint32 = 3 // count of go routinue

func (o *obj) init() error{
	tickets, err := NewGoTickets(concurrency)
	if err != nil {
		return errors.Wrapf(err, "Initialization aborted.")
	}
	o.tickets = tickets
	return nil
}

func (o *obj) asyncCall(i int) {
	o.tickets.Take() //this call will wait until one ticket can be acquired.
	go func() {
		defer func() {
			// do something like recover
			//if p := recover();p != nil {
			//
			//}

			o.tickets.Return() // return back the ticket
		}()
		log.GlobalLogger.Debug(i)
		time.Sleep(1 * time.Second)
	}()
}

func TestNewGoTickets(t *testing.T) {
	convey.Convey("example for using go tickets.\n", t, func() {
		o := obj{}
		if err := o.init(); err != nil {
			log.GlobalLogger.Error(err)
			os.Exit(-1)
		}
		for i := 0; i < 20; i++{
			o.asyncCall(i)
		}
	})
}
