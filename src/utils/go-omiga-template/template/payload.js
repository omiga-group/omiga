const { GoGenerator } = require("@asyncapi/modelina");
const { File } = require("@asyncapi/generator-react-sdk");

export default async function ({ asyncapi, params }) {
  const generator = new GoGenerator({
    presets: [
      {
        struct: {
          field({ fieldName, field, renderer, type, model: { required } }) {
            const formattedFieldName = renderer.nameField(fieldName, field);
            const fieldType = renderer.renderType(field);

            const description = field.originalInput
              ? ` // ${field.originalInput["description"]}`
              : "";

            const isRequired = required
              ? required.findIndex((item) => item === fieldName) !== -1
              : false;
            const unrequiredMark = !isRequired ? "*" : "";

            const format = field.originalInput
              ? field.originalInput["format"]
              : "";
            let finalFieldType = fieldType.startsWith("*")
              ? fieldType.substring(1)
              : fieldType;

            switch (format) {
              case "date-time":
                finalFieldType = "time.Time";
                break;

              case "uuid":
                finalFieldType = "ID";
                break;

              default:
                break;
            }

            const tag =
              isRequired
                ? `\`json:"${fieldName}"\``
                : `\`json:"${fieldName},omitempty"\``

            return `${formattedFieldName} ${unrequiredMark}${finalFieldType} ${tag} ${description}`;
          },
        },
      },
    ],
  });
  const models = await generator.generate(asyncapi);

  let payloadContent = `
// Code generated by go-omiga-template, DO NOT EDIT.

package ${params.packageName}

import (
	"time"

	"github.com/google/uuid"
)

type ID uuid.UUID
`;

  models.forEach((model) => {
    payloadContent += `
    ${model.dependencies.join("\n")}
    ${model.result}
    `;
  });

  return <File name="payloads_gen.go">{payloadContent}</File>;
}
