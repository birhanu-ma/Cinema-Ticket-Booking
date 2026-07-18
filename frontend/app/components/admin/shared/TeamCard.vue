<script setup>
import { computed } from "vue";

const props = defineProps({
  person: {
    type: Object,
    required: true,
  },
});

const role = computed(() => {
  if ("movies_aggregate" in props.person) return "Director";
  if ("movie_stars_aggregate" in props.person) return "Star";

  return props.person.role || "Director";
});
</script>

<template>
  <div
    class="flex bg-gray-950 border border-gray-900 rounded-3xl overflow-hidden shadow-2xl max-w-xl w-full group hover:border-gray-800 transition-all duration-300"
  >
    <div
      class="w-36 sm:w-40 relative flex-shrink-0 overflow-hidden bg-gray-900"
    >
      <img
        :src="person.photo_url || person.photo"
        class="w-full h-full object-cover object-top group-hover:scale-105 transition-transform duration-500"
        :alt="`${person.name} profile portrait photo`"
      />

      <div
        class="absolute inset-0 bg-gradient-to-r from-transparent via-transparent to-gray-950/40"
      ></div>
    </div>

    <AdminSharedTeamDetail :person="person" :role="role" />
  </div>
</template>
