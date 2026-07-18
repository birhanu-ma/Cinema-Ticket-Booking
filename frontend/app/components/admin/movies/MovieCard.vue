<script setup>
import { computed } from "vue";

const props = defineProps({
  movie: {
    type: Object,
    required: true,
  },
});

const thumbnail = computed(() => {
  return (
    props.movie.movie_images.find((img) => img.is_featured)?.image_url ||
    props.movie.movie_images[0]?.image_url ||
    "https://placehold.co/500x300?text=No+Image"
  );
});

const genre = computed(() =>
  props.movie.movie_genres.map((g) => g.genre.name).join(", ")
);
</script>

<template>
  <div
    class="flex bg-gray-950 border border-gray-900 rounded-3xl overflow-hidden shadow-2xl  w-full group hover:border-gray-800 transition-all duration-300"
  >
    <div
      class="w-40 sm:w-44 relative flex-shrink-0 overflow-hidden bg-gray-900"
    >
      <img
        :src="thumbnail"
        class="w-full h-full object-cover group-hover:scale-105 transition-transform duration-500"
        alt="Movie Poster Landscape Artwork"
      />
      <div
        class="absolute inset-0 bg-gradient-to-r from-transparent via-transparent to-gray-950/40"
      ></div>
    </div>

    <AdminMoviesMovieDetail
      :movie-id="movie.id"
      :title="movie.title"
      :genre="genre"
      :duration="movie.duration"
    />
  </div>
</template>