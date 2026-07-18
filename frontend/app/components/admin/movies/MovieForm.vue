<script setup>
import { useForm, useField } from "vee-validate";
import { toTypedSchema } from "@vee-validate/zod";
import * as z from "zod";

const props = defineProps({
  directorsList: {
    type: Array,
    required: true,
  },

  genresList: {
    type: Array,
    required: true,
  },

  starsList: {
    type: Array,
    required: true,
  },
});

const validationSchema = toTypedSchema(
  z.object({
    title: z
      .string({
        required_error: "Movie title is required",
      })
      .min(1),

    description: z
      .string({
        required_error: "Description is required",
      })
      .min(10),

    duration: z
      .number({
        required_error: "Duration is required",
      })
      .positive(),

    thumbnailUrl: z
      .string({
        required_error: "Thumbnail URL is required",
      })
      .url(),

    directorId: z
      .string({
        required_error: "Please select a director",
      })
      .uuid(),

    genreId: z
      .string({
        required_error: "Please select a genre",
      })
      .uuid(),

    starId: z
      .string({
        required_error: "Please select a main star",
      })
      .uuid(),
  }),
);

const { handleSubmit, isSubmitting, errors, resetForm } = useForm({
  validationSchema,

  initialValues: {
    title: "",
    description: "",
    duration: 120,
    directorId: "",
    genreId: "",
    starId: "",
    thumbnailUrl: "",
  },
});

const { value: title } = useField("title");
const { value: description } = useField("description");
const { value: duration } = useField("duration");
const { value: thumbnailUrl } = useField("thumbnailUrl");
const { value: directorId } = useField("directorId");
const { value: genreId } = useField("genreId");
const { value: starId } = useField("starId");

const emit = defineEmits(["submitMovie"]);

const onSubmit = handleSubmit((values) => {
  const graphqlPayload = {
    title: values.title,

    description: values.description,

    duration: values.duration,

    director_id: values.directorId,

    movie_genres: {
      data: [
        {
          genre_id: values.genreId,
        },
      ],
    },

    movie_stars: {
      data: [
        {
          star_id: values.starId,
        },
      ],
    },

    movie_images: {
      data: [
        {
          image_url: values.thumbnailUrl,
          is_featured: true,
        },
      ],
    },
  };

  emit("submitMovie", graphqlPayload);

  resetForm();
});
</script>

<template>
  <div
    class="w-full max-w-2xl bg-gray-950 border border-gray-900 rounded-3xl p-6 shadow-2xl text-white"
  >
    <div class="mb-6">
      <h2 class="text-xl font-bold tracking-wide">Add Catalog Movie</h2>

      <p class="text-xs text-gray-500 mt-1">
        Select and attach relational keys from master tables dynamically.
      </p>
    </div>

    <form @submit.prevent="onSubmit" class="space-y-4">
      <div class="grid grid-cols-3 gap-4">
        <div class="col-span-2 flex flex-col gap-1.5">
          <label
            class="text-xs font-bold text-gray-400 uppercase tracking-wider"
          >
            Movie Title
          </label>

          <input
            v-model="title"
            type="text"
            placeholder="e.g., Inception"
            :class="[
              'w-full bg-gray-900 border rounded-xl px-4 py-3 text-sm focus:outline-none focus:border-lime-400',
              errors.title ? 'border-red-500' : 'border-gray-800',
            ]"
          />

          <span v-if="errors.title" class="text-xs text-red-500">
            {{ errors.title }}
          </span>
        </div>

        <div class="flex flex-col gap-1.5">
          <label
            class="text-xs font-bold text-gray-400 uppercase tracking-wider"
          >
            Duration (min)
          </label>

          <input
            v-model.number="duration"
            type="number"
            :class="[
              'w-full bg-gray-900 border rounded-xl px-4 py-3 text-sm focus:outline-none focus:border-lime-400',
              errors.duration ? 'border-red-500' : 'border-gray-800',
            ]"
          />

          <span v-if="errors.duration" class="text-xs text-red-500">
            {{ errors.duration }}
          </span>
        </div>
      </div>

      <div class="flex flex-col gap-1.5">
        <label class="text-xs text-gray-400 font-bold uppercase tracking-wider">
          Featured Thumbnail URL
        </label>

        <input
          v-model="thumbnailUrl"
          type="text"
          placeholder="https://images.unsplash.com/..."
          :class="[
            'w-full bg-gray-900 border rounded-xl px-4 py-3 text-sm focus:outline-none focus:border-lime-400',
            errors.thumbnailUrl ? 'border-red-500' : 'border-gray-800',
          ]"
        />

        <span v-if="errors.thumbnailUrl" class="text-xs text-red-500">
          {{ errors.thumbnailUrl }}
        </span>
      </div>

      <div class="flex flex-col gap-1.5">
        <label class="text-xs font-bold text-gray-400 uppercase tracking-wider">
          Movie Director
        </label>

        <select
          v-model="directorId"
          :class="[
            'w-full bg-gray-900 border rounded-xl px-4 py-3 text-sm focus:outline-none focus:border-lime-400 cursor-pointer',
            errors.directorId ? 'border-red-500' : 'border-gray-800',
          ]"
        >
          <option value="" disabled>Choose the primary director...</option>

          <option v-for="dir in directorsList" :key="dir.id" :value="dir.id">
            {{ dir.name }}
          </option>
        </select>

        <span v-if="errors.directorId" class="text-xs text-red-500">
          {{ errors.directorId }}
        </span>
      </div>

      <div class="flex flex-col gap-1.5">
        <label class="text-xs font-bold text-gray-400 uppercase tracking-wider">
          Movie Genre
        </label>

        <select
          v-model="genreId"
          :class="[
            'w-full bg-gray-900 border rounded-xl px-4 py-3 text-sm focus:outline-none focus:border-lime-400 cursor-pointer',
            errors.genreId ? 'border-red-500' : 'border-gray-800',
          ]"
        >
          <option value="" disabled>Choose the primary genre...</option>

          <option v-for="gen in genresList" :key="gen.id" :value="gen.id">
            {{ gen.name }}
          </option>
        </select>

        <span v-if="errors.genreId" class="text-xs text-red-500">
          {{ errors.genreId }}
        </span>
      </div>

      <div class="flex flex-col gap-1.5">
        <label class="text-xs font-bold text-gray-400 uppercase tracking-wider">
          Lead Star
        </label>

        <select
          v-model="starId"
          :class="[
            'w-full bg-gray-900 border rounded-xl px-4 py-3 text-sm focus:outline-none focus:border-lime-400 cursor-pointer',
            errors.starId ? 'border-red-500' : 'border-gray-800',
          ]"
        >
          <option value="" disabled>Choose the primary actor...</option>

          <option v-for="star in starsList" :key="star.id" :value="star.id">
            {{ star.name }}
          </option>
        </select>

        <span v-if="errors.starId" class="text-xs text-red-500">
          {{ errors.starId }}
        </span>
      </div>

      <div class="flex flex-col gap-1.5">
        <label class="text-xs font-bold text-gray-400 uppercase tracking-wider">
          Movie Description
        </label>

        <textarea
          v-model="description"
          rows="3"
          placeholder="Cinematic plot narrative..."
          :class="[
            'w-full bg-gray-900 border rounded-xl px-4 py-3 text-sm focus:outline-none focus:border-lime-400 resize-none',
            errors.description ? 'border-red-500' : 'border-gray-800',
          ]"
        ></textarea>

        <span v-if="errors.description" class="text-xs text-red-500">
          {{ errors.description }}
        </span>
      </div>

      <div class="pt-2">
        <button
          type="submit"
          :disabled="isSubmitting"
          class="w-full bg-lime-400 hover:bg-lime-300 disabled:bg-gray-800 text-black font-bold text-sm py-3 rounded-xl shadow-lg transition-all cursor-pointer"
        >
          {{ isSubmitting ? "Processing relations..." : "Save Movie" }}
        </button>
      </div>
    </form>
  </div>
</template>
