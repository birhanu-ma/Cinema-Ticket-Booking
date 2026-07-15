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
    class="flex bg-gray-950 border border-gray-900 rounded-3xl overflow-hidden shadow-2xl  w-full group hover:border-gray-800 transition-all duration-300"
  >
    <div
      class="w-40 sm:w-44 relative flex-shrink-0 overflow-hidden bg-gray-900"
    >
      <img
        :src="movie.thumbnail"
        class="w-full h-full object-cover group-hover:scale-105 transition-transform duration-500"
        alt="Movie Poster Landscape Artwork"
      />
      <div
        class="absolute inset-0 bg-gradient-to-r from-transparent via-transparent to-gray-950/40"
      ></div>
    </div>

    <AdminMoviesMovieDetail
      :title="movie.title"
      :genre="movie.genre"
      :duration="movie.duration"
    />
  </div>
</template>
