<script setup>
const activeFilter = ref(null);
const filters = ["Genres", "Rating"];
</script>

<template>

  <div
    class="flex gap-3 overflow-x-auto pb-4 scrollbar-hide [&::-webkit-scrollbar]:hidden [-ms-overflow-style:none] [scrollbar-width:none]"
  >
    <FilterOption
      v-for="filter in filters"
      :key="filter"
      :label="filter"
      :is-active="activeFilter === filter"
      @click="activeFilter = activeFilter === filter ? null : filter"
    />
  </div>
</template>
