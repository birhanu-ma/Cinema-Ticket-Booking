<script setup>
import { ref } from "vue";
import { useRoute } from "#app";
import { gql } from "@apollo/client/core";

definePageMeta({
  layout: "admin",
});

const route = useRoute();
const { $apollo } = useNuxtApp();

const hall = ref(null);

const GET_HALL = gql`
  query GetHall($id: uuid!) {
    cinema_halls_by_pk(id: $id) {
      id
      name
      capacity

      cinema {
        id
        name
        location
      }

      seats(order_by: [{ row_name: asc }, { seat_number: asc }]) {
        id
        seat_number
        row_name
        type
      }
    }
  }
`;

try {
  const { data } = await $apollo.query({
    query: GET_HALL,
    variables: {
      id: route.params.id,
    },
    fetchPolicy: "network-only",
  });

  hall.value = data.cinema_halls_by_pk;
} catch (error) {
  console.error("Failed loading hall:", error);
}
</script>

<template>
  <div
    class="min-h-screen bg-linear-to-t px-4 from-[#51751f] to-transparent text-white p-8"
  >
    <div class="max-w-5xl mx-auto">
      <NuxtLink
        to="/admin/cinemaHall"
        class="text-lime-400 hover:text-lime-300 text-sm"
      >
        ← Back to Cinema Halls
      </NuxtLink>

      <div
        v-if="hall"
        class="mt-6 bg-gray-950 border border-gray-800 rounded-3xl p-8"
      >
        <div class="flex items-center justify-between">
          <div>
            <h1 class="text-3xl font-bold">
              {{ hall.name }}
            </h1>

            <p class="text-gray-400 mt-2">
              {{ hall.cinema.name }} • {{ hall.cinema.location }}
            </p>
          </div>

          <div
            class="bg-lime-400/10 text-lime-400 px-5 py-3 rounded-2xl text-center"
          >
            <p class="text-xs uppercase">Capacity</p>
            <p class="text-2xl font-bold">{{ hall.capacity }}</p>
          </div>
        </div>

        <div class="mt-10">
          <h2 class="text-lg font-semibold mb-6">
            Seat Layout ({{ hall.seats.length }} Seats)
          </h2>

          <div class="grid grid-cols-10 gap-3">
            <div
              v-for="seat in hall.seats"
              :key="seat.id"
              class="rounded-xl border border-lime-400 bg-lime-400/10 py-3 text-center"
            >
              <p class="font-semibold text-lime-400">
                {{ seat.row_name }}{{ seat.seat_number }}
              </p>

              <p class="text-[10px] uppercase text-gray-400 mt-1">
                {{ seat.type }}
              </p>
            </div>
          </div>
        </div>
      </div>

      <div v-else class="mt-12 text-center text-gray-500 animate-pulse">
        Loading cinema hall...
      </div>
    </div>
  </div>
</template>
