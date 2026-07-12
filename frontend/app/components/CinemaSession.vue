<script setup>
import { ref } from "vue";

const sessionData = ref({
  movieTitle: "Mavka: Forest Song",
  sessionsCount: 14,
  time: "10:30",
  date: "21 Mar",
  hallName: "Lux Hall",
  seats: [
    { id: 1, available: true },
    { id: 2, available: true },
    { id: 3, available: false },
    { id: 4, available: true },
    { id: 5, available: true },
    { id: 6, available: false },
    { id: 7, available: true },
    { id: 8, available: true },
    { id: 9, available: true },
    { id: 10, available: false },
    { id: 11, available: true },
    { id: 12, available: true },
  ],
});
</script>

<template>
  <div class="bg-gray-950 max-w-md w-full p-7 rounded-3xl border border-gray-900/50 shadow-2xl">
    <div class="flex justify-between items-start mb-6">
      <div>
        <h2 class="text-white text-xl font-bold">
          {{ sessionData.movieTitle }}
        </h2>
        <p class="text-gray-400 text-sm mt-0.5">
          {{ sessionData.sessionsCount }} sessions
        </p>
      </div>

      <div class="bg-lime-400 px-3 py-1.5 rounded-xl flex items-center gap-2 shadow-md shadow-lime-400/10">
        <span class="text-black text-sm">📅</span>
        <div class="text-black font-bold leading-tight">
          <div class="text-xs font-extrabold">{{ sessionData.time }}</div>
          <div class="text-[9px] uppercase tracking-wider opacity-80">
            {{ sessionData.date }}
          </div>
        </div>
      </div>
    </div>

    <CinemaHall :hall-name="sessionData.hallName" :rows="sessionData.seats" />
  </div>
</template>