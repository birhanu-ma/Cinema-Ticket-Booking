<script setup>
import { ref } from "vue";
import { gql } from "@apollo/client/core";

definePageMeta({
  layout: "admin",
});

const { $apollo } = useNuxtApp();

const GET_DIRECTORS = gql`
  query GetDirectors {
    directors(order_by: { name: asc }) {
      id
      name
      bio
      photo_url

      movies_aggregate {
        aggregate {
          count
        }
      }
    }
  }
`;

const directors = ref([]);

try {
  const { data } = await $apollo.query({
    query: GET_DIRECTORS,
    fetchPolicy: "network-only",
  });

  directors.value = data.directors;
} catch (err) {
  console.error(err);
}
</script>

<template>
  <div
    class="flex bg-linear-to-t px-4 from-[#51751f] to-transparent flex-col gap-6 h-[calc(100vh-4rem)] overflow-hidden"
  >
    <div>
      <h1 class="text-2xl font-bold tracking-wide text-gray-100">
        Manage Directors
      </h1>

      <p class="text-xs text-gray-500 mt-0.5">
        Active creative leads catalog logs
      </p>
    </div>

    <div class="flex gap-10 flex-1 min-h-0 overflow-hidden">
      <div
        class="w-full overflow-y-auto pb-12 [&::-webkit-scrollbar]:hidden [-ms-overflow-style:none] [scrollbar-width:none]"
      >
        <div class="grid grid-cols-1 xl:grid-cols-2 gap-6">
          <AdminSharedTeamCard
            v-for="director in directors"
            :key="director.id"
            :person="director"
          />
        </div>
      </div>

      <div class="flex flex-col gap-6 sticky top-0 h-fit flex-shrink-0 pb-12">
        <AdminSharedAddCard
          class="w-150 h-60 rounded-lg"
          subheading="Cinema"
          title="Add Director"
          button-text="Create Director"
          theme="green"
          to="/admin/directors/new"
        />
      </div>
    </div>
  </div>
</template>
