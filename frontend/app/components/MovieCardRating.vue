<script setup>
defineProps({
  rating: {
    type: Number,
    default: 0
  }
})
</script>

<template>
  <div class="flex items-center gap-1">
    <span 
      v-for="star in 5" 
      :key="star"
      class="text-xl"
      :class="star <= rating ? 'text-yellow-400' : 'text-gray-600'"
    >
      ★
    </span>
  </div>
</template>