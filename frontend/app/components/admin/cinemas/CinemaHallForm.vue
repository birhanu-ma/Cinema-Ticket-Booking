<script setup>
import { useForm, useField } from "vee-validate";
import { toTypedSchema } from "@vee-validate/zod";
import * as z from "zod";

defineProps({
  cinemasList: {
    type: Array,
    required: true,
  },
});

const validationSchema = toTypedSchema(
  z.object({
    cinemaId: z
      .string({
        required_error: "Please select a cinema",
      })
      .uuid("Please select a valid cinema"),

    name: z
      .string({
        required_error: "Hall name is required",
      })
      .min(1, "Hall name is required"),

    capacity: z
      .number({
        required_error: "Capacity is required",
      })
      .positive("Capacity must be greater than zero"),
  }),
);

const { handleSubmit, errors, resetForm, isSubmitting } = useForm({
  validationSchema,

  initialValues: {
    cinemaId: "",
    name: "",
    capacity: 100,
  },
});

const { value: cinemaId } = useField("cinemaId");
const { value: name } = useField("name");
const { value: capacity } = useField("capacity");

const emit = defineEmits(["submitHall"]);

const onSubmit = handleSubmit((values) => {
  console.log("Submitting hall:", values);

  emit("submitHall", {
    cinema_id: values.cinemaId,
    name: values.name,
    capacity: values.capacity,
  });

  resetForm();
});
</script>

<template>
  <div
    class="w-full max-w-2xl mx-auto bg-gray-950 border border-gray-900 rounded-3xl p-6 shadow-2xl text-white"
  >
    <div class="mb-6">
      <h2 class="text-xl font-bold tracking-wide">Create Cinema Hall</h2>

      <p class="text-xs text-gray-500 mt-1">
        Assign a hall to a cinema location.
      </p>
    </div>

    <form @submit.prevent="onSubmit" class="space-y-5">
      <!-- Cinema -->
      <div class="flex flex-col gap-1.5">
        <label class="text-xs font-bold text-gray-400 uppercase tracking-wider">
          Cinema
        </label>

        <select
          v-model="cinemaId"
          :class="[
            'w-full bg-gray-900 border rounded-xl px-4 py-3 text-sm focus:outline-none focus:border-lime-400 cursor-pointer',
            errors.cinemaId ? 'border-red-500' : 'border-gray-800',
          ]"
        >
          <option value="" disabled>Select Cinema...</option>

          <option
            v-for="cinema in cinemasList"
            :key="cinema.id"
            :value="cinema.id"
          >
            {{ cinema.name }}
          </option>
        </select>

        <span v-if="errors.cinemaId" class="text-xs text-red-500">
          {{ errors.cinemaId }}
        </span>
      </div>

      <!-- Hall Name -->
      <div class="flex flex-col gap-1.5">
        <label class="text-xs font-bold text-gray-400 uppercase tracking-wider">
          Hall Name
        </label>

        <input
          v-model="name"
          type="text"
          placeholder="Hall 1"
          :class="[
            'w-full bg-gray-900 border rounded-xl px-4 py-3 text-sm focus:outline-none focus:border-lime-400',
            errors.name ? 'border-red-500' : 'border-gray-800',
          ]"
        />

        <span v-if="errors.name" class="text-xs text-red-500">
          {{ errors.name }}
        </span>
      </div>

      <!-- Capacity -->
      <div class="flex flex-col gap-1.5">
        <label class="text-xs font-bold text-gray-400 uppercase tracking-wider">
          Capacity
        </label>

        <input
          v-model.number="capacity"
          type="number"
          min="1"
          :class="[
            'w-full bg-gray-900 border rounded-xl px-4 py-3 text-sm focus:outline-none focus:border-lime-400',
            errors.capacity ? 'border-red-500' : 'border-gray-800',
          ]"
        />

        <span v-if="errors.capacity" class="text-xs text-red-500">
          {{ errors.capacity }}
        </span>
      </div>

      <div class="pt-2">
        <button
          type="submit"
          :disabled="isSubmitting"
          class="w-full bg-lime-400 hover:bg-lime-300 disabled:bg-gray-800 text-black font-bold text-sm py-3 rounded-xl transition-all cursor-pointer"
        >
          {{ isSubmitting ? "Saving..." : "Save Hall" }}
        </button>
      </div>
    </form>
  </div>
</template>
