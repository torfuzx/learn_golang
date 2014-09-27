package main

// The closure does a non-blocking send, which is achieved by using the send operation in
// a select statement with a default case. If the send cannot go through immediately, the
// default case will be selected. Making the send non-blocking guarantees that non of the
// goroutines launched in the loop wil hang around.

// However, if the result arrives before
// the main function has made it to the receive, the send could fail since no one is
// ready.

// To fix this race condition, make sure to buffer the channel ch, guarantteeing tha
// the first send has a place to put the value. This ensures the send will always
// succeed, and the first value to arrive will be retrieved regardless of the order of
// execution

func Quermamy(conns []Conn, query string) Result {
	ch := make(chan Result, 1)
	for _, conn := range conns {
		go func(c Conn) {
			select {
			case ch <- c.DoQuery(query):
			default:
			}
		}(conn)
	}
}
