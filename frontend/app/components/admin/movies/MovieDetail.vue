<script setup>
defineProps({
  // Accept the ID from the database to construct the route
  movieId: {
    type: [String, Number],
    required: true,
  },
  title: {
    type: String,
    required: true,
  },
  genre: {
    type: String,
    required: true,
  },
  duration: {
    type: [Number, String], 
    required: true,
  },
});
</script>

<template>
  <NuxtLink
    :to="`/admin/movies/${movieId}`"
    class="flex-1 p-5 flex cursor-pointer flex-col justify-between hover:bg-gray-900/40 rounded-r-3xl transition-all duration-300 group block"
  >
    <div>
      <div class="flex items-center justify-between mb-2">
        <span
          class="text-[10px] bg-lime-400/10 text-lime-400 font-bold px-2.5 py-0.5 rounded-md uppercase tracking-wider"
        >
          {{ genre }}
        </span>

        <button
          @click.stop.prevent="console.log('Context menu triggered')"
          class="text-gray-500 hover:text-white transition-colors text-lg leading-none cursor-pointer p-1"
        >
          ⋮
        </button>
      </div>

      <h3
        class="text-white text-lg font-bold leading-snug tracking-wide line-clamp-2 group-hover:text-lime-400 transition-colors"
      >
        {{ title }}
      </h3>
    </div>

    <div class="w-full h-px bg-gray-900 my-4"></div>

    <div class="flex items-center justify-between gap-4">
      <div class="flex flex-col gap-0.5">
        <span
          class="text-[10px] text-gray-500 font-bold uppercase tracking-wider"
          >Duration</span
        >
        <div
          class="flex items-center gap-1.5 text-gray-300 text-xs font-semibold"
        >
          <span class="text-gray-400">🕒</span>
          <span>{{
            typeof duration === "number" ? `${duration} min` : duration
          }}</span>
        </div>
      </div>
    </div>
  </NuxtLink>
</template>
