<script setup>
import { ref, computed } from 'vue';

const director = ref({
  name: "Gina Prince-Bythewood",
  role: "Director",
  image: null 
});

const imageUrl = computed(() => {
  if (director.value.image) return director.value.image;
  const name = director.value.name || 'director';
  return `https://api.dicebear.com/9.x/adventurer/svg?seed=${encodeURIComponent(name)}`;
});
</script>

<template>
  <div class="flex flex-col">
  
    <div class="bg-gray-950 p-4 rounded-3xl flex flex-col items-center text-center w-32 shrink-0">
      <img 
        :src="imageUrl" 
        :alt="director.name"
        class="w-20 h-20 rounded-full object-cover mb-4 border-2 border-gray-800"
      />
      
      <div class="space-y-0.5">
        <h3 class="text-white font-bold text-sm truncate w-24">
          {{ director.name }}
        </h3>
        <p class="text-lime-400 text-xs">
          {{ director.role }}
        </p>
      </div>
    </div>
  </div>
</template>