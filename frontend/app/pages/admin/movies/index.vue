<script setup>
definePageMeta({
  layout: "admin",
});

const premiumMoviesList = [
  {
    title: "Shazam! Fury of the Gods",
    genre: "Action/Adventure",
    duration: "130 min",
    thumbnail:
      "https://images.unsplash.com/photo-1536440136628-849c177e76a1?w=500",
  },
  {
    title: "Mavka: Forest Song",
    genre: "Animation/Fantasy",
    duration: 99,
    thumbnail:
      "https://images.unsplash.com/photo-1489599849927-2ee91cede3ba?w=500",
  },
  {
    title: "Shazam! Fury of the Gods",
    genre: "Action/Adventure",
    duration: "130 min",
    thumbnail:
      "https://images.unsplash.com/photo-1536440136628-849c177e76a1?w=500",
  },
  {
    title: "Mavka: Forest Song",
    genre: "Animation/Fantasy",
    duration: 99,
    thumbnail:
      "https://images.unsplash.com/photo-1489599849927-2ee91cede3ba?w=500",
  },
  {
    title: "Shazam! Fury of the Gods",
    genre: "Action/Adventure",
    duration: "130 min",
    thumbnail:
      "https://images.unsplash.com/photo-1536440136628-849c177e76a1?w=500",
  },
  {
    title: "Mavka: Forest Song",
    genre: "Animation/Fantasy",
    duration: 99,
    thumbnail:
      "https://images.unsplash.com/photo-1489599849927-2ee91cede3ba?w=500",
  },
  {
    title: "Shazam! Fury of the Gods",
    genre: "Action/Adventure",
    duration: "130 min",
    thumbnail:
      "https://images.unsplash.com/photo-1536440136628-849c177e76a1?w=500",
  },
  {
    title: "Mavka: Forest Song",
    genre: "Animation/Fantasy",
    duration: 99,
    thumbnail:
      "https://images.unsplash.com/photo-1489599849927-2ee91cede3ba?w=500",
  },
  {
    title: "Shazam! Fury of the Gods",
    genre: "Action/Adventure",
    duration: "130 min",
    thumbnail:
      "https://images.unsplash.com/photo-1536440136628-849c177e76a1?w=500",
  },
  {
    title: "Mavka: Forest Song",
    genre: "Animation/Fantasy",
    duration: 99,
    thumbnail:
      "https://images.unsplash.com/photo-1489599849927-2ee91cede3ba?w=500",
  },
  {
    title: "Mavka: Forest Song",
    genre: "Animation/Fantasy",
    duration: 99,
    thumbnail:
      "https://images.unsplash.com/photo-1489599849927-2ee91cede3ba?w=500",
  },
  {
    title: "Mavka: Forest Song",
    genre: "Animation/Fantasy",
    duration: 99,
    thumbnail:
      "https://images.unsplash.com/photo-1489599849927-2ee91cede3ba?w=500",
  },
];
</script>

<template>
  <div
    class="flex bg-linear-to-t from-[#51751f] to-transparent flex-col gap-6 h-[calc(100vh-4rem)] overflow-hidden"
  >
    <div>
      <h1 class="text-2xl font-bold tracking-wide text-gray-100">
        Manage Movies
      </h1>

      <p class="text-xs text-gray-500 mt-0.5">
        Active catalog listings created by admin
      </p>
    </div>

    <div class="flex gap-10 flex-1 min-h-0 overflow-hidden">
      <div
        class="w-full space-y-6 overflow-y-auto pb-12 [&::-webkit-scrollbar]:hidden [-ms-overflow-style:none] [scrollbar-width:none]"
      >
        <AdminMoviesMovieCard
          v-for="(movieData, i) in premiumMoviesList"
          :key="i"
          :movie="movieData"
        />
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
