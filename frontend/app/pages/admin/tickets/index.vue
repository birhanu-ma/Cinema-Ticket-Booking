<script setup>
definePageMeta({
  layout: "admin",
});

const dynamicTickets = [
  {
    ticketId: "ST081",
    title: "Shazam! Fury of the Gods - IMAX 3D Premiere Screening",
    status: "Inprogress",
    completionDate: "20-03-2026",
    thumbnail:
      "https://images.unsplash.com/photo-1536440136628-849c177e76a1?w=500",
  },
  {
    ticketId: "ST082",
    title: "Mavka: Forest Song - Premiere Screening",
    status: "Inprogress",
    completionDate: "21-03-2026",
    thumbnail:
      "https://images.unsplash.com/photo-1489599849927-2ee91cede3ba?w=500",
  },
  {
    ticketId: "ST081",
    title: "Shazam! Fury of the Gods - IMAX 3D Premiere Screening",
    status: "Inprogress",
    completionDate: "20-03-2026",
    thumbnail:
      "https://images.unsplash.com/photo-1536440136628-849c177e76a1?w=500",
  },
  {
    ticketId: "ST082",
    title: "Mavka: Forest Song - Premiere Screening",
    status: "Inprogress",
    completionDate: "21-03-2026",
    thumbnail:
      "https://images.unsplash.com/photo-1489599849927-2ee91cede3ba?w=500",
  },
  {
    ticketId: "ST081",
    title: "Shazam! Fury of the Gods - IMAX 3D Premiere Screening",
    status: "Inprogress",
    completionDate: "20-03-2026",
    thumbnail:
      "https://images.unsplash.com/photo-1536440136628-849c177e76a1?w=500",
  },
  {
    ticketId: "ST082",
    title: "Mavka: Forest Song - Premiere Screening",
    status: "Inprogress",
    completionDate: "21-03-2026",
    thumbnail:
      "https://images.unsplash.com/photo-1489599849927-2ee91cede3ba?w=500",
  },
  {
    ticketId: "ST083",
    title: "Avatar: The Way of Water",
    status: "Inprogress",
    completionDate: "22-03-2026",
    thumbnail:
      "https://images.unsplash.com/photo-1536440136628-849c177e76a1?w=500",
  },
  {
    ticketId: "ST084",
    title: "Interstellar - IMAX Revival",
    status: "Inprogress",
    completionDate: "25-03-2026",
    thumbnail:
      "https://images.unsplash.com/photo-1489599849927-2ee91cede3ba?w=500",
  },
];

const premiumCustomers = [
  "https://images.unsplash.com/photo-1534528741775-53994a69daeb?w=100",
  "https://images.unsplash.com/photo-1507003211169-0a1dd7228f2d?w=100",
  "https://images.unsplash.com/photo-1494790108377-be9c29b29330?w=100",
  "https://images.unsplash.com/photo-1500648767791-00dcc994a43e?w=100",
];
</script>

<template>
  <div
    class="flex bg-linear-to-t from-[#51751f] to-transparent flex-col gap-6 h-[calc(100vh-4rem)] overflow-hidden"
  >
    <div>
      <h1 class="text-2xl font-bold tracking-wide text-gray-100">
        Ticket Overview Logs
      </h1>
    </div>

    <div class="flex gap-10 flex-1 min-h-0 overflow-hidden">
      <div
        class="w-full overflow-y-auto pb-12 [&::-webkit-scrollbar]:hidden [-ms-overflow-style:none] [scrollbar-width:none]"
      >
        <div class="flex flex-col gap-4">
          <AdminTicketsTicketStatus
            v-for="singleTicket in dynamicTickets"
            :key="singleTicket.ticketId"
            :ticket="singleTicket"
          />
        </div>
      </div>

      <div class="flex flex-col gap-6 sticky top-0 h-fit flex-shrink-0 pb-12">
        <div class="flex gap-6">
          <AdminTicketsTicketStatCard />

          <AdminTicketsTicketStatCard bg-color="bg-[#96d13c]" />
        </div>

        <div class="flex gap-6">
          <AdminTicketsTicketStatCard bg-color="bg-[#2c6e59]" />
          <AdminTicketsTicketStatCard bg-color="bg-[#adadad]" />
        </div>

        <AdminMoviesAddMovie class="w-150 h-60 rounded-lg" />
      </div>
    </div>
  </div>
</template>
