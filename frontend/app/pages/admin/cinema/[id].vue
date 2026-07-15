<script setup>
const menuItems = [
  { label: "Movies", icon: "🏠" },
  { label: "Directors", icon: "⭐" },
  { label: "Stars", icon: "🎁" },
  { label: "Tickets", icon: "🛍️" },
];

const handleCinemaInsert = async (cleanPayload) => {
  console.log(
    "Ready for standard Hasura GraphQL Insert Mutation:",
    cleanPayload,
  );

  /
};
</script>

<template>
  <div
    class="h-screen w-screen flex flex-row bg-gray-900 overflow-hidden text-white"
  >
    <div class="flex-none">
      <Sidebar :menu-items="menuItems" />
    </div>

    <div
      class="flex-1 overflow-y-auto p-8 flex flex-col items-center justify-center gap-6"
    >
      <AdminCinemaForm @submit-cinema="handleCinemaInsert" />
    </div>
  </div>
</template>
