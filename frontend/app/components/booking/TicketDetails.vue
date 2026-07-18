<script setup>
defineProps({
  ticket: {
    type: Object,
    required: true,
  },
});
</script>

<template>
  <div class="bg-black p-5 flex gap-5 items-center">
    <div class="flex-1">
      <h2 class="text-2xl font-bold text-white">
        {{ ticket.movieTitle }}
      </h2>

      <p class="text-gray-400 mt-2">
        {{
          ticket.schedule
            ? new Date(ticket.schedule.start_time).toLocaleString([], {
                day: "numeric",
                month: "short",
                hour: "2-digit",
                minute: "2-digit",
              })
            : "No Schedule"
        }}
      </p>

      <div class="mt-5 space-y-2 text-gray-300">
        <div class="flex justify-between">
          <span> Cinema </span>

          <strong class="text-white">
            {{ ticket.schedule?.hall?.cinema?.name }}
          </strong>
        </div>

        <div class="flex justify-between">
          <span> Hall </span>

          <strong class="text-white">
            {{ ticket.schedule?.hall?.name }}
          </strong>
        </div>

        <div class="flex justify-between">
          <span> Seat </span>

          <strong class="text-lime-400">
            {{ ticket.seat ? ticket.seat.seat_number : "Select Seat" }}
          </strong>
        </div>
      </div>

      <div class="mt-5 flex justify-between items-center">
        <span class="text-3xl font-bold text-white">
          {{ ticket.price }}
          ETB
        </span>

        <button
          class="bg-lime-400 text-black px-6 py-3 rounded-full font-bold hover:bg-lime-500"
        >
          Buy Tickets
        </button>
      </div>
    </div>

    <img
      :src="ticket.thumbnail"
      class="w-32 h-44 rounded-xl object-cover"
      alt="Movie"
    />
  </div>
</template>
