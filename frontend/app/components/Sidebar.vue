<script setup>
import { h } from "vue";

const activeIndex = ref(0);

const menuItems = [
  { label: "Home", icon: "🏠" },
  { label: "Bookmark", icon: "⭐" },
  // { label: "Bonuses", icon: "🎁" },
  // { label: "Food And Drinks", icon: "🛍️" },
  // { label: "Reminder", icon: "🕒", badge: 3 },
];
</script>
<template>
  <nav class="w-64 h-screen bg-gray-950 p-4 flex flex-col">
    <ClientOnly>
      <div class="pt-2">
        <img src="~/assets/image.png" alt="Logo" class="mb-6" />
      </div>

      <div class="space-y-2">
        <SidebarItem
          v-for="(item, index) in menuItems"
          :key="item.label"
          :is-active="activeIndex === index"
          @click="activeIndex = index"
          v-bind="item"
        />
      </div>

      <button
        class="mt-auto border border-red-600 cursor-pointer bg-transparent text-red-500 hover:bg-red-600/10 transition-colors rounded-lg py-2 pl-10 text-left"
      >
        Logout
      </button>
    </ClientOnly>
  </nav>
</template>
