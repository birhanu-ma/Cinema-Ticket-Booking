<script setup>
import { ref } from "vue";
import { useRoute } from "#app";
import { gql } from "@apollo/client/core";

definePageMeta({
  layout: "admin",
});

const route = useRoute();
const { $apollo } = useNuxtApp();

const GET_DIRECTOR = gql`
  query GetDirector($id: uuid!) {
    directors_by_pk(id: $id) {
      id
      name
      bio
      photo_url

      movies_aggregate {
        aggregate {
          count
        }
      }

      movies(limit: 5) {
        id
      }

      created_at
    }
  }
`;

const isLoading = ref(true);
const director = ref(null);

try {
  const { data } = await $apollo.query({
    query: GET_DIRECTOR,
    variables: {
      id: route.params.id,
    },
  });

  director.value = data.directors_by_pk;
} catch (err) {
  console.error(err);
} finally {
  isLoading.value = false;
}

const formatDate = (date) => {
  if (!date) return "Unknown";

  return new Date(date).toLocaleDateString("en-US", {
    month: "long",
    year: "numeric",
  });
};

const imageUrl = (director) => {
  if (director?.photo_url) return director.photo_url;

  return `https://api.dicebear.com/9.x/adventurer/svg?seed=${encodeURIComponent(
    director?.name || "director",
  )}`;
};
</script>

<template>
  <div
    class="min-h-screen w-full bg-linear-to-t from-[#51751f] to-transparent text-white p-8 flex flex-col items-center"
  >
    <div class="w-full max-w-4xl space-y-6">
      <div class="flex items-center justify-between">
        <NuxtLink
          to="/admin/directors"
          class="text-xs text-lime-400 hover:underline flex items-center gap-1"
        >
          ← Back to Directors Directory
        </NuxtLink>

        <NuxtLink
          v-if="director"
          :to="`/admin/directors/form/${director.id}`"
          class="bg-gray-800 hover:bg-gray-750 text-xs font-semibold px-4 py-2 rounded-xl border border-gray-700 transition-colors cursor-pointer"
        >
          Edit Profile Nodes
        </NuxtLink>
      </div>

      <div
        v-if="isLoading"
        class="w-full text-center py-24 animate-pulse text-gray-500 text-sm"
      >
        Retrieving filmmaker profile schema...
      </div>

      <div
        v-else-if="director"
        class="bg-gray-950 border border-gray-900 rounded-3xl p-6 md:p-8 shadow-2xl flex flex-col md:flex-row gap-8 items-start"
      >
        <div
          class="w-full md:w-56 h-72 flex-shrink-0 bg-gray-900 rounded-2xl overflow-hidden border border-gray-800 shadow-inner relative group"
        >
          <img
            :src="imageUrl(director)"
            class="w-full h-full object-cover group-hover:scale-102 transition-transform duration-500"
            :alt="director.name"
          />
        </div>

        <div class="flex-1 space-y-6">
          <div>
            <div class="flex items-center gap-2 mb-2">
              <span
                class="text-[10px] bg-lime-400/10 text-lime-400 font-bold px-2.5 py-0.5 rounded-md uppercase tracking-wider"
              >
                Director
              </span>

              <span
                class="text-[10px] bg-gray-800 text-gray-400 font-bold px-2.5 py-0.5 rounded-md uppercase tracking-wider"
              >
                ID: {{ director.id }}
              </span>
            </div>

            <h1 class="text-3xl font-extrabold tracking-wide text-gray-100">
              {{ director.name }}
            </h1>
          </div>

          <div class="w-full h-px bg-gray-900"></div>

          <div class="space-y-2">
            <h3
              class="text-xs font-bold text-gray-500 uppercase tracking-wider"
            >
              Biography Overview
            </h3>

            <p class="text-gray-300 text-sm leading-relaxed">
              {{
                director.bio ||
                "No biographic description submitted for this filmmaker."
              }}
            </p>
          </div>

          <div class="w-full h-px bg-gray-900"></div>

          <div class="grid grid-cols-2 gap-4 pt-1">
            <div
              class="bg-gray-900/50 p-4 rounded-2xl border border-gray-900/60"
            >
              <span
                class="text-[10px] text-gray-500 font-bold uppercase tracking-wider block mb-1"
              >
                Cataloged Films
              </span>

              <span class="text-xl font-extrabold text-lime-400">
                {{ director.movies_aggregate.aggregate.count }}
              </span>
            </div>

            <div
              class="bg-gray-900/50 p-4 rounded-2xl border border-gray-900/60"
            >
              <span
                class="text-[10px] text-gray-500 font-bold uppercase tracking-wider block mb-1"
              >
                Database Registered
              </span>

              <span class="text-sm font-semibold text-gray-200">
                {{ formatDate(director.created_at) }}
              </span>
            </div>
          </div>
        </div>
      </div>

      <div v-else class="w-full text-center py-24 text-gray-600 text-sm">
        Director catalog entry not located.
      </div>
    </div>
  </div>
</template>
