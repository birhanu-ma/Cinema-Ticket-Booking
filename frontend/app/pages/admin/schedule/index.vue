<script setup>


const activeMovies = [
  { id: "10000000-0000-0000-0000-000000000000", title: "Shazam! Fury of the Gods" },
  { id: "20000000-0000-0000-0000-000000000000", title: "Interstellar" }
]

const activeCinemas = [
  { id: "88888888-4444-4444-4444-ffffffffffff", name: "IMAX Theater Room 1" },
  { id: "99999999-4444-4444-4444-ffffffffffff", name: "Dolby Prime VIP Lounge" }
]

const handleScheduleInsert = async (cleanPayload) => {
  console.log("Ready for your Hasura GraphQL insert payload:", cleanPayload)
  

}
</script>

<template>
  <div class="h-screen w-screen flex flex-row bg-gray-900 overflow-hidden text-white">
    <div class="flex-none">
      <Sidebar :menu-items="menuItems" />
    </div>

    <div class="flex-1 overflow-y-auto p-8 flex flex-col items-center justify-center gap-6">
      <AdminScheduleForm 
        :movies-list="activeMovies"
        :cinemas-list="activeCinemas"
        @submit-schedule="handleScheduleInsert"
      />
    </div>
  </div>
</template>