package requestHandler

import (
	"errors"

	graphql "github.com/nautilus/graphql"
	"github.com/vektah/gqlparser/v2/ast"

	trainGenerated "gatewayCore/train/graph/generated"

	trainGraph "gatewayCore/train/graph"

	flightGenerated "gatewayCore/flight/graph/generated"

	flightGraph "gatewayCore/flight/graph"
)

type ExternalServiceInformation struct {
	ServiceName string
	ServiceUrl  string
}

type ExternalService map[string]ExternalServiceInformation

func introspectExternalServiceSchema(externalGraphqlUrl string) (*ast.Schema, error) {

	externalRequestQuereyer := graphql.NewSingleRequestQueryer(externalGraphqlUrl)

	introspectedSchemaAST, err := graphql.IntrospectAPI(externalRequestQuereyer)

	if err != nil {
		return nil, errors.New("introspection failed")
	}

	return introspectedSchemaAST, nil

}

func InternalServiceSchemaEntry(internalServiceAST *ast.Schema, serviceName string) *SchemaRegistry {

	global := GetSchemaRegistry()

	global.PushServiceSchemaInSchemaRegistry(internalServiceAST, serviceName, true, nil)

	return global

}

// Returns globalMapper and boolean value to identify if we were successful in getting the introspection
func ExternalServiceSchemaEntry(externalServiceURL string, serviceName string) (*SchemaRegistry, bool) {

	global := GetSchemaRegistry()

	externalSchemaAST, err := introspectExternalServiceSchema(externalServiceURL)

	if err != nil {
		return global, false
	}

	global.PushServiceSchemaInSchemaRegistry(externalSchemaAST, serviceName, false, &externalServiceURL)

	return global, true
}

func InternalServiceEntry() {

	flightServiceAST := flightGenerated.NewExecutableSchema(flightGenerated.Config{Resolvers: &flightGraph.Resolver{}}).Schema()

	InternalServiceSchemaEntry(flightServiceAST, "flight")

	trainServiceAST := trainGenerated.NewExecutableSchema(trainGenerated.Config{Resolvers: &trainGraph.Resolver{}}).Schema()

	InternalServiceSchemaEntry(trainServiceAST, "train")

}

func ExternalServiceEntry() {

	external := ExternalService{
		"search": ExternalServiceInformation{
			ServiceName: "search",
			ServiceUrl:  "http://localhost:8000/search",
		}}

	//external := ExternalService{}

	for _, val := range external {
		ExternalServiceSchemaEntry(val.ServiceUrl, val.ServiceName)
	}

}
