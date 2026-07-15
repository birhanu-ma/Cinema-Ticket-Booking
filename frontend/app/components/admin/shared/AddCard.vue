<script setup>
import { computed } from "vue";
import { useRouter } from "vue-router"; 

const props = defineProps({
  subheading: {
    type: String,
    required: true,
  },
  title: {
    type: String,
    required: true,
  },
  buttonText: {
    type: String,
    required: true,
  },
  to: {
    type: String,
    default: "",
  },
  theme: {
    type: String,
    default: "green", 
    validator: (value) =>
      ["green", "lime", "blue", "purple", "gray"].includes(value),
  },
});

const emit = defineEmits(["create"]);
const router = useRouter();

const handleClick = () => {
  emit("create");
  if (props.to) {
    router.push(props.to);
  }
};

const buttonColorClass = computed(() => {
  const themes = {
    green: "bg-[#2f6f5a] hover:bg-[#245646] text-white shadow-[#2f6f5a]/10",
    lime: "bg-lime-400 hover:bg-lime-300 text-black shadow-lime-400/10",
    blue: "bg-blue-600 hover:bg-blue-500 text-white shadow-blue-600/10",
    purple: "bg-purple-600 hover:bg-purple-500 text-white shadow-purple-600/10",
    gray: "bg-gray-800 hover:bg-gray-700 text-white shadow-gray-800/10",
  };
  return themes[props.theme] || themes.green;
});
</script>

<template>
  <div
    class="relative overflow-hidden w-70 bg-white rounded-lg p-2 shadow-xl shadow-gray-100/50 border border-gray-100/40 select-none flex flex-col justify-between"
  >
    <div
      class="absolute -top-10 left-1/4 w-36 h-36 rounded-full bg-[#f5ebe0] opacity-60 blur-xl pointer-events-none"
    ></div>
    <div
      class="absolute -bottom-16 left-1/4 w-40 h-40 rounded-full bg-[#e8f1ee] opacity-70 blur-xl pointer-events-none"
    ></div>

    <div class="relative z-10 flex flex-col items-start gap-2">
      <span class="text-gray-400 text-xs font-bold uppercase tracking-wider">
        {{ subheading }}
      </span>

      <h3
        class="text-gray-900 text-base font-bold leading-snug tracking-normal"
      >
        {{ title }}
      </h3>
    </div>

    <div class="relative z-10 pt-4">
      <button
        @click="handleClick"
        type="button"
        class="font-semibold text-xs tracking-wide uppercase px-5 py-3 rounded-xl transition-all duration-200 active:scale-95 shadow-md cursor-pointer"
        :class="buttonColorClass"
      >
        {{ buttonText }}
      </button>
    </div>
  </div>
</template>
