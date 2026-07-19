<script setup>
import { ref, watch } from "vue";
import { gql } from "@apollo/client/core";

const props = defineProps({
  movieId: {
    type: String,
    required: true,
  },

  genres: {
    type: Array,
    default: () => [],
  },
});

const { $apollo } = useNuxtApp();

const movies = ref([]);
const loading = ref(false);
const error = ref("");

const GET_RELATED_MOVIES = gql`
  query GetRelatedMovies($genreIds: [uuid!], $movieId: uuid!) {
    movies(
      where: {
        id: { _neq: $movieId }
        movie_genres: { genre_id: { _in: $genreIds } }
      }
      limit: 30
      order_by: { created_at: desc }
    ) {
      id
      title

      movie_images {
        image_url
        is_featured
      }

      movie_genres {
        genre {
          id
          name
        }
      }

      ratings_aggregate {
        aggregate {
          avg {
            rating
          }
        }
      }
    }
  }
`;

async function loadMovies() {
  if (!props.genres.length) {
    movies.value = [];
    return;
  }

  loading.value = true;
  error.value = "";

  const genreIds = props.genres.map((g) => g.genre.id);

  try {
    const { data } = await $apollo.query({
      query: GET_RELATED_MOVIES,
      variables: {
        genreIds,
        movieId: props.movieId,
      },
      fetchPolicy: "network-only",
    });

    // Dedupe client-side, since joining through movie_genres can return
    // the same movie multiple times when it matches more than one genre
    const seen = new Set();
    const deduped = [];

    for (const movie of data.movies) {
      if (!seen.has(movie.id)) {
        seen.add(movie.id);
        deduped.push(movie);
      }
    }

    movies.value = deduped.slice(0, 9);
  } catch (err) {
    console.error("Failed to load related movies:", err);
    error.value = err.message || "Failed to load related movies";
  } finally {
    loading.value = false;
  }
}

watch(() => props.genres, loadMovies, { immediate: true });
</script>

<template>
  <h1 class="text-2xl font-bold text-white mb-2 tracking-wide">
    Watch Related Schedules
  </h1>
  <div v-if="error" class="text-red-400 text-sm my-4">
    {{ error }}
  </div>

  <div
    v-else-if="movies.length > 0"
    class="grid grid-cols-2 md:grid-cols-4 lg:grid-cols-9 my-10 gap-4"
  >
    <MovieCard v-for="movie in movies" :key="movie.id" :movie="movie" />
  </div>
</template>
