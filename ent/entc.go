//go:build ignore
// +build ignore

package main

import (
	"log"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

func main() {
	// customWhere := gen.MustParse(gen.NewTemplate("../templates/where_input.tmpl").
	// 	Funcs(gen.Funcs).
	// 	Funcs(entgql.TemplateFuncs).
	// 	ParseFiles("../templates/where_input.tmpl"))
	//
	// entgqlTemplates := []*gen.Template{
	// 	entgql.CollectionTemplate,
	// 	entgql.EnumTemplate,
	// 	entgql.NodeTemplate,
	// 	entgql.PaginationTemplate,
	// 	entgql.TransactionTemplate,
	// 	entgql.EdgeTemplate,
	// 	customWhere,
	// }

	ex, err := entgql.NewExtension(
		entgql.WithWhereFilters(true),
		entgql.WithSchemaPath("../ent.graphql"),
		entgql.WithConfigPath("../gqlgen.yml"),
		// entgql.WithTemplates(entgqlTemplates...),
	)
	if err != nil {
		log.Fatalf("creating entgql extension: %v", err)
	}
	err = entc.Generate("./schema", &gen.Config{},
		entc.Extensions(ex),
	)
	if err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}
