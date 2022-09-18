const { File } = require("@asyncapi/generator-react-sdk");

export default async function ({ asyncapi, params }) {
  if (asyncapi.channels().length === 0) {
    return null;
  }

  const payloadContent = `
// Code generated by go-omiga-template, DO NOT EDIT.

package ${params.packageName}

import (
	"github.com/gobuffalo/packr/v2"
)

const TopicName = "${Object.keys(asyncapi.channels())[0]}"

func GetJsonSchema() (string, error) {
	return packr.New("schema","./schema").FindString("./jsonschema.json")
}

func GetDereferencedJsonSchema() (string, error) {
	return packr.New("schema","./schema").FindString("./dereferenced-jsonschema.json")
}

func GetAvroSchema() (string, error) {
	return packr.New("schema","./schema").FindString("./avro.json")
}
`;

  return <File name="metadata_eventgen.go">{payloadContent}</File>;
}
