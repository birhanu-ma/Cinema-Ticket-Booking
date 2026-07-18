<script setup>
import { ref } from "vue";
import { gql } from "@apollo/client/core";

definePageMeta({
  layout: "admin",
});

const { $apollo } = useNuxtApp();

const GET_MOVIES = gql`
  query GetAdminMovies {
    movies(order_by: { created_at: desc }) {
      id
      title
      duration

      director {
        id
        name
      }

      movie_genres {
        genre {
          id
          name
        }
      }

      movie_images {
        image_url
        is_featured
      }

      schedules_aggregate {
        aggregate {
          count
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

const movies = ref([]);
const loading = ref(true);

try {
  const { data } = await $apollo.query({
    query: GET_MOVIES,
    fetchPolicy: "network-only",
  });

  movies.value = data.movies;
} catch (err) {
  console.error(err);
} finally {
  loading.value = false;
}
</script>

<template>
  <div
    class="flex bg-linear-to-t px-4 from-[#51751f] to-transparent flex-col gap-6 h-[calc(100vh-4rem)] overflow-hidden"
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
          v-for="movie in movies"
          :key="movie.id"
          :movie="movie"
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
        <AdminSharedAddCard
          class="w-150 h-60 rounded-lg"
          subheading="Cinema hall"
          title="Add a New cinema hall"
          button-text="Create cinema hall"
          theme="green"
          to="/admin/cinemahall/form"
        />
      </div>
    </div>
  </div>
</template>
