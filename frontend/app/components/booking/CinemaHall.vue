<script setup>
import { ref } from "vue";

const props = defineProps({
  hallName: String,

  seats: {
    type: Array,
    required: true,
  },
});

const selectedSeat = ref(null);

const toggleSeat = (seat) => {
  if (!seat.is_available) {
    return;
  }

  selectedSeat.value = selectedSeat.value === seat.id ? null : seat.id;
};
</script>

<template>
  <div class="flex flex-col items-center w-full">
    <div class="w-full h-1 bg-gray-700 rounded mb-3"></div>

    <p class="text-gray-500 uppercase text-xs tracking-widest mb-4">Screen</p>

    <h3 class="text-lime-400 font-semibold mb-5">
      {{ hallName }}
    </h3>

    <div class="grid grid-cols-6 gap-3">
      <button
        v-for="seat in seats"
        :key="seat.id"
        @click="toggleSeat(seat)"
        :disabled="!seat.is_available"
        class="w-10 h-10 rounded-xl text-xs font-bold transition"
        :class="
          selectedSeat === seat.id
            ? 'bg-lime-400 text-black scale-110'
            : seat.is_available
              ? 'bg-gray-800 text-gray-300 hover:bg-gray-700'
              : 'bg-gray-900 text-gray-600 cursor-not-allowed'
        "
      >
        {{ seat.id.slice(0, 2) }}
      </button>
    </div>

    <div class="flex gap-5 text-xs text-gray-400 mt-6">
      <div class="flex gap-2 items-center">
        <div class="w-3 h-3 bg-gray-800 rounded"></div>

        Available
      </div>

      <div class="flex gap-2 items-center">
        <div class="w-3 h-3 bg-lime-400 rounded"></div>

        Selected
      </div>

      <div class="flex gap-2 items-center">
        <div class="w-3 h-3 bg-gray-900 rounded"></div>

        Taken
      </div>
    </div>
  </div>
</template>
