<script setup>
import { ref } from "vue";
import { useRoute } from "#app";
import { gql } from "graphql-tag";

const route = useRoute();
const { $apollo } = useNuxtApp();

const movie = ref(null);
const movies = ref([]);
const loading = ref(true);
const error = ref("");

const GET_MOVIE = gql`
  query GetMovie($id: uuid!) {
    movies_by_pk(id: $id) {
      id
      title
      description
      duration
      created_at

      director {
        id
        name
        bio
        photo_url
      }

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

      movie_stars {
        star {
          id
          name
          bio
          photo_url
        }
      }

      schedules {
        id
        price
        start_time

        hall {
          id
          name
          capacity

          cinema {
            id
            name
            location
          }
        }

        schedule_seats {
          id
          seat_id
          is_available
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

    movies(
      where: { id: { _neq: $id } }
      limit: 9
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

const loadMovie = async () => {
  try {
    loading.value = true;
    error.value = "";

    console.log("========== ROUTE ID ==========");
    console.log(route.params.id);

    const { data } = await $apollo.query({
      query: GET_MOVIE,

      variables: {
        id: route.params.id,
      },

      fetchPolicy: "network-only",
    });

    console.log("========== FULL RESPONSE ==========");
    console.log(data);

    console.log("========== MOVIE OBJECT ==========");
    console.log(data.movies_by_pk);

    console.log("========== SCHEDULES ==========");
    console.log(data.movies_by_pk?.schedules);

    console.log("========== SCHEDULE SEATS ==========");
    console.log(
      data.movies_by_pk?.schedules?.map((schedule) => ({
        scheduleId: schedule.id,

        seats: schedule.schedule_seats,
      })),
    );

    movie.value = data.movies_by_pk;

    movies.value = data.movies;

    if (!movie.value) {
      error.value = "Movie not found";
    }
  } catch (err) {
    console.log("========== MOVIE LOAD ERROR ==========");
    console.error(err);

    error.value = err.message;
  } finally {
    loading.value = false;
  }
};

await loadMovie();
</script>

<template>
  <div
    class="p-10 bg-linear-to-t from-[#51751f] to-transparent h-screen overflow-y-auto"
  >
    <div v-if="loading" class="text-white text-xl">Loading movie...</div>

    <div v-else-if="error" class="bg-red-900 text-white rounded-lg p-6">
      <h2 class="font-bold text-xl mb-2">Error loading movie</h2>

      <p>
        {{ error }}
      </p>
    </div>

    <div v-else-if="movie" class="flex flex-col gap-10">
      <div class="flex flex-row gap-8">
        <!--
          flex-1 + min-w-0: this column always fills whatever space is
          left over after the fixed-width sidebar, instead of shrinking
          or growing to fit its own content (description length, related
          movie count, etc.) - that content-driven sizing was why the
          column's width differed between movies.
        -->
        <div class="flex flex-col gap-5 flex-1 min-w-0">
          <MovieWatchThriller :movie="movie" />

          <div class="flex flex-row gap-8">
            <MovieDirector :director="movie.director" />

            <MovieStars :stars="movie.movie_stars" />
          </div>

          <MovieAboutTheMovie :movie="movie" class="min-h-44 max-w-285" />

          <MovieRelatedMovie
            :movie-id="movie.id"
            :genres="movie.movie_genres"
          />
        </div>

        <!--
          w-96 + shrink-0: fixed-width sidebar that never resizes based
          on its own content or the left column's content. Adjust w-96
          to match your actual intended sidebar width - this is a
          starting value, not a measured one.
        -->
        <div class="flex flex-col gap-6 w-96 shrink-0">
          <BookingCinemaSession :movie="movie" />

          <BookingTicketCard :movie="movie" />
          <UiStars :movie-id="movie.id" />
        </div>
      </div>
    </div>
  </div>
</template>
