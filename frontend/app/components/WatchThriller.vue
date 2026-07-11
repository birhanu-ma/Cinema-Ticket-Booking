<!-- app/components/WatchThriller.vue -->

<!-- app/pages/index.vue -->
<script setup>
// This data would typically come from your Apollo GraphQL query
const movie = [
  {
    id: 1,
    title: "Shazam! Fury of the Gods",
    genre: "Action/Adventure",
    rating: 4, // Number of stars
    posterUrl: "https://image_dab641.png", // Replace with your actual Hasura URL
  },

  // Add more movie objects here...
];
</script>
<template>
  <div
    v-if="movie"
    class="relative w-full h-[400px] rounded-3xl overflow-hidden shadow-2xl"
  >
    <img
      :src="movie?.backdropUrl"
      class="w-full h-full object-cover"
      alt="Movie Backdrop"
    />

    <div
      class="absolute inset-0 bg-gradient-to-t from-black/80 to-transparent"
    ></div>

    <ThrillerDetails
      :movie="movie"
      @watch="console.log('Playing trailer for', movie.title)"
    />
  </div>

  <div
    v-else
    class="w-full h-[400px] bg-gray-900 animate-pulse rounded-3xl"
  ></div>
</template>
