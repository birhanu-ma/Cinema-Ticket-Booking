<script setup>
import { ref, onMounted } from "vue";
import { gql } from "@apollo/client/core";
import { useRouter } from "vue-router";

definePageMeta({
  layout: "admin",
});

const { $apollo } = useNuxtApp();
const router = useRouter();

const halls = ref([]);

const GET_HALLS = gql`
  query {
    cinema_halls {
      id
      name
    }
  }
`;

const INSERT_SEATS = gql`
  mutation InsertSeats($objects: [seats_insert_input!]!) {
    insert_seats(objects: $objects) {
      affected_rows
    }
  }
`;

onMounted(async () => {
  const { data } = await $apollo.query({
    query: GET_HALLS,
  });

  halls.value = data.cinema_halls;
});

const handleSeatSubmit = async (payload) => {
  try {
    console.log(payload);

    const { data } = await $apollo.mutate({
      mutation: INSERT_SEATS,

      variables: {
        objects: payload,
      },
    });

    console.log(data);

    router.push("/admin/seat");
  } catch (err) {
    console.error(err);
  }
};
</script>

<template>
  <div
    class="min-h-screen bg-linear-to-t from-[#51751f] to-transparent text-white p-8"
  >
    <NuxtLink to="/admin/seat" class="text-lime-400 text-xs">
      ← Back Seats
    </NuxtLink>

    <h1 class="text-2xl font-bold mt-3">Create Seats</h1>

    <AdminSeatsSeatForm :halls-list="halls" @submit-seats="handleSeatSubmit" />
  </div>
</template>
