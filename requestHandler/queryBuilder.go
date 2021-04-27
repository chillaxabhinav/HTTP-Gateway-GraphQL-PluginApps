package requestHandler

import (
	"github.com/vektah/gqlparser/v2/ast"
)

func BuildQuery(
	operationName,
	parentType string,
	variables ast.VariableDefinitionList,
	selectionSet ast.SelectionSet,
	// fragmentDefinitions ast.FragmentDefinitionList,
) *ast.QueryDocument {
	// build up an operation for the query
	operation := &ast.OperationDefinition{
		VariableDefinitions: variables,
		Name:                operationName,
	}

	// assign the right operation
	switch parentType {
	case "Mutation":
		operation.Operation = ast.Mutation
	case "Subscription":
		operation.Operation = ast.Subscription
	default:
		operation.Operation = ast.Query
	}

	// if we are querying an operation all we need to do is add the selection set at the root
	if parentType == "Query" || parentType == "Mutation" || parentType == "Subscription" {
		operation.SelectionSet = selectionSet
		operation.VariableDefinitions = variables
	}

	// add the operation to a QueryDocument
	return &ast.QueryDocument{
		Operations: ast.OperationList{operation},
		// Fragments:  fragmentDefinitions,
	}
}
