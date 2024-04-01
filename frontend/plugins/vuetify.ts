import { createVuetify } from 'vuetify';
import { aliases, mdi } from 'vuetify/iconsets/mdi';
import { VDataTable } from 'vuetify/labs/VDataTable';
import { VDatePicker } from 'vuetify/labs/VDatePicker';
// import { VStepper } from 'vuetify/labs/VStepper';
import { VuetifyDateAdapter } from 'vuetify/labs/date/adapters/vuetify';
import '@mdi/font/css/materialdesignicons.css';

export default defineNuxtPlugin((nuxtApp) => {
  const vuetify = createVuetify({
    ssr: false,
    components: {
      VDataTable,
      VDatePicker,
      // VStepper,
    },
    theme: {
      defaultTheme: 'light',
      themes: {
        light: {
          dark: false,
          colors: {
            primary: '#000000',
            'primary-hover': '#444444',
            secondary: '#F5F5F5',
            'secondary-hover': '#E5E5E5',
            action: '#F93C00',
            'action-hover': '#DB0D00',
            success: '#00A200',
            error: '#C20000',
            link: '#2D68F6',
            'link-hover': '#2D68F6',
            'pale-pink': '#FFD7D2',
            gold: '#A5825F',
            mint: '#A5D2A0',
            'pale-blue': '#CDE6FF',
            'dark-blue': '#3C5064',
            yellow: '#FFA622',
            //
            brown: '6D3F14',
            khaki: '8E9352',
            darkBlue: '#4474F3',
            // light theme
            'primary-label': '#000000',
            'primary-label-hover': '#888888',
            'secondary-label': '#888888',
            'secondary-label-hover': '#000000',
            disabled: '#E5E5E5',
            separator: '#E5E5E5',
            'badge-foreground': '#FFFFFF',
            background: '#FFFFFF',
            'secondary-background': '#F5F5F5',
          },
        },
      },
    },
    icons: {
      defaultSet: 'mdi',
      aliases,
      sets: {
        mdi,
      },
    },
    locale: {
      locale: 'ru',
    },
    date: {
      adapter: VuetifyDateAdapter,
    },
    defaults: {
      VBtn: {
        color: 'primary',
        size: 'large',
        variant: 'flat',
        class: 'text-none',
      },
    },
  });

  nuxtApp.vueApp.use(vuetify);
});
