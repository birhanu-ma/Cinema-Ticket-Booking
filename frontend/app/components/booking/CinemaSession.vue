<script setup>
import { computed } from "vue";

const props = defineProps({
  movie: {
    type: Object,
    required: true,
  },
});

const schedule = computed(() => {
  return props.movie?.schedules?.[0] ?? null;
});

const sessionInfo = computed(() => {
  if (!schedule.value) {
    return null;
  }

  const seats = schedule.value.schedule_seats ?? [];

  return {
    cinemaName: schedule.value.hall?.cinema?.name ?? "Unknown Cinema",

    cinemaLocation: schedule.value.hall?.cinema?.location ?? "Unknown Location",

    hallName: schedule.value.hall?.name ?? "Unknown Hall",

    capacity: schedule.value.hall?.capacity ?? 0,

    totalSeats: seats.length,

    availableSeats: seats.filter((seat) => seat.is_available).length,

    price: schedule.value.price,

    time: new Date(schedule.value.start_time).toLocaleTimeString([], {
      hour: "2-digit",
      minute: "2-digit",
    }),

    date: new Date(schedule.value.start_time).toLocaleDateString([], {
      day: "numeric",
      month: "short",
    }),
  };
});
</script>

<template>
  <div
    v-if="sessionInfo"
    class="bg-gray-950 border border-gray-800 rounded-3xl p-7 max-w-md w-full shadow-xl"
  >
    <div class="flex justify-between items-start mb-6">
      <div>
        <h2 class="text-white text-xl font-bold">
          {{ sessionInfo.cinemaName }}
        </h2>

        <p class="text-gray-400 text-sm">
          {{ sessionInfo.cinemaLocation }}
        </p>
      </div>

      <div class="bg-lime-400 text-black rounded-xl px-3 py-2 text-center">
        <p class="font-bold text-sm">
          {{ sessionInfo.time }}
        </p>

        <p class="text-xs">
          {{ sessionInfo.date }}
        </p>
      </div>
    </div>

    <div class="space-y-3 text-gray-300">
      <div class="flex justify-between">
        <span> Hall </span>

        <span class="font-semibold">
          {{ sessionInfo.hallName }}
        </span>
      </div>

      <div class="flex justify-between">
        <span> Capacity </span>

        <span>
          {{ sessionInfo.capacity }}
        </span>
      </div>

      <div class="flex justify-between">
        <span> Available Seats </span>

        <span class="text-lime-400 font-bold">
          {{ sessionInfo.availableSeats }}
          /
          {{ sessionInfo.totalSeats }}
        </span>
      </div>

      <div class="flex justify-between">
        <span> Price </span>

        <span> {{ sessionInfo.price }} ETB </span>
      </div>
    </div>

    <div class="mt-6">
      <CinemaHall
        :hall-name="sessionInfo.hallName"
        :seats="schedule.schedule_seats"
      />
    </div>
  </div>

  <div v-else class="bg-gray-950 text-gray-400 rounded-xl p-6">
    No sessions available.
  </div>
</template>
