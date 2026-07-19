<script setup>
defineProps({
  seatId: {
    type: String,
    required: true,
  },

  seatNumber: {
    type: [String, Number],
    default: "-",
  },

  row: {
    type: String,
    default: "",
  },

  type: {
    type: String,
    default: "standard",
  },

  hall: {
    type: String,
    default: "",
  },

  cinema: {
    type: String,
    default: "",
  },
});

const typeClass = (type) => {
  switch (type?.toLowerCase()) {
    case "vip":
      return "bg-yellow-400/10 text-yellow-400";

    case "couple":
      return "bg-pink-400/10 text-pink-400";

    case "wheelchair":
      return "bg-blue-400/10 text-blue-400";

    default:
      return "bg-lime-400/10 text-lime-400";
  }
};
</script>

<template>
  <NuxtLink
    :to="`/admin/seat/${seatId}`"
    class="flex-1 p-5 flex flex-col justify-between hover:bg-gray-900/40 rounded-r-3xl transition-all duration-300 group"
  >
    <div>
      <div class="flex items-center justify-between mb-3">
        <span
          class="text-[10px] bg-blue-400/10 text-blue-400 font-bold px-2.5 py-1 rounded-md uppercase tracking-wider"
        >
          Seat
        </span>

     
      </div>

      <h3
        class="text-white text-lg font-bold group-hover:text-blue-400 transition-colors"
      >
        Seat {{ row }}{{ seatNumber }}
      </h3>

      <p class="text-gray-400 text-sm mt-2">
        {{ hall }}

        <span v-if="cinema"> • {{ cinema }} </span>
      </p>
    </div>

    <div class="w-full h-px bg-gray-900 my-4"></div>

    <div class="flex items-center justify-between">
      <div>
        <span
          class="text-[10px] text-gray-500 font-bold uppercase tracking-wider"
        >
          Seat Type
        </span>

        <p
          class="inline-block mt-1 px-3 py-1 rounded-lg text-xs font-bold capitalize"
          :class="typeClass(type)"
        >
          {{ type }}
        </p>
      </div>
    </div>
  </NuxtLink>
</template>
