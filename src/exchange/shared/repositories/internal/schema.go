// Code generated by ent, DO NOT EDIT.

//go:build tools
// +build tools

// Package internal holds a loadable version of the latest schema.
package internal

const Schema = `{"Schema":"github.com/omiga-group/omiga/src/exchange/shared/repositories/schema","Package":"github.com/omiga-group/omiga/src/exchange/shared/repositories","Schemas":[{"name":"Exchange","config":{"Table":""}},{"name":"Outbox","config":{"Table":""},"fields":[{"name":"timestamp","type":{"Type":2,"Ident":"","PkgPath":"time","PkgName":"","Nillable":false,"RType":null},"position":{"Index":0,"MixedIn":false,"MixinIndex":0}},{"name":"topic","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"position":{"Index":1,"MixedIn":false,"MixinIndex":0}},{"name":"key","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"position":{"Index":2,"MixedIn":false,"MixinIndex":0}},{"name":"payload","type":{"Type":5,"Ident":"","PkgPath":"","PkgName":"","Nillable":true,"RType":null},"position":{"Index":3,"MixedIn":false,"MixinIndex":0}},{"name":"headers","type":{"Type":3,"Ident":"map[string]string","PkgPath":"","PkgName":"","Nillable":true,"RType":{"Name":"","Ident":"map[string]string","Kind":21,"PkgPath":"","Methods":{}}},"position":{"Index":4,"MixedIn":false,"MixinIndex":0}},{"name":"retry_count","type":{"Type":12,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"position":{"Index":5,"MixedIn":false,"MixinIndex":0}},{"name":"status","type":{"Type":6,"Ident":"outbox.Status","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"enums":[{"N":"Pending","V":"PENDING"},{"N":"Succeeded","V":"SUCCEEDED"},{"N":"Failed","V":"FAILED"}],"position":{"Index":6,"MixedIn":false,"MixinIndex":0}},{"name":"last_retry","type":{"Type":12,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"optional":true,"position":{"Index":7,"MixedIn":false,"MixinIndex":0}}],"indexes":[{"fields":["last_retry"]},{"fields":["status"]}]}],"Features":["privacy","entql","schema/snapshot","sql/schemaconfig","sql/lock","sql/modifier","sql/execquery","sql/upsert","sql/versioned-migration"]}`
