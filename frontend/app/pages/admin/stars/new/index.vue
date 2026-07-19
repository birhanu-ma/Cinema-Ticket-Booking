<script setup>
import { gql } from "@apollo/client/core";
import { useRouter } from "vue-router";

definePageMeta({
  layout: "admin",
});

const { $apollo } = useNuxtApp();
const router = useRouter();

const INSERT_STAR = gql`
  mutation InsertStar($object: stars_insert_input!) {
    insert_stars_one(object: $object) {
      id
      name
    }
  }
`;

const handleStarSubmit = async (cleanPayload) => {
  try {
    await $apollo.mutate({
      mutation: INSERT_STAR,

      variables: {
        object: cleanPayload,
      },
    });

    router.push("/admin/stars");
  } catch (err) {
    console.error("Star creation failed:", err);
  }
};
</script>

<template>
  <div
    class="h-screen w-full flex flex-col items-center overflow-y-auto p-8 gap-6 bg-linear-to-t from-[#51751f] to-transparent text-white"
  >
    <div class="w-full max-w-lg flex flex-col gap-1 text-left">
      <NuxtLink
        to="/admin/stars"
        class="text-xs text-lime-400 hover:underline w-fit"
      >
        ← Back to Catalog Logs
      </NuxtLink>

      <h1 class="text-2xl font-bold tracking-wide mt-2">Add New Star</h1>
    </div>

    <AdminStarsStarForm @submit-star="handleStarSubmit" />
  </div>
</template>
