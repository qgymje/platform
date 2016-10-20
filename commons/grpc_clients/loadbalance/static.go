package loadbalance

import "google.golang.org/grpc/naming"

// StaticResolver implements a gRPC resolver/watcher that simply returns
// a list of address, then blocks.
type StaticResolver struct {
	addr []*naming.Update
}

// NewStaticResolver initializes and returns a new StaticResolver
func NewStaticResolver(addr ...string) *StaticResolver {
	r := &StaticResolver{}
	for _, a := range addr {
		r.addr = append(r.addr, &naming.Update{Op: naming.Addr, Addr: a})
	}
	return r
}

// Resolve creates a watcher for target. The watcher interface is implementd
// by StaticResolver as well, see Next and Close.
func (r *StaticResolver) Resolve(target string) (naming.Watcher, error) {
	return r, nil
}

// Next returns the list of address once, then blocks on consecutive calls
func (r *StaticResolver) Next() ([]*naming.Update, error) {
	if r.addr != nil {
		updates := r.addr
		r.addr = nil
		return updates, nil
	}
	infinite := make(chan struct{})
	<-infinite
	return nil, nil
}

// Close is a no-op for StaticResolver
func (r *StaticResolver) Close() {}
