import type { NeurocapsuleGetParams } from 'api-gen';

export const useColor = () => {
  const colorSchemList = [{
    id: 0,
    value: 'contrast',
    title: 'Контрастые',
  }, {
    id: 1,
    value: 'monochrome',
    title: 'Монохромные',
  }, {
    id: 2,
    value: 'classic',
    title: 'Чёрно-белые',
  }];

  const colorList = [{
    id: 0,
    value: 'brown',
  }, {
    id: 1,
    value: 'khaki',
  }, {
    id: 2,
    value: 'darkBlue',
  }];

  const colorSchem = ref<NeurocapsuleGetParams['color_scheme'] | null>(null);
  const color = ref<NeurocapsuleGetParams['color_base'] | 'blackWhite'>('blackWhite');

  return {
    colorSchemList,
    colorList,
    colorSchem,
    color,
  };
};
