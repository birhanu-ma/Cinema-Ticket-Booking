<script setup>
import { ref } from "vue";
import { gql } from "@apollo/client/core";

definePageMeta({
  layout: "admin",
});

const { $apollo } = useNuxtApp();

const GET_GENRES = gql`
  query GetGenres {
    genres(order_by: { name: asc }) {
      id
      name

      movie_genres_aggregate {
        aggregate {
          count
        }
      }
    }
  }
`;

const genres = ref([]);

try {
  const { data } = await $apollo.query({
    query: GET_GENRES,
  });

  genres.value = data.genres;
} catch (err) {
  console.error(err);
}
</script>

<template>
  <div
    class="flex px-4 bg-linear-to-t from-[#51751f] to-transparent flex-col gap-6 h-[calc(100vh-4rem)] overflow-hidden"
  >
    <div>
      <h1 class="text-2xl font-bold tracking-wide text-gray-100">
        Manage Genres
      </h1>

      <p class="mt-0.5 text-xs text-gray-500">Movie genre catalog logs</p>
    </div>

    <div class="flex gap-10 flex-1 min-h-0 overflow-hidden">
      <div
        class="w-full overflow-y-auto pb-12 [&::-webkit-scrollbar]:hidden [-ms-overflow-style:none] [scrollbar-width:none]"
      >
        <div class="grid grid-cols-1 xl:grid-cols-2 gap-6">
          <div
            v-for="genre in genres"
            :key="genre.id"
            class="flex items-center justify-between gap-4 rounded-lg bg-gray-950/60 border border-gray-800 px-5 py-4"
          >
            <div>
              <h3 class="text-gray-100 font-semibold">{{ genre.name }}</h3>

              <p class="mt-0.5 text-xs text-gray-500">
                {{ genre.movie_genres_aggregate.aggregate.count }} movie(s)
              </p>
            </div>
          </div>
        </div>
      </div>

      <div class="flex flex-col gap-6 sticky top-0 h-fit flex-shrink-0 pb-12">
        <AdminSharedAddCard
          class="w-150 h-60 rounded-lg"
          subheading="genre"
          title="Add a New Genre"
          button-text="Create Genre"
          theme="green"
          to="/admin/genre/new"
        />
      </div>
    </div>
  </div>
</template>
