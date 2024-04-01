import type { Product } from 'api-gen';

export const useAddToCard = () => {
  const addToCart = (products: Product[]) => {
    const url = `https://www.lamoda.ru/checkout/cart/?force_add=${products.map(item => item.full_sku).join(',')}`;

    window.open(url, '_blank');
  };

  return { addToCart };
};
