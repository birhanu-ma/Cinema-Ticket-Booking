<script setup>
const activeIndex = ref(0);

const menuItems = [
  { label: "Home", icon: "🏠" },
  { label: "Favorite", icon: "⭐" },
  { label: "Bonuses", icon: "🎁" },
  { label: "Food And Drinks", icon: "🛍️" },
  { label: "Reminder", icon: "🕒", badge: 3 },
];
</script>

<template>
  <nav class="w-64 h-screen bg-gray-950 p-4 space-y-2">
    <ClientOnly>
      <div class="pt-2">
        <img src="~/assets/image.png" alt="Logo" class="mb-6" />
      </div>

      <SidebarItem
        v-for="(item, index) in menuItems"
        :key="item.label"
        :is-active="activeIndex === index"
        @click="activeIndex = index"
        v-bind="item"
      />
    </ClientOnly>
  </nav>
</template>
