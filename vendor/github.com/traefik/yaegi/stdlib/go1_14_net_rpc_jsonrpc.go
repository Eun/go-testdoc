// Code generated by 'yaegi extract net/rpc/jsonrpc'. DO NOT EDIT.

// +build go1.14,!go1.15

package stdlib

import (
	"net/rpc/jsonrpc"
	"reflect"
)

func init() {
	Symbols["net/rpc/jsonrpc"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"Dial":           reflect.ValueOf(jsonrpc.Dial),
		"NewClient":      reflect.ValueOf(jsonrpc.NewClient),
		"NewClientCodec": reflect.ValueOf(jsonrpc.NewClientCodec),
		"NewServerCodec": reflect.ValueOf(jsonrpc.NewServerCodec),
		"ServeConn":      reflect.ValueOf(jsonrpc.ServeConn),
	}
}
