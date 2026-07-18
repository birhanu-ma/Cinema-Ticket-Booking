<script setup>
const props = defineProps({
  scheduleId: {
    type: String,
    required: true,
  },

  movieTitle: {
    type: String,
    default: "Unknown Movie",
  },

  cinemaName: {
    type: String,
    default: "Unknown Cinema",
  },

  hallName: {
    type: String,
    default: "Unknown Hall",
  },

  startTime: {
    type: String,
    default: null,
  },

  price: {
    type: [Number, String],
    default: 0,
  },
});

const formatDate = (date) => {
  if (!date) return "-";

  return new Date(date).toLocaleString("en-US", {
    dateStyle: "medium",
    timeStyle: "short",
  });
};
</script>

<template>
  <NuxtLink
    :to="`/admin/schedule/${scheduleId}`"
    class="flex-1 p-5 flex flex-col justify-between hover:bg-gray-900/40 rounded-r-3xl transition-all duration-300 cursor-pointer group"
  >
    <div>
      <div class="flex items-center justify-between mb-3">
        <span
          class="text-[10px] bg-purple-400/10 text-purple-400 font-bold px-2.5 py-1 rounded-md uppercase"
        >
          Schedule
        </span>

        <button
          @click.stop.prevent
          class="text-gray-500 hover:text-white text-lg"
        >
          ⋮
        </button>
      </div>

      <h3
        class="text-white text-lg font-bold group-hover:text-purple-400 transition"
      >
        {{ movieTitle }}
      </h3>

      <p class="text-gray-400 text-sm mt-2">🏢 {{ cinemaName }}</p>

      <p class="text-gray-400 text-sm">🎬 {{ hallName }}</p>
    </div>

    <div class="w-full h-px bg-gray-900 my-4"></div>

    <div class="flex items-center justify-between">
      <div>
        <span class="text-[10px] text-gray-500 uppercase font-bold">
          Show Time
        </span>

        <p class="text-gray-200 text-xs font-semibold">
          📅 {{ formatDate(startTime) }}
        </p>
      </div>

      <div class="text-right">
        <span class="text-[10px] text-gray-500 uppercase font-bold">
          Price
        </span>

        <p class="text-lime-400 font-bold text-sm">{{ price }} ETB</p>
      </div>
    </div>
  </NuxtLink>
</template>
