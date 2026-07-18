<script setup>
defineProps({
  stars: {
    type: Array,
    default: () => [],
  },
});
</script>

<template>
  <div class="flex gap-4 overflow-x-auto pb-4 scrollbar-hide">
    <MovieStarCard
      v-for="star in stars"
      :key="star.star.id"
      :actor="star.star"
    />
  </div>
</template>