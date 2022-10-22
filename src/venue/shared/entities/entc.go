// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

//go:build ignore
// +build ignore

package main

import (
	"log"
	"path"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

func main() {
	ex, err := entgql.NewExtension(
		entgql.WithWhereFilters(true),
		entgql.WithSchemaPath(path.Join("..", "..", "..", "..", "api-definitions", "graphql", "omiga", "venue", "V1", "ent.graphql")),
		entgql.WithConfigPath(path.Join("..", "graphql", "gqlgen.yml")),
	)
	if err != nil {
		log.Fatalf("creating entgql extension: %v", err)
	}

	if err = entc.Generate(
		"./schema",
		&gen.Config{
			Features: gen.AllFeatures,
		},
		entc.Extensions(ex)); err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}
