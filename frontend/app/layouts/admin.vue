<script setup>
const menuItems = [
  { label: "Cinema", icon: "", to: "/admin/cinema" },
  { label: "Cinema Hall", icon: "", to: "/admin/cinemaHall" },
  { label: "Seat", icon: "", to: "/admin/seat" },

  { label: "Movies", icon: "", to: "/admin/movies" },
  { label: "Directors", icon: "", to: "/admin/directors" },
  { label: "Stars", icon: "", to: "/admin/stars" },
  { label: "Genre", icon: "", to: "/admin/genre" },

  { label: "Schedule", icon: "", to: "/admin/schedule" },

  { label: "Tickets", icon: "", to: "/admin/tickets" },
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
