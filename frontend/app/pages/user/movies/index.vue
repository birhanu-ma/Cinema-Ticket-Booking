<script setup>
definePageMeta({
  layout: "user",
});

const movies = [
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
    class="flex h-screen bg-gray-900 bg-gradient-to-t from-[#51751f] to-transparent overflow-hidden"
  >
    <div class="flex-1 overflow-y-auto px-6 pb-4">
      <div
        class="sticky top-0 z-20 bg-gray-900 bg-gradient-to-t from-[#51751f] to-transparent pb-4"
      >
        <SearchBar class="mb-4" />
        <SearchFilter />
      </div>

      <div
        class="grid pt-4 grid-cols-2 gap-4 md:grid-cols-4 lg:grid-cols-6 xl:grid-cols-8"
      >
        <MovieCard
          v-for="(movie, index) in movies"
          :key="index"
          :movie="movie"
        />
      </div>
    </div>
  </div>
</template>
