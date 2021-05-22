/**
* Author: JeffreyBool
* Date: 2021/5/22
* Time: 11:00 下午
* Software: GoLand
 */

package server

import (
	"context"
	"time"

	"github.com/google/uuid"
)

var (
	DefaultAddress                 = ":0"
	DefaultName                    = "go.micro.server"
	DefaultVersion                 = "latest"
	DefaultId                      = uuid.New().String()
	DefaultServer           Server = newRpcServer()
	DefaultRouter                  = newRpcRouter()
	DefaultRegisterCheck           = func(context.Context) error { return nil }
	DefaultRegisterInterval        = time.Second * 30
	DefaultRegisterTTL             = time.Second * 90

	// NewServer creates a new server
	NewServer func(...Option) Server = newRpcServer
)
