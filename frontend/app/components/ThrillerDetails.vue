<script setup>
defineProps({
  movie: {
    type: Object,
    required: true
  }
})

const emit = defineEmits(['watch'])
</script>

<template>
  <div class="absolute bottom-6 left-6 flex items-end justify-between w-[95%]">
    <button 
      @click="$emit('watch')"
      class="bg-lime-400 text-black px-6 py-3 rounded-full font-bold flex items-center gap-2 hover:bg-lime-500 transition-colors"
    >
      Watch the trailer <span>▶</span>
    </button>

    <div class="text-white text-3xl font-black tracking-widest">
      {{ movie.title }}
    </div>
  </div>
</template>