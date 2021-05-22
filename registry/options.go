/**
* Author: JeffreyBool
* Date: 2021/5/22
* Time: 10:48 下午
* Software: GoLand
 */

package registry

import (
	"context"
	"crypto/tls"
	"time"
)

type Options struct {
	Addrs     []string
	Timeout   time.Duration
	Secure    bool
	TLSConfig *tls.Config
	// Other options for implementations of the interface
	// can be stored in a context
	Context context.Context
}

type RegisterOptions struct {
	TTL time.Duration
	// Other options for implementations of the interface
	// can be stored in a context
	Context context.Context
}

type WatchOptions struct {
	// Specify a service to watch
	// If blank, the watch is for all services
	Service string
	// Other options for implementations of the interface
	// can be stored in a context
	Context context.Context
}
type DeregisterOptions struct {
	Context context.Context
}

type GetOptions struct {
	Context context.Context
}

type ListOptions struct {
	Context context.Context
}
