<script setup>
import { computed } from "vue";

const props = defineProps({
  cinema: {
    type: Object,
    required: true,
  },
});

const image = computed(() => {
  return props.cinema.image_url || "https://placehold.co/500x300?text=Cinema";
});
</script>

<template>
  <div
    class="flex bg-gray-950 border border-gray-900 rounded-3xl overflow-hidden shadow-2xl w-full group hover:border-gray-800 transition-all duration-300"
  >
    <div
      class="w-40 sm:w-44 relative flex-shrink-0 overflow-hidden bg-gray-900"
    >
      <img
        :src="image"
        alt="Cinema"
        class="w-full h-full object-cover group-hover:scale-105 transition-transform duration-500"
      />

      <div
        class="absolute inset-0 bg-gradient-to-r from-transparent to-gray-950/40"
      ></div>
    </div>

    <AdminCinemasCinemaDetail
      :cinema-id="cinema.id"
      :name="cinema.name"
      :location="cinema.location"
    />
  </div>
</template>
