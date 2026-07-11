<script setup>
import { ref, computed } from 'vue';

const dummyDescription = "This is an example of a long movie description that exceeds the character limit. In a real-world scenario, this would be fetched from an API, but for testing purposes, we are using this hardcoded string to ensure that our 'Read More' and 'Read Less' logic works perfectly within the Vue component.";

const props = defineProps({
  description: {
    type: String,
    default: dummyDescription 
  },
  maxLength: {
    type: Number,
    default: 150
  }
});

const isExpanded = ref(false);

const displayText = computed(() => {
  const text = props.description || '';
  
  if (isExpanded.value || text.length <= props.maxLength) {
    return text;
  }
  return text.slice(0, props.maxLength) + '...';
});
</script>

<template>
  <div class="text-gray-300 leading-relaxed">
    <h3 class="text-white font-bold text-lg mb-2">Description</h3>
    
    <p>
      {{ displayText }}
      <button 
        v-if="description?.length > maxLength"
        @click="isExpanded = !isExpanded"
        class="text-lime-400 font-bold ml-2 hover:underline focus:outline-none"
      >
        {{ isExpanded ? 'Read Less' : 'Read More' }}
      </button>
    </p>
  </div>
</template>