<script setup>
import { ref, onMounted, computed } from "vue";
import { useRoute, useRouter } from "vue-router";

definePageMeta({
  layout: "admin",
});

const route = useRoute();
const router = useRouter();
const ticketId = route.params.id;

const isEditMode = computed(() => ticketId !== "new");
const isLoading = ref(false);
const targetTicketData = ref(null);

onMounted(async () => {
  if (isEditMode.value) {
    isLoading.value = true;
    try {
      await new Promise((resolve) => setTimeout(resolve, 400));

      targetTicketData.value = {
        schedule_id: "3fa85f64-5717-4562-b3fc-2c963f66afa6",
        user_id: "77777777-8888-9999-0000-aaaaaaaaaaaa",
        status: "Confirmed",
      };
    } catch (err) {
      console.error("Failed downloading ticket record from database:", err);
    } finally {
      isLoading.value = false;
    }
  }
});

const handleTicketSubmit = async (cleanPayload) => {
  if (isEditMode.value) {
    console.log(
      `Executing Hasura UPDATE Mutation for Ticket ID (${ticketId}):`,
      cleanPayload,
    );
  } else {
    console.log(
      "Executing Hasura INSERT Mutation for New Ticket:",
      cleanPayload,
    );
  }

  router.push("/admin/tickets");
};
</script>

<template>
  <div
    class="h-full w-full flex flex-col items-center justify-center p-8 gap-6 bg-linear-to-t from-[#51751f] to-transparent text-white"
  >
    <div class="w-full max-w-lg flex flex-col gap-1 text-left">
      <NuxtLink
        to="/admin/tickets"
        class="text-xs text-lime-400 hover:underline w-fit"
      >
        ← Back to Catalog Logs
      </NuxtLink>
      <h1 class="text-2xl font-bold tracking-wide mt-2">
        {{ isEditMode ? "Edit Ticket Details" : "Add New Ticket" }}
      </h1>
    </div>

    <div v-if="isLoading" class="text-sm text-gray-500 animate-pulse py-12">
      Syncing ticket database records...
    </div>

    <AdminTicketsTicketForm
      v-else
      :initial-data="targetTicketData"
      @submit-ticket="handleTicketSubmit"
    />
  </div>
</template>
