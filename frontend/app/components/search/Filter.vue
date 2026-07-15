<script setup>
import { ref, computed, watch } from "vue";

const activeFilter = ref(null);
const filters = ["Genres", "Rating", "Schedule"];

const emit = defineEmits(["filterChange", "scheduleChange"]);


const availableGenres = [
  "Action",
  "Comedy",
  "Drama",
  "Sci-Fi",
  "Horror",
  "Romance",
];
const selectedGenres = ref([]);

const genresLabel = computed(() => {
  if (selectedGenres.value.length === 0) return "Genres";
  if (selectedGenres.value.length <= 2)
    return `Genres: ${selectedGenres.value.join(", ")}`;
  return `Genres (${selectedGenres.value.length})`;
});

const availableRatings = ["All Ratings", "0", "1", "2", "3", "4"];
const selectedRating = ref("All Ratings");

const ratingLabel = computed(() => {
  return selectedRating.value === "All Ratings"
    ? "Rating"
    : `Rating: ${selectedRating.value}`;
});

const selectRating = (rating) => {
  selectedRating.value = rating;
  activeFilter.value = null; 
};


const activePreset = ref("upcoming");
const dateFrom = ref("");
const dateTo = ref("");

const scheduleLabel = computed(() => {
  if (!dateFrom.value && !dateTo.value) return "Schedule";
  if (activePreset.value && activePreset.value !== "custom") {
    return `Schedule: ${activePreset.value.charAt(0).toUpperCase() + activePreset.value.slice(1)}`;
  }
  return "Schedule: Custom";
});

const applyPreset = (presetType) => {
  activePreset.value = presetType;
  const now = new Date();
  const start = new Date(
    now.getFullYear(),
    now.getMonth(),
    now.getDate(),
    0,
    0,
    0,
  );
  let end = new Date();

  switch (presetType) {
    case "today":
      end = new Date(
        now.getFullYear(),
        now.getMonth(),
        now.getDate(),
        23,
        59,
        59,
      );
      break;
    case "weekend":
      const currentDay = now.getDay();
      const daysToFriday = (5 - currentDay + 7) % 7;
      const friday = new Date(now.setDate(now.getDate() + daysToFriday));
      start.setTime(friday.setHours(0, 0, 0, 0));

      const sunday = new Date(friday.setDate(friday.getDate() + 2));
      end.setTime(sunday.setHours(23, 59, 59, 999));
      break;
    case "week":
      end = new Date(now.setDate(now.getDate() + 7));
      end.setHours(23, 59, 59, 999);
      break;
    case "upcoming":
    default:
      start.setTime(new Date().getTime());
      end = new Date(now.setFullYear(now.getFullYear() + 1));
      break;
  }

  dateFrom.value = new Date(start.getTime() - start.getTimezoneOffset() * 60000)
    .toISOString()
    .slice(0, 16);
  dateTo.value = new Date(end.getTime() - end.getTimezoneOffset() * 60000)
    .toISOString()
    .slice(0, 16);
};


watch(
  [selectedGenres, selectedRating],
  () => {
    emit("filterChange", {
      genres: selectedGenres.value,
      rating:
        selectedRating.value === "All Ratings" ? null : selectedRating.value,
    });
  },
  { deep: true },
);


let debounceTimer = null;
watch([dateFrom, dateTo], ([newFrom, newTo]) => {
  clearTimeout(debounceTimer);
  debounceTimer = setTimeout(() => {
    emit("scheduleChange", {
      _gte: newFrom ? new Date(newFrom).toISOString() : null,
      _lte: newTo ? new Date(newTo).toISOString() : null,
    });
  }, 150);
});


applyPreset("upcoming");
</script>

<template>
  <div class="relative w-full">
    <div
      class="flex gap-3 overflow-x-auto pb-4 scrollbar-hide [&::-webkit-scrollbar]:hidden [-ms-overflow-style:none] [scrollbar-width:none]"
    >
      <SearchFilterOption
        v-for="filter in filters"
        :key="filter"
        :label="
          filter === 'Schedule'
            ? scheduleLabel
            : filter === 'Genres'
              ? genresLabel
              : ratingLabel
        "
        :is-active="activeFilter === filter"
        @click="activeFilter = activeFilter === filter ? null : filter"
      />
    </div>

    <div class="relative w-full">
      <div
        v-if="activeFilter === 'Genres'"
        class="absolute left-0 mt-1 z-50 w-56 bg-gray-950 border border-gray-900 rounded-3xl p-4 shadow-2xl text-white"
      >
        <div
          class="text-[10px] text-gray-500 uppercase font-bold tracking-widest mb-3 border-b border-gray-900 pb-2"
        >
          Filter Genres
        </div>
        <div class="flex flex-col gap-2.5 max-h-52 overflow-y-auto pr-1">
          <label
            v-for="genre in availableGenres"
            :key="genre"
            class="flex items-center gap-3 text-sm font-medium text-gray-300 hover:text-white cursor-pointer select-none"
          >
            <input
              type="checkbox"
              :value="genre"
              v-model="selectedGenres"
              class="w-4 h-4 rounded border-gray-800 bg-gray-900 text-lime-400 focus:ring-0 checked:bg-lime-400 accent-lime-400 cursor-pointer"
            />
            <span>{{ genre }}</span>
          </label>
        </div>
      </div>

      <div
        v-if="activeFilter === 'Rating'"
        class="absolute left-24 mt-1 z-50 w-52 bg-gray-950 border border-gray-900 rounded-3xl p-2 shadow-2xl text-white"
      >
        <div
          class="text-[10px] text-gray-500 uppercase font-bold tracking-widest px-3 pt-2 pb-1.5 border-b border-gray-900 mb-1"
        >
          Select Rating
        </div>
        <div class="flex flex-col gap-0.5">
          <button
            v-for="rating in availableRatings"
            :key="rating"
            type="button"
            @click="selectRating(rating)"
            class="w-full text-left px-3 py-2 text-sm font-medium rounded-xl transition-colors cursor-pointer"
            :class="
              selectedRating === rating
                ? 'bg-lime-400 text-black font-semibold'
                : 'text-gray-300 hover:bg-gray-900 hover:text-white'
            "
          >
            {{ rating }}
          </button>
        </div>
      </div>

      <div
        v-if="activeFilter === 'Schedule'"
        class="absolute left-48 mt-1 z-50 w-[440px] bg-gray-950 border border-gray-900 rounded-3xl p-5 shadow-2xl text-white flex flex-col gap-4"
      >
        <div class="flex flex-col gap-1.5">
          <span
            class="text-[10px] font-bold text-gray-500 uppercase tracking-widest"
            >Presets</span
          >
          <div class="flex flex-wrap gap-1.5">
            <button
              v-for="preset in ['upcoming', 'today', 'weekend', 'week']"
              :key="preset"
              type="button"
              @click="applyPreset(preset)"
              :class="[
                'px-3 py-1.5 rounded-full text-xs font-bold uppercase tracking-wider cursor-pointer transition-colors',
                activePreset === preset
                  ? 'bg-lime-400 text-black'
                  : 'bg-gray-900 text-gray-400 hover:text-white',
              ]"
            >
              {{ preset }}
            </button>
          </div>
        </div>

        <div class="flex items-center gap-3 border-t border-gray-900 pt-3">
          <div class="flex flex-col gap-1 w-full">
            <span
              class="text-[10px] font-bold text-gray-500 uppercase tracking-widest"
              >From</span
            >
            <input
              v-model="dateFrom"
              type="datetime-local"
              @input="activePreset = 'custom'"
              class="w-full bg-gray-900 border border-gray-800 rounded-xl px-3 py-2 text-xs text-white scheme-dark focus:outline-none focus:border-lime-400"
            />
          </div>
          <span class="text-gray-700 text-xs pt-4 select-none">to</span>
          <div class="flex flex-col gap-1 w-full">
            <span
              class="text-[10px] font-bold text-gray-500 uppercase tracking-widest"
              >To</span
            >
            <input
              v-model="dateTo"
              type="datetime-local"
              @input="activePreset = 'custom'"
              class="w-full bg-gray-900 border border-gray-800 rounded-xl px-3 py-2 text-xs text-white scheme-dark focus:outline-none focus:border-lime-400"
            />
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
