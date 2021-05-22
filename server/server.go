/**
* Author: JeffreyBool
* Date: 2021/5/22
* Time: 10:26 下午
* Software: GoLand
 */

package server

import (
	"context"

	"github.com/go-hexo/go-hexo/registry"
)

// Server is a simple micro server abstraction
type Server interface {
	// Init Initialise options
	Init(...Option) error
	// Options Retrieve the options
	Options() Options
	// Handle Register a handler
	Handle(Handler) error
	// NewHandler Create a new handler
	NewHandler(interface{}, ...HandlerOption) Handler
	// Start Start the server
	Start() error
	// Stop Stop the server
	Stop() error
	// String Server implementation
	String() string
}

// Router handle serving messages
type Router interface {
	// ProcessMessage processes a message
	ProcessMessage(context.Context, Message) error
	// ServeRequest processes a request to completion
	ServeRequest(context.Context, Request, Response) error
}

// Message is an async message interface
type Message interface {
	// Topic of the message
	Topic() string
	// Payload The decoded payload value
	Payload() interface{}
	// ContentType The content type of the payload
	ContentType() string
	// Header The raw headers of the message
	Header() map[string]string
	// Body The raw body of the message
	Body() []byte
	// Codec used to decode the message
	Codec() codec.Reader
}

// Request is a synchronous request interface
type Request interface {
	// Service name requested
	Service() string
	// Method The action requested
	Method() string
	// Endpoint name requested
	Endpoint() string
	// ContentType Content type provided
	ContentType() string
	// Header of the request
	Header() map[string]string
	// Body is the initial decoded value
	Body() interface{}
	// Read the un decoded request body
	Read() ([]byte, error)
	// Codec The encoded message stream
	Codec() codec.Reader
	// Stream Indicates whether its a stream
	Stream() bool
}

// Response is the response writer for unencoded messages
type Response interface {
	// Codec Encoded writer
	Codec() codec.Writer
	// WriteHeader Write the header
	WriteHeader(map[string]string)
	// Write write a response directly to the client
	Write([]byte) error
}

// Stream represents a stream established with a client.
// A stream can be bidirectional which is indicated by the request.
// The last error will be left in Error().
// EOF indicates end of the stream.
type Stream interface {
	Context() context.Context
	Request() Request
	Send(interface{}) error
	Recv(interface{}) error
	Error() error
	Close() error
}

// Handler interface represents a request handler. It's generated
// by passing any type of public concrete object with endpoints into server.NewHandler.
// Most will pass in a struct.
//
// Example:
//
//      type Greeter struct {}
//
//      func (g *Greeter) Hello(context, request, response) error {
//              return nil
//      }
//
type Handler interface {
	Name() string
	Handler() interface{}
	Endpoints() []*registry.Endpoint
	Options() HandlerOptions
}

type Option func(*Options)
