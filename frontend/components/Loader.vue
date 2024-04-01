<template>
  <div ref="loaderElement" :class="$style.root" />
</template>

<script setup lang="ts">
import lottie from 'lottie-web';

const anim = ref<ReturnType<typeof lottie.loadAnimation> | null>(null);
const loaderElement = ref<HTMLElement | null>(null);

onMounted(() => {
  if (loaderElement.value) {
    anim.value = lottie.loadAnimation({
      container: loaderElement.value,
      renderer: 'svg',
      loop: true,
      autoplay: true,
      path: '/Loader.json',
    });
  }
});
</script>

<style module>
  .root {
    max-width: 1200px;
    margin: 0 auto;
  }
</style>
