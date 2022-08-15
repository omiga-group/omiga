const { GoGenerator } = require("@asyncapi/modelina");
const { File } = require("@asyncapi/generator-react-sdk");

export default async function ({ asyncapi, params }) {
  const payloadContent = `
// Code generated by go-omiga-template, DO NOT EDIT.

package ${params.packageName}

//go:generate mockgen -source=consumer_gen.go -destination=mock/mock-consumer_gen.go
//go:generate mockgen -source=handler_gen.go -destination=mock/mock-handler_gen.go
//go:generate mockgen -source=producer_gen.go -destination=mock/mock-producer_gen.go

`;

  return <File name="generate_gen.go">{payloadContent}</File>;
}