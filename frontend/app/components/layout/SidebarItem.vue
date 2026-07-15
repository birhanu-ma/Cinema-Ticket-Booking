<script setup>
defineProps({
  icon: String,
  label: String,
  isActive: Boolean,
  badge: [String, Number]
});
</script>

<template>
  <div
    class="flex items-center gap-4 px-4 py-3 rounded-2xl transition-colors"
    :class="isActive
      ? 'bg-lime-400 text-black'
      : 'text-gray-400 hover:text-white'"
  >
    <span>{{ icon }}</span>
    <span class="flex-1">{{ label }}</span>

    <span
      v-if="badge"
      class="bg-lime-400 text-black px-2 rounded-full text-xs"
    >
      {{ badge }}
    </span>
  </div>
</template>