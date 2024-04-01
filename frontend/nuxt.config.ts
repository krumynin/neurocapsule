// https://nuxt.com/docs/api/configuration/nuxt-config
import { resolve } from 'path';
import vuetify from 'vite-plugin-vuetify';

export default defineNuxtConfig({
  ssr: false,
  typescript: {
    shim: false,
  },
  devServer: {
    https: false,
    port: 8081,
  },
  css: ['@/styles/global.scss', '@/styles/settings.scss'],
  vite: {
    server: {
      proxy: {
        '/jsonrpc/v1': {
          target: 'http://localhost:8080',
          changeOrigin: true,
        },
      },
    },
    ssr: {
      noExternal: ['vuetify'],
    },
    resolve: {
      alias: {
        'api-gen': resolve(__dirname, 'api-gen'),
        composables: resolve(__dirname, 'composables'),
        stores: resolve(__dirname, 'stores'),
        types: resolve(__dirname, 'types'),
      },
    },
  },
  modules: [
    [
      '@pinia/nuxt',
      {
        autoImports: [
          'defineStore',
        ],
      },
    ],
    '@vueuse/nuxt',
    async (options, nuxt) => {
      nuxt.hooks.hook('vite:extendConfig', (config) => {
        config.plugins!.push(vuetify());
      });
    },
  ],
  app: {
    head: {
      title: 'Lamoda Нейрокапсула ',
      link: [{ rel: 'icon', type: 'image/svg+xml', href: '/favicon.svg' }],
      meta: [
        { charset: 'utf-8' },
        { name: 'viewport', content: 'width=device-width, initial-scale=1' },
        { hid: 'description', name: 'description', content: 'Lamoda Нейрокапсула' },
      ],
      htmlAttrs: {
        lang: 'ru',
      },
    },
    baseURL: '/',
  },
});
