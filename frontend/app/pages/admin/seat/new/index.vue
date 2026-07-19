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
    class="h-screen w-full flex flex-col items-center overflow-y-auto p-8 gap-6 bg-linear-to-t from-[#51751f] to-transparent text-white"
  >
    <div class="w-full max-w-2xl text-left">
      <NuxtLink to="/admin/seat" class="text-lime-400 text-xs hover:underline">
        ← Back Seats
      </NuxtLink>
    </div>

    <AdminSeatsSeatForm :halls-list="halls" @submit-seats="handleSeatSubmit" />
  </div>
</template>
