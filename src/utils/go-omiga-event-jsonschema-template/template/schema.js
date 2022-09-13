const { File, render } = require("@asyncapi/generator-react-sdk");

export default async function ({ asyncapi, params }) {
  if (asyncapi.channels().length === 0) {
    return null;
  }

  const messages = Object.entries(asyncapi._json.components.messages);
  const jsonSchema = getJsonSchema(messages[0]);

  return [<File name="schema.json">{render(jsonSchema)}</File>];
}

const captializeFirstChar = (str) => str.charAt(0).toUpperCase() + str.slice(1);
const getDefinitionReference = (name) => ({ $ref: `#/definitions/${name}` });

// Push types with the `$id` property into `definitions`
const consolidateDefinitions = (obj, defs, schemaReferences) => {
  for (const key in obj) {
    if (Object.hasOwnProperty.call(obj, key)) {
      const item = obj[key];
      const parserSchemaId = item["x-parser-schema-id"];

      const itemId = item.$id;
      if (itemId) {
        obj[key] = getDefinitionReference(itemId);
        defs[itemId] = item;
      } else if (parserSchemaId && item.type === "object") {
        // if there is a schema that is reused but has no $id, then use the first property name
        let schemaRef = schemaReferences[parserSchemaId];
        if (!schemaRef) {
          if (defs[key]) {
            throw new Error(
              `Key already defined: ${key}. You might need to set the $id for this one`
            );
          }
          // Use the name of the key that this type is attached
          schemaRef = getDefinitionReference(key);
          schemaReferences[parserSchemaId] = schemaRef;
          defs[key] = item;
        }
        obj[key] = schemaRef;
      }

      if (item.properties) {
        consolidateDefinitions(item.properties, defs, schemaReferences);
      }

      if (item.items) {
        handleArrayType(item, defs);
      }
    }
  }
};

const handleArrayType = (item, defs, schemaReferences) => {
  const arrayType = item.items;
  if (arrayType) {
    const itemsId = arrayType.$id;
    if (itemsId) {
      defs[itemsId] = arrayType;
      item.items = getDefinitionReference(itemsId);
    }

    if (arrayType.properties) {
      consolidateDefinitions(arrayType.properties, defs, schemaReferences);
    }
  }
};

const getJsonSchema = ([messageName, message]) => {
  messageName = captializeFirstChar(messageName);

  const properties = message.payload.properties;
  const definitions = {};
  const schemaReferences = {};

  consolidateDefinitions(properties, definitions, schemaReferences);

  // go through all properties. If one of the properties has $id, then push that to defs
  // and replace with reference.

  var schema = {
    $schema: "http://json-schema.org/draft-04/schema#",
    title: messageName,
    description: message.description,
    ...message.payload,
    definitions: definitions,
  };

  const strJson = JSON.stringify(
    schema,
    (key, value) => (key === "x-parser-schema-id" ? undefined : value),
    2
  );

  return strJson;
};