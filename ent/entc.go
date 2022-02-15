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
	customPagination := gen.MustParse(gen.NewTemplate("template/pagination.tmpl").
		Funcs(gen.Funcs).
		Funcs(entgql.TemplateFuncs).
		ParseFiles("../templates/pagination.tmpl"))

	entgqlTemplates := []*gen.Template{
		entgql.CollectionTemplate,
		entgql.EnumTemplate,
		entgql.GlobalIDTemplate,
		entgql.NodeTemplate,
		customPagination,
		entgql.TransactionTemplate,
		entgql.EdgeTemplate,
	}

	ex, err := entgql.NewExtension(
		entgql.WithWhereFilters(true),
		entgql.WithSchemaPath("../ent.graphql"),
		entgql.WithConfigPath("../gqlgen.yml"),
		entgql.WithTemplates(entgqlTemplates...),
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
