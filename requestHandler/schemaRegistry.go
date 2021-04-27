package requestHandler

import (
	"github.com/vektah/gqlparser/v2/ast"
)

type ServiceInfo struct {
	Internal      bool
	Introspection *ast.Schema
	ServiceUrl    *string
	ServiceName   string
}

type SchemaRegistry map[string]ServiceInfo

var globalSchemaRegistry SchemaRegistry = SchemaRegistry{}

func (global SchemaRegistry) PushServiceSchemaInSchemaRegistry(
	internalServiceAST *ast.Schema,
	serviceName string,
	internal bool,
	serviceUrl *string) *SchemaRegistry {

	serviceInfoStr := ServiceInfo{
		Introspection: internalServiceAST,
		ServiceName:   serviceName,
		Internal:      internal,
		ServiceUrl:    serviceUrl,
	}

	global[serviceName] = serviceInfoStr

	return &global

}

func GetSchemaRegistry() *SchemaRegistry {
	return &globalSchemaRegistry
}
