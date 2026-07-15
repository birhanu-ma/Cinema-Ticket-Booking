<script setup>
import { useRoute } from "#app";

defineProps({
  menuItems: Array,
});

const route = useRoute();

const isItemActive = (path) => {
  return (
    route.path === path ||
    route.path.startsWith(path + "/")
  );
};
</script>

<template>
  <nav class="w-64 h-screen bg-gray-950 p-4 flex flex-col">
    <div class="pt-2">
      <img src="~/assets/image.png" alt="Logo" class="mb-6" />
    </div>

    <div class="space-y-2">
      <NuxtLink
        v-for="item in menuItems"
        :key="item.label"
        :to="item.to"
      >
        <LayoutSidebarItem
          :icon="item.icon"
          :label="item.label"
          :badge="item.badge"
          :is-active="isItemActive(item.to)"
        />
      </NuxtLink>
    </div>

    <button
      class="mt-auto border border-red-600 rounded-lg py-2 text-red-500"
    >
      Logout
    </button>
  </nav>
</template>