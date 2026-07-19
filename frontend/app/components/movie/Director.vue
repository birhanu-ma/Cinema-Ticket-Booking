<script setup>
import { computed } from "vue";

const props = defineProps({
  director: {
    type: Object,
    required: false,
    default: null,
  },
});

const imageUrl = computed(() => {
  if (props.director?.photo_url) {
    return props.director.photo_url;
  }

  return `https://api.dicebear.com/9.x/adventurer/svg?seed=${encodeURIComponent(
    props.director?.name || "director",
  )}`;
});
</script>

<template>
  <div v-if="director" class="flex flex-col">
    <div
      class="bg-gray-950 p-4 rounded-3xl flex flex-col items-center text-center w-32 shrink-0"
    >
      <img
        :src="imageUrl"
        :alt="director.name"
        class="w-20 h-20 rounded-full object-cover mb-4 border-2 border-gray-800"
      />

      <div class="space-y-0.5">
        <h3 class="text-white font-bold text-sm truncate w-24">
          {{ director.name }}
        </h3>

        <p class="text-lime-400 text-xs">Director</p>
      </div>
    </div>
  </div>

  <!-- No director on this movie (director_id is nullable) - show a
  lightweight placeholder instead of rendering nothing at all, so the
  layout doesn't visibly shift where this card would've been. -->
  <div v-else class="flex flex-col">
    <div
      class="bg-gray-950 p-4 rounded-3xl flex flex-col items-center text-center w-32 shrink-0 border border-dashed border-gray-800"
    >
      <div
        class="w-20 h-20 rounded-full mb-4 border-2 border-gray-800 bg-gray-900 flex items-center justify-center text-gray-600 text-xs"
      >
        N/A
      </div>

      <div class="space-y-0.5">
        <h3 class="text-gray-500 font-bold text-sm truncate w-24">Unknown</h3>
        <p class="text-gray-600 text-xs">Director</p>
      </div>
    </div>
  </div>
</template>
