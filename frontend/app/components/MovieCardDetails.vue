<script setup>
defineProps({
  title: {
    type: String,
    required: true
  },
  genre: {
    type: String,
    required: true
  }
})
</script>

<template>
  <div>
    <h3 class="text-lg font-semibold text-white truncate leading-tight">
      {{ title }}
    </h3>
    <p class="text-sm text-gray-400 mt-1">
      {{ genre }}
    </p>
  </div>
</template>