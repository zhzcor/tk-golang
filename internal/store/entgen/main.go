package main

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"fmt"
	"log"
)

func main() {
	err := entc.Generate("./schema", &gen.Config{
		Hooks: []gen.Hook{
			TagFields("json"),
		},
		Target:  "./ent",
		Package: "tkserver/internal/store/ent",
	}, entc.TemplateDir("./enttemplate"), entc.Annotations(entsql.Annotation{OnDelete: entsql.NoAction}))
	if err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}

}

// TagFields tags all fields defined in the schema with the given struct-tag.
func TagFields(name string) gen.Hook {
	return func(next gen.Generator) gen.Generator {
		return gen.GenerateFunc(func(g *gen.Graph) error {
			for _, node := range g.Nodes {
				for _, field := range node.Fields {
					field.StructTag = fmt.Sprintf("%s:%q", name, field.Name)
				}
			}
			return next.Generate(g)
		})
	}
}
