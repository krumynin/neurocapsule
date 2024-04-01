<template>
  <TransitionGroup name="slide">
    <template v-for="(item, index) in list">
      <span
        v-if="index === count"
        :key="item"
      >
        {{ item }}
      </span>
    </template>
  </TransitionGroup>
</template>

<script setup lang="ts">
const list = [
  'Разогреваем нейроны',
  'Подключаем все базы',
  'Готовим нейрокапсулу',
];

const count = ref(0);
const timer = ref<number | null>(null);

const update = () => {
  timer.value = window.setTimeout(() => {
    if (count.value >= list.length - 1) {
      count.value = 0;
    } else {
      count.value += 1;
    }
    update();
  }, 2000);
};

onMounted(update);
onBeforeUnmount(() => {
  timer.value = null;
});
</script>

<style>
  .slide-enter-from {
    opacity: 0;
    transform: translateY(30px)
  }
  .slide-enter-active {
    transition: opacity .8s ease,
                transform .8s ease,
                color 5s ease
  }

  .slide-leave-active {
    position: absolute;
  }
</style>
