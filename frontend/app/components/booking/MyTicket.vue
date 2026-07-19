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

const ticketPrice = (ticket) => {
  const payment = ticket.payments?.[0];

  if (payment?.amount) {
    return `${payment.amount} ${payment.currency ?? "ETB"}`;
  }

  if (ticket.schedule?.price) {
    return `${ticket.schedule.price} ETB`;
  }

  return "—";
};

const ticketReference = (ticket) => {
  return (
    ticket.payments?.[0]?.transaction_ref || ticket.booking_reference || "—"
  );
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
    class="relative flex w-full max-w-3xl bg-white rounded-2xl shadow-sm overflow-hidden border border-lime-100"
  >
    <div class="w-28 shrink-0 bg-lime-100 relative">
      <img
        v-if="ticket.schedule?.movie?.movie_images?.[0]?.image_url"
        :src="ticket.schedule.movie.movie_images[0].image_url"
        :alt="ticket.schedule.movie.title"
        class="w-full h-full object-cover opacity-90"
      />

      <div
        v-else
        class="w-full h-full flex items-center justify-center text-lime-400 text-xs px-2 text-center"
      >
        No poster
      </div>
    </div>

    <div class="relative w-0 border-l-2 border-dashed border-lime-200">
      <span class="absolute -top-3 -left-3 w-6 h-6 rounded-full bg-lime-50" />
      <span
        class="absolute -bottom-3 -left-3 w-6 h-6 rounded-full bg-lime-50"
      />
    </div>

    <div class="flex-1 p-4 flex flex-col justify-between">
      <div>
        <div class="flex items-start justify-between gap-2">
          <h3 class="text-base font-semibold leading-tight text-lime-950">
            {{ ticket.schedule?.movie?.title || "Movie" }}
          </h3>

          <span
            class="shrink-0 px-2.5 py-1 rounded-full text-xs font-bold capitalize"
            :class="ticketStatusClasses(ticket.status)"
          >
            {{ ticket.status }}
          </span>
        </div>

        <p class="text-sm font-medium text-lime-800 mt-2">
          Issued for:
          {{ ticket.user?.name || "Guest" }}
        </p>

        <p class="text-sm text-lime-700 mt-1">
          {{ ticket.schedule?.hall?.cinema?.name }}
          &middot;
          {{ ticket.schedule?.hall?.name }}
        </p>

        <p class="text-sm text-lime-700">
          {{ formattedShowtime(ticket.schedule?.start_time) }}
        </p>
      </div>

      <div class="mt-4 pt-3 border-t border-dashed border-lime-200">
        <div class="flex items-end justify-between">
          <div>
            <p class="text-[10px] uppercase tracking-wide text-lime-500">
              Seat
            </p>

            <p class="text-lg font-bold font-mono text-lime-600">
              {{ seatLabel(ticket) || "—" }}
            </p>
          </div>

          <div class="text-right">
            <p class="text-[10px] uppercase tracking-wide text-lime-500">
              Price
            </p>

            <p class="text-lg font-bold text-lime-700">
              {{ ticketPrice(ticket) }}
            </p>
          </div>
        </div>

        <div class="mt-3">
          <p class="text-[10px] uppercase tracking-wide text-lime-500">
            Reference
          </p>

          <p class="text-xs font-mono text-lime-700 break-all">
            {{ ticketReference(ticket) }}
          </p>
        </div>
      </div>
    </div>
  </div>
</template>
