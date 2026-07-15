<script setup>
import { watch } from "vue";
import { useForm, useField } from "vee-validate";
import { toTypedSchema } from "@vee-validate/zod";
import * as z from "zod";

const props = defineProps({
  initialData: {
    type: Object,
    default: null,
  },
});

const activeSchedules = [
  {
    id: "3fa85f64-5717-4562-b3fc-2c963f66afa6",
    details: "Shazam! - Room 1 (March 20, 2026 @ 18:00)",
  },
];

const activeUsers = [
  {
    id: "11111111-2222-3333-4444-555555555555",
    name: "John Doe (john@example.com)",
  },
  {
    id: "77777777-8888-9999-0000-aaaaaaaaaaaa",
    name: "Jane Smith (jane@example.com)",
  },
];

const statusOptions = ["Inprogress", "Confirmed", "Cancelled"];

const validationSchema = toTypedSchema(
  z.object({
    scheduleId: z
      .string({ required_error: "Select a showtime schedule" })
      .uuid("Invalid Schedule UUID"),
    userId: z
      .string({ required_error: "Assign this ticket to a customer account" })
      .uuid("Invalid User UUID"),
    status: z.string({ required_error: "Select a status" }),
  }),
);

const { handleSubmit, isSubmitting, errors, resetForm, setValues } = useForm({
  validationSchema,
  initialValues: {
    scheduleId: "",
    userId: "",
    status: "Inprogress",
  },
});

const { value: scheduleId } = useField("scheduleId");
const { value: userId } = useField("userId");
const { value: status } = useField("status");

watch(
  () => props.initialData,
  (newData) => {
    if (newData) {
      setValues({
        scheduleId: newData.schedule_id || "",
        userId: newData.user_id || "",
        status: newData.status || "Inprogress",
      });
    }
  },
  { immediate: true },
);

const emit = defineEmits(["submitTicket"]);
const onSubmit = handleSubmit((values) => {
  const cleanPayload = {
    schedule_id: values.scheduleId,
    user_id: values.userId,
    status: values.status,
  };

  emit("submitTicket", cleanPayload);
  resetForm();
});
</script>

<template>
  <div
    class="w-full max-w-lg bg-gray-950 border border-gray-900 rounded-3xl p-6 shadow-2xl text-white"
  >
    <div class="mb-5">
      <h2 class="text-xl font-bold tracking-wide">
        {{ initialData ? "Update Ticket Record" : "Generate Database Ticket" }}
      </h2>
      <p class="text-xs text-gray-500 mt-0.5">
        {{
          initialData
            ? "Modify relationships for this ticket profile record."
            : "Relational mapping for Schedules, Users, and Status nodes."
        }}
      </p>
    </div>

    <form @submit.prevent="onSubmit" class="space-y-4">
      <div class="flex flex-col gap-1.5">
        <label class="text-xs font-bold text-gray-400 uppercase tracking-wider"
          >Showtime Schedule</label
        >
        <select
          v-model="scheduleId"
          :class="[
            'w-full bg-gray-900 border rounded-xl px-4 py-3 text-sm focus:outline-none transition-colors',
            errors.scheduleId
              ? 'border-red-500'
              : 'border-gray-800 focus:border-lime-400',
          ]"
        >
          <option value="" disabled selected class="text-gray-600">
            Select a movie time slot...
          </option>
          <option v-for="sch in activeSchedules" :key="sch.id" :value="sch.id">
            {{ sch.details }}
          </option>
        </select>
        <span v-if="errors.scheduleId" class="text-xs text-red-500">{{
          errors.scheduleId
        }}</span>
      </div>

      <div class="flex flex-col gap-1.5">
        <label class="text-xs font-bold text-gray-400 uppercase tracking-wider"
          >Assign to Registered User</label
        >
        <select
          v-model="userId"
          :class="[
            'w-full bg-gray-900 border rounded-xl px-4 py-3 text-sm focus:outline-none transition-colors',
            errors.userId
              ? 'border-red-500'
              : 'border-gray-800 focus:border-lime-400',
          ]"
        >
          <option value="" disabled selected class="text-gray-600">
            Select the buyer account...
          </option>
          <option v-for="user in activeUsers" :key="user.id" :value="user.id">
            {{ user.name }}
          </option>
        </select>
        <span v-if="errors.userId" class="text-xs text-red-500">{{
          errors.userId
        }}</span>
      </div>

      <div class="flex flex-col gap-1.5">
        <label class="text-xs font-bold text-gray-400 uppercase tracking-wider"
          >Ticket Status</label
        >
        <select
          v-model="status"
          class="w-full bg-gray-900 border border-gray-800 rounded-xl px-4 py-3 text-sm focus:outline-none focus:border-lime-400 text-white"
        >
          <option v-for="state in statusOptions" :key="state" :value="state">
            {{ state }}
          </option>
        </select>
      </div>

      <button
        type="submit"
        :disabled="isSubmitting"
        class="w-full bg-lime-400 hover:bg-lime-300 disabled:bg-gray-800 disabled:text-gray-600 text-black font-bold text-sm py-3 rounded-xl shadow-lg transition-all mt-2 cursor-pointer"
      >
        {{
          isSubmitting
            ? "Processing..."
            : initialData
            ? "Update Ticket Reference"
            : "Commit Ticket Record to Hasura"
        }}
      </button>
    </form>
  </div>
</template>