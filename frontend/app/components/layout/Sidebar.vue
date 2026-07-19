<script setup>
import { computed } from "vue";
import { useRoute } from "#app";

defineProps({
  menuItems: Array,
});

const route = useRoute();

const { logout } = useAuth();

const isLoggedIn = computed(() => {
  if (import.meta.server) return false;

  return !!localStorage.getItem("auth-token");
});

const isItemActive = (path) => {
  return route.path === path || route.path.startsWith(path + "/");
};
</script>

<template>
  <nav class="w-64 h-screen bg-gray-950 p-4 flex flex-col">
    <div class="pt-2">
      <img src="~/assets/image.png" alt="Logo" class="mb-6" />
    </div>

    <div class="space-y-2">
      <NuxtLink v-for="item in menuItems" :key="item.label" :to="item.to">
        <LayoutSidebarItem
          :icon="item.icon"
          :label="item.label"
          :badge="item.badge"
          :is-active="isItemActive(item.to)"
        />
      </NuxtLink>
    </div>

    <button
      v-if="isLoggedIn"
      @click="logout"
      class="mt-auto border cursor-pointer border-red-600 rounded-lg py-2 text-red-500"
    >
      Logout
    </button>

    <NuxtLink v-else to="/auth/signup" class="mt-auto">
      <button
        class="w-full rounded-lg bg-lime-400 cursor-pointer py-2 font-semibold text-black"
      >
        Signin
      </button>
    </NuxtLink>
  </nav>
</template>
