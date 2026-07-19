<script setup>
import { ref, computed } from "vue";

const props = defineProps({
  movie: {
    type: Object,
    required: true,
  },

  maxLength: {
    type: Number,
    default: 150,
  },
});

const isExpanded = ref(false);

const displayText = computed(() => {
  const text = props.movie?.description || "";

  if (isExpanded.value || text.length <= props.maxLength) {
    return text;
  }

  return text.slice(0, props.maxLength) + "...";
});
</script>

<template>
  <section class="bg-zinc-900 border border-zinc-800 rounded-xl p-6">
    <h2 class="text-white text-xl font-semibold mb-4">About the Movie</h2>

    <p class="text-gray-400 leading-7">
      {{ displayText }}

      <button
        v-if="movie?.description?.length > maxLength"
        @click="isExpanded = !isExpanded"
        class="ml-2 text-lime-400 font-medium cursor-pointer hover:text-lime-300 transition-colors"
      >
        {{ isExpanded ? "Read Less" : "Read More" }}
      </button>
    </p>
  </section>
</template>
