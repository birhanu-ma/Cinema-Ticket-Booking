<script setup>
import { ref, onMounted } from "vue";
import { gql } from "@apollo/client/core";
import { useRouter } from "vue-router";

definePageMeta({
  layout: "admin",
});

const { $apollo } = useNuxtApp();

const router = useRouter();

const cinemas = ref([]);

const isLoading = ref(false);

const GET_CINEMAS = gql`
  query {
    cinemas(order_by: { name: asc }) {
      id

      name
    }
  }
`;

const INSERT_HALL = gql`
  mutation InsertHall($object: cinema_halls_insert_input!) {
    insert_cinema_halls_one(object: $object) {
      id
    }
  }
`;

onMounted(async () => {
  isLoading.value = true;

  try {
    const { data } = await $apollo.query({
      query: GET_CINEMAS,
    });

    cinemas.value = data.cinemas;
  } finally {
    isLoading.value = false;
  }
});

const handleHallSubmit = async (payload) => {
  await $apollo.mutate({
    mutation: INSERT_HALL,

    variables: {
      object: payload,
    },
  });

  router.push("/admin/cinema-halls");
};
</script>

<template>
  <div
    class="min-h-screen bg-linear-to-t from-[#51751f] to-transparent text-white p-8"
  >
    <NuxtLink to="/admin/cinema-halls" class="text-xs text-lime-400">
      ← Back Halls
    </NuxtLink>

    <h1 class="text-2xl font-bold mt-3">Add Cinema Hall</h1>

    <div v-if="isLoading">Loading cinemas...</div>

    <AdminCinemasCinemaHallForm
      v-else
      :cinemas-list="cinemas"
      @submit-hall="handleHallSubmit"
    />
  </div>
</template>
