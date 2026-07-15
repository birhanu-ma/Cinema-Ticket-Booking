<script setup>

const currentTicket = {
  movieTitle: "Mavka: Forest Song",
  date: "21st Mar, 10:30",
  hall: "Lux",
  seats: "B7, B8",
  price: 26,
  thumbnail:
    "https://images.unsplash.com/photo-1536440136628-849c177e76a1?w=400",
};
</script>

<template>
  <div
    class="flex bg-gray-950 rounded-3xl overflow-hidden shadow-xl h-64 max-w-lg"
  >
    <BookingTicketDetails :ticket="currentTicket" />

    <div class="w-40 relative flex-shrink-0">
      <img
        :src="currentTicket.thumbnail"
        class="w-full h-full object-cover"
        alt="Movie Thumbnail"
      />
    </div>
  </div>
</template>
