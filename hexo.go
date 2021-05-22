/**
* Author: JeffreyBool
* Date: 2021/5/19
* Time: 12:53 上午
* Software: GoLand
 */

package hexo

// Service is an interface that wraps the lower level libraries
// within go-hexo. Its a convenience method for building
// and initialising services.
type Service interface {
	// Name The service name
	Name() string
	// Init initialises options
	Init(...Option)
	// Options returns the current options
	Options() Options
	// Client is used to call services
	Client() client.Client
	// Server is for handling requests and events
	Server() server.Server
	// Run the service
	Run() error
	// String he service implementation
	String() string
}
