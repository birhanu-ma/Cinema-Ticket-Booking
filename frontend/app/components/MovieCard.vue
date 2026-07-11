<script setup>
defineProps({
  movie: {
    type: Object,
    required: true,
  },
});
</script>

<template>

  <NuxtLink
    :to="`/movie/${movie.id}`"
    class="block bg-gray-950 rounded-2xl shadow-xl overflow-hidden w-30 transform transition-transform hover:scale-105 hover:shadow-2xl cursor-pointer"
  >
    <div class="p-2 pb-0">
      <img
        :src="movie.posterUrl"
        :alt="`Poster for ${movie.title}`"
        class="w-full h-20 object-cover rounded-xl"
      />
    </div>

    <div class="p-2 space-y-3">
      <MovieCardDetails :title="movie.title" :genre="movie.genre" />

      <MovieCardRating :rating="movie.rating" />
    </div>
  </NuxtLink>
</template>
