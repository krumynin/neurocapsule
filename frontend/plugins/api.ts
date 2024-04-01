import { NeurocapsuleGetApi } from 'api-gen/data/NeurocapsuleGetApi';

interface IApi {
  main: NeurocapsuleGetApi;
}

export class Api implements IApi {
  main: NeurocapsuleGetApi;

  constructor () {
    this.main = new NeurocapsuleGetApi({ baseURL: '/jsonrpc/v1' });
  }
}

export default defineNuxtPlugin(() => {
  const api = new Api();

  return {
    provide: {
      api,
    },
  };
});
