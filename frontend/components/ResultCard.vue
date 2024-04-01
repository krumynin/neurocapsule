<template>
  <div>
    <div class="mb-10 font-size-h2">
      {{ title }}
    </div>

    <div class="d-flex justify-space-between flex-wrap">
      <div
        v-for="item in products"
        :key="item.full_sku"
        class="d-flex py-4 mb-6"
        :class="$style.width"
      >
        <div class="mr-4">
          <VCheckbox
            v-model="selected"
            class="base-checkbox"
            :value="item"
            density="compact"
          />
        </div>
        <ProductCard :item="item" />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import type { Product } from 'api-gen';

const props = defineProps<{
  title: string;
  products: Product[];
  selected: Product[];
}>();

const emit = defineEmits<{
  'update:selected': [Product[]];
}>();

const selected = computed({
  get () {
    return props.selected;
  },
  set (value) {
    emit('update:selected', value);
  },
});
</script>

<style module>

  .input {
    width: 542px;
  }

  .color {
    width: 56px !important;
    height: 56px !important;
    display: flex;
    justify-content: center;
    align-items: center;
  }

  .summary {
    width: 448px;
  }

  .width {
    width: 100% !important;
  }

  @media screen and (min-width: 960px) {
    .width {
      width: 50% !important;
    }
  }
</style>
