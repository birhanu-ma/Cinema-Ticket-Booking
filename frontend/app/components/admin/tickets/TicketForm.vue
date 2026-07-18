<script setup>
import { ref, onMounted } from "vue";
import { gql } from "@apollo/client/core";
import { useForm, useField } from "vee-validate";

import { toTypedSchema } from "@vee-validate/zod";

import * as z from "zod";

const { $apollo } = useNuxtApp();

const schedules = ref([]);

const users = ref([]);

const GET_DATA = gql`
  query {
    schedules {
      id

      start_time

      movie {
        title
      }
    }

    users {
      id

      name

      email
    }
  }
`;

onMounted(async () => {
  try {
    const { data } = await $apollo.query({
      query: GET_DATA,
    });

    schedules.value = data.schedules;

    users.value = data.users;
  } catch (err) {
    console.error(err);
  }
});

const validationSchema = toTypedSchema(
  z.object({
    scheduleId: z
      .string({
        required_error: "Select schedule",
      })
      .uuid(),

    userId: z
      .string({
        required_error: "Select user",
      })
      .uuid(),

    status: z.string(),
  }),
);

const {
  handleSubmit,

  isSubmitting,

  errors,

  resetForm,
} = useForm({
  validationSchema,

  initialValues: {
    scheduleId: "",

    userId: "",

    status: "pending",
  },
});

const { value: scheduleId } = useField("scheduleId");

const { value: userId } = useField("userId");

const { value: status } = useField("status");

const emit = defineEmits(["submitTicket"]);

const onSubmit = handleSubmit((values) => {
  emit(
    "submitTicket",

    {
      schedule_id: values.scheduleId,

      user_id: values.userId,

      status: values.status,
    },
  );

  resetForm();
});
</script>

<template>
  <div
    class="w-full max-w-lg bg-gray-950 border border-gray-900 rounded-3xl p-6 shadow-2xl text-white"
  >
    <div class="mb-5">
      <h2 class="text-xl font-bold tracking-wide">Generate Database Ticket</h2>

      <p class="text-xs text-gray-500 mt-0.5">
        Relational mapping for schedules, users, and ticket status.
      </p>
    </div>

    <form @submit.prevent="onSubmit" class="space-y-4">
      <div class="flex flex-col gap-1.5">
        <label class="text-xs font-bold text-gray-400 uppercase tracking-wider">
          Showtime Schedule
        </label>

        <select
          v-model="scheduleId"
          class="w-full bg-gray-900 border border-gray-800 rounded-xl px-4 py-3 text-sm focus:outline-none focus:border-lime-400"
        >
          <option value="" disabled>Select movie schedule...</option>

          <option v-for="sch in schedules" :key="sch.id" :value="sch.id">
            {{ sch.movie?.title }} -
            {{ new Date(sch.start_time).toLocaleString() }}
          </option>
        </select>

        <span v-if="errors.scheduleId" class="text-xs text-red-500">
          {{ errors.scheduleId }}
        </span>
      </div>

      <div class="flex flex-col gap-1.5">
        <label class="text-xs font-bold text-gray-400 uppercase tracking-wider">
          Customer
        </label>

        <select
          v-model="userId"
          class="w-full bg-gray-900 border border-gray-800 rounded-xl px-4 py-3 text-sm focus:outline-none focus:border-lime-400"
        >
          <option value="" disabled>Select customer...</option>

          <option v-for="user in users" :key="user.id" :value="user.id">
            {{ user.name }} ({{ user.email }})
          </option>
        </select>

        <span v-if="errors.userId" class="text-xs text-red-500">
          {{ errors.userId }}
        </span>
      </div>

      <div class="flex flex-col gap-1.5">
        <label class="text-xs font-bold text-gray-400 uppercase tracking-wider">
          Ticket Status
        </label>

        <select
          v-model="status"
          class="w-full bg-gray-900 border border-gray-800 rounded-xl px-4 py-3 text-sm focus:outline-none focus:border-lime-400"
        >
          <option value="pending">Pending</option>

          <option value="confirmed">Confirmed</option>

          <option value="cancelled">Cancelled</option>
        </select>
      </div>

      <button
        type="submit"
        :disabled="isSubmitting"
        class="w-full bg-lime-400 hover:bg-lime-300 disabled:bg-gray-800 disabled:text-gray-600 text-black font-bold text-sm py-3 rounded-xl shadow-lg transition-all mt-2 cursor-pointer"
      >
        {{ isSubmitting ? "Processing..." : "Create Ticket " }}
      </button>
    </form>
  </div>
</template>
