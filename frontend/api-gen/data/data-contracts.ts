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

export interface BaseResponse {
  jsonrpc: string;
  id: string;
  error?: Error;
}

export interface Error {
  code: number;
  /** Human readable error description for debug usage */
  message: string;
  /** Arbitary additional error-related data */
  data?: object;
}

export interface NeurocapsuleResponse {
  neurocapsule: {
    bottom: {
      trousers: Product;
      jeans: Product;
      sweatpants: Product;
    };
    top1: {
      shirt1: Product;
      shirt2: Product;
      tShirt1: Product;
      tShirt2: Product;
    };
    top2: {
      hoodie: Product;
      sweater: Product;
      blazer: Product;
    };
  };
}

export interface Product {
  sku: string;
  full_sku: string;
  name: string;
  price_amount: number;
  thumbnail: string;
  gallery: string;
  color: string;
}

export interface NeurocapsuleGetParams {
  jsonrpc: string;
  id: string;
  price_segment: 'low' | 'middle' | 'high';
  color_scheme: 'contrast' | 'monochrome' | 'classic';
  color_base: 'brown' | 'khaki' | 'darkBlue' | 'blackWhite';
  /** Comma-separated list of values sku: QE456EMBT231,MP002XM4F7C9,MP002XM101FC */
  excluded_sku_list?: string;
}
