package requestHandler

import (
	"github.com/vektah/gqlparser/v2/ast"

	"github.com/gin-gonic/gin"
)

type MergedSchemaMap *ast.Schema

var MyMergedSchema MergedSchemaMap

func Init() gin.HandlerFunc {
	return func(c *gin.Context) {
		InternalServiceEntry()

		ExternalServiceEntry()

		merged, _ := SchemaMerge()

		MyMergedSchema = merged
	}
}
