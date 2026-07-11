<!-- app/components/Filter.vue -->
<script setup>
const activeFilter = ref(null);
const filters = ["Release year", "Genres", "Price", "Rating", "Age rating"];
</script>

<template>
  <!-- 
    - flex: Layouts items horizontally
    - gap-3: Spacing between items
    - overflow-x-auto: Enables horizontal scrolling
    - scrollbar-hide: Hides scrollbar (requires tailwind-scrollbar-hide plugin)
    - [-ms-overflow-style:none]: Standard CSS scrollbar hiding for IE/Edge
    - [scrollbar-width:none]: Standard CSS scrollbar hiding for Firefox
  -->
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
