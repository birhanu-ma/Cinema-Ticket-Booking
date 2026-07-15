<script setup>
const menuItems = [
  { label: "Home", icon: "", to: "/user/home" },
  { label: "Movies", icon: "", to: "/user/movies" },
  { label: "Bookmark", icon: "", to: "/user/bookmark" },
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
