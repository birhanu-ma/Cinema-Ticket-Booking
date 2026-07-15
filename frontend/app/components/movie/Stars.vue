<script setup>
const cast = ref([
  { name: "Chiwetel Ejiofor", role: "Clark", image: null }, 
  { name: "Charlize Theron", role: "Andy", image: null },
  {
    name: "KiKi Layne",
    role: "Nile",
    image: "https://i.pravatar.cc/150?u=kiki",
  }, 
  { name: "Matthias Schoenaerts", role: "Booker", image: null },
]);
</script>

<template>
  <div class="flex gap-4 overflow-x-auto pb-4 scrollbar-hide">
    <MovieStarCard v-for="(actor, index) in cast" :key="index" :actor="actor" />
  </div>
</template>
