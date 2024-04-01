/**
 * @file Конфигурация swagger-typescript-api
 * @link https://github.com/acacode/swagger-typescript-api
 */

import { generateApi } from 'swagger-typescript-api';
import { resolve, dirname } from 'path';
import { fileURLToPath } from 'url';

const __dirname = dirname(fileURLToPath(import.meta.url));

const onPrepareConfig = (currentConfiguration) => {
  currentConfiguration.routes.combined = currentConfiguration.routes.combined.map(item => ({
    ...item,
    moduleName: `${item.moduleName}Api`,
  }));

  return currentConfiguration;
};

generateApi({
  url: 'http://localhost:8081/server.yaml',
  input: resolve(process.cwd(), './server.yaml'),
  output: resolve(__dirname, 'data'),
  httpClientType: 'axios',
  singleHttpClient: false,
  modular: true,
  extractRequestParams: true,
  extractRequestBody: true,
  addReadonly: true,
  generateUnionEnums: true,
  // extractEnums: true,
  prettier: {
    tabWidth: 2,
    trailingComma: 'all',
    parser: 'typescript',
    singleQuote: true,
  },
  hooks: {
    onPrepareConfig,
  },
});
