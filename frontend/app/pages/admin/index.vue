<script setup>
import { ref } from "vue";
import { gql } from "@apollo/client/core";

definePageMeta({
  layout: "admin",
});

const { $apollo } = useNuxtApp();

const GET_TICKETS = gql`
  query GetTickets {
    tickets(order_by: { created_at: desc }) {
      id
      status
      created_at

      user {
        id
        name
        email
      }

      schedule {
        id
        start_time

        movie {
          id
          title

          movie_images(where: { is_featured: { _eq: true } }, limit: 1) {
            image_url
          }
        }
      }

      payments(limit: 1, order_by: { created_at: desc }) {
        id
        status
        created_at
      }
    }
  }
`;

const tickets = ref([]);

try {
  const { data } = await $apollo.query({
    query: GET_TICKETS,
  });

  tickets.value = data.tickets;
  console.log("Tickets:", tickets.value);
} catch (err) {
  console.error("GraphQL Error:", err);
}
</script>
<template>
  <div
    class="flex px-4 bg-linear-to-t from-[#51751f] to-transparent flex-col gap-6 h-[calc(100vh-4rem)] overflow-hidden"
  >
    <div>
      <h1 class="text-2xl font-bold tracking-wide text-gray-100">
        Ticket Overview Logs
      </h1>
    </div>

    <div class="flex gap-10 flex-1 min-h-0 overflow-hidden">
      <div
        class="w-full overflow-y-auto pb-12 [&::-webkit-scrollbar]:hidden [-ms-overflow-style:none] [scrollbar-width:none]"
      >
        <div class="flex flex-col gap-4">
          <AdminTicketsTicketStatus
            v-for="ticket in tickets"
            :key="ticket.id"
            :ticket="ticket"
          />
        </div>
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
       
      </div>
    </div>
  </div>
</template>
