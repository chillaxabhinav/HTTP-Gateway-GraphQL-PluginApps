package requestHandler

import (
	"fmt"

	"github.com/vektah/gqlparser/v2"

	"github.com/vektah/gqlparser/v2/ast"
)

type ServiceQueryMapping map[string][]*ast.QueryDocument

// type SingleQueryBuilder struct {
// 	operationName       string // OperationType --> operation.name in loop
// 	parentType          string // OperationType --> operationType variable
// 	variables           ast.VariableDefinitionList
// 	selectionSet        ast.SelectionSet
// 	fragmentDefinitions ast.FragmentDefinitionList
// }

func GetPlan(requestContext *RequestContext) {

	parsedQuery, err := gqlparser.LoadQuery(MyMergedSchema, requestContext.Query)

	if err != nil {
		fmt.Println(err)
		return
	}

	for _, operation := range parsedQuery.Operations {

		// get the type for the operation
		operationType := "Query"

		switch operation.Operation {
		case ast.Mutation:
			operationType = "Mutation"
		case ast.Subscription:
			operationType = "Subscription"
		}

		for _, selectionSetSingle := range operation.SelectionSet {

			myQuery := BuildQuery(operation.Name, operationType, operation.VariableDefinitions, []ast.Selection{selectionSetSingle})

			fmt.Println(myQuery.Operations[0].SelectionSet[0])

			// validationFail := validator.Validate(globalSchemaRegistry["train"].Introspection, myQuery)

			// if validationFail == nil {
			// 	fmt.Println("Validation Pass")
			// }

			// for _, service := range globalSchemaRegistry {

			// 	fmt.Println("serviceName ", service.ServiceName)

			// 	validationFailed := validator.Validate(service.Introspection, myQuery)

			// 	if validationFailed == nil {
			// 		if len(globalServiceQueryMapping) == 0 {
			// 			globalServiceQueryMapping[service.ServiceName] = append(globalServiceQueryMapping[service.ServiceName], myQuery)
			// 		} else {
			// 			_, ok := globalServiceQueryMapping[service.ServiceName]
			// 			if !ok {
			// 				globalServiceQueryMapping[service.ServiceName] = append(globalServiceQueryMapping[service.ServiceName], myQuery)
			// 			} else {
			// 				globalServiceQueryMapping[service.ServiceName] = append(globalServiceQueryMapping[service.ServiceName], myQuery)
			// 			}
			// 		}
			// 	}

			// }

		}

	}

}
