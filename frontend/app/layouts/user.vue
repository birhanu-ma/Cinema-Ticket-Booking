<script setup>
const menuItems = [
  { label: "Movies", icon: "", to: "/user/movies" },
  { label: "Bookmark", icon: "", to: "/user/bookmarks" },
  { label: "Directors", icon: "", to: "/user/directors" },
  { label: "Genre", icon: "", to: "/user/genres" },
  { label: "My Ticket", icon: "", to: "/user/tickets" },
];
</script>

<template>
  <div class="h-screen flex bg-gray-900 text-white">
    <LayoutSidebar :menu-items="menuItems" />

    <main class="flex-1 overflow-y-auto p-8">
      <slot />
    </main>
  </div>
</template>
