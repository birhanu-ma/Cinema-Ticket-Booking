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

const GET_RELATED_MOVIES = gql`
  query GetRelatedMovies($genreIds: [uuid!], $movieId: uuid!) {
    movies(
      where: {
        id: { _neq: $movieId }
        movie_genres: {
          genre_id: { _in: $genreIds }
        }
      }
      limit: 9
      distinct_on: id
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
  if (!props.genres.length) return;

  const genreIds = props.genres.map(g => g.genre.id);

  const { data } = await $apollo.query({
    query: GET_RELATED_MOVIES,
    variables: {
      genreIds,
      movieId: props.movieId,
    },
    fetchPolicy: "network-only",
  });

  movies.value = data.movies;
}

watch(
  () => props.genres,
  loadMovies,
  { immediate: true }
);
</script>

<template>
  <div
    class="grid grid-cols-2 md:grid-cols-4 lg:grid-cols-9 my-10 gap-4"
  >
    <MovieCard
      v-for="movie in movies"
      :key="movie.id"
      :movie="movie"
    />
  </div>
</template>