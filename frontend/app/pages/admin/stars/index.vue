<script setup>
import { ref } from "vue";
import { gql } from "@apollo/client/core";

definePageMeta({
  layout: "admin",
});

const { $apollo } = useNuxtApp();

const GET_STARS = gql`
  query GetStars {
    stars(order_by: { name: asc }) {
      id
      name
      bio
      photo_url

      movie_stars_aggregate {
        aggregate {
          count
        }
      }
    }
  }
`;

const stars = ref([]);

try {
  const { data } = await $apollo.query({
    query: GET_STARS,
  });

  stars.value = data.stars;
} catch (err) {
  console.error(err);
}
</script>

<template>
  <div
    class="flex px-4 bg-linear-to-t from-[#51751f] to-transparent flex-col gap-6 h-[calc(100vh-4rem)] overflow-hidden"
  >
    <div>
      <h1 class="text-2xl font-bold tracking-wide text-gray-100">
        Manage Stars
      </h1>

      <p class="mt-0.5 text-xs text-gray-500">
        Active cast members catalog logs
      </p>
    </div>

    <div class="flex gap-10 flex-1 min-h-0 overflow-hidden">
      <div
        class="w-full overflow-y-auto pb-12 [&::-webkit-scrollbar]:hidden [-ms-overflow-style:none] [scrollbar-width:none]"
      >
        <div class="grid grid-cols-1 xl:grid-cols-2 gap-6">
          <AdminSharedTeamCard
            v-for="star in stars"
            :key="star.id"
            :person="{
              ...star,
              role: 'Star',
              description: `${star.movie_stars_aggregate.aggregate.count} movie appearance(s).`,
            }"
          />
        </div>
      </div>

      <div class="flex flex-col gap-6 sticky top-0 h-fit flex-shrink-0 pb-12">
   <AdminSharedAddCard
          class="w-150 h-60 rounded-lg"
          subheading="star"
          title="Add a New Star"
          button-text="Create Star"
          theme="green"
          to="/admin/stars/new"
        />      </div>
    </div>
  </div>
</template>
