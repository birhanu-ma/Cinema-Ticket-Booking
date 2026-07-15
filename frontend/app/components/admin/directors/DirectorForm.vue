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

const validationSchema = toTypedSchema(
  z.object({
    name: z
      .string({ required_error: "Director name is required" })
      .min(2, "Name must be at least 2 characters long"),
    bio: z.string().nullable().optional(),
    photoUrl: z
      .string()
      .url("Please provide a valid image URL link (e.g., https://...)")
      .or(z.literal("")) 
      .nullable()
      .optional(),
  }),
);

const { handleSubmit, isSubmitting, errors, resetForm, setValues } = useForm({
  validationSchema,
  initialValues: {
    name: "",
    bio: "",
    photoUrl: "",
  },
});

const { value: name } = useField("name");
const { value: bio } = useField("bio");
const { value: photoUrl } = useField("photoUrl");

watch(
  () => props.initialData,
  (newData) => {
    if (newData) {
      setValues({
        name: newData.name || "",
        bio: newData.bio || "",
        photoUrl: newData.photo_url || "", 
      });
    }
  },
  { immediate: true },
);

const emit = defineEmits(["submitDirector"]);

const onSubmit = handleSubmit((values) => {
  const cleanPayload = {
    name: values.name,
    bio: values.bio || null,
    photo_url: values.photoUrl || null, 
  };

  emit("submitDirector", cleanPayload);
  resetForm();
});
</script>

<template>
  <div
    class="w-full max-w-lg bg-gray-950 border border-gray-900 rounded-3xl p-6 shadow-2xl text-white"
  >
    <div class="mb-6">
      <h2 class="text-xl font-bold tracking-wide">
        {{ initialData ? "Edit Director Profile" : "Register New Director" }}
      </h2>
      <p class="text-xs text-gray-500 mt-1">
        {{
          initialData
            ? "Modify this filmmaker's metadata records within the database catalog."
            : "Add filmmaker metadata reference nodes to the active system catalog."
        }}
      </p>
    </div>

    <form @submit.prevent="onSubmit" class="space-y-5">
      <div class="flex flex-col gap-1.5">
        <label class="text-xs font-bold text-gray-400 uppercase tracking-wider">
          Director Full Name <span class="text-red-500">*</span>
        </label>
        <input
          v-model="name"
          type="text"
          placeholder="e.g., Christopher Nolan"
          :class="[
            'w-full bg-gray-900 border rounded-xl px-4 py-3 text-sm focus:outline-none transition-colors placeholder:text-gray-600',
            errors.name
              ? 'border-red-500 focus:border-red-500'
              : 'border-gray-800 focus:border-lime-400',
          ]"
        />
        <span
          v-if="errors.name"
          class="text-xs font-medium text-red-500 mt-0.5"
        >
          {{ errors.name }}
        </span>
      </div>

      <div class="flex flex-col gap-1.5">
        <label class="text-xs font-bold text-gray-400 uppercase tracking-wider">
          Profile Portrait Image URL
          <span class="text-gray-600 text-[10px]">(Optional)</span>
        </label>
        <input
          v-model="photoUrl"
          type="text"
          placeholder="https://images.unsplash.com/photo-..."
          :class="[
            'w-full bg-gray-900 border rounded-xl px-4 py-3 text-sm focus:outline-none transition-colors placeholder:text-gray-600',
            errors.photoUrl
              ? 'border-red-500 focus:border-red-500'
              : 'border-gray-800 focus:border-lime-400',
          ]"
        />
        <span
          v-if="errors.photoUrl"
          class="text-xs font-medium text-red-500 mt-0.5"
        >
          {{ errors.photoUrl }}
        </span>
      </div>

      <div class="flex flex-col gap-1.5">
        <label class="text-xs font-bold text-gray-400 uppercase tracking-wider">
          Biography Overview
          <span class="text-gray-600 text-[10px]">(Optional)</span>
        </label>
        <textarea
          v-model="bio"
          rows="4"
          placeholder="Write a brief professional summary about their stylistic film history..."
          class="w-full bg-gray-900 border border-gray-800 rounded-xl px-4 py-3 text-sm focus:outline-none focus:border-lime-400 transition-colors placeholder:text-gray-600 resize-none"
        ></textarea>
      </div>

      <div class="pt-2">
        <button
          type="submit"
          :disabled="isSubmitting"
          class="w-full bg-lime-400 hover:bg-lime-300 disabled:bg-gray-800 disabled:text-gray-600 text-black font-bold text-sm py-3 px-4 rounded-xl shadow-lg hover:shadow-lime-400/10 transition-all duration-200 active:scale-[0.99] cursor-pointer"
        >
          {{
            isSubmitting
              ? "Saving Metadata..."
              : initialData
                ? "Update Director Profile"
                : "Save Director Profile"
          }}
        </button>
      </div>
    </form>
  </div>
</template>
