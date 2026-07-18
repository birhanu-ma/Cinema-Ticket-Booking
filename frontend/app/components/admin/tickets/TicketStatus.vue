<script setup>
import { computed } from "vue";

const props = defineProps({
  ticket: {
    type: Object,
    required: true,
  },
});

const thumbnail = computed(() => {
  return (
    props.ticket.movie_image ||
    props.ticket.schedule?.movie?.movie_images?.[0]?.image_url ||
    "https://api.dicebear.com/9.x/shapes/svg?seed=ticket"
  );
});
</script>

<template>
  <div
    class="flex bg-gray-950 border border-gray-900 rounded-3xl overflow-hidden shadow-2xl group hover:border-gray-800 transition-all duration-300"
  >
    <div
      class="w-40 sm:w-48 relative flex-shrink-0 overflow-hidden bg-gray-900"
    >
      <img
        :src="thumbnail"
        class="w-full h-full object-cover group-hover:scale-105 transition-transform duration-500"
        alt="Movie Thumbnail Poster"
      />

      <div
        class="absolute inset-0 bg-gradient-to-r from-transparent to-gray-950/20"
      ></div>
    </div>

    <AdminTicketsTicketStatusDetail :ticket="ticket" />
  </div>
</template>
