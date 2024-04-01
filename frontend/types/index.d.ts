import type { Api } from 'plugins/api';
import dayjs from 'dayjs';

declare module '#app' {
  interface NuxtApp {
    $api: Api;
    $dayjs: typeof dayjs;
  }
}
declare module 'vue' {
  interface ComponentCustomProperties {
    $api: Api;
    $dayjs: typeof dayjs;
  }
}