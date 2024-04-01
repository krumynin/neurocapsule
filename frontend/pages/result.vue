<template>
  <VSheet class="pa-6">
    <TheHeader>
      <template
        v-if="!pending"
        #count
      >
        {{ productsCount }} товаров
      </template>
      <template
        v-if="pending"
        #default
      >
        <LoaderText />
      </template>
    </TheHeader>

    <Loader v-if="pending" />

    <template v-else>
      <ResultCard
        v-if="top2"
        v-model:selected="selected"
        title="Верх"
        :products="top2"
        class="mb-10"
      />

      <ResultCard
        v-if="top1"
        v-model:selected="selected"
        title="База"
        :products="top1"
        class="mb-10"
      />

      <ResultCard
        v-if="bottom"
        v-model:selected="selected"
        title="Низ"
        :products="bottom"
        class="mb-10"
      />

      <Summary
        :count="selected.length"
        :price="summaryPrice"
      />

      <VBtn
        class="mr-md-6 mr-lg-6 mr-xl-6 mr-xxl-6"
        :class="$style.button"
        :loading="pending"
        variant="outlined"
        :block="width <= 960"
        @click="refresh"
      >
        Попробовать снова
      </VBtn>

      <VBtn
        :loading="pending"
        :disabled="!selected.length"
        :block="width <= 960"
        @click="addToCart(selected)"
      >
        Купить капсулу
      </VBtn>
    </template>
  </VSheet>
</template>

<script setup lang="ts">
import { useDisplay } from 'vuetify';

import type { Product } from 'api-gen';

const { neurocapsule, pending, refresh } = useGetNeurocapsule();
const { addToCart } = useAddToCard();
const { width } = useDisplay();

const selected = ref<Product[]>([]);

watch(neurocapsule, (value) => {
  if (value) {
    selected.value = [
      ...Object.values(value.top1).filter(Boolean),
      ...Object.values(value.top2).filter(Boolean),
      ...Object.values(value.bottom).filter(Boolean),
    ];
  }
});

const top1 = computed(() => neurocapsule.value?.top1 ? Object.values(neurocapsule.value.top1).filter(Boolean) : null);
const top2 = computed(() => neurocapsule.value?.top2 ? Object.values(neurocapsule.value.top2).filter(Boolean) : null);
const bottom = computed(() => neurocapsule.value?.bottom ? Object.values(neurocapsule.value.bottom).filter(Boolean) : null);

const summaryPrice = computed(() => selected.value.reduce((acc, cur) => acc + cur.price_amount, 0));

const productsCount = computed(() => {
  let count = 0;

  if (top1.value) {
    count += top1.value.length;
  }

  if (top2.value) {
    count += top2.value.length;
  }

  if (bottom.value) {
    count += bottom.value.length;
  }

  return count;
});
</script>

<style module>
  .button {
    margin-bottom: 16px;
  }

  @media screen and (min-width: 960px) {
    .button {
      margin-bottom: 0;
    }
  }
</style>
