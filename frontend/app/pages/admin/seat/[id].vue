<script setup>
import { ref } from "vue";
import { useRoute } from "#app";
import { gql } from "@apollo/client/core";

definePageMeta({
  layout: "admin",
});

const route = useRoute();
const { $apollo } = useNuxtApp();

const seat = ref(null);
const isLoading = ref(true);

const GET_SEAT = gql`
  query GetSeat($id: uuid!) {
    seats_by_pk(id: $id) {
      id
      seat_number
      row_name
      type
      created_at

      cinema_hall {
        id
        name
        capacity

        cinema {
          id
          name
          location
        }
      }
    }
  }
`;

try {
  const { data } = await $apollo.query({
    query: GET_SEAT,
    variables: {
      id: route.params.id,
    },
    fetchPolicy: "network-only",
  });

  seat.value = data.seats_by_pk;
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
          to="/admin/seat"
          class="text-xs text-lime-400 hover:underline"
        >
          ← Back to Seats
        </NuxtLink>

        <NuxtLink
          v-if="seat"
          :to="`/admin/seat/form/${seat.id}`"
          class="bg-gray-800 hover:bg-gray-700 text-xs px-4 py-2 rounded-xl border border-gray-700"
        >
          Edit Seat
        </NuxtLink>
      </div>

      <div
        v-if="isLoading"
        class="text-center py-20 animate-pulse text-gray-500"
      >
        Loading seat...
      </div>

      <div
        v-else-if="seat"
        class="bg-gray-950 border border-gray-900 rounded-3xl p-8 shadow-2xl"
      >
        <div class="flex items-center justify-between">
          <div>
            <span
              class="bg-blue-500/10 text-blue-400 text-xs px-3 py-1 rounded-lg uppercase"
            >
              {{ seat.type }}
            </span>

            <h1 class="text-4xl font-bold mt-4">
              Seat {{ seat.row_name }}{{ seat.seat_number }}
            </h1>
          </div>

          <div
            class="w-24 h-24 rounded-2xl bg-gray-900 border border-gray-800 flex items-center justify-center text-4xl"
          >
            💺
          </div>
        </div>

        <div class="h-px bg-gray-900 my-8"></div>

        <div class="grid grid-cols-2 gap-6">
          <div class="bg-gray-900/50 rounded-2xl p-5">
            <p class="text-xs uppercase text-gray-500 mb-2">Cinema</p>

            <p class="font-semibold text-lg">
              {{ seat.cinema_hall.cinema.name }}
            </p>
          </div>

          <div class="bg-gray-900/50 rounded-2xl p-5">
            <p class="text-xs uppercase text-gray-500 mb-2">Hall</p>

            <p class="font-semibold text-lg">
              {{ seat.cinema_hall.name }}
            </p>
          </div>

          <div class="bg-gray-900/50 rounded-2xl p-5">
            <p class="text-xs uppercase text-gray-500 mb-2">Capacity</p>

            <p class="font-semibold text-lg">
              {{ seat.cinema_hall.capacity }}
            </p>
          </div>

          <div class="bg-gray-900/50 rounded-2xl p-5">
            <p class="text-xs uppercase text-gray-500 mb-2">Location</p>

            <p class="font-semibold text-lg">
              {{ seat.cinema_hall.cinema.location }}
            </p>
          </div>
        </div>

        <div class="h-px bg-gray-900 my-8"></div>

        <div class="grid grid-cols-3 gap-4">
          <div class="bg-gray-900/50 rounded-2xl p-5 text-center">
            <span class="text-xs uppercase text-gray-500 block mb-2">
              Row
            </span>

            <span class="text-2xl font-bold text-lime-400">
              {{ seat.row_name }}
            </span>
          </div>

          <div class="bg-gray-900/50 rounded-2xl p-5 text-center">
            <span class="text-xs uppercase text-gray-500 block mb-2">
              Seat Number
            </span>

            <span class="text-2xl font-bold text-lime-400">
              {{ seat.seat_number }}
            </span>
          </div>

          <div class="bg-gray-900/50 rounded-2xl p-5 text-center">
            <span class="text-xs uppercase text-gray-500 block mb-2">
              Type
            </span>

            <span class="text-2xl font-bold text-lime-400 capitalize">
              {{ seat.type }}
            </span>
          </div>
        </div>
      </div>

      <div v-else class="text-center py-20 text-gray-500">Seat not found.</div>
    </div>
  </div>
</template>
