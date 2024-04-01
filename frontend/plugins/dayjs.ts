import dayjs, { extend, locale } from 'dayjs';
import customParseFormat from 'dayjs/plugin/customParseFormat';
import localizedFormat from 'dayjs/plugin/localizedFormat';
import quarterOfYear from 'dayjs/plugin/quarterOfYear';
import ru from 'dayjs/locale/ru';

export default defineNuxtPlugin(() => {
  locale(ru);
  extend(customParseFormat);
  extend(localizedFormat);
  extend(quarterOfYear);

  return {
    provide: {
      dayjs,
    },
  };
});
