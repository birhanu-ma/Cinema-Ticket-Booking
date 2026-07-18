<script setup>
import { ref, computed, watch } from "vue";
import { useRoute } from "#app";
import { gql } from "@apollo/client/core";

definePageMeta({
  layout: "user",
});

const route = useRoute();
const { $apollo } = useNuxtApp();

const GET_DIRECTOR_MOVIES = gql`
  query GetDirectorMovies($directorId: uuid!, $where: movies_bool_exp!) {
    directors_by_pk(id: $directorId) {
      id
      name
      bio
      photo_url
    }

    movies(
      order_by: { created_at: desc }
      where: { _and: [{ director_id: { _eq: $directorId } }, $where] }
    ) {
      id
      title

      movie_genres {
        genre {
          name
        }
      }

      movie_images {
        image_url
        is_featured
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

const director = ref(null);
const movies = ref([]);
const loading = ref(true);
const error = ref("");

const searchQuery = ref("");

const activeFilters = ref({
  genres: [],
  rating: null,
});

const activeSchedule = ref({
  _gte: null,
  _lte: null,
});

const filteredMovies = computed(() => {
  if (!activeFilters.value.rating) {
    return movies.value;
  }

  const minRating = Number(activeFilters.value.rating);

  return movies.value.filter((movie) => {
    const avg = Number(movie.ratings_aggregate.aggregate.avg.rating ?? 0);
    return avg >= minRating;
  });
});

const buildWhere = () => {
  const conditions = [];

  if (searchQuery.value.trim()) {
    conditions.push({
      title: { _ilike: `%${searchQuery.value.trim()}%` },
    });
  }

  if (activeFilters.value.genres.length > 0) {
    conditions.push({
      movie_genres: {
        genre: { name: { _in: activeFilters.value.genres } },
      },
    });
  }

  if (activeSchedule.value._gte || activeSchedule.value._lte) {
    const startTime = {};

    if (activeSchedule.value._gte) {
      startTime._gte = activeSchedule.value._gte;
    }

    if (activeSchedule.value._lte) {
      startTime._lte = activeSchedule.value._lte;
    }

    conditions.push({
      schedules: { start_time: startTime },
    });
  }

  if (conditions.length === 0) {
    return {};
  }

  return { _and: conditions };
};

const fetchMovies = async () => {
  loading.value = true;
  error.value = "";

  try {
    const { data } = await $apollo.query({
      query: GET_DIRECTOR_MOVIES,
      variables: {
        directorId: route.params.id,
        where: buildWhere(),
      },
      fetchPolicy: "network-only",
    });

    director.value = data.directors_by_pk;
    movies.value = data.movies;

    if (!director.value) {
      error.value = "Director not found";
    }
  } catch (err) {
    console.error(err);
    error.value = err.message || "Failed to load movies";
  } finally {
    loading.value = false;
  }
};

const handleFilterChange = (filters) => {
  activeFilters.value = filters;
};

const handleScheduleChange = (schedule) => {
  activeSchedule.value = schedule;
};

let searchDebounce = null;
watch(searchQuery, () => {
  clearTimeout(searchDebounce);
  searchDebounce = setTimeout(fetchMovies, 300);
});

watch(
  activeFilters,
  () => {
    fetchMovies();
  },
  { deep: true },
);

watch(
  activeSchedule,
  () => {
    fetchMovies();
  },
  { deep: true },
);

await fetchMovies();
</script>

<template>
  <div
    class="min-h-screen bg-gray-900 bg-gradient-to-t from-[#51751f] to-transparent px-6 py-10"
  >
    <NuxtLink
      to="/user/directors"
      class="text-gray-400 hover:text-lime-400 text-sm mb-6 inline-block"
    >
      ← All Directors
    </NuxtLink>

    <div v-if="loading && !director" class="text-white text-center pt-12">
      Loading movies...
    </div>

    <div v-else-if="error" class="bg-red-900 text-white rounded-lg p-6">
      {{ error }}
    </div>

    <template v-else>
      <div class="flex items-center gap-5 mb-6">
        <div
          class="w-20 h-20 rounded-full overflow-hidden bg-gray-900 shrink-0"
        >
          <img
            v-if="director?.photo_url"
            :src="director.photo_url"
            :alt="director.name"
            class="w-full h-full object-cover"
          />
          <div
            v-else
            class="w-full h-full flex items-center justify-center text-gray-600 text-2xl font-bold"
          >
            {{ director?.name?.charAt(0) }}
          </div>
        </div>

        <div>
          <h1 class="text-3xl font-bold text-white">
            {{ director?.name }}
          </h1>
          <p v-if="director?.bio" class="text-gray-400 text-sm mt-1 max-w-xl">
            {{ director.bio }}
          </p>
        </div>
      </div>

      <div class="mb-4">
        <SearchBar v-model="searchQuery" class="mb-4" />
        <SearchFilter
          @filter-change="handleFilterChange"
          @schedule-change="handleScheduleChange"
        />
      </div>

      <div v-if="loading" class="text-white text-center pt-12">
        Loading movies...
      </div>

      <div
        v-else-if="filteredMovies.length === 0"
        class="text-gray-400 text-center pt-12"
      >
        No movies found for this director.
      </div>

      <div
        v-else
        class="grid pt-4 grid-cols-2 gap-4 md:grid-cols-4 lg:grid-cols-6 xl:grid-cols-8"
      >
        <MovieCard
          v-for="movie in filteredMovies"
          :key="movie.id"
          :movie="movie"
        />
      </div>
    </template>
  </div>
</template>
