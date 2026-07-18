<script setup>
import { gql } from "@apollo/client/core";
import { useRouter } from "vue-router";

definePageMeta({
  layout: "admin",
});

const { $apollo } = useNuxtApp();
const router = useRouter();

const INSERT_DIRECTOR = gql`
  mutation InsertDirector($object: directors_insert_input!) {
    insert_directors_one(object: $object) {
      id
      name
    }
  }
`;

const handleDirectorSubmit = async (cleanPayload) => {
  try {
    await $apollo.mutate({
      mutation: INSERT_DIRECTOR,
      variables: {
        object: cleanPayload,
      },
    });

    router.push("/admin/directors");
  } catch (err) {
    console.error("Director creation failed:", err);
  }
};
</script>

<template>
  <div
    class="h-screen w-full flex flex-col items-center overflow-y-auto p-8 gap-6 bg-linear-to-t from-[#51751f] to-transparent text-white"
  >
    <div class="w-full max-w-lg text-left">
      <NuxtLink
        to="/admin/directors"
        class="text-xs text-lime-400 hover:underline"
      >
        ← Back to Directors Directory
      </NuxtLink>
    </div>

    <AdminDirectorsDirectorForm @submit-director="handleDirectorSubmit" />
  </div>
</template>
