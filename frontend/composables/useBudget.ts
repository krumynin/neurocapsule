import type { NeurocapsuleGetParams } from 'api-gen';

export const useBudget = () => {
  const budgetList = [{
    id: 0,
    value: 'low',
    title: 'Бюджетный (до 13 000 ₽)',
  }, {
    id: 1,
    value: 'middle',
    title: 'Средний (до 30 000 ₽)',
  }, {
    id: 2,
    value: 'high',
    title: 'Премиальный (от 30 000 ₽)',
  }];

  const budget = ref<NeurocapsuleGetParams['price_segment'] | null>(null);

  return { budgetList, budget };
};
