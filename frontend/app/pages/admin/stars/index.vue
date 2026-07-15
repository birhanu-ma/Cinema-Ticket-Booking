<script setup>
definePageMeta({
  layout: "admin",
});

const leadStar = [
  {
    name: "Zachary Levi",
    photo: "https://images.unsplash.com/photo-1500648767791-00dcc994a43e?w=400",
    bio: "American actor famously known for his expressive comedic presence.",
    role: "Star",
    description:
      "Portrays the primary heroic character alter-ego, delivering high-energy emotional depth and executing core action-sequence stunts throughout the film.",
  },
  {
    name: "Zachary Levi",
    photo: "https://images.unsplash.com/photo-1500648767791-00dcc994a43e?w=400",
    bio: "American actor famously known for his expressive comedic presence.",
    role: "Star",
    description:
      "Portrays the primary heroic character alter-ego, delivering high-energy emotional depth and executing core action-sequence stunts throughout the film.",
  },
  {
    name: "Zachary Levi",
    photo: "https://images.unsplash.com/photo-1500648767791-00dcc994a43e?w=400",
    bio: "American actor famously known for his expressive comedic presence.",
    role: "Star",
    description:
      "Portrays the primary heroic character alter-ego, delivering high-energy emotional depth and executing core action-sequence stunts throughout the film.",
  },
  {
    name: "Zachary Levi",
    photo: "https://images.unsplash.com/photo-1500648767791-00dcc994a43e?w=400",
    bio: "American actor famously known for his expressive comedic presence.",
    role: "Star",
    description:
      "Portrays the primary heroic character alter-ego, delivering high-energy emotional depth and executing core action-sequence stunts throughout the film.",
  },
  {
    name: "Zachary Levi",
    photo: "https://images.unsplash.com/photo-1500648767791-00dcc994a43e?w=400",
    bio: "American actor famously known for his expressive comedic presence.",
    role: "Star",
    description:
      "Portrays the primary heroic character alter-ego, delivering high-energy emotional depth and executing core action-sequence stunts throughout the film.",
  },
  {
    name: "Zachary Levi",
    photo: "https://images.unsplash.com/photo-1500648767791-00dcc994a43e?w=400",
    bio: "American actor famously known for his expressive comedic presence.",
    role: "Star",
    description:
      "Portrays the primary heroic character alter-ego, delivering high-energy emotional depth and executing core action-sequence stunts throughout the film.",
  },
  {
    name: "Zachary Levi",
    photo: "https://images.unsplash.com/photo-1500648767791-00dcc994a43e?w=400",
    bio: "American actor famously known for his expressive comedic presence.",
    role: "Star",
    description:
      "Portrays the primary heroic character alter-ego, delivering high-energy emotional depth and executing core action-sequence stunts throughout the film.",
  },
  {
    name: "Zachary Levi",
    photo: "https://images.unsplash.com/photo-1500648767791-00dcc994a43e?w=400",
    bio: "American actor famously known for his expressive comedic presence.",
    role: "Star",
    description:
      "Portrays the primary heroic character alter-ego, delivering high-energy emotional depth and executing core action-sequence stunts throughout the film.",
  },
  {
    name: "Zachary Levi",
    photo: "https://images.unsplash.com/photo-1500648767791-00dcc994a43e?w=400",
    bio: "American actor famously known for his expressive comedic presence.",
    role: "Star",
    description:
      "Portrays the primary heroic character alter-ego, delivering high-energy emotional depth and executing core action-sequence stunts throughout the film.",
  },
  {
    name: "Zachary Levi",
    photo: "https://images.unsplash.com/photo-1500648767791-00dcc994a43e?w=400",
    bio: "American actor famously known for his expressive comedic presence.",
    role: "Star",
    description:
      "Portrays the primary heroic character alter-ego, delivering high-energy emotional depth and executing core action-sequence stunts throughout the film.",
  },
  {
    name: "Zachary Levi",
    photo: "https://images.unsplash.com/photo-1500648767791-00dcc994a43e?w=400",
    bio: "American actor famously known for his expressive comedic presence.",
    role: "Star",
    description:
      "Portrays the primary heroic character alter-ego, delivering high-energy emotional depth and executing core action-sequence stunts throughout the film.",
  },
  {
    name: "Zachary Levi",
    photo: "https://images.unsplash.com/photo-1500648767791-00dcc994a43e?w=400",
    bio: "American actor famously known for his expressive comedic presence.",
    role: "Star",
    description:
      "Portrays the primary heroic character alter-ego, delivering high-energy emotional depth and executing core action-sequence stunts throughout the film.",
  },
  {
    name: "Zachary Levi",
    photo: "https://images.unsplash.com/photo-1500648767791-00dcc994a43e?w=400",
    bio: "American actor famously known for his expressive comedic presence.",
    role: "Star",
    description:
      "Portrays the primary heroic character alter-ego, delivering high-energy emotional depth and executing core action-sequence stunts throughout the film.",
  },
];
</script>

<template>
  <div
    class="flex bg-linear-to-t from-[#51751f] to-transparent flex-col gap-6 h-[calc(100vh-4rem)] overflow-hidden"
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
            v-for="(star, index) in leadStar"
            :key="index"
            :person="star"
          />
        </div>
      </div>

      <div class="flex flex-col gap-6 sticky top-0 h-fit flex-shrink-0 pb-12">
        <AdminStarsAddStar class="w-150 h-60 rounded-lg" />
      </div>
    </div>
  </div>
</template>
