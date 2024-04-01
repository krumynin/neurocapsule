/* eslint-disable */
/* tslint:disable */
/*
 * ---------------------------------------------------------------
 * ## THIS FILE WAS GENERATED VIA SWAGGER-TYPESCRIPT-API        ##
 * ##                                                           ##
 * ## AUTHOR: acacode                                           ##
 * ## SOURCE: https://github.com/acacode/swagger-typescript-api ##
 * ---------------------------------------------------------------
 */

import { BaseResponse, NeurocapsuleGetParams, NeurocapsuleResponse } from './data-contracts';
import { HttpClient, RequestParams } from './http-client';

export class NeurocapsuleGetApi<SecurityDataType = unknown> extends HttpClient<SecurityDataType> {
  /**
   * @description Get a neurocapsule based on body parameters
   *
   * @name NeurocapsuleGet
   * @request GET:/neurocapsule.get
   */
  neurocapsuleGet = (query: NeurocapsuleGetParams, params: RequestParams = {}) =>
    this.request<
      BaseResponse & {
        result?: NeurocapsuleResponse;
      },
      any
    >({
      path: `/neurocapsule.get`,
      method: 'GET',
      query: query,
      format: 'json',
      ...params,
    });
}
