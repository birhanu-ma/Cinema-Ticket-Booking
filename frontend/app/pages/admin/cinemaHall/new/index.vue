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

  router.push("/admin/cinemaHall");
};
</script>

<template>
  <div
    class="h-screen w-full flex flex-col items-center overflow-y-auto p-8 gap-6 bg-linear-to-t from-[#51751f] to-transparent text-white"
  >
    <div class="w-full max-w-2xl text-left">
      <NuxtLink
        to="/admin/cinema-halls"
        class="text-xs text-lime-400 hover:underline"
      >
        ← Back Halls
      </NuxtLink>
    </div>

    <div v-if="isLoading">Loading cinemas...</div>

    <AdminCinemasCinemaHallForm
      v-else
      :cinemas-list="cinemas"
      @submit-hall="handleHallSubmit"
    />
  </div>
</template>
