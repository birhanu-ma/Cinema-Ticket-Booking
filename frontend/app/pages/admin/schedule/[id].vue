<script setup>
import { ref } from "vue";
import { useRoute } from "#app";
import { gql } from "@apollo/client/core";

definePageMeta({
  layout: "admin",
});

const route = useRoute();

const { $apollo } = useNuxtApp();

const schedule = ref(null);
const loading = ref(true);

const GET_SCHEDULE = gql`
  query GetSchedule($id: uuid!) {
    schedules_by_pk(id: $id) {
      id
      price
      start_time
      end_time
      language
      format

      movie {
        id
        title
        duration
      }

      hall {
        id
        name
        capacity

        cinema {
          id
          name
          location
        }
      }

      schedule_seats {
        id
        is_available
        locked_by
        locked_at

        seat {
          id
          seat_number
          row_name
          type
        }
      }

      tickets {
        id
        status
        booking_reference
      }
    }
  }
`;

try {
  const { data } = await $apollo.query({
    query: GET_SCHEDULE,
    variables: {
      id: route.params.id,
    },
    fetchPolicy: "network-only",
  });

  schedule.value = data.schedules_by_pk;
} catch (error) {
  console.error("Loading schedule failed:", error);
} finally {
  loading.value = false;
}

const formatDate = (value) => {
  if (!value) return "-";
  return new Date(value).toLocaleString();
};
</script>

<template>
  <div
    class="min-h-screen bg-linear-to-t from-[#51751f] to-transparent text-white p-8"
  >
    <div class="max-w-5xl mx-auto">
      <NuxtLink to="/admin" class="text-lime-400 text-sm"> ← Back </NuxtLink>

      <div v-if="loading" class="mt-20 text-center text-gray-500">
        Loading...
      </div>

      <div
        v-else-if="schedule"
        class="mt-8 bg-gray-900 rounded-3xl p-8 space-y-8"
      >
        <div class="flex justify-between">
          <div>
            <h1 class="text-3xl font-bold">
              {{ schedule.movie.title }}
            </h1>
            <p class="text-gray-500 text-sm">
              {{ schedule.movie.duration }} minutes
            </p>
          </div>

          <div class="text-right">
            <p class="text-lime-400 text-2xl font-bold">
              {{ schedule.price }} ETB
            </p>
            <p class="text-gray-400 text-sm">
              {{ formatDate(schedule.start_time) }}
            </p>
          </div>
        </div>

        <hr class="border-gray-800" />

        <div class="grid md:grid-cols-2 gap-5">
          <div class="bg-gray-800 rounded-xl p-5">
            <p class="text-gray-500 text-xs">HALL</p>
            <h2 class="font-bold text-xl">
              {{ schedule.hall.name }}
            </h2>
            <p class="text-gray-400">Capacity: {{ schedule.hall.capacity }}</p>
          </div>

          <div class="bg-gray-800 rounded-xl p-5">
            <p class="text-gray-500 text-xs">CINEMA</p>
            <h2 class="font-bold text-xl">
              {{ schedule.hall.cinema.name }}
            </h2>
            <p class="text-gray-400">
              {{ schedule.hall.cinema.location }}
            </p>
          </div>
        </div>

        <div>
          <h2 class="font-bold mb-4">Seats</h2>

          <div
            v-if="schedule.schedule_seats?.length"
            class="grid grid-cols-6 gap-3"
          >
            <div
              v-for="ss in schedule.schedule_seats"
              :key="ss.id"
              :class="
                ss.is_available
                  ? 'bg-lime-500/20 border-lime-400'
                  : 'bg-red-500/20 border-red-400'
              "
              class="border rounded-xl p-3 text-center"
            >
              {{ ss.seat.row_name }}{{ ss.seat.seat_number }}
            </div>
          </div>

          <p v-else class="text-gray-500">No seat assigned</p>
        </div>

        <div v-if="schedule.tickets?.length">
          <h2 class="font-bold mb-4">Tickets</h2>
          <div class="space-y-2">
            <div
              v-for="ticket in schedule.tickets"
              :key="ticket.id"
              class="bg-gray-800 rounded-xl p-3 flex justify-between"
            >
              <span>{{ ticket.booking_reference }}</span>
              <span class="text-gray-400">{{ ticket.status }}</span>
            </div>
          </div>
        </div>
      </div>

      <div v-else class="text-center text-gray-500">Schedule not found</div>
    </div>
  </div>
</template>
