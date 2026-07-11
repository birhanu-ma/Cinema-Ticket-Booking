<script setup>
defineProps({
  hallName: String,
  rows: { 
    type: Array, 
    required: true 
  }
})
</script>

<template>
  <div class="mt-8 flex flex-col items-center">
    <span class="text-lime-400">{{ hallName }}</span>
    
    <div class="grid grid-cols-4 gap-6 mt-8">
      <div v-for="seat in rows" :key="seat.id">
      </div>
    </div>
  </div>
</template>