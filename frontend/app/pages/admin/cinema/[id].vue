<script setup>
import { ref } from "vue";
import { useRoute } from "#app";
import { gql } from "@apollo/client/core";

definePageMeta({
  layout: "admin",
});

const route = useRoute();

const { $apollo } = useNuxtApp();

const cinema = ref(null);
const isLoading = ref(true);

const GET_CINEMA = gql`
  query GetCinema($id: uuid!) {
    cinemas_by_pk(id: $id) {
      id
      name
      location
      created_at

      cinema_halls {
        id
        name
        capacity
      }
    }
  }
`;

try {
  const { data } = await $apollo.query({
    query: GET_CINEMA,

    variables: {
      id: route.params.id,
    },

    fetchPolicy: "network-only",
  });

  cinema.value = data.cinemas_by_pk;
} catch (err) {
  console.error(err);
} finally {
  isLoading.value = false;
}
</script>

<template>
  <div
    class="min-h-screen bg-linear-to-t from-[#51751f] to-transparent text-white p-8 flex justify-center"
  >
    <div class="max-w-4xl w-full">
      <NuxtLink to="/admin/cinemas" class="text-lime-400 text-xs">
        ← Back Cinemas
      </NuxtLink>

      <div
        v-if="cinema"
        class="mt-6 bg-gray-950 border border-gray-800 rounded-3xl p-8 space-y-8"
      >
        <div>
          <h1 class="text-3xl font-bold">
            {{ cinema.name }}
          </h1>

          <p class="text-gray-400">
            {{ cinema.location }}
          </p>
        </div>

        <div class="grid grid-cols-2 gap-5">
          <div class="bg-gray-900 rounded-2xl p-5">
            <p class="text-xs text-gray-500">Halls</p>

            <p class="text-2xl font-bold">
              {{ cinema.cinema_halls.length }}
            </p>
          </div>

          <div class="bg-gray-900 rounded-2xl p-5">
            <p class="text-xs text-gray-500">Total Capacity</p>

            <p class="text-2xl font-bold">
              {{ cinema.cinema_halls.reduce((a, b) => a + b.capacity, 0) }}
            </p>
          </div>
        </div>

        <h2 class="text-gray-400 uppercase text-xs">Cinema Halls</h2>

        <div class="space-y-3">
          <NuxtLink
            v-for="hall in cinema.cinema_halls"
            :key="hall.id"
            :to="`/admin/cinema-halls/${hall.id}`"
            class="block bg-gray-900 p-5 rounded-2xl hover:bg-gray-800"
          >
            <p class="font-bold">
              {{ hall.name }}
            </p>

            <p class="text-gray-400 text-sm">Capacity {{ hall.capacity }}</p>
          </NuxtLink>
        </div>
      </div>
    </div>
  </div>
</template>
