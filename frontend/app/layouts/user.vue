<script setup>
import { computed } from "vue";

const { loggedIn } = useAuth();

const menuItems = computed(() => {
  const items = [
    { label: "Movies", icon: "", to: "/user/movies" },
    { label: "Directors", icon: "", to: "/user/directors" },
    { label: "Genre", icon: "", to: "/user/genres" },
  ];

  if (loggedIn.value) {
    items.push(
      { label: "Bookmark", icon: "", to: "/user/bookmarks" },
      { label: "My Ticket", icon: "", to: "/user/tickets" },
    );
  }

  return items;
});
</script>

<template>
  <div class="h-screen flex bg-gray-900 text-white">
    <LayoutSidebar :menu-items="menuItems" />

    <main class="flex-1 overflow-y-auto p-8">
      <slot />
    </main>
  </div>
</template>