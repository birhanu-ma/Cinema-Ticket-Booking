<script setup>
import { ref } from "vue";
import { gql } from "@apollo/client/core";

definePageMeta({
  layout: "admin",
});

const { $apollo } = useNuxtApp();

const cinemas = ref([]);
const loading = ref(true);

const GET_CINEMAS = gql`
  query GetAdminCinemas {
    cinemas(order_by: { created_at: desc }) {
      id
      name
      location
    }
  }
`;

try {
  const { data } = await $apollo.query({
    query: GET_CINEMAS,
    fetchPolicy: "network-only",
  });

  cinemas.value = data.cinemas;
} catch (err) {
  console.error("Failed loading cinemas:", err);
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
        Manage Cinemas
      </h1>

      <p class="text-xs text-gray-500 mt-0.5">
        Cinema locations and auditoriums
      </p>
    </div>

    <div class="flex gap-10 flex-1 min-h-0 overflow-hidden">
      <!-- Left Scroll Area -->
      <div
        class="w-full space-y-6 overflow-y-auto pb-12 [&::-webkit-scrollbar]:hidden [-ms-overflow-style:none] [scrollbar-width:none]"
      >
        <div
          v-if="loading"
          class="text-sm text-gray-500 animate-pulse py-12 text-center"
        >
          Loading cinemas...
        </div>

        <AdminCinemasCinemaStatus
          v-else
          v-for="cinema in cinemas"
          :key="cinema.id"
          :cinema="cinema"
        />
      </div>

      <!-- Right Sticky Sidebar -->
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
          title="Add a New Cinema"
          button-text="Create Cinema"
          theme="green"
          to="/admin/cinema/form"
        />
      </div>
    </div>
  </div>
</template>
