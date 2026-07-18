<script setup>
import { ref, watch } from "vue";
import { gql } from "@apollo/client/core";

const props = defineProps({
  scheduleId: {
    type: String,
    required: true,
  },
});

const { $apollo } = useNuxtApp();

const isBookmarked = ref(false);
const bookmarkId = ref(null);
const loading = ref(false);

const CHECK_BOOKMARK = gql`
  query CheckBookmark($scheduleId: uuid!) {
    bookmarks(where: { schedule_id: { _eq: $scheduleId } }) {
      id
    }
  }
`;

const ADD_BOOKMARK = gql`
  mutation AddBookmark($scheduleId: uuid!) {
    insert_bookmarks_one(
      object: { schedule_id: $scheduleId }
      on_conflict: {
        constraint: bookmarks_user_id_schedule_id_key
        update_columns: []
      }
    ) {
      id
    }
  }
`;

const REMOVE_BOOKMARK = gql`
  mutation RemoveBookmark($id: uuid!) {
    delete_bookmarks_by_pk(id: $id) {
      id
    }
  }
`;

const checkStatus = async () => {
  try {
    const { data } = await $apollo.query({
      query: CHECK_BOOKMARK,
      variables: { scheduleId: props.scheduleId },
      fetchPolicy: "network-only",
    });

    const existing = data.bookmarks?.[0];

    if (existing) {
      isBookmarked.value = true;
      bookmarkId.value = existing.id;
    } else {
      isBookmarked.value = false;
      bookmarkId.value = null;
    }
  } catch (err) {
    console.error("Failed to check bookmark status:", err);
  }
};

const toggleBookmark = async () => {
  if (loading.value) return;

  loading.value = true;

  try {
    if (isBookmarked.value && bookmarkId.value) {
      await $apollo.mutate({
        mutation: REMOVE_BOOKMARK,
        variables: { id: bookmarkId.value },
      });

      isBookmarked.value = false;
      bookmarkId.value = null;
    } else {
      const { data } = await $apollo.mutate({
        mutation: ADD_BOOKMARK,
        variables: { scheduleId: props.scheduleId },
      });

      isBookmarked.value = true;
      bookmarkId.value = data.insert_bookmarks_one.id;
    }
  } catch (err) {
    console.error("Failed to toggle bookmark:", err);
  } finally {
    loading.value = false;
  }
};

watch(
  () => props.scheduleId,
  () => {
    checkStatus();
  },
  { immediate: true },
);
</script>

<template>
  <button
    @click="toggleBookmark"
    :disabled="loading"
    type="button"
    class="w-10 h-10 flex items-center justify-center rounded-full border transition-all duration-200 cursor-pointer disabled:opacity-50"
    :class="
      isBookmarked
        ? 'bg-lime-400 border-lime-400 text-black'
        : 'bg-gray-900 border-gray-800 text-gray-400 hover:border-lime-400 hover:text-lime-400'
    "
  >
    <span class="text-lg">{{ isBookmarked ? "★" : "☆" }}</span>
  </button>
</template>
