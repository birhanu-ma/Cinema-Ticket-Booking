<script setup>
import { ref } from "vue";
import { gql } from "@apollo/client/core";

definePageMeta({
  layout: "user",
});

const { $apollo } = useNuxtApp();

const GET_BOOKMARKS = gql`
  query GetBookmarks {
    bookmarks(order_by: { created_at: desc }) {
      id
      schedule_id

      schedule {
        id
        start_time
        price

        movie {
          id
          title

          movie_images(order_by: { is_featured: desc }, limit: 1) {
            image_url
          }
        }

        hall {
          name
          cinema {
            name
          }
        }
      }
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

const bookmarks = ref([]);
const loading = ref(true);
const error = ref("");

const fetchBookmarks = async () => {
  loading.value = true;
  error.value = "";

  try {
    const { data } = await $apollo.query({
      query: GET_BOOKMARKS,
      fetchPolicy: "network-only",
    });

    bookmarks.value = data.bookmarks;
  } catch (err) {
    console.error(err);
    error.value = err.message || "Failed to load bookmarks";
  } finally {
    loading.value = false;
  }
};

const removeBookmark = async (id) => {
  try {
    await $apollo.mutate({
      mutation: REMOVE_BOOKMARK,
      variables: { id },
    });

    bookmarks.value = bookmarks.value.filter((b) => b.id !== id);
  } catch (err) {
    console.error("Failed to remove bookmark:", err);
  }
};

await fetchBookmarks();
</script>

<template>
  <div
    class="min-h-screen bg-gray-900 bg-gradient-to-t from-[#51751f] to-transparent px-6 py-10"
  >
    <h1 class="text-3xl font-bold text-white mb-8">My Bookmarks</h1>

    <div v-if="loading" class="text-white text-center pt-12">
      Loading bookmarks...
    </div>

    <div v-else-if="error" class="bg-red-900 text-white rounded-lg p-6">
      {{ error }}
    </div>

    <div
      v-else-if="bookmarks.length === 0"
      class="text-gray-400 text-center pt-12"
    >
      You haven't bookmarked any showtimes yet.
    </div>

    <div v-else class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-5">
      <NuxtLink
        v-for="bookmark in bookmarks"
        :key="bookmark.id"
        :to="`/user/movies/${bookmark.schedule?.movie?.id}`"
        class="group bg-gray-950 border border-gray-900 rounded-2xl overflow-hidden flex hover:border-lime-400 transition-colors duration-200"
      >
        <div class="w-24 shrink-0 bg-black">
          <img
            v-if="bookmark.schedule?.movie?.movie_images?.[0]?.image_url"
            :src="bookmark.schedule.movie.movie_images[0].image_url"
            :alt="bookmark.schedule.movie.title"
            class="w-full h-full object-cover"
          />
        </div>

        <div class="flex-1 p-4 flex flex-col justify-between">
          <div>
            <h3
              class="text-white font-semibold text-sm group-hover:text-lime-400 transition-colors"
            >
              {{ bookmark.schedule?.movie?.title }}
            </h3>

            <p class="text-gray-400 text-xs mt-1">
              {{ bookmark.schedule?.hall?.cinema?.name }} ·
              {{ bookmark.schedule?.hall?.name }}
            </p>

            <p class="text-gray-500 text-xs">
              {{
                new Date(bookmark.schedule?.start_time).toLocaleString([], {
                  day: "numeric",
                  month: "short",
                  hour: "2-digit",
                  minute: "2-digit",
                })
              }}
            </p>
          </div>

          <div class="flex items-center justify-between mt-3">
            <span class="text-lime-400 font-bold text-sm">
              {{ bookmark.schedule?.price }} ETB
            </span>

            <button
              type="button"
              @click.prevent.stop="removeBookmark(bookmark.id)"
              class="text-xs text-gray-500 hover:text-red-400 cursor-pointer"
            >
              Remove
            </button>
          </div>
        </div>
      </NuxtLink>
    </div>
  </div>
</template>
