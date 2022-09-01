const { GoGenerator } = require("@asyncapi/modelina");
const { File } = require("@asyncapi/generator-react-sdk");

export default async function ({ asyncapi, params }) {
  let foundTimeTypeField = false;
  let foundIDTypeField = false;

  const generator = new GoGenerator({
    presets: [
      {
        struct: {
          field({ fieldName, field, renderer, model: { required } }) {
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
                foundTimeTypeField = true;

                break;

              case "int32":
                finalFieldType = "int32";

                break;

              case "int64":
                finalFieldType = "int64";

                break;

              case "uuid":
                finalFieldType = "ID";
                foundIDTypeField = true;

                break;

              default:
                break;
            }

            const tag = isRequired
              ? `\`json:"${fieldName}"\``
              : `\`json:"${fieldName},omitempty"\``;

            if (field.type === "array") {
              finalFieldType = "[]" + finalFieldType.substring("[]*".length);
            }

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

`;

  if (foundTimeTypeField && foundIDTypeField) {
    payloadContent =
      payloadContent +
      `
    import (
      "time"
    
      "github.com/google/uuid"
    )
    
    type ID uuid.UUID

    `;
  } else if (foundTimeTypeField && !foundIDTypeField) {
    payloadContent =
      payloadContent +
      `
    import "time"

    `;
  } else if (!foundTimeTypeField && foundIDTypeField) {
    payloadContent =
      payloadContent +
      `
    import "github.com/google/uuid"
    
    type ID uuid.UUID

    `;
  }

  models.forEach((model) => {
    let result = model.result;

    if (model.model.type === "string" && model.model.enum) {
      result = result.split("\n").reduce((reduction, line) => {
        if (line.indexOf(` = \"`) === -1) {
          return `${reduction}\n${line}`;
        }

        if (line.indexOf(` ${model.modelName} = \"`) !== -1) {
          return `${reduction}\n${line}`;
        }

        const updatedLine = line.replace(' = "', ` ${model.modelName} = \"`);

        return `${reduction}\n${updatedLine}`;
      }, "");
    }

    payloadContent += `
    ${model.dependencies.join("\n")}
    ${result}
    `;
  });

  return <File name="payloads_eventgen.go">{payloadContent}</File>;
}
