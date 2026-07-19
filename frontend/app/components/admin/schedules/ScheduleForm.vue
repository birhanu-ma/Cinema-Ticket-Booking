<script setup>
import { useForm, useField } from "vee-validate";
import { toTypedSchema } from "@vee-validate/zod";
import * as z from "zod";

const props = defineProps({
  moviesList: {
    type: Array,
    required: true,
  },

  hallsList: {
    type: Array,
    required: true,
  },
});

const validationSchema = toTypedSchema(
  z.object({
    moviesId: z.string().uuid("Please select a movie"),

    hallId: z.string().uuid("Please select a cinema hall"),

    price: z.number().positive("Price must be greater than zero"),

    startTime: z.string().min(1, "Please select a date"),
  }),
);

const { handleSubmit, errors, resetForm, isSubmitting } = useForm({
  validationSchema,

  initialValues: {
    moviesId: "",

    hallId: "",

    price: 100,

    startTime: "",
  },
});

const { value: moviesId } = useField("moviesId");

const { value: hallId } = useField("hallId");

const { value: price } = useField("price");

const { value: startTime } = useField("startTime");

const emit = defineEmits(["submitSchedule"]);

const onSubmit = handleSubmit((values) => {
  const payload = {
    movies_id: values.moviesId,

    hall_id: values.hallId,

    price: Number(values.price),

    // date only
    start_time: values.startTime,
  };

  console.log("Schedule payload:", payload);

  emit("submitSchedule", payload);

  resetForm();
});
</script>

<template>
  <div
    class="w-full max-w-lg bg-gray-950 border border-gray-900 rounded-3xl p-6 text-white"
  >
    <h2 class="text-xl font-bold mb-6">Create Movie Schedule</h2>

    <form @submit.prevent="onSubmit" class="space-y-5">
      <!-- MOVIE -->

      <div>
        <label class="text-xs uppercase text-gray-400"> Movie </label>

        <select
          v-model="moviesId"
          class="w-full mt-2 bg-gray-900 border border-gray-800 rounded-xl px-4 py-3"
        >
          <option value="">Select Movie</option>

          <option
            v-for="movie in props.moviesList"
            :key="movie.id"
            :value="movie.id"
          >
            {{ movie.title }}
          </option>
        </select>

        <p v-if="errors.moviesId" class="text-red-500 text-xs mt-1">
          {{ errors.moviesId }}
        </p>
      </div>

      <!-- HALL -->

      <div>
        <label class="text-xs uppercase text-gray-400"> Cinema Hall </label>

        <select
          v-model="hallId"
          class="w-full mt-2 bg-gray-900 border border-gray-800 rounded-xl px-4 py-3"
        >
          <option value="">Select Hall</option>

          <option
            v-for="hall in props.hallsList"
            :key="hall.id"
            :value="hall.id"
          >
            {{ hall.name }}
          </option>
        </select>

        <p v-if="errors.hallId" class="text-red-500 text-xs mt-1">
          {{ errors.hallId }}
        </p>
      </div>

      <div>
        <label class="text-xs uppercase text-gray-400"> Ticket Price </label>

        <input
          v-model.number="price"
          type="number"
          min="1"
          class="w-full mt-2 bg-gray-900 border border-gray-800 rounded-xl px-4 py-3"
        />

        <p v-if="errors.price" class="text-red-500 text-xs mt-1">
          {{ errors.price }}
        </p>
      </div>

      <div>
        <label class="text-xs uppercase text-gray-400"> Show Date </label>

        <input
          v-model="startTime"
          type="date"
          class="w-full mt-2 bg-gray-900 border border-gray-800 rounded-xl px-4 py-3 text-white"
        />

        <p class="text-xs text-gray-500 mt-1">Select show date only</p>

        <p v-if="errors.startTime" class="text-red-500 text-xs mt-1">
          {{ errors.startTime }}
        </p>
      </div>

      <button
        type="submit"
        :disabled="isSubmitting"
        class="w-full bg-lime-400 cursor-pointer text-black font-bold py-3 rounded-xl"
      >
        {{ isSubmitting ? "Saving..." : "Create Schedule" }}
      </button>
    </form>
  </div>
</template>
