/**
* Author: JeffreyBool
* Date: 2021/5/22
* Time: 10:44 下午
* Software: GoLand
 */

package server

// HandlerOption handler option.
type HandlerOption func(*HandlerOptions)

// HandlerOptions handler options.
type HandlerOptions struct {
	Internal bool
	Metadata map[string]map[string]string
}

// EndpointMetadata is a Handler option that allows metadata to be added to
// individual endpoints.
func EndpointMetadata(name string, md map[string]string) HandlerOption {
	return func(o *HandlerOptions) {
		o.Metadata[name] = md
	}
}

// InternalHandler Internal Handler options specifies that a handler is not advertised
// to the discovery system. In the future this may also limit request
// to the internal network or authorised user.
func InternalHandler(b bool) HandlerOption {
	return func(o *HandlerOptions) {
		o.Internal = b
	}
}
