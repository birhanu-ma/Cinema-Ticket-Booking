<script setup>
import { ref } from "vue";
import { gql } from "@apollo/client/core";

definePageMeta({
  layout: "admin",
});

const { $apollo } = useNuxtApp();

const schedules = ref([]);
const loading = ref(true);

const GET_SCHEDULES = gql`
  query GetAdminSchedules {
    schedules(order_by: { start_time: desc }) {
      id
      price
      start_time
      language
      format

      movie {
        title
      }

      hall {
        id
        name

        cinema {
          id
          name
        }
      }

      tickets_aggregate {
        aggregate {
          count
        }
      }
    }
  }
`;

try {
  const { data } = await $apollo.query({
    query: GET_SCHEDULES,
    fetchPolicy: "network-only",
  });

  schedules.value = data.schedules;
} catch (err) {
  console.error("Failed loading schedules:", err);
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
        Manage Schedules
      </h1>

      <p class="text-xs text-gray-500 mt-0.5">
        Movie showtimes and availability
      </p>
    </div>

    <div class="flex gap-10 flex-1 min-h-0 overflow-hidden">
      <!-- Left -->
      <div
        class="w-full space-y-6 overflow-y-auto pb-12 [&::-webkit-scrollbar]:hidden [-ms-overflow-style:none] [scrollbar-width:none]"
      >
        <div
          v-if="loading"
          class="text-center text-gray-500 py-12 animate-pulse"
        >
          Loading schedules...
        </div>

        <AdminSchedulesScheduleStatus
          v-else
          v-for="schedule in schedules"
          :key="schedule.id"
          :schedule="schedule"
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
          subheading="Schedule "
          title="Add a New Schedule"
          button-text="Create schedule"
          theme="green"
          to="/admin/schedule/form"
        />
      </div>
    </div>
  </div>
</template>
