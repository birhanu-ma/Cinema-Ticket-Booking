<script setup>
import { ref, computed, watch } from "vue";
import { gql } from "@apollo/client/core";

const props = defineProps({
  movieId: {
    type: String,
    required: true,
  },
});

const { $apollo } = useNuxtApp();

const myRating = ref(null);
const hoverStar = ref(0);
const selectedStar = ref(0);
const review = ref("");

const loading = ref(false);
const submitting = ref(false);
const errorMsg = ref("");

const GET_MY_RATING = gql`
  query GetMyRating($movieId: uuid!) {
    ratings(where: { movie_id: { _eq: $movieId } }, limit: 1) {
      id
      rating
      review
      user_id
    }
  }
`;

const SUBMIT_RATING = gql`
  mutation SubmitRating($movieId: uuid!, $rating: Int!, $review: String) {
    insert_ratings_one(
      object: { movie_id: $movieId, rating: $rating, review: $review }
    ) {
      id
      rating
      review
    }
  }
`;

const checkMyRating = async () => {
  loading.value = true;

  try {
    const { data } = await $apollo.query({
      query: GET_MY_RATING,
      variables: { movieId: props.movieId },
      fetchPolicy: "network-only",
    });

    const own = data.ratings?.find((r) => r.user_id) ?? data.ratings?.[0];

    myRating.value = own ?? null;
  } catch (err) {
    console.error("Failed to check rating:", err);
  } finally {
    loading.value = false;
  }
};

const submitRating = async () => {
  errorMsg.value = "";

  if (selectedStar.value === 0) {
    errorMsg.value = "Please select a star rating";
    return;
  }

  submitting.value = true;

  try {
    const { data } = await $apollo.mutate({
      mutation: SUBMIT_RATING,
      variables: {
        movieId: props.movieId,
        rating: selectedStar.value,
        review: review.value.trim() || null,
      },
    });

    myRating.value = data.insert_ratings_one;
  } catch (err) {
    console.error("Failed to submit rating:", err);

    const message =
      err?.graphQLErrors?.[0]?.message || err?.message || "";

    if (message.includes("unique") || message.includes("duplicate")) {
      errorMsg.value = "You've already rated this movie";
      checkMyRating();
    } else {
      errorMsg.value = "Failed to submit rating, please try again";
    }
  } finally {
    submitting.value = false;
  }
};

watch(
  () => props.movieId,
  () => {
    checkMyRating();
  },
  { immediate: true },
);
</script>

<template>
  <div class="bg-gray-950 border border-gray-900 rounded-2xl p-5">
    <h3 class="text-white font-bold mb-4">Your Rating</h3>

    <div v-if="loading" class="text-gray-500 text-sm">Loading...</div>

    <div v-else-if="myRating">
      <div class="flex gap-1 mb-2">
        <span
          v-for="n in 5"
          :key="n"
          class="text-2xl"
          :class="n <= myRating.rating ? 'text-lime-400' : 'text-gray-700'"
        >
          ★
        </span>
      </div>

      <p v-if="myRating.review" class="text-gray-300 text-sm mt-2">
        {{ myRating.review }}
      </p>

      <p class="text-gray-500 text-xs mt-3">
        You've already rated this movie
      </p>
    </div>

    <div v-else>
      <div class="flex gap-1 mb-3">
        <button
          v-for="n in 5"
          :key="n"
          type="button"
          @click="selectedStar = n"
          @mouseenter="hoverStar = n"
          @mouseleave="hoverStar = 0"
          class="text-3xl cursor-pointer transition-colors"
          :class="
            n <= (hoverStar || selectedStar)
              ? 'text-lime-400'
              : 'text-gray-700'
          "
        >
          ★
        </button>
      </div>

      <textarea
        v-model="review"
        rows="3"
        placeholder="Write a review (optional)"
        class="w-full bg-gray-900 border border-gray-800 rounded-xl px-3 py-2 text-sm text-white placeholder:text-gray-600 focus:outline-none focus:border-lime-400 resize-none"
      ></textarea>

      <p v-if="errorMsg" class="text-red-400 text-xs mt-2">
        {{ errorMsg }}
      </p>

      <button
        @click="submitRating"
        :disabled="submitting"
        class="mt-3 w-full bg-lime-400 text-black font-bold py-2.5 rounded-full text-sm hover:bg-lime-500 disabled:opacity-50 cursor-pointer"
      >
        {{ submitting ? "Submitting..." : "Submit Rating" }}
      </button>
    </div>
  </div>
</template>