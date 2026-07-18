<script setup>
import { ref, onMounted } from "vue";
import { gql } from "@apollo/client/core";
import { useRouter } from "vue-router";

definePageMeta({
  layout: "admin",
});

const { $apollo } = useNuxtApp();

const router = useRouter();

const movies = ref([]);

const halls = ref([]);

const loading = ref(true);

const GET_DATA = gql`
  query {
    movies(order_by: { title: asc }) {
      id
      title
    }

    cinema_halls(order_by: { name: asc }) {
      id
      name
    }
  }
`;

const INSERT_SCHEDULE = gql`
  mutation CreateSchedule($object: schedules_insert_input!) {
    insert_schedules_one(object: $object) {
      id
      hall_id
    }
  }
`;

const GET_SEATS = gql`
  query GetSeats($hallId: uuid!) {
    seats(where: { hall_id: { _eq: $hallId } }) {
      id
    }
  }
`;

const INSERT_SCHEDULE_SEATS = gql`
  mutation CreateScheduleSeats($objects: [schedule_seats_insert_input!]!) {
    insert_schedule_seats(objects: $objects) {
      affected_rows
    }
  }
`;

onMounted(async () => {
  try {
    const { data } = await $apollo.query({
      query: GET_DATA,
    });

    movies.value = data.movies;

    halls.value = data.cinema_halls;
  } catch (error) {
    console.error("Loading failed:", error);
  } finally {
    loading.value = false;
  }
});

const handleScheduleSubmit = async (payload) => {
  try {
    console.log("Schedule payload:", payload);


    const scheduleResult = await $apollo.mutate({
      mutation: INSERT_SCHEDULE,

      variables: {
        object: payload,
      },
    });

    const schedule = scheduleResult.data.insert_schedules_one;

    console.log("Created schedule:", schedule);


    const seatsResult = await $apollo.query({
      query: GET_SEATS,

      variables: {
        hallId: schedule.hall_id,
      },

      fetchPolicy: "network-only",
    });

    const seats = seatsResult.data.seats;

    console.log("Hall seats:", seats);


    const scheduleSeats = seats.map((seat) => ({
      schedule_id: schedule.id,

      seat_id: seat.id,

      is_available: true,
    }));

    console.log("Insert schedule seats:", scheduleSeats);

    await $apollo.mutate({
      mutation: INSERT_SCHEDULE_SEATS,

      variables: {
        objects: scheduleSeats,
      },
    });

    console.log("Schedule seats created");

    router.push("/admin/schedules");
  } catch (error) {
    console.error("Creation failed:", error);
  }
};
</script>

<template>
  <div
    class="min-h-screen bg-linear-to-t from-[#51751f] to-transparent text-white p-8"
  >
    <NuxtLink to="/admin/schedules" class="text-lime-400 text-xs">
      ← Back Schedules
    </NuxtLink>

    <h1 class="text-2xl font-bold mt-3">Create Schedule</h1>

    <div v-if="loading" class="mt-10 text-gray-400">Loading...</div>

    <AdminSchedulesScheduleForm
      v-else
      :movies-list="movies"
      :halls-list="halls"
      @submit-schedule="handleScheduleSubmit"
    />
  </div>
</template>
