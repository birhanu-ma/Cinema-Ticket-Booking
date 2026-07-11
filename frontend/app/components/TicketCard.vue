<script setup>
defineProps({
  ticket: {
    type: Object,
    required: true,
  },
});

const ticket = [
  {
    movieTitle: "Mavka: Forest Song",
    date: "21st Mar, 10:30",
    hall: "Lux",
    seats: "B7, B8",
    price: 26,
    thumbnail: "https://i.pravatar.cc/150?u=kiki",
  },
];
</script>

<template>
  <div class="flex bg-gray-950 rounded-3xl overflow-hidden shadow-xl">
    <TicketDetails :ticket="ticket" />

    <div class="w-32 relative">
      <img
        :src="ticket.thumbnail"
        class="w-full h-full object-cover"
        alt="Movie Thumbnail"
      />
    </div>
  </div>
</template>
