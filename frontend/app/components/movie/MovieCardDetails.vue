<script setup>
defineProps({
  title: {
    type: String,
    required: true,
  },
  genre: {
    type: String,
    required: true,
  },
  movieId: {
    type: [String, Number],
    required: true,
  },
});
</script>

<template>
  <RouterLink :to="`/user/movies/${movieId}`">
    <div>
      <h3 class="text-white font-semibold text-lg truncate">
        {{ title }}
      </h3>

      <p class="text-gray-400 text-sm mt-1">
        {{ genre }}
      </p>
    </div>
  </RouterLink>
</template>