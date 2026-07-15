<script setup>
import { ref, onMounted, computed } from "vue";
import { useRoute, useRouter } from "vue-router";

definePageMeta({
  layout: "admin",
});

const route = useRoute();
const router = useRouter();
const directorId = route.params.id;

const isEditMode = computed(() => directorId !== "new");
const isLoading = ref(false);
const targetDirectorData = ref(null);

onMounted(async () => {
  if (isEditMode.value) {
    isLoading.value = true;
    try {
      
      await new Promise((resolve) => setTimeout(resolve, 400));

      targetDirectorData.value = {
        name: "Christopher Nolan",
        bio: "Academy Award-winning filmmaker known for nonlinear storytelling and practical effects.",
        photo_url:
          "https://images.unsplash.com/photo-1560250097-0b93528c311a?w=400",
      };
    } catch (err) {
      console.error("Failed downloading filmmaker database record:", err);
    } finally {
      isLoading.value = false;
    }
  }
});

const handleDirectorSubmit = async (cleanPayload) => {
  if (isEditMode.value) {
    console.log(
      `Executing Hasura UPDATE Mutation for Director ID (${directorId}):`,
      cleanPayload,
    );

  } else {
    console.log(
      "Executing Hasura INSERT Mutation for New Director:",
      cleanPayload,
    );
  }

  router.push("/admin/directors");
};
</script>

<template>
  <div
    class="h-full w-full flex flex-col items-center justify-center p-8 gap-6 bg-linear-to-t from-[#51751f] to-transparent text-white"
  >
    <div class="w-full max-w-lg text-left">
      <NuxtLink
        to="/admin/directors"
        class="text-xs text-lime-400 hover:underline"
      >
        ← Back to Directors Directory
      </NuxtLink>
    </div>

    <div v-if="isLoading" class="text-sm text-gray-500 animate-pulse py-12">
      Syncing director profile nodes...
    </div>

    <AdminDirectorsDirectorForm
      v-else
      :initial-data="targetDirectorData"
      @submit-director="handleDirectorSubmit"
    />
  </div>
</template>
