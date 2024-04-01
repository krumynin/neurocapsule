<template>
  <div :class="[$style.root, 'mb-14']">
    <div class="d-flex justify-space-between mb-2">
      <div>
        {{ count }} товар(а) на сумму
      </div>
      <div>
        {{ price }} ₽
      </div>
    </div>
    <div class="mb-1 d-flex justify-space-between font-size-h2">
      <div>
        Итого
      </div>
      <div>
        {{ price }} ₽
      </div>
    </div>
    <div class="text-secondary-label font-size-body-3">
      Без учета возможной стоимости доставки
    </div>
  </div>
</template>

<script setup lang="ts">
defineProps<{
  count: number;
  price: number;
}>();
</script>

<style module>
  .root {
    width: 100%;
  }

  @media screen and (min-width: 960px) {
    .root {
      width: 448px;
    }
  }
</style>
