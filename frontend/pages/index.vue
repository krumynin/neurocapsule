<template>
  <VSheet class="pa-6">
    <TheHeader>
      <template #default>
        Чтобы мы могли точнее подобрать капсулу, нужно ввести несколько простых параметров
      </template>
    </TheHeader>

    <div class="mb-6">
      <span class="mr-4 font-size-h2">
        Какой размер вы носите?
      </span>
      <a
        href="https://www.lamoda.ru/about/guides/?utm=footer&amp;from=footer"
        class="text-primary"
        target="_blank"
      >
        Таблица размеров
      </a>
    </div>

    <div class="mb-6">
      <div class="mb-4 font-size-body-1">
        Худи, свитшоты, пиджаки, кардиганы
      </div>
      <VChipGroup
        v-model="outerwearSize"
        class="base-chip-group"
        column
      >
        <VChip
          v-for="item in outerwearClothesList"
          :key="item.id"
          :value="item.title"
          size="x-large"
          selected-class="bg-primary"
        >
          {{ item.title }}
        </VChip>
      </VChipGroup>
    </div>

    <div class="mb-6">
      <div class="mb-4 font-size-body-1">
        Футболки и рубашки
      </div>
      <VChipGroup
        v-model="shirtSize"
        class="base-chip-group"
        column
      >
        <VChip
          v-for="item in shirtSizeList"
          :key="item.id"
          :value="item.title"
          size="x-large"
          selected-class="bg-primary"
        >
          {{ item.title }}
        </VChip>
      </VChipGroup>
    </div>

    <div class="mb-6">
      <div class="mb-4 font-size-body-1">
        Брюки
      </div>
      <VChipGroup
        v-model="trousersSize"
        class="base-chip-group"
        column
      >
        <VChip
          v-for="item in trousersSizeList"
          :key="item.id"
          :value="item.title"
          size="x-large"
          selected-class="bg-primary"
        >
          {{ item.title }}
        </VChip>
      </VChipGroup>
    </div>

    <div class="mb-10">
      <div class="mb-4 font-size-body-1">
        Джинсы
      </div>
      <VChipGroup
        v-model="jeansSize"
        class="base-chip-group"
        column
      >
        <VChip
          v-for="item in jeansClothesList"
          :key="item.id"
          :value="item.title"
          size="x-large"
          selected-class="bg-primary"
        >
          {{ item.title }}
        </VChip>
      </VChipGroup>
    </div>

    <div class="mb-6">
      <div class="mb-4 font-size-body-1">
        Бюджет на капсулу
      </div>
      <VChipGroup
        v-model="budget"
        class="base-chip-group"
        column
      >
        <VChip
          v-for="item in budgetList"
          :key="item.id"
          :value="item.value"
          size="x-large"
          selected-class="bg-primary"
        >
          {{ item.title }}
        </VChip>
      </VChipGroup>
    </div>

    <div>
      <div class="mb-4 font-size-body-1">
        Какие сочетания цветов вам больше нравятся?
      </div>
      <VChipGroup
        v-model="colorSchem"
        class="base-chip-group"
        column
      >
        <VChip
          v-for="item in colorSchemList"
          :key="item.id"
          :value="item.value"
          size="x-large"
          selected-class="bg-primary"
        >
          {{ item.title }}
        </VChip>
      </VChipGroup>
    </div>

    <div
      v-if="colorSchem && colorSchem !== 'classic'"
      class="mt-6"
    >
      <div class="mb-4 font-size-body-1">
        Какие сочетания цветов вам больше нравятся?
      </div>
      <VChipGroup
        v-model="color"
        class="base-chip-group"
        filter
        column
      >
        <VChip
          v-for="item in colorList"
          :key="item.id"
          :value="item.value"
          :class="[$style.color, `bg-${item.value}`, 'color']"
        />
      </VChipGroup>
    </div>

    <VBtn
      class="mt-16"
      :loading="false"
      :disabled="!isButtomEnable"
      @click="result"
    >
      Подобрать капсулу
    </VBtn>
  </VSheet>
</template>

<script setup lang="ts">
const {
  colorSchemList,
  colorList,
  colorSchem,
  color,
} = useColor();

const { budgetList, budget } = useBudget();

const {
  outerwearClothesList,
  shirtSizeList,
  trousersSizeList,
  jeansClothesList,
  outerwearSize,
  jeansSize,
  trousersSize,
  shirtSize,
} = useSizes();

const isButtomEnable = computed(() => (
  outerwearSize.value
  && jeansSize.value
  && trousersSize.value
  && shirtSize.value
  && budget.value
  && colorSchem.value
  && color.value
));

const router = useRouter();

const result = () => {
  router.push({
    path: '/result',
    query: {
      outerwear: outerwearSize.value,
      jeans: jeansSize.value,
      trousers: trousersSize.value,
      shirt: shirtSize.value,
      budget: budget.value,
      color_schem: colorSchem.value,
      color: color.value,
    },
  });
};
</script>

<style module>
  .color {
    width: 56px !important;
    height: 56px !important;
    display: flex;
    justify-content: center;
    align-items: center;
  }
</style>

<style>
  .color .v-chip__filter {
    margin: 0 !important;
    margin-inline-start: 0 !important;
    margin-inline-end: 0 !important;
  }
</style>
