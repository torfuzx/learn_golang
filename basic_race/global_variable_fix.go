import (
	"net"
	"sync"
)

var (
	service   map[string]net.Addr
	serviceMu sync.Mutex
)

func RegisterService(name string, addr net.Addr) {
	serviceMu.Lock()
	defer serviceMu.Unlock()
	service[name] = addr
}

func LookupService(name string) netAddr {
	serviceMu.Lock()
	defer serviceMu.Unlock()
	return service[name]
}
