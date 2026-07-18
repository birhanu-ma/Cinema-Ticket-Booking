<script setup>
defineProps({
  hall: {
    type: Object,
    required: true,
  },
});
</script>

<template>
  <div
    class="flex bg-gray-950 border border-gray-900 rounded-3xl overflow-hidden shadow-2xl w-full group hover:border-gray-800 transition-all duration-300"
  >
    <div
      class="w-40 sm:w-44 flex-shrink-0 bg-gray-900 flex items-center justify-center text-5xl"
    >
      🎬
    </div>

    <AdminCinemasCinemaHallDetail
      :hall-id="hall.id"
      :name="hall.name"
      :capacity="hall.capacity"
      :cinema-name="hall.cinema?.name"
    />
  </div>
</template>
