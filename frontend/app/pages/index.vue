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
  <div
    class="h-screen flex bg-linear-to-t from-[#51751f] to-transparent flex-row bg-gray-900 overflow-hidden"
  >
    <div class="flex-none">
      <Sidebar />
    </div>

    <div class="flex-1 overflow-y-auto px-6">
      <div
        class="sticky top-0 rounded-md mb-4 bg-linear-to-t from-[#51751f] to-transparent z-20 bg-gray-900 pb-4"
      >
        <SearchBar class="mb-4" />
        <Filter />
      </div>

      <div class="grid grid-cols-2 md:grid-cols-4 lg:grid-cols-6 gap-4">
        <MovieCard v-for="movie in movies" :key="movie.id" :movie="movie" />
      </div>
    </div>
  </div>
</template>
