<script setup>
definePageMeta({
  layout: "admin",
});

const movieDirector = [
  {
    name: "Christopher Nolan",
    photo: "https://images.unsplash.com/photo-1560250097-0b93528c311a?w=400",
    bio: "Academy Award-winning filmmaker known for nonlinear storytelling.",
    role: "Director",
    description:
      "He directed this thriller utilizing practical IMAX cameras to heighten environmental tension and maximize structural screen real-estate visibility.",
  },
  {
    name: "Christopher Nolan",
    photo: "https://images.unsplash.com/photo-1560250097-0b93528c311a?w=400",
    bio: "Academy Award-winning filmmaker known for nonlinear storytelling.",
    role: "Director",
    description:
      "He directed this thriller utilizing practical IMAX cameras to heighten environmental tension and maximize structural screen real-estate visibility.",
  },
  {
    name: "Christopher Nolan",
    photo: "https://images.unsplash.com/photo-1560250097-0b93528c311a?w=400",
    bio: "Academy Award-winning filmmaker known for nonlinear storytelling.",
    role: "Director",
    description:
      "He directed this thriller utilizing practical IMAX cameras to heighten environmental tension and maximize structural screen real-estate visibility.",
  },
  {
    name: "Christopher Nolan",
    photo: "https://images.unsplash.com/photo-1560250097-0b93528c311a?w=400",
    bio: "Academy Award-winning filmmaker known for nonlinear storytelling.",
    role: "Director",
    description:
      "He directed this thriller utilizing practical IMAX cameras to heighten environmental tension and maximize structural screen real-estate visibility.",
  },
  {
    name: "Christopher Nolan",
    photo: "https://images.unsplash.com/photo-1560250097-0b93528c311a?w=400",
    bio: "Academy Award-winning filmmaker known for nonlinear storytelling.",
    role: "Director",
    description:
      "He directed this thriller utilizing practical IMAX cameras to heighten environmental tension and maximize structural screen real-estate visibility.",
  },
  {
    name: "Christopher Nolan",
    photo: "https://images.unsplash.com/photo-1560250097-0b93528c311a?w=400",
    bio: "Academy Award-winning filmmaker known for nonlinear storytelling.",
    role: "Director",
    description:
      "He directed this thriller utilizing practical IMAX cameras to heighten environmental tension and maximize structural screen real-estate visibility.",
  },
  {
    name: "Christopher Nolan",
    photo: "https://images.unsplash.com/photo-1560250097-0b93528c311a?w=400",
    bio: "Academy Award-winning filmmaker known for nonlinear storytelling.",
    role: "Director",
    description:
      "He directed this thriller utilizing practical IMAX cameras to heighten environmental tension and maximize structural screen real-estate visibility.",
  },
  {
    name: "Christopher Nolan",
    photo: "https://images.unsplash.com/photo-1560250097-0b93528c311a?w=400",
    bio: "Academy Award-winning filmmaker known for nonlinear storytelling.",
    role: "Director",
    description:
      "He directed this thriller utilizing practical IMAX cameras to heighten environmental tension and maximize structural screen real-estate visibility.",
  },
  {
    name: "Christopher Nolan",
    photo: "https://images.unsplash.com/photo-1560250097-0b93528c311a?w=400",
    bio: "Academy Award-winning filmmaker known for nonlinear storytelling.",
    role: "Director",
    description:
      "He directed this thriller utilizing practical IMAX cameras to heighten environmental tension and maximize structural screen real-estate visibility.",
  },
  {
    name: "Christopher Nolan",
    photo: "https://images.unsplash.com/photo-1560250097-0b93528c311a?w=400",
    bio: "Academy Award-winning filmmaker known for nonlinear storytelling.",
    role: "Director",
    description:
      "He directed this thriller utilizing practical IMAX cameras to heighten environmental tension and maximize structural screen real-estate visibility.",
  },
  {
    name: "Christopher Nolan",
    photo: "https://images.unsplash.com/photo-1560250097-0b93528c311a?w=400",
    bio: "Academy Award-winning filmmaker known for nonlinear storytelling.",
    role: "Director",
    description:
      "He directed this thriller utilizing practical IMAX cameras to heighten environmental tension and maximize structural screen real-estate visibility.",
  },
  {
    name: "Christopher Nolan",
    photo: "https://images.unsplash.com/photo-1560250097-0b93528c311a?w=400",
    bio: "Academy Award-winning filmmaker known for nonlinear storytelling.",
    role: "Director",
    description:
      "He directed this thriller utilizing practical IMAX cameras to heighten environmental tension and maximize structural screen real-estate visibility.",
  },
  {
    name: "Christopher Nolan",
    photo: "https://images.unsplash.com/photo-1560250097-0b93528c311a?w=400",
    bio: "Academy Award-winning filmmaker known for nonlinear storytelling.",
    role: "Director",
    description:
      "He directed this thriller utilizing practical IMAX cameras to heighten environmental tension and maximize structural screen real-estate visibility.",
  },
  {
    name: "Christopher Nolan",
    photo: "https://images.unsplash.com/photo-1560250097-0b93528c311a?w=400",
    bio: "Academy Award-winning filmmaker known for nonlinear storytelling.",
    role: "Director",
    description:
      "He directed this thriller utilizing practical IMAX cameras to heighten environmental tension and maximize structural screen real-estate visibility.",
  },
  {
    name: "Christopher Nolan",
    photo: "https://images.unsplash.com/photo-1560250097-0b93528c311a?w=400",
    bio: "Academy Award-winning filmmaker known for nonlinear storytelling.",
    role: "Director",
    description:
      "He directed this thriller utilizing practical IMAX cameras to heighten environmental tension and maximize structural screen real-estate visibility.",
  },
  {
    name: "Christopher Nolan",
    photo: "https://images.unsplash.com/photo-1560250097-0b93528c311a?w=400",
    bio: "Academy Award-winning filmmaker known for nonlinear storytelling.",
    role: "Director",
    description:
      "He directed this thriller utilizing practical IMAX cameras to heighten environmental tension and maximize structural screen real-estate visibility.",
  },
];
</script>

<template>
  <div
    class="flex bg-linear-to-t from-[#51751f] to-transparent flex-col gap-6 h-[calc(100vh-4rem)] overflow-hidden"
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
            v-for="(director, index) in movieDirector"
            :key="index"
            :person="director"
          />
        </div>
      </div>

      <div class="flex flex-col gap-6 sticky top-0 h-fit flex-shrink-0 pb-12">
        <AdminDirectorsAddDirector class="w-150 h-60 rounded-lg" />
      </div>
    </div>
  </div>
</template>
