<script setup>
import { computed } from 'vue';

const props = defineProps({
  actor: {
    type: Object,
    required: true, 
    default: () => ({ name: 'Unknown', role: 'N/A', image: null })
  }
});

const imageUrl = computed(() => {
  if (props.actor.image) return props.actor.image;
  const name = props.actor.name || 'placeholder';
  return `https://api.dicebear.com/9.x/adventurer/svg?seed=${encodeURIComponent(name)}`;
});
</script>

<template>
  <div class="bg-gray-950 p-4 rounded-3xl flex flex-col items-center text-center w-32 shrink-0">
    <img 
      :src="imageUrl" 
      :alt="actor.name"
      class="w-20 h-20 rounded-full object-cover mb-4 border-2 border-gray-800"
    />
    
    <div class="space-y-0.5">
      <h3 class="text-white font-bold text-sm truncate w-24">
        {{ actor.name }}
      </h3>
      <p class="text-lime-400 text-xs">
        {{ actor.role }}
      </p>
    </div>
  </div>
</template>