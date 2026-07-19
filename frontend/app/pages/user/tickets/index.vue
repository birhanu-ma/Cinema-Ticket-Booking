<script setup>
import { ref } from "vue";
import { gql } from "@apollo/client/core";

definePageMeta({
  layout: "user",
});

const { $apollo } = useNuxtApp();

const GET_MY_TICKETS = gql`
  query GetMyTickets {
    tickets(
      where: { status: { _eq: "confirmed" } }
      order_by: { created_at: desc }
    ) {
      id
      booking_reference
      status

      user {
        name
      }

      schedule_seat {
        seat {
          row_name
          seat_number
        }
      }

      schedule {
        id
        start_time
        price

        movie {
          id
          title

          movie_images(order_by: { is_featured: desc }, limit: 1) {
            image_url
          }
        }

        hall {
          id
          name

          cinema {
            id
            name
          }
        }
      }

      payments(limit: 1, order_by: { created_at: desc }) {
        id
        transaction_ref
        amount
        currency
      }
    }
  }
`;

const tickets = ref([]);
const loading = ref(true);
const error = ref("");

try {
  const { data } = await $apollo.query({
    query: GET_MY_TICKETS,
    fetchPolicy: "network-only",
  });

  tickets.value = data.tickets;
} catch (err) {
  console.error(err);
  error.value = err.message || "Failed to load tickets";
} finally {
  loading.value = false;
}
</script>

<template>
  <div
    class="min-h-screen bg-gradient-to-t from-[#51751f] to-transparent px-4 py-12"
  >
    <h1 class="text-3xl font-bold text-gray-100 0 mb-8 text-center">
      My Tickets
    </h1>

    <div v-if="loading" class="text-lime-800 text-center pt-12">
      Loading tickets...
    </div>

    <div
      v-else-if="error"
      class="bg-red-100 text-red-700 rounded-lg p-6 max-w-2xl mx-auto"
    >
      {{ error }}
    </div>

    <div
      v-else-if="tickets.length === 0"
      class="text-lime-700 text-center pt-12"
    >
      You don't have any confirmed tickets yet.
    </div>

    <div v-else class="flex flex-col items-center gap-6">
      <BookingMyTicket
        v-for="ticket in tickets"
        :key="ticket.id"
        :ticket="ticket"
      />
    </div>
  </div>
</template>
