<template>
  <div>
    <div class="mb-2 font-size-h1">
      Нейрокапсула
      <span
        v-if="$slots.count"
        class="text-secondary-label font-size-h2 text-no-wrap"
      >
        <slot name="count" />
      </span>
    </div>
    <div class="mb-10 text-secondary-label font-size-h2">
      <slot name="default" />
    </div>
  </div>
</template>
