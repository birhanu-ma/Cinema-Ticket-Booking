<script setup>
import { ref } from "vue";
import { useRoute } from "#app";
import { gql } from "@apollo/client/core";

definePageMeta({
  layout: "admin",
});

const route = useRoute();
const { $apollo } = useNuxtApp();

const GET_MOVIE = gql`
  query GetAdminMovie($id: uuid!) {
    movies_by_pk(id: $id) {
      id
      title
      description
      duration
      created_at

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

const movie = ref(null);
const isLoading = ref(true);

try {
  const { data } = await $apollo.query({
    query: GET_MOVIE,
    variables: {
      id: route.params.id,
    },
    fetchPolicy: "network-only",
  });

  movie.value = data.movies_by_pk;
} catch (err) {
  console.error(err);
} finally {
  isLoading.value = false;
}
</script>

<template>
  <div
    class="min-h-screen w-full bg-linear-to-t from-[#51751f] to-transparent text-white p-8 flex flex-col items-center"
  >
    <div class="w-full max-w-4xl space-y-6">
      <div class="flex items-center justify-between">
        <NuxtLink
          to="/admin/movies"
          class="text-xs text-lime-400 hover:underline flex items-center gap-1"
        >
          ← Back to Movies Catalog
        </NuxtLink>

        <NuxtLink
          v-if="movie"
          :to="`/admin/movies/form/${movie.id}`"
          class="bg-gray-800 hover:bg-gray-750 text-xs font-semibold px-4 py-2 rounded-xl border border-gray-700 transition-colors cursor-pointer"
        >
          Edit Film Metadata
        </NuxtLink>
      </div>

      <div
        v-if="isLoading"
        class="w-full text-center py-24 animate-pulse text-gray-500 text-sm"
      >
        Retrieving database media nodes...
      </div>

      <div
        v-else-if="movie"
        class="bg-gray-950 border border-gray-900 rounded-3xl p-6 md:p-8 shadow-2xl flex flex-col md:flex-row gap-8 items-start"
      >
        <div
          class="w-full md:w-56 h-80 flex-shrink-0 bg-gray-900 rounded-2xl overflow-hidden border border-gray-800 shadow-inner relative group"
        >
          <img
            :src="
              movie.movie_images.find((i) => i.is_featured)?.image_url ||
              movie.movie_images[0]?.image_url ||
              'https://images.unsplash.com/photo-1440404653325-ab127d49abc1?w=400'
            "
            class="w-full h-full object-cover group-hover:scale-102 transition-transform duration-500"
            :alt="movie.title"
          />
        </div>

        <div class="flex-1 space-y-6">
          <div>
            <div class="flex flex-wrap items-center gap-2 mb-2">
              <span
                class="text-[10px] bg-lime-400/10 text-lime-400 font-bold px-2.5 py-0.5 rounded-md uppercase tracking-wider"
              >
                {{ movie.movie_genres.map((g) => g.genre.name).join(", ") }}
              </span>

              <span
                class="text-[10px] bg-purple-500/10 text-purple-400 font-bold px-2.5 py-0.5 rounded-md uppercase tracking-wider"
              >
                Rating:
                {{ movie.ratings_aggregate.aggregate.avg.rating ?? "N/A" }}
              </span>

              <span
                class="text-[10px] bg-gray-800 text-gray-400 font-bold px-2.5 py-0.5 rounded-md uppercase tracking-wider"
              >
                ID: {{ movie.id }}
              </span>
            </div>

            <h1 class="text-3xl font-extrabold tracking-wide text-gray-100">
              {{ movie.title }}
            </h1>
          </div>

          <div class="w-full h-px bg-gray-900"></div>

          <div class="space-y-2">
            <h3
              class="text-xs font-bold text-gray-500 uppercase tracking-wider"
            >
              Synopsis
            </h3>

            <p class="text-gray-300 text-sm leading-relaxed">
              {{ movie.description || "No synopsis submitted for this film." }}
            </p>
          </div>

          <div class="w-full h-px bg-gray-900"></div>

          <div class="grid grid-cols-3 gap-4 pt-1">
            <div
              class="bg-gray-900/50 p-4 rounded-2xl border border-gray-900/60 text-center"
            >
              <span
                class="text-[10px] text-gray-500 font-bold uppercase tracking-wider block mb-1"
              >
                Director
              </span>

              <span class="text-xs font-semibold text-gray-200 block truncate">
                {{ movie.director?.name || "Unknown" }}
              </span>
            </div>

            <div
              class="bg-gray-900/50 p-4 rounded-2xl border border-gray-900/60 text-center"
            >
              <span
                class="text-[10px] text-gray-500 font-bold uppercase tracking-wider block mb-1"
              >
                Duration
              </span>

              <span class="text-xs font-semibold text-lime-400 block">
                {{ movie.duration }} min
              </span>
            </div>

            <div
              class="bg-gray-900/50 p-4 rounded-2xl border border-gray-900/60 text-center"
            >
              <span
                class="text-[10px] text-gray-500 font-bold uppercase tracking-wider block mb-1"
              >
                Release Year
              </span>

              <span class="text-xs font-semibold text-gray-200 block">
                {{ new Date(movie.created_at).getFullYear() }}
              </span>
            </div>
          </div>
        </div>
      </div>

      <div v-else class="w-full text-center py-24 text-gray-600 text-sm">
        Film catalog entry not located.
      </div>
    </div>
  </div>
</template>
