<script setup>
defineProps({
  cinemaId: {
    type: String,
    required: true,
  },

  name: {
    type: String,
    required: true,
  },

  location: {
    type: String,
    default: "Unknown location",
  },
});
</script>

<template>
  <NuxtLink
    :to="`/admin/cinema/${cinemaId}`"
    class="flex-1 p-5 flex flex-col justify-between hover:bg-gray-900/40 rounded-r-3xl transition-all duration-300 cursor-pointer group"
  >
    <div>
      <div class="flex justify-between items-center mb-3">
        <span
          class="text-[10px] bg-lime-400/10 text-lime-400 font-bold px-2.5 py-1 rounded-md uppercase"
        >
          Cinema
        </span>

        <button
          @click.stop.prevent
          class="text-gray-500 hover:text-white text-lg"
        >
          ⋮
        </button>
      </div>

      <h3
        class="text-white text-lg font-bold group-hover:text-lime-400 transition"
      >
        {{ name }}
      </h3>
    </div>

    <div class="w-full h-px bg-gray-900 my-4"></div>

    <div>
      <span class="text-[10px] text-gray-500 uppercase font-bold">
        Location
      </span>

      <p class="text-gray-300 text-sm mt-1">📍 {{ location }}</p>
    </div>
  </NuxtLink>
</template>
