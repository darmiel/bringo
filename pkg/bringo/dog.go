package bringo

import (
	"fmt"
	"time"
)

// expireDog renews token when they are to Expires
type expireDog struct {
	Expires time.Time
	ticker  *time.Ticker
	closer  chan bool
}

func newExpireDog(expiresIn int) (dog *expireDog) {
	dog = &expireDog{
		Expires: time.Now().Add((time.Duration(expiresIn - 10)) * time.Second),
		ticker:  time.NewTicker(time.Second),
		closer:  make(chan bool),
	}
	// start Expires Dog
	go dog.listen()
	return
}

func (d *expireDog) Stop() {
	if d.ticker != nil {
		d.ticker.Stop()
		d.ticker = nil
	}
	d.closer <- true
}

func (d *expireDog) listen() {
	for {
		select {
		case <-d.ticker.C:
			d.Check()
		case <-d.closer:
			return
		}
	}
}

func (d *expireDog) Check() {
	if d.Expires.Before(time.Now()) {
		return
	}
	fmt.Println("[d] token expired. renew!")
	// TODO: renew token
}
