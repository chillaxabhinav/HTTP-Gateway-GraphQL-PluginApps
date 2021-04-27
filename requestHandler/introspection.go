package requestHandler

import "encoding/json"

type IntrospectionQueryResult struct {
	Schema *IntrospectionQuerySchema `json:"__schema"`
}

type IntrospectionQuerySchema struct {
	QueryType        IntrospectionQueryRootType    `json:"queryType"`
	MutationType     *IntrospectionQueryRootType   `json:"mutationType"`
	SubscriptionType *IntrospectionQueryRootType   `json:"subscriptionType"`
	Types            []IntrospectionQueryFullType  `json:"types"`
	Directives       []IntrospectionQueryDirective `json:"directives"`
}

type IntrospectionQueryDirective struct {
	Name        string                    `json:"name"`
	Description string                    `json:"description"`
	Locations   []string                  `json:"locations"`
	Args        []IntrospectionInputValue `json:"arg"`
}

type IntrospectionQueryRootType struct {
	Name string `json:"name"`
}

type IntrospectionQueryFullTypeField struct {
	Name              string                    `json:"name"`
	Description       string                    `json:"description"`
	Args              []IntrospectionInputValue `json:"args"`
	Type              IntrospectionTypeRef      `json:"type"`
	IsDeprecated      bool                      `json:"isDeprecated"`
	DeprecationReason string                    `json:"deprecationReason"`
}

type IntrospectionQueryFullType struct {
	Kind          string                             `json:"kind"`
	Name          string                             `json:"name"`
	Description   string                             `json:"description"`
	InputFields   []IntrospectionInputValue          `json:"inputFields"`
	Interfaces    []IntrospectionTypeRef             `json:"interfaces"`
	PossibleTypes []IntrospectionTypeRef             `json:"possibleTypes"`
	Fields        []IntrospectionQueryFullTypeField  `json:"fields"`
	EnumValues    []IntrospectionQueryEnumDefinition `json:"enumValues"`
}

type IntrospectionQueryEnumDefinition struct {
	Name              string `json:"name"`
	Description       string `json:"description"`
	IsDeprecated      bool   `json:"isDeprecated"`
	DeprecationReason string `json:"deprecationReason"`
}

type IntrospectionInputValue struct {
	Name         string               `json:"name"`
	Description  string               `json:"description"`
	DefaultValue string               `json:"defaultValue"`
	Type         IntrospectionTypeRef `json:"type"`
}

type IntrospectionTypeRef struct {
	Kind   string                `json:"kind"`
	Name   string                `json:"name"`
	OfType *IntrospectionTypeRef `json:"ofType"`
}

func MakeIntroSpectionQueryResponse() *IntrospectionQuerySchema {

	myMarshal, _ := json.Marshal(MyMergedSchema)

	myRes := IntrospectionQuerySchema{}

	err := json.Unmarshal(myMarshal, &myRes)

	if err != nil {
		panic(err)
	}

	return &myRes

}
