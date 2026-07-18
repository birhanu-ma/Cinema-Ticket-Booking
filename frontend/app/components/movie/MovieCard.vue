<script setup>
import { computed } from "vue";

const props = defineProps({
  movie: {
    type: Object,
    required: true,
  },
});

const featuredImage = computed(() => {
  const featured = props.movie.movie_images.find(
    (image) => image.is_featured,
  );

  return (
    featured?.image_url ||
    props.movie.movie_images[0]?.image_url ||
    "https://via.placeholder.com/300x400?text=No+Image"
  );
});

const genres = computed(() =>
  props.movie.movie_genres
    .map((genre) => genre.genre.name)
    .join(", "),
);

const rating = computed(() => {
  return Number(
    props.movie.ratings_aggregate.aggregate.avg.rating ?? 0,
  );
});
</script>

<template>
  <div
    class="bg-gray-900 rounded-xl overflow-hidden shadow-lg hover:scale-105 transition-transform duration-300 cursor-pointer"
  >
    <img
      :src="featuredImage"
      :alt="movie.title"
      class="w-full h-36 object-cover"
    />

    <div class="p-4">
      <MovieCardDetails
        :movie-id="movie.id"
        :title="movie.title"
        :genre="genres"
      />

      <MovieCardRating
        :rating="rating"
        class="mt-3"
      />
    </div>
  </div>
</template>