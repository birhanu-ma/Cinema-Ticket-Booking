<script setup>
import { ref, onMounted, computed } from "vue";
import { useRoute, useRouter } from "vue-router";

definePageMeta({
  layout: "admin",
});

const route = useRoute();
const router = useRouter();
const movieId = route.params.id;

const isEditMode = computed(() => movieId !== "new");
const isLoading = ref(false);
const targetMovieData = ref(null);

const directors = ref([]);
const genres = ref([]);
const stars = ref([]);

onMounted(async () => {
  isLoading.value = true;
  try {
    await new Promise((resolve) => setTimeout(resolve, 400));

    directors.value = [
      { id: "d1111111-1111-1111-1111-111111111111", name: "Christopher Nolan" },
      { id: "d2222222-2222-2222-2222-222222222222", name: "Denis Villeneuve" },
    ];
    genres.value = [
      { id: "g1111111-1111-1111-1111-111111111111", name: "Sci-Fi" },
      { id: "g2222222-2222-2222-2222-222222222222", name: "Thriller" },
    ];
    stars.value = [
      { id: "s1111111-1111-1111-1111-111111111111", name: "Leonardo DiCaprio" },
      { id: "s2222222-2222-2222-2222-222222222222", name: "Timothée Chalamet" },
    ];

    if (isEditMode.value) {
      targetMovieData.value = {
        title: "Inception",
        description:
          "A thief who steals corporate secrets through the use of dream-sharing technology...",
        duration: 148,
        director_id: "d1111111-1111-1111-1111-111111111111",
        movie_genres: [{ genre_id: "g2222222-2222-2222-2222-222222222222" }],
        movie_stars: [{ star_id: "s1111111-1111-1111-1111-111111111111" }],
        movie_images: [
          {
            image_url:
              "https://images.unsplash.com/photo-1536440136628-849c177e76a1",
          },
        ],
      };
    }
  } catch (err) {
    console.error("Error setting up dynamic relational nodes:", err);
  } finally {
    isLoading.value = false;
  }
});

const handleMovieSubmit = async (graphqlPayload) => {
  if (isEditMode.value) {
    console.log(
      `Executing Hasura UPDATE Mutation for Movie ID (${movieId}):`,
      graphqlPayload,
    );
  } else {
    console.log(
      "Executing Hasura INSERT Mutation for New Movie:",
      graphqlPayload,
    );
  }

  router.push("/admin/movies");
};
</script>

<template>
  <div
    class="h-full w-full flex flex-col items-center justify-center p-8 gap-6 bg-linear-to-t from-[#51751f] to-transparent text-white"
  >
    <div class="w-full max-w-2xl flex flex-col gap-1 text-left">
      <NuxtLink
        to="/admin/movies"
        class="text-xs text-lime-400 hover:underline w-fit"
      >
        ← Back to Catalog Logs
      </NuxtLink>
      <h1 class="text-2xl font-bold tracking-wide mt-2">
        {{ isEditMode ? "Edit Movie Details" : "Add New Movie" }}
      </h1>
    </div>

    <div v-if="isLoading" class="text-sm text-gray-500 animate-pulse py-12">
      Linking dynamic relational schemas...
    </div>

    <AdminMoviesMovieForm
      v-else
      :directors-list="directors"
      :genres-list="genres"
      :stars-list="stars"
      :initial-data="targetMovieData"
      @submit-movie="handleMovieSubmit"
    />
  </div>
</template>
