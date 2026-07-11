<script setup>
defineProps({
  label: String,
  isActive: Boolean
})
defineEmits(['click'])
</script>

<template>
  <button 
    @click="$emit('click')"
    class="flex items-center gap-2 px-4 py-2 rounded-full border transition-all duration-200"
    :class="isActive 
      ? 'bg-lime-400 text-black border-lime-400' 
      : 'bg-gray-950 text-gray-400 border-gray-800 hover:border-gray-600'"
  >
    <span>{{ label }}</span>
    <span class="text-xs">▼</span>
  </button>
</template>