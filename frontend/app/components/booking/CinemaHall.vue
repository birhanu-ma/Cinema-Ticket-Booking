<script setup>
import { ref } from 'vue'

defineProps({
  hallName: String,
  rows: { 
    type: Array, 
    required: true 
  }
})

const selectedSeatId = ref(null)

const toggleSeat = (seat) => {
  if (!seat.available) return
  
  if (selectedSeatId.value === seat.id) {
    selectedSeatId.value = null
  } else {
    selectedSeatId.value = seat.id
  }
}
</script>

<template>
  <div class="mt-6 flex flex-col items-center w-full">
    <div class="w-full h-1 bg-gradient-to-r from-transparent via-gray-700 to-transparent mb-2 rounded"></div>
    <span class="text-gray-500 text-[10px] tracking-widest uppercase mb-6">Screen</span>

    <span class="text-lime-400 font-medium text-sm mb-4">{{ hallName }}</span>
    <p class="text-lime-400 font-medium text-sm mb-4">Select available seat </p>

    
    <div class="w-full overflow-x-auto pb-4 scrollbar-hide snap-x">
      <div class="inline-flex gap-4 px-2">
        <button 
          v-for="seat in rows" 
          :key="seat.id"
          :disabled="!seat.available"
          @click="toggleSeat(seat)"
          :class="[
            'w-11 h-11 rounded-xl font-semibold text-xs flex items-center justify-center transition-all duration-200 snap-center flex-shrink-0',
            selectedSeatId === seat.id 
              ? 'bg-lime-400 text-black shadow-lg shadow-lime-400/20 scale-105' 
              : '',
            seat.available && selectedSeatId !== seat.id
              ? 'bg-gray-800 text-gray-300 hover:bg-gray-700 hover:text-white' 
              : '',
            !seat.available 
              ? 'bg-gray-900 text-gray-600 cursor-not-allowed border border-dashed border-gray-800' 
              : ''
          ]"
        >
          {{ seat.id }}
        </button>
      </div>
    </div>

    <div class="flex gap-4 mt-4 text-[11px] text-gray-400">
      <div class="flex items-center gap-1.5">
        <div class="w-3 h-3 bg-gray-800 rounded-md"></div> Available
      </div>
      <div class="flex items-center gap-1.5">
        <div class="w-3 h-3 bg-lime-400 rounded-md"></div> Selected
      </div>
      <div class="flex items-center gap-1.5">
        <div class="w-3 h-3 bg-gray-900 border border-dashed border-gray-800 rounded-md"></div> Taken
      </div>
    </div>
  </div>
</template>

<style scoped>
.scrollbar-hide::-webkit-scrollbar {
  display: none;
}
.scrollbar-hide {
  -ms-overflow-style: none;
  scrollbar-width: none;
}
</style>