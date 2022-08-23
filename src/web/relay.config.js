module.exports = {
    // ...
    // Configuration options accepted by the `relay-compiler` command-line tool and `babel-plugin-relay`.
    src: "./src",
    schema: "../../api-definitions/graphql/omiga/schema.graphql",
    exclude: ["**/node_modules/**", "**/__mocks__/**", "**/__generated__/**"],
    language: "typescript"
  }
