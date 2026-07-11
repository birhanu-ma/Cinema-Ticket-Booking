<script setup>
const sessionData = ref({
  movieTitle: 'Mavka: Forest Song',
  sessionsCount: 14,
  time: '10:30',
  date: '21 Mar',
  hallName: 'Lux Hall',
  seats: [
    { id: 1, available: true }, { id: 2, available: true }, { id: 3, available: false }, { id: 4, available: true },
    { id: 5, available: true }, { id: 6, available: false }, { id: 7, available: true }, { id: 8, available: true }
  ]
})
</script>

<template>
  <div class="bg-gray-950 p-7 rounded-3xl">
    <div class="flex justify-between items-start mb-6">
      <div>
        <h2 class="text-white text-xl font-bold">{{ sessionData.movieTitle }}</h2>
        <p class="text-gray-400 text-sm">{{ sessionData.sessionsCount }} sessions</p>
      </div>
      
      <div class="bg-lime-400 px-3 py-2 rounded-xl flex items-center gap-2">
        <span class="text-black">📅</span>
        <div class="text-black font-bold">
          <div class="text-xs">{{ sessionData.time }}</div>
          <div class="text-[10px]">{{ sessionData.date }}</div>
        </div>
      </div>
    </div>

    <CinemaHall 
      :hall-name="sessionData.hallName"
      :rows="sessionData.seats"
    />
  </div>
</template>