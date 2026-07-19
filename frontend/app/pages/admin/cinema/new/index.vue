<script setup>
import { gql } from "@apollo/client/core";
import { useRouter } from "vue-router";

definePageMeta({
  layout: "admin",
});

const { $apollo } = useNuxtApp();

const router = useRouter();

const INSERT_CINEMA = gql`
  mutation InsertCinema($object: cinemas_insert_input!) {
    insert_cinemas_one(object: $object) {
      id
    }
  }
`;

const handleCinemaSubmit = async (payload) => {
  try {
    await $apollo.mutate({
      mutation: INSERT_CINEMA,

      variables: {
        object: payload,
      },
    });

    router.push("/admin/cinema");
  } catch (err) {
    console.error("Cinema creation failed:", err);
  }
};
</script>

<template>
  <div
    class="h-screen w-full flex flex-col items-center p-8 gap-6 bg-linear-to-t from-[#51751f] to-transparent text-white overflow-y-auto"
  >
    <div class="w-full max-w-2xl">
      <NuxtLink to="/admin/cinemas" class="text-xs text-lime-400">
        ← Back Cinemas
      </NuxtLink>

      <h1 class="text-2xl font-bold mt-3">Add New Cinema</h1>
    </div>

    <AdminCinemasCinemaForm @submit-cinema="handleCinemaSubmit" />
  </div>
</template>
