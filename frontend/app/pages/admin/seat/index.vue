<script setup>
import { ref } from "vue";
import { gql } from "@apollo/client/core";

definePageMeta({
  layout: "admin",
});

const { $apollo } = useNuxtApp();

const seats = ref([]);
const loading = ref(true);

const GET_SEATS = gql`
  query GetAdminSeats {
    seats(order_by: { created_at: desc }) {
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
    query: GET_SEATS,
    fetchPolicy: "network-only",
  });

  seats.value = data.seats;
} catch (err) {
  console.error("Failed loading seats:", err);
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
        Manage Seats
      </h1>

      <p class="text-xs text-gray-500 mt-0.5">Cinema hall seat inventory</p>
    </div>

    <div class="flex gap-10 flex-1 min-h-0 overflow-hidden">
      <div
        class="w-full space-y-6 overflow-y-auto pb-12 [&::-webkit-scrollbar]:hidden [-ms-overflow-style:none] [scrollbar-width:none]"
      >
        <div
          v-if="loading"
          class="text-center text-gray-500 py-12 animate-pulse"
        >
          Loading seat inventory...
        </div>

        <AdminSeatsSeatStatus
          v-else
          v-for="seat in seats"
          :key="seat.id"
          :seat="seat"
        />
      </div>

      <!-- Right -->
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
          class="w-150 h-60 rounded-lg"
          subheading="Cinema"
          title="Add a New seat"
          button-text="Create seat"
          theme="green"
          to="/admin/seat/form"
        />

      </div>
    </div>
  </div>
</template>
