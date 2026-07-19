<script setup>
import { ref, onMounted } from "vue";
import { gql } from "@apollo/client/core";
import { useRouter } from "vue-router";

definePageMeta({
  layout: "admin",
});

const { $apollo } = useNuxtApp();
const router = useRouter();

const isLoading = ref(false);

const directors = ref([]);
const genres = ref([]);
const stars = ref([]);

const GET_MOVIE_FORM_DATA = gql`
  query GetMovieFormData {
    directors(order_by: { name: asc }) {
      id
      name
    }

    genres(order_by: { name: asc }) {
      id
      name
    }

    stars(order_by: { name: asc }) {
      id
      name
    }
  }
`;

const INSERT_MOVIE = gql`
  mutation InsertMovie($object: movies_insert_input!) {
    insert_movies_one(object: $object) {
      id
    }
  }
`;

onMounted(async () => {
  isLoading.value = true;

  try {
    const { data } = await $apollo.query({
      query: GET_MOVIE_FORM_DATA,
    });

    directors.value = data.directors;

    genres.value = data.genres;

    stars.value = data.stars;
  } catch (err) {
    console.error("Failed loading movie relations:", err);
  } finally {
    isLoading.value = false;
  }
});

const handleMovieSubmit = async (payload) => {
  try {
    await $apollo.mutate({
      mutation: INSERT_MOVIE,

      variables: {
        object: payload,
      },
    });

    router.push("/admin/movies");
  } catch (err) {
    console.error("Movie creation failed:", err);
  }
};
</script>

<template>
  <div
  class="h-screen w-full flex flex-col items-center p-8 gap-6 bg-linear-to-t from-[#51751f] to-transparent text-white overflow-y-auto"
>
    <div class="w-full max-w-2xl flex flex-col gap-1 text-left">
      <NuxtLink
        to="/admin/movies"
        class="text-xs text-lime-400 hover:underline w-fit"
      >
        ← Back to Catalog Logs
      </NuxtLink>

      <h1 class="text-2xl font-bold tracking-wide mt-2">Add New Movie</h1>
    </div>

    <div v-if="isLoading" class="text-sm text-gray-500 animate-pulse py-12">
      Loading movie relations...
    </div>

    <AdminMoviesMovieForm
      v-else
      :directors-list="directors"
      :genres-list="genres"
      :stars-list="stars"
      @submit-movie="handleMovieSubmit"
    />
  </div>
</template>
