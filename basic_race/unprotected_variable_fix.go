import (
	"fmt"
	"os"
	"sync/atomic"
	"time"
)

type Watchdog struct {
	last int64
}

func (w *Watchdog) KeepAlive() {
	atomic.StoreInt64(&w.last, time.Now().UnixNano())
}

func (w *Watchdog) Start() {
	go func() {
		for {
			time.Sleep(time.Second)
			if atomic.LoadInt64(&w.last) < time.Now().Add(-10*time.Second).UnixNano() {
				fmt.Println("No keepalives for 10 seconds, dying...")
				os.Exit(1)
			}
		}
	}()
}
