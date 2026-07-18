<script setup>

const handleGenreInsert = async (cleanPayload) => {
  console.log("Payload data packages ready for Hasura background transaction operations:", cleanPayload)
  
}
</script>

<template>
  <div class="h-screen w-screen flex flex-row bg-linear-to-t from-[#51751f] to-transparent overflow-hidden text-white">
    <div class="flex-none">
      <Sidebar :menu-items="menuItems" />
    </div>

    <div class="flex-1 overflow-y-auto p-8 flex flex-col items-center justify-center gap-6">
      <AdminGenreForm @submit-genre="handleGenreInsert" />
    </div>
  </div>
</template>