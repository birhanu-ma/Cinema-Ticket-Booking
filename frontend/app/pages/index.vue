<script setup>
import { ref, computed, watch } from "vue";
import { gql } from "@apollo/client/core";

definePageMeta({
  layout: "user",
});

const { $apollo } = useNuxtApp();

const GET_MOVIES = gql`
  query GetMovies($where: movies_bool_exp) {
    movies(order_by: { created_at: desc }, where: $where) {
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
      query: GET_MOVIES,
      variables: {
        where: buildWhere(),
      },
      fetchPolicy: "network-only",
    });

    movies.value = data.movies;
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
    class="flex h-screen bg-gradient-to-t from-[#51751f] to-transparent overflow-hidden"
  >
    <div class="flex-1 overflow-y-auto px-6 pb-4">
      <div
        class="sticky top-0 z-20 bg-gray-900 bg-gradient-to-t from-[#51751f] to-transparent pb-4"
      >
        <SearchBar v-model="searchQuery" class="mb-4" />
        <SearchFilter
          @filter-change="handleFilterChange"
          @schedule-change="handleScheduleChange"
        />
      </div>

      <div v-if="loading" class="text-white text-center pt-12">
        Loading movies...
      </div>

      <div v-else-if="error" class="bg-red-900 text-white rounded-lg p-6 mt-4">
        {{ error }}
      </div>

      <div
        v-else-if="filteredMovies.length === 0"
        class="text-gray-400 text-center pt-12"
      >
        No movies match your search.
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
    </div>
  </div>
</template>
