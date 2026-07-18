<script setup>
defineProps({
  schedule: {
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
      🎥
    </div>

    <AdminSchedulesScheduleDetail
      :schedule-id="schedule.id"
      :movie-title="schedule.movie?.title"
      :cinema-name="schedule.hall?.cinema?.name"
      :hall-name="schedule.hall?.name"
      :start-time="schedule.start_time"
      :price="schedule.price"
    />
  </div>
</template>
