<script setup>
defineProps({
  title: {
    type: String,
    default: "coming soon",
  },
  count: {
    type: [Number, String],
    required: false,
  },
  bgColor: {
    type: String,
    default: "bg-amber-500",
  },
  icon: {
    type: String,
    default: "🎫",
  },
  users: {
    type: Array,
    default: () => [],
  },
});
</script>

<template>
  <div
    :class="[
      'relative p-6 rounded-3xl w-72 h-44 flex flex-col justify-between overflow-hidden shadow-lg select-none',
      bgColor,
    ]"
  >
    <div
      class="absolute -top-10 -right-10 w-32 h-32 rounded-full bg-white/10 pointer-events-none"
    ></div>
    <div
      class="absolute -top-4 -right-4 w-24 h-24 rounded-full bg-white/5 pointer-events-none"
    ></div>

    <div class="z-10">
      <div
        class="w-12 h-12 rounded-xl bg-white/20 flex items-center justify-center text-xl shadow-inner backdrop-blur-xs"
      >
        <span>{{ icon }}</span>
      </div>
    </div>

    <div class="flex items-end justify-between z-10">
      <div class="text-white">
        <p class="text-sm font-medium tracking-wide opacity-90 mb-1">
          {{ title }}
        </p>
        <h2 class="text-4xl font-extrabold tracking-tight leading-none">
          {{ count }}
        </h2>
      </div>

      <div v-if="users.length > 0" class="flex -space-x-3 overflow-hidden">
        <img
          v-for="(avatar, i) in users.slice(0, 4)"
          :key="i"
          :src="avatar"
          class="inline-block h-9 w-9 rounded-full ring-2 ring-white/20 object-cover object-center shadow-md transform hover:-translate-y-1 transition-transform duration-200 cursor-pointer"
          alt="User Profile Thumbnail"
        />
      </div>
    </div>
  </div>
</template>
