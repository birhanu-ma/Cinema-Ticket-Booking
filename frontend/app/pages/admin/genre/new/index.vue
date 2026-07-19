<script setup>
import { gql } from "@apollo/client/core";

definePageMeta({
  layout: "admin",
});

const { $apollo } = useNuxtApp();

const INSERT_GENRE = gql`
  mutation InsertGenre($object: genres_insert_input!) {
    insert_genres_one(object: $object) {
      id
      name
    }
  }
`;

const handleGenreInsert = async (cleanPayload) => {
  try {
    await $apollo.mutate({
      mutation: INSERT_GENRE,
      variables: {
        object: cleanPayload,
      },
    });

    await navigateTo("/admin/genre");
  } catch (err) {
    console.error(err);

    const message =
      err?.graphQLErrors?.[0]?.message ||
      err?.networkError?.result?.errors?.[0]?.message ||
      err?.message ||
      "Failed to create genre";

    alert(message);
  }
};
</script>

<template>
  <div
    class="h-screen w-screen flex flex-row bg-linear-to-t from-[#51751f] to-transparent overflow-hidden text-white"
  >
    <div
      class="flex-1 overflow-y-auto p-8 flex flex-col items-center justify-center gap-6"
    >
      <AdminGenreForm @submit-genre="handleGenreInsert" />
    </div>
  </div>
</template>
