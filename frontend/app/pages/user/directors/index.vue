<script setup>
import { ref } from "vue";
import { gql } from "@apollo/client/core";

definePageMeta({
  layout: "user",
});

const { $apollo } = useNuxtApp();

const GET_DIRECTORS = gql`
  query GetDirectors {
    directors(order_by: { name: asc }) {
      id
      name
      photo_url
      movies_aggregate {
        aggregate {
          count
        }
      }
    }
  }
`;

const directors = ref([]);
const loading = ref(true);
const error = ref("");

try {
  const { data } = await $apollo.query({
    query: GET_DIRECTORS,
    fetchPolicy: "network-only",
  });

  directors.value = data.directors;
} catch (err) {
  console.error(err);
  error.value = err.message || "Failed to load directors";
} finally {
  loading.value = false;
}
</script>

<template>
  <div
    class="min-h-screen bg-gray-900 bg-gradient-to-t from-[#51751f] to-transparent px-6 py-10"
  >
    <h1 class="text-3xl font-bold text-white mb-8">Browse by Director</h1>

    <div v-if="loading" class="text-white text-center pt-12">
      Loading directors...
    </div>

    <div v-else-if="error" class="bg-red-900 text-white rounded-lg p-6">
      {{ error }}
    </div>

    <div
      v-else
      class="grid grid-cols-2 gap-4 md:grid-cols-4 lg:grid-cols-6 xl:grid-cols-8"
    >
      <NuxtLink
        v-for="director in directors"
        :key="director.id"
        :to="`/user/directors/${director.id}`"
        class="group bg-gray-950 border border-gray-900 rounded-2xl overflow-hidden hover:border-lime-400 transition-all duration-200"
      >
        <div class="w-full aspect-square bg-gray-900">
          <img
            v-if="director.photo_url"
            :src="director.photo_url"
            :alt="director.name"
            class="w-full h-full object-cover"
          />
          <div
            v-else
            class="w-full h-full flex items-center justify-center text-gray-600 text-3xl font-bold"
          >
            {{ director.name?.charAt(0) }}
          </div>
        </div>

        <div class="p-4">
          <p
            class="text-white font-semibold text-sm truncate group-hover:text-lime-400 transition-colors"
          >
            {{ director.name }}
          </p>

          <p class="text-gray-500 text-xs mt-1">
            {{ director.movies_aggregate.aggregate.count }} movies
          </p>
        </div>
      </NuxtLink>
    </div>
  </div>
</template>