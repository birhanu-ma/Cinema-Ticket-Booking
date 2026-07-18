<script setup>
import { ref } from "vue";
import { gql } from "@apollo/client/core";

definePageMeta({
  layout: "user",
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
const loading = ref(true);
const error = ref("");

try {
  const { data } = await $apollo.query({
    query: GET_GENRES,
    fetchPolicy: "network-only",
  });

  genres.value = data.genres;
} catch (err) {
  console.error(err);
  error.value = err.message || "Failed to load genres";
} finally {
  loading.value = false;
}
</script>

<template>
  <div
    class="min-h-screen bg-gradient-to-t from-[#51751f] to-transparent px-6 py-10"
  >
    <h1 class="text-3xl font-bold text-white mb-8">Browse by Genre</h1>

    <div v-if="loading" class="text-white text-center pt-12">
      Loading genres...
    </div>

    <div v-else-if="error" class="bg-red-900 text-white rounded-lg p-6">
      {{ error }}
    </div>

    <div
      v-else
      class="grid grid-cols-2 gap-4 md:grid-cols-3 lg:grid-cols-4 xl:grid-cols-5"
    >
      <NuxtLink
        v-for="genre in genres"
        :key="genre.id"
        :to="`/user/genres/${genre.id}`"
        class="group bg-gray-950 border border-gray-900 rounded-2xl p-6 flex flex-col items-start gap-2 hover:border-lime-400 hover:bg-gray-900 transition-all duration-200"
      >
        <span
          class="text-xl font-bold text-white group-hover:text-lime-400 transition-colors"
        >
          {{ genre.name }}
        </span>

        <span class="text-xs text-gray-500 uppercase tracking-widest">
          {{ genre.movie_genres_aggregate.aggregate.count }} movies
        </span>
      </NuxtLink>
    </div>
  </div>
</template>
