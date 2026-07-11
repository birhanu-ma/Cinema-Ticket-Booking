<script setup>
defineProps({
  icon: String, 
  label: String,
  isActive: Boolean,
  badge: [String, Number]
})
</script>

<template>
  <div 
    class="flex items-center gap-4 px-4 py-3 rounded-2xl cursor-pointer transition-colors"
    :class="isActive ? 'bg-lime-400 text-black' : 'text-gray-400 hover:text-white'"
  >
    <span class="text-xl">{{ icon }}</span>
    
    <span class="font-medium flex-1">{{ label }}</span>

    <span 
      v-if="badge" 
      class="bg-lime-400 text-black text-xs font-bold px-2 py-0.5 rounded-full"
    >
      {{ badge }}
    </span>
  </div>
</template>