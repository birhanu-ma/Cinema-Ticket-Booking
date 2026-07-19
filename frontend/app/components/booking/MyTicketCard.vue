<script setup>
defineProps({
  ticket: {
    type: Object,
    required: true,
  },
});

const formattedShowtime = (isoString) => {
  if (!isoString) return "";

  return new Date(isoString).toLocaleString(undefined, {
    weekday: "short",
    month: "short",
    day: "numeric",
    hour: "numeric",
    minute: "2-digit",
  });
};

const seatLabel = (ticket) => {
  const seat = ticket.schedule_seat?.seat;

  if (!seat) return null;

  return `${seat.row_name}${seat.seat_number}`;
};

const ticketStatusClasses = (ticketStatus) => {
  if (ticketStatus === "confirmed") {
    return "bg-lime-400 text-black";
  }

  if (ticketStatus === "failed") {
    return "bg-red-500 text-white";
  }

  return "bg-gray-700 text-lime-300";
};
</script>

<template>
  <div
    class="relative flex bg-gray-950 rounded-2xl shadow-xl overflow-hidden border border-gray-900"
  >
    <!-- Poster / left stub -->
    <div class="w-28 shrink-0 bg-black relative">
      <img
        v-if="ticket.schedule?.movie?.movie_images?.[0]?.image_url"
        :src="ticket.schedule.movie.movie_images[0].image_url"
        :alt="ticket.schedule.movie.title"
        class="w-full h-full object-cover opacity-90"
      />
      <div
        v-else
        class="w-full h-full flex items-center justify-center text-gray-600 text-xs px-2 text-center"
      >
        No poster
      </div>
    </div>

    <!-- Perforation between stub and body -->
    <div class="relative w-0 border-l-2 border-dashed border-gray-800">
      <span class="absolute -top-3 -left-3 w-6 h-6 rounded-full bg-black" />
      <span class="absolute -bottom-3 -left-3 w-6 h-6 rounded-full bg-black" />
    </div>

    <!-- Main ticket body -->
    <div class="flex-1 p-4 flex flex-col justify-between">
      <div>
        <div class="flex items-start justify-between gap-2">
          <h3 class="text-base font-semibold leading-tight text-white">
            {{ ticket.schedule?.movie?.title || "Movie" }}
          </h3>

          <span
            class="shrink-0 px-2.5 py-1 rounded-full text-xs font-bold capitalize"
            :class="ticketStatusClasses(ticket.status)"
          >
            {{ ticket.status }}
          </span>
        </div>

        <p class="text-sm text-gray-400 mt-1">
          {{ ticket.schedule?.hall?.cinema?.name }} &middot;
          {{ ticket.schedule?.hall?.name }}
        </p>

        <p class="text-sm text-gray-400">
          {{ formattedShowtime(ticket.schedule?.start_time) }}
        </p>
      </div>

      <div
        class="flex items-end justify-between mt-4 pt-3 border-t border-dashed border-gray-800"
      >
        <div>
          <p class="text-[10px] uppercase tracking-wide text-gray-500">
            Seat
          </p>
          <p class="text-lg font-bold font-mono text-lime-400">
            {{ seatLabel(ticket) || "—" }}
          </p>
        </div>

        <p class="text-[10px] font-mono text-gray-600 tracking-wider">
          #{{ ticket.id.slice(0, 8) }}
        </p>
      </div>
    </div>
  </div>
</template>