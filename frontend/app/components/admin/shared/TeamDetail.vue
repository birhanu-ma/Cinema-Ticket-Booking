<script setup>
import { computed } from "vue";

const props = defineProps({
  id: {
    type: [String, Number],
    required: true,
  },
  name: {
    type: String,
    required: true,
  },
  bio: {
    type: String,
    default: "No biography available.",
  },
  role: {
    type: String,
    required: true,
    validator: (value) => ["director", "star"].includes(value.toLowerCase()),
  },
  description: {
    type: String,
    required: true,
  },
});

const destinationRoute = computed(() => {
  const normalizedRole = props.role.toLowerCase();
  if (normalizedRole === "director") {
    return `/admin/directors/${props.id}`;
  } else if (normalizedRole === "star") {
    return `/admin/stars/${props.id}`;
  }
  return "#";
});
</script>

<template>
  <NuxtLink
    :to="destinationRoute"
    class="flex-1 p-5 flex flex-col justify-between hover:bg-gray-900/40 rounded-r-3xl transition-all duration-300 group block cursor-pointer"
  >
    <div>
      <div class="flex items-center justify-between mb-2">
        <span
          :class="[
            'text-[10px] font-bold px-2.5 py-0.5 rounded-md uppercase tracking-wider',
            role.toLowerCase() === 'director'
              ? 'bg-purple-500/10 text-purple-400'
              : 'bg-blue-500/10 text-blue-400',
          ]"
        >
          {{ role }}
        </span>

        <button
          @click.stop.prevent="console.log('Context menu triggered')"
          class="text-gray-500 hover:text-white transition-colors text-lg leading-none p-1 cursor-pointer"
        >
          ⋮
        </button>
      </div>

      <h3
        class="text-white text-xl font-bold leading-snug tracking-wide group-hover:text-lime-400 transition-colors"
      >
        {{ name }}
      </h3>

      <p class="text-gray-400 text-xs mt-1 line-clamp-2 italic">"{{ bio }}"</p>
    </div>

    <div class="w-full h-px bg-gray-900 my-3"></div>

    <div class="flex flex-col gap-1">
      <span
        class="text-[10px] text-gray-500 font-bold uppercase tracking-wider"
      >
        {{
          role.toLowerCase() === "director" ? "Direction Insight" : "Cast Notes"
        }}
      </span>

      <p class="text-gray-300 text-xs font-normal leading-relaxed line-clamp-2">
        {{ description }}
      </p>
    </div>
  </NuxtLink>
</template>
