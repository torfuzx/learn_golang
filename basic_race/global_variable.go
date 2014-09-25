import "net"

var service map[string]net.Addr

func RegisterService(name string, add net.Addr) {
	service[name] = addr
}

func LookupService(name string) net.Addr {
	return service[name]
}
