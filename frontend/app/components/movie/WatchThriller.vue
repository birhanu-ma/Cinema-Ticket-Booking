<script setup>
const props = defineProps({
  movie: {
    type: Object,
    required: true,
  },
});

const featuredImage = computed(() => {
  return (
    props.movie?.movie_images?.find((img) => img.is_featured)?.image_url ||
    props.movie?.movie_images?.[0]?.image_url ||
    "https://placehold.co/1200x600?text=Movie"
  );
});
</script>

<template>
  <div
    v-if="movie"
    class="relative w-287 h-[460px] rounded-3xl overflow-hidden shadow-2xl bg-gray-950"
  >
    <img
      :src="featuredImage"
      class="w-full h-full object-cover"
      :alt="movie.title"
    />

    <div
      class="absolute inset-0 bg-gradient-to-t from-black/90 via-black/40 to-transparent"
    ></div>

    <MovieThrillerDetails
      :movie="movie"
      @watch="console.log('Playing trailer for', movie.title)"
    />
  </div>

  <div
    v-else
    class="max-w-4xl w-full h-[400px] bg-gray-900 animate-pulse rounded-3xl"
  ></div>
</template>