/**
* Author: JeffreyBool
* Date: 2021/5/22
* Time: 11:15 下午
* Software: GoLand
 */

package server

import (
	"sync"
)

type rpcServer struct {
	sync.RWMutex
	opts Options
}

func newRpcServer() Server {
	return &rpcServer{}
}

func (s *rpcServer) Init(opts ...Option) error {
	s.Lock()
	defer s.Unlock()

}

func (s *rpcServer) Options() Options {
	s.RLock()
	opts := s.opts
	s.RUnlock()
	return opts
}

func (s *rpcServer) Handle(handler Handler) error {
	panic("implement me")
}

func (s *rpcServer) NewHandler(handler interface{}, option ...HandlerOption) Handler {
	panic("implement me")
}

func (s *rpcServer) Start() error {
	panic("implement me")
}

func (s *rpcServer) Stop() error {
	panic("implement me")
}

func (s *rpcServer) String() string {
	panic("implement me")
}
