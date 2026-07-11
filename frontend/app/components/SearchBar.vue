<script setup>
const props = defineProps(["modelValue"]);
const emit = defineEmits(["update:modelValue"]);
</script>

<template>
  <div class="relative w-full max-w-2xl">
    <span class="absolute left-4 top-1/2 -translate-y-1/2 text-gray-400">
      🔍
    </span>


    <input
      :value="modelValue"
      @input="$emit('update:modelValue', $event.target.value)"
      type="text"
      placeholder="Type to Search..."
      class="w-full bg-transparent backdrop-blur-sm text-white pl-12 pr-4 py-3 rounded-full border border-white/10 focus:outline-none focus:border-lime-400 focus:ring-1 focus:ring-lime-400/30 transition-all placeholder:text-gray-500"
    />
  </div>
</template>
