<script setup>
import { ref } from "vue";
import { useForm, useField, useFieldArray } from "vee-validate";
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
    title: z.string({ required_error: "Movie title is required" }).min(1),

    description: z
      .string({ required_error: "Description is required" })
      .min(10),

    duration: z.number({ required_error: "Duration is required" }).positive(),

    directorId: z.string({ required_error: "Please select a director" }).uuid(),

    genreIds: z.array(z.string().uuid()).min(1, "Select at least one genre"),

    starIds: z.array(z.string().uuid()).min(1, "Select at least one star"),

    images: z
      .array(
        z.object({
          url: z.string().url(),
          isFeatured: z.boolean(),
        }),
      )
      .min(1, "Upload at least one image")
      .refine(
        (imgs) => imgs.filter((i) => i.isFeatured).length === 1,
        "Exactly one image must be marked as featured",
      ),
  }),
);

const { handleSubmit, isSubmitting, errors, resetForm } = useForm({
  validationSchema,

  initialValues: {
    title: "",
    description: "",
    duration: 120,
    directorId: "",
    genreIds: [],
    starIds: [],
    images: [],
  },
});

const { value: title } = useField("title");
const { value: description } = useField("description");
const { value: duration } = useField("duration");
const { value: directorId } = useField("directorId");
const { value: genreIds } = useField("genreIds");
const { value: starIds } = useField("starIds");

const {
  fields: imageFields,
  push: pushImage,
  remove: removeImage,
} = useFieldArray("images");

const addImageSlot = () => {
  pushImage({ url: "", isFeatured: imageFields.value.length === 0 });
};

const setFeatured = (index) => {
  imageFields.value.forEach((field, i) => {
    field.value.isFeatured = i === index;
  });
};

const toggleGenre = (genreId) => {
  const current = genreIds.value || [];

  if (current.includes(genreId)) {
    genreIds.value = current.filter((id) => id !== genreId);
  } else {
    genreIds.value = [...current, genreId];
  }
};

const toggleStar = (starId) => {
  const current = starIds.value || [];

  if (current.includes(starId)) {
    starIds.value = current.filter((id) => id !== starId);
  } else {
    starIds.value = [...current, starId];
  }
};

const emit = defineEmits(["submitMovie"]);

const onSubmit = handleSubmit((values) => {
  const graphqlPayload = {
    title: values.title,
    description: values.description,
    duration: values.duration,
    director_id: values.directorId,

    movie_genres: {
      data: values.genreIds.map((genre_id) => ({ genre_id })),
    },

    movie_stars: {
      data: values.starIds.map((star_id) => ({ star_id })),
    },

    movie_images: {
      data: values.images.map((img) => ({
        image_url: img.url,
        is_featured: img.isFeatured,
      })),
    },
  };

  emit("submitMovie", graphqlPayload);

  resetForm();
});

// Start with one image slot by default
addImageSlot();
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

      <!-- Images -->
      <div class="flex flex-col gap-2">
        <div class="flex items-center justify-between">
          <label
            class="text-xs font-bold text-gray-400 uppercase tracking-wider"
          >
            Movie Images
          </label>

          <button
            type="button"
            @click="addImageSlot"
            class="text-xs text-lime-400 hover:text-lime-300 cursor-pointer font-semibold"
          >
            + Add Image
          </button>
        </div>

        <div class="grid grid-cols-2 gap-4">
          <div
            v-for="(field, index) in imageFields"
            :key="field.key"
            class="flex flex-col gap-2 bg-gray-900/40 border border-gray-900 rounded-xl p-3"
          >
            <AdminImageUploader
              v-model="field.value.url"
              :label="`Image ${index + 1}`"
            />

            <div class="flex items-center justify-between">
              <label
                class="flex items-center gap-2 text-xs text-gray-400 cursor-pointer"
              >
                <input
                  type="radio"
                  name="featuredImage"
                  :checked="field.value.isFeatured"
                  @change="setFeatured(index)"
                  class="accent-lime-400 cursor-pointer"
                />
                Set as featured
              </label>

              <button
                type="button"
                @click="removeImage(index)"
                class="text-xs text-gray-600 hover:text-red-400 cursor-pointer"
              >
                Remove
              </button>
            </div>
          </div>
        </div>

        <span v-if="errors.images" class="text-xs text-red-500">
          {{ errors.images }}
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
          <option value="" disabled>Choose the director...</option>

          <option v-for="dir in directorsList" :key="dir.id" :value="dir.id">
            {{ dir.name }}
          </option>
        </select>

        <span v-if="errors.directorId" class="text-xs text-red-500">
          {{ errors.directorId }}
        </span>
      </div>

      <!-- Genres (multi-select) -->
      <div class="flex flex-col gap-1.5">
        <label class="text-xs font-bold text-gray-400 uppercase tracking-wider">
          Genres
        </label>

        <div
          class="flex flex-wrap gap-2 bg-gray-900 border border-gray-800 rounded-xl p-3"
        >
          <button
            v-for="genre in genresList"
            :key="genre.id"
            type="button"
            @click="toggleGenre(genre.id)"
            class="px-3 py-1.5 rounded-full text-xs font-semibold cursor-pointer transition-colors"
            :class="
              genreIds?.includes(genre.id)
                ? 'bg-lime-400 text-black'
                : 'bg-gray-800 text-gray-400 hover:text-white'
            "
          >
            {{ genre.name }}
          </button>
        </div>

        <span v-if="errors.genreIds" class="text-xs text-red-500">
          {{ errors.genreIds }}
        </span>
      </div>

      <!-- Stars (multi-select) -->
      <div class="flex flex-col gap-1.5">
        <label class="text-xs font-bold text-gray-400 uppercase tracking-wider">
          Stars
        </label>

        <div
          class="flex flex-wrap gap-2 bg-gray-900 border border-gray-800 rounded-xl p-3"
        >
          <button
            v-for="star in starsList"
            :key="star.id"
            type="button"
            @click="toggleStar(star.id)"
            class="px-3 py-1.5 rounded-full text-xs font-semibold cursor-pointer transition-colors"
            :class="
              starIds?.includes(star.id)
                ? 'bg-lime-400 text-black'
                : 'bg-gray-800 text-gray-400 hover:text-white'
            "
          >
            {{ star.name }}
          </button>
        </div>

        <span v-if="errors.starIds" class="text-xs text-red-500">
          {{ errors.starIds }}
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
