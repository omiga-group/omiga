interface GlobalConfig {
  GRAPHQL_API_ENDPOINT: string;
}

declare global {
  interface Window {
    _env_: GlobalConfig;
  }
}

export {};
