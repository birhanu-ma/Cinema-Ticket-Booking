<script setup>
import { useForm, useField } from "vee-validate";
import { toTypedSchema } from "@vee-validate/zod";
import * as z from "zod";

const validationSchema = toTypedSchema(
  z.object({
    name: z
      .string({
        required_error: "Cinema name is required",
      })
      .min(2),

    location: z
      .string({
        required_error: "Cinema location is required",
      })
      .min(3),
  }),
);

const { handleSubmit, errors, resetForm, isSubmitting } = useForm({
  validationSchema,

  initialValues: {
    name: "",
    location: "",
  },
});

const { value: name } = useField("name");

const { value: location } = useField("location");

const emit = defineEmits(["submitCinema"]);

const onSubmit = handleSubmit((values) => {
  emit("submitCinema", {
    name: values.name,

    location: values.location,
  });

  resetForm();
});
</script>

<template>
  <div
    class="w-full max-w-xl bg-gray-950 border border-gray-900 rounded-3xl p-6 text-white"
  >
    <h2 class="text-xl font-bold mb-6">Add Cinema</h2>

    <form @submit.prevent="onSubmit" class="space-y-4">
      <div>
        <label> Cinema Name </label>

        <input v-model="name" placeholder="Century Cinema" class="input" />

        <span v-if="errors.name" class="text-red-500 text-xs">
          {{ errors.name }}
        </span>
      </div>

      <div>
        <label> Location </label>

        <input
          v-model="location"
          placeholder="Bole Addis Ababa"
          class="input"
        />

        <span v-if="errors.location" class="text-red-500 text-xs">
          {{ errors.location }}
        </span>
      </div>

      <button class="w-full bg-lime-400 text-black py-3 rounded-xl font-bold">
        {{ isSubmitting ? "Saving" : "Save Cinema" }}
      </button>
    </form>
  </div>
</template>
