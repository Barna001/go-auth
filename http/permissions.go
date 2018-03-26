package http

import (
	"reflect"
)

type Permission string

const (
	GetAllUser Permission = "GetAllUser"
	GetOneUser Permission = "GetOneUser"
)

func (permission Permission) String() string {
	return reflect.ValueOf(permission).String()
}
