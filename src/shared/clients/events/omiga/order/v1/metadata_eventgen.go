
// Code generated by go-omiga-template, DO NOT EDIT.

package orderv1

import (
	_ "embed"
)

const TopicName = "order.v1.event"

//go:embed schema/jsonschema.json
var jsonschema string

//go:embed schema/dereferenced-jsonschema.json
var dereferencedJsonschema string

//go:embed schema/avro.avsc
var avro string

func GetJsonSchema() string {
  return jsonschema
}

func GetDereferencedJsonSchema() string {
  return dereferencedJsonschema
}

func GetAvroSchema() string {
  return avro
}
