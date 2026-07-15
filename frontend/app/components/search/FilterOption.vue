<script setup>
defineProps({
  label: String,
  isActive: Boolean,
});
defineEmits(["click"]);
</script>

<template>
  <button
    @click="$emit('click')"
    type="button"
    class="flex items-center gap-2 px-5 py-2.5 rounded-full border text-xs font-semibold tracking-wide transition-all duration-200 cursor-pointer shrink-0"
    :class="
      isActive
        ? 'bg-lime-400 text-black border-lime-400 shadow-lg shadow-lime-400/10'
        : 'bg-gray-950 text-gray-400 border-gray-800 hover:border-gray-600 hover:text-gray-200'
    "
  >
    <span>{{ label }}</span>
    <span
      :class="[
        'text-[10px] transition-transform duration-200',
        isActive ? 'rotate-180' : '',
      ]"
      >▼</span
    >
  </button>
</template>
