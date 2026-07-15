<script setup>
defineProps({
  movie: {
    type: Object,
    required: true,
  },
});
</script>

<template>
  <div
    class="bg-gray-900 rounded-xl overflow-hidden shadow-lg hover:scale-105 transition-transform duration-300 cursor-pointer"
  >
    <img
      :src="movie.posterUrl"
      :alt="movie.title"
      class="w-full h-36 object-cover"
    />

    <div class="p-4">
      <MovieCardDetails :title="movie.title" :genre="movie.genre" />

      <MovieCardRating :rating="movie.rating" class="mt-3" />
    </div>
  </div>
</template>
