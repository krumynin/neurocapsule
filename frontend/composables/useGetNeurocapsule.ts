import type { NeurocapsuleGetParams } from 'api-gen';

export const useGetNeurocapsule = () => {
  const api = useApi();
  const getQuery = useGetRouteQuery();
  const router = useRouter();

  const {
    data: neurocapsule,
    pending,
    error,
    refresh,
  } = useLazyAsyncData('neurocapsule', async () => {
    const budget = getQuery<NeurocapsuleGetParams['price_segment']>('budget');
    const color = getQuery<NeurocapsuleGetParams['color_base'] | null>('color');
    const colorSchem = getQuery<NeurocapsuleGetParams['color_scheme']>('color_schem');

    if (typeof budget === 'undefined' || typeof color === 'undefined' || typeof colorSchem === 'undefined') {
      throw new TypeError('something went wrong');
    }

    const { data } = await api.main.neurocapsuleGet({
      jsonrpc: '2.0',
      id: '12345',
      price_segment: budget,
      color_scheme: colorSchem,
      color_base: color || 'brown',
      excluded_sku_list: '',
    });

    if (data.error) {
      throw new TypeError(data.error.message);
    }

    return data.result?.neurocapsule || null;
  });

  watch(error, (value) => {
    if (value) {
      router.push({
        path: '/error',
      });
    }
  });

  return { neurocapsule, pending, refresh };
};
