<script setup>
defineProps({
  hallId: {
    type: String,
    required: true,
  },

  name: {
    type: String,
    required: true,
  },

  capacity: {
    type: Number,
    default: 0,
  },

  cinemaName: {
    type: String,
    default: "Unknown Cinema",
  },
});
</script>

<template>
  <NuxtLink
    :to="`/admin/cinemaHall/${hallId}`"
    class="flex-1 p-5 flex flex-col justify-between hover:bg-gray-900/40 rounded-r-3xl transition-all duration-300 cursor-pointer group"
  >
    <div>
      <div class="flex justify-between items-center mb-3">
        <span
          class="text-[10px] bg-blue-400/10 text-blue-400 font-bold px-2.5 py-1 rounded-md uppercase"
        >
          Hall
        </span>

   
      </div>

      <h3
        class="text-white text-lg font-bold group-hover:text-blue-400 transition"
      >
        {{ name }}
      </h3>

      <p class="text-gray-400 text-sm mt-2">🏢 {{ cinemaName }}</p>
    </div>

    <div class="w-full h-px bg-gray-900 my-4"></div>

    <div>
      <span class="text-[10px] text-gray-500 uppercase font-bold">
        Capacity
      </span>

      <p class="text-gray-200 text-sm font-semibold mt-1">
        🪑 {{ capacity }} Seats
      </p>
    </div>
  </NuxtLink>
</template>
