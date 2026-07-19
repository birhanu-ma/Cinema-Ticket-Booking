<script setup>
import { ref } from "vue";
import { gql } from "@apollo/client/core";

definePageMeta({
  layout: "admin",
});

const { $apollo } = useNuxtApp();

const halls = ref([]);
const loading = ref(true);

const GET_HALLS = gql`
  query GetAdminCinemaHalls {
    cinema_halls(order_by: { created_at: desc }) {
      id
      name
      capacity

      cinema {
        id
        name
        location
      }

      seats_aggregate {
        aggregate {
          count
        }
      }
    }
  }
`;

try {
  const { data } = await $apollo.query({
    query: GET_HALLS,
    fetchPolicy: "network-only",
  });

  halls.value = data.cinema_halls;
} catch (err) {
  console.error(err);
} finally {
  loading.value = false;
}
</script>

<template>
  <div
    class="flex bg-linear-to-t px-4 from-[#51751f] to-transparent flex-col gap-6 h-[calc(100vh-4rem)] overflow-hidden"
  >
    <div>
      <h1 class="text-2xl font-bold tracking-wide text-gray-100">
        Manage Cinema Halls
      </h1>

      <p class="text-xs text-gray-500 mt-0.5">Auditoriums and seat layouts</p>
    </div>

    <div class="flex gap-10 flex-1 min-h-0 overflow-hidden">
      <div
        class="w-full space-y-6 overflow-y-auto pb-12 [&::-webkit-scrollbar]:hidden [-ms-overflow-style:none] [scrollbar-width:none]"
      >
        <div
          v-if="loading"
          class="text-center text-gray-500 py-12 animate-pulse"
        >
          Loading cinema halls...
        </div>

        <AdminCinemasCinemaHallStatus
          v-else
          v-for="hall in halls"
          :key="hall.id"
          :hall="hall"
        />
      </div>

      <div class="flex flex-col gap-6 sticky top-0 h-fit flex-shrink-0 pb-12">
        <div class="flex gap-6">
          <AdminTicketsTicketStatCard />

          <AdminTicketsTicketStatCard bg-color="bg-[#96d13c]" />
        </div>

        <div class="flex gap-6">
          <AdminTicketsTicketStatCard bg-color="bg-[#2c6e59]" />

          <AdminTicketsTicketStatCard bg-color="bg-[#adadad]" />
        </div>
        <AdminSharedAddCard
          class="w-150 h-60 rounded-lg mx-auto"
          subheading="Cinema hall"
          title="Add a New cinema hall"
          button-text="Create cinema hall"
          theme="green"
          to="/admin/cinemahall/new"
        />
      </div>
    </div>
  </div>
</template>
