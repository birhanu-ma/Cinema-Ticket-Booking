<script setup>
const movies = [
  {
    id: 1,
    title: "Shazam! Fury of the Gods",
    genre: "Action/Adventure",
    rating: 4, 
    posterUrl: "https://i.pravatar.cc/150?u=kiki", 
  },
  {
    id: 1,
    title: "Shazam! Fury of the Gods",
    genre: "Action/Adventure",
    rating: 4, 
    posterUrl: "https://i.pravatar.cc/150?u=kiki", 
  },
  {
    id: 1,
    title: "Shazam! Fury of the Gods",
    genre: "Action/Adventure",
    rating: 4, 
    posterUrl: "https://i.pravatar.cc/150?u=kiki", 
  },
  {
    id: 1,
    title: "Shazam! Fury of the Gods",
    genre: "Action/Adventure",
    rating: 4, 
    posterUrl: "https://i.pravatar.cc/150?u=kiki", 
  },
  {
    id: 1,
    title: "Shazam! Fury of the Gods",
    genre: "Action/Adventure",
    rating: 4, 
    posterUrl: "https://i.pravatar.cc/150?u=kiki", 
  },
  {
    id: 1,
    title: "Shazam! Fury of the Gods",
    genre: "Action/Adventure",
    rating: 4, 
    posterUrl: "https://i.pravatar.cc/150?u=kiki", 
  },
  {
    id: 1,
    title: "Shazam! Fury of the Gods",
    genre: "Action/Adventure",
    rating: 4, 
    posterUrl: "https://i.pravatar.cc/150?u=kiki", 
  },
  {
    id: 1,
    title: "Shazam! Fury of the Gods",
    genre: "Action/Adventure",
    rating: 4, 
    posterUrl: "https://i.pravatar.cc/150?u=kiki", 
  },
  {
    id: 1,
    title: "Shazam! Fury of the Gods",
    genre: "Action/Adventure",
    rating: 4, 
    posterUrl: "https://i.pravatar.cc/150?u=kiki", 
  },
  {
    id: 1,
    title: "Shazam! Fury of the Gods",
    genre: "Action/Adventure",
    rating: 4, 
    posterUrl: "https://i.pravatar.cc/150?u=kiki", 
  },
  {
    id: 1,
    title: "Shazam! Fury of the Gods",
    genre: "Action/Adventure",
    rating: 4, 
    posterUrl: "https://i.pravatar.cc/150?u=kiki", 
  },
  {
    id: 1,
    title: "Shazam! Fury of the Gods",
    genre: "Action/Adventure",
    rating: 4, 
    posterUrl: "https://i.pravatar.cc/150?u=kiki", 
  },
  {
    id: 1,
    title: "Shazam! Fury of the Gods",
    genre: "Action/Adventure",
    rating: 4, 
    posterUrl: "https://i.pravatar.cc/150?u=kiki", 
  },
  {
    id: 1,
    title: "Shazam! Fury of the Gods",
    genre: "Action/Adventure",
    rating: 4, 
    posterUrl: "https://i.pravatar.cc/150?u=kiki", 
  },
  {
    id: 1,
    title: "Shazam! Fury of the Gods",
    genre: "Action/Adventure",
    rating: 4, 
    posterUrl: "https://i.pravatar.cc/150?u=kiki", 
  },
  {
    id: 1,
    title: "Shazam! Fury of the Gods",
    genre: "Action/Adventure",
    rating: 4, 
    posterUrl: "https://i.pravatar.cc/150?u=kiki", 
  },
];
</script>
<template>
  <div class="p-10 bg-gray-900 min-h-screen">
    <div class="flex flex-row gap-2">
      <div class="flex flex-col gap-4">
        <WatchThriller />
        <div class="flex gap-2 flex-row">
          <Director />
          <Stars />
        </div>
      </div>
      <div class="flex flex-col gap-6">
        <CinemaSession />
        <TicketCard />
      </div>
    </div>

    <AboutTheMovie />
    <div class="grid grid-cols-2 md:grid-cols-4 lg:grid-cols-7 my-10 gap-4">
      <MovieCard v-for="movie in movies" :key="movie.id" :movie="movie" />
    </div>
  </div>
</template>
