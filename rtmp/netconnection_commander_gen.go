// Automatically generated by go generate; DO NOT EDIT.

package rtmp

import (
	"context"
)

type NetConnectionCommander interface {
	Connect(ctx context.Context, commandObject map[string]interface{}, optionalUserArguments map[string]interface{}) error
	ConnectResult(ctx context.Context, properties map[string]interface{}, information map[string]interface{}) error
	ConnectError(ctx context.Context, properties map[string]interface{}, information map[string]interface{}) error
	Call(ctx context.Context, procedureName string, transactionID uint32, commandObject map[string]interface{}, optionalArguments map[string]interface{}) error
	CallResponse(ctx context.Context, commandName string, transactionID uint32, commandObject map[string]interface{}, response map[string]interface{}) error
	CreateStream(ctx context.Context, transactionID uint32, commandObject map[string]interface{}) error
	CreateStreamResult(ctx context.Context, transactionID uint32, commandObject map[string]interface{}, streamID uint32) error
	CreateStreamError(ctx context.Context, transactionID uint32, commandObject map[string]interface{}, streamID uint32) error
}
