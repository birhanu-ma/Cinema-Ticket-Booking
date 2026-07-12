<script setup>
const movieData = {
  id: 1,
  title: "Shazam! Fury of the Gods",
  genre: "Action/Adventure",
  rating: 4,
  backdropUrl:
    "https://images.unsplash.com/photo-1536440136628-849c177e76a1?w=1200",
};
</script>

<template>
  <div
    v-if="movieData"
    class="relative w-287  h-[460px] rounded-3xl overflow-hidden shadow-2xl bg-gray-950"
  >
    <img
      :src="movieData.backdropUrl"
      class="w-full h-full object-cover"
      alt="Movie Backdrop"
    />

    <div
      class="absolute inset-0 bg-gradient-to-t from-black/90 via-black/40 to-transparent"
    ></div>

    <ThrillerDetails
      :movie="movieData"
      @watch="console.log('Playing trailer for', movieData.title)"
    />
  </div>

  <div
    v-else
    class="max-w-4xl w-full h-[400px] bg-gray-900 animate-pulse rounded-3xl"
  ></div>
</template>
