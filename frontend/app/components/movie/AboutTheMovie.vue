<script setup>
import { ref, computed } from "vue";

const dummyDescription =
  "This is an example of a long movie description that exceeds the character limit. In a real-world scenario, this would be fetched from an API, but for testing purposes, we are using this hardcoded string to ensure that our 'Read More' and 'Read Less' logic works perfectly within the Vue component.";

const props = defineProps({
  description: {
    type: String,
    default: dummyDescription,
  },
  maxLength: {
    type: Number,
    default: 150,
  },
});

const isExpanded = ref(false);

const displayText = computed(() => {
  const text = props.description || "";

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
        v-if="description?.length > maxLength"
        @click="isExpanded = !isExpanded"
        class="ml-2 text-lime-400 font-medium hover:text-lime-300 transition-colors"
      >
        {{ isExpanded ? "Read Less" : "Read More" }}
      </button>
    </p>
  </section>
</template>
