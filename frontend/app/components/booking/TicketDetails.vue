<script setup>
defineProps({
  ticket: {
    type: Object,
    required: true
  }
})
</script>

<template>
  <div class="flex-1 p-2 flex flex-col justify-between">
    <div>
      <h2 class="text-2xl font-bold text-white">{{ ticket.movieTitle }}</h2>
      <p class="text-gray-400 mt-1">{{ ticket.date }}</p>
      
      <div class="mt-6 space-y-2 text-gray-300">
        <div class="flex justify-between">
          <span>Hall:</span>
          <span class="font-bold text-white">{{ ticket.hall }}</span>
        </div>
        <div class="flex justify-between">
          <span>Seats:</span>
          <span class="font-bold text-white">{{ ticket.seats }}</span>
        </div>
      </div>
    </div>

    <div class="flex items-center justify-between my-4">
      <span class="text-3xl mx-2 font-bold text-white">${{ ticket.price }}</span>
      <button class="bg-lime-400 text-black px-6 py-2 rounded-full font-bold hover:bg-lime-500">
        Buy tickets
      </button>
    </div>
  </div>
</template>