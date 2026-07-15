<script setup>
const movies = [
  {
    id: 3,
    title: "Avatar: The Way of Water",
    genre: "Sci-Fi/Adventure",
    rating: 5,
    posterUrl: "https://i.pravatar.cc/300?u=3",
  },
  {
    id: 1,
    title: "Shazam! Fury of the Gods",
    genre: "Action/Adventure",
    rating: 4,
    posterUrl: "https://i.pravatar.cc/300?u=1",
  },
  {
    id: 2,
    title: "John Wick: Chapter 4",
    genre: "Action/Thriller",
    rating: 5,
    posterUrl: "https://i.pravatar.cc/300?u=2",
  },
  {
    id: 3,
    title: "Avatar: The Way of Water",
    genre: "Sci-Fi/Adventure",
    rating: 5,
    posterUrl: "https://i.pravatar.cc/300?u=3",
  },
  {
    id: 4,
    title: "Black Panther: Wakanda Forever",
    genre: "Action/Fantasy",
    rating: 4,
    posterUrl: "https://i.pravatar.cc/300?u=4",
  },
  {
    id: 5,
    title: "The Super Mario Bros. Movie",
    genre: "Animation/Comedy",
    rating: 4,
    posterUrl: "https://i.pravatar.cc/300?u=5",
  },
  {
    id: 6,
    title: "Guardians of the Galaxy Vol. 3",
    genre: "Action/Sci-Fi",
    rating: 5,
    posterUrl: "https://i.pravatar.cc/300?u=6",
  },
  {
    id: 7,
    title: "Oppenheimer",
    genre: "Drama/History",
    rating: 5,
    posterUrl: "https://i.pravatar.cc/300?u=7",
  },
  {
    id: 3,
    title: "Avatar: The Way of Water",
    genre: "Sci-Fi/Adventure",
    rating: 5,
    posterUrl: "https://i.pravatar.cc/300?u=3",
  },
  {
    id: 1,
    title: "Shazam! Fury of the Gods",
    genre: "Action/Adventure",
    rating: 4,
    posterUrl: "https://i.pravatar.cc/300?u=1",
  },
  {
    id: 2,
    title: "John Wick: Chapter 4",
    genre: "Action/Thriller",
    rating: 5,
    posterUrl: "https://i.pravatar.cc/300?u=2",
  },
  {
    id: 3,
    title: "Avatar: The Way of Water",
    genre: "Sci-Fi/Adventure",
    rating: 5,
    posterUrl: "https://i.pravatar.cc/300?u=3",
  },
  {
    id: 4,
    title: "Black Panther: Wakanda Forever",
    genre: "Action/Fantasy",
    rating: 4,
    posterUrl: "https://i.pravatar.cc/300?u=4",
  },
  {
    id: 5,
    title: "The Super Mario Bros. Movie",
    genre: "Animation/Comedy",
    rating: 4,
    posterUrl: "https://i.pravatar.cc/300?u=5",
  },
  {
    id: 6,
    title: "Guardians of the Galaxy Vol. 3",
    genre: "Action/Sci-Fi",
    rating: 5,
    posterUrl: "https://i.pravatar.cc/300?u=6",
  },
  {
    id: 7,
    title: "Oppenheimer",
    genre: "Drama/History",
    rating: 5,
    posterUrl: "https://i.pravatar.cc/300?u=7",
  },
  {
    id: 3,
    title: "Avatar: The Way of Water",
    genre: "Sci-Fi/Adventure",
    rating: 5,
    posterUrl: "https://i.pravatar.cc/300?u=3",
  },
  {
    id: 1,
    title: "Shazam! Fury of the Gods",
    genre: "Action/Adventure",
    rating: 4,
    posterUrl: "https://i.pravatar.cc/300?u=1",
  },
  {
    id: 2,
    title: "John Wick: Chapter 4",
    genre: "Action/Thriller",
    rating: 5,
    posterUrl: "https://i.pravatar.cc/300?u=2",
  },
  {
    id: 3,
    title: "Avatar: The Way of Water",
    genre: "Sci-Fi/Adventure",
    rating: 5,
    posterUrl: "https://i.pravatar.cc/300?u=3",
  },
  {
    id: 4,
    title: "Black Panther: Wakanda Forever",
    genre: "Action/Fantasy",
    rating: 4,
    posterUrl: "https://i.pravatar.cc/300?u=4",
  },
  {
    id: 5,
    title: "The Super Mario Bros. Movie",
    genre: "Animation/Comedy",
    rating: 4,
    posterUrl: "https://i.pravatar.cc/300?u=5",
  },
  {
    id: 6,
    title: "Guardians of the Galaxy Vol. 3",
    genre: "Action/Sci-Fi",
    rating: 5,
    posterUrl: "https://i.pravatar.cc/300?u=6",
  },
  {
    id: 7,
    title: "Oppenheimer",
    genre: "Drama/History",
    rating: 5,
    posterUrl: "https://i.pravatar.cc/300?u=7",
  },
  {
    id: 3,
    title: "Avatar: The Way of Water",
    genre: "Sci-Fi/Adventure",
    rating: 5,
    posterUrl: "https://i.pravatar.cc/300?u=3",
  },
  {
    id: 1,
    title: "Shazam! Fury of the Gods",
    genre: "Action/Adventure",
    rating: 4,
    posterUrl: "https://i.pravatar.cc/300?u=1",
  },
  {
    id: 2,
    title: "John Wick: Chapter 4",
    genre: "Action/Thriller",
    rating: 5,
    posterUrl: "https://i.pravatar.cc/300?u=2",
  },
  {
    id: 3,
    title: "Avatar: The Way of Water",
    genre: "Sci-Fi/Adventure",
    rating: 5,
    posterUrl: "https://i.pravatar.cc/300?u=3",
  },
  {
    id: 4,
    title: "Black Panther: Wakanda Forever",
    genre: "Action/Fantasy",
    rating: 4,
    posterUrl: "https://i.pravatar.cc/300?u=4",
  },
  {
    id: 5,
    title: "The Super Mario Bros. Movie",
    genre: "Animation/Comedy",
    rating: 4,
    posterUrl: "https://i.pravatar.cc/300?u=5",
  },
  {
    id: 6,
    title: "Guardians of the Galaxy Vol. 3",
    genre: "Action/Sci-Fi",
    rating: 5,
    posterUrl: "https://i.pravatar.cc/300?u=6",
  },
  {
    id: 7,
    title: "Oppenheimer",
    genre: "Drama/History",
    rating: 5,
    posterUrl: "https://i.pravatar.cc/300?u=7",
  },
  {
    id: 8,
    title: "Barbie",
    genre: "Comedy/Fantasy",
    rating: 4,
    posterUrl: "https://i.pravatar.cc/300?u=8",
  },
];
</script>
<template>
  <div
    class="p-10 bg-linear-to-t from-[#51751f] to-transparent h-screen overflow-y-auto"
  >
    <div class="flex flex-row gap-8">
      <div class="flex flex-col gap-4">
        <MovieWatchThriller />

        <div class="flex gap-2 flex-row">
          <MovieDirector />
          <MovieStars />
        </div>
      </div>

      <div class="flex flex-col gap-6">
        <BookingCinemaSession />
        <BookingTicketCard />
      </div>
    </div>

    <MovieAboutTheMovie />

    <div class="grid grid-cols-2 md:grid-cols-4 lg:grid-cols-9 my-10 gap-4">
      <MovieCard v-for="(movie, index) in movies" :key="index" :movie="movie" />
    </div>
  </div>
</template>
