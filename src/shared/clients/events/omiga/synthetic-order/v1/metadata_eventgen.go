
// Code generated by go-omiga-template, DO NOT EDIT.

package syntheticorderv1

import (
	"github.com/gobuffalo/packr/v2"
)

const TopicName = "synthetic.order.v1.event"

func GetJsonSchema() (string, error) {
	return packr.New("schema","./schema").FindString("./jsonschema.json")
}

func GetDereferencedJsonSchema() (string, error) {
	return packr.New("schema","./schema").FindString("./dereferenced-jsonschema.json")
}

func GetAvroSchema() (string, error) {
	return packr.New("schema","./schema").FindString("./avro.avsc")
}
