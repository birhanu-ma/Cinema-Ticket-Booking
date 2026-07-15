<script setup>
import { ref, onMounted, computed } from "vue";
import { useRoute, useRouter } from "vue-router";

definePageMeta({
  layout: "admin",
});

const route = useRoute();
const router = useRouter();
const starId = route.params.id;

const isEditMode = computed(() => starId !== "new");
const isLoading = ref(false);
const targetStarData = ref(null);

onMounted(async () => {
  if (isEditMode.value) {
    isLoading.value = true;
    try {
      await new Promise((resolve) => setTimeout(resolve, 400));

      targetStarData.value = {
        name: "Zachary Levi",
        bio: "An American actor best known for starring as Chuck Bartowski in the series Chuck and as the title character in Shazam!.",
        photo_url:
          "https://images.unsplash.com/photo-1560250097-0b93528c311a?w=400",
      };
    } catch (err) {
      console.error("Failed downloading star record from database:", err);
    } finally {
      isLoading.value = false;
    }
  }
});

const handleStarSubmit = async (cleanPayload) => {
  if (isEditMode.value) {
    console.log(
      `Executing Hasura UPDATE Mutation for Star ID (${starId}):`,
      cleanPayload,
    );
  } else {
    console.log("Executing Hasura INSERT Mutation for New Star:", cleanPayload);
  }

  router.push("/admin/stars");
};
</script>

<template>
  <div
    class="h-full w-full flex flex-col items-center justify-center p-8 gap-6 bg-linear-to-t from-[#51751f] to-transparent text-white"
  >
    <div class="w-full max-w-lg flex flex-col gap-1 text-left">
      <NuxtLink
        to="/admin/stars"
        class="text-xs text-lime-400 hover:underline w-fit"
      >
        ← Back to Catalog Logs
      </NuxtLink>
      <h1 class="text-2xl font-bold tracking-wide mt-2">
        {{ isEditMode ? "Edit Star Details" : "Add New Star" }}
      </h1>
    </div>

    <div v-if="isLoading" class="text-sm text-gray-500 animate-pulse py-12">
      Syncing star catalog profiles...
    </div>

    <AdminStarsStarForm
      v-else
      :initial-data="targetStarData"
      @submit-star="handleStarSubmit"
    />
  </div>
</template>
