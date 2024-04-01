<template>
  <a
    class="d-flex text-primary text-decoration-none"
    :href="`https://www.lamoda.ru/p/${item.sku}`"
    target="_blank"
  >
    <VImg
      class="bg-secondary mr-6"
      width="116"
      height="168"
      :src="`https://a.lmcdn.ru/img389x562/${item.gallery}`"
    />
    <div>
      <div class="mb-1">
        {{ item.name }}
      </div>
      <div class="mb-1">
        40 RUS (S INT)
      </div>
      <div class="mb-4">
        {{ item.color }}
      </div>
      <div>
        {{ item.price_amount }} ₽
      </div>
    </div>
  </a>
</template>

<script setup lang="ts">
import { Product } from 'api-gen';

defineProps<{
  item: Product;
}>();
</script>

<style module>
</style>
