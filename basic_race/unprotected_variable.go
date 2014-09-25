import "time"

type WatchDog struct {
	last int64
}

func (w *WatchDog) KeepAlive() {
	w.last = time.Now().UnixNano()
}

func (w *WatchDog) Start() {
    go func () {
        for {
            time.Sleep(time.Second)
            // second conflicting access
            if w.last < time.Now().Add(-10*time.Second).UnixNano() {
                fmt.Println("No keepalives for 10 seconds, dying...")
                os.Exit(1)
            }
        }
    }
}
