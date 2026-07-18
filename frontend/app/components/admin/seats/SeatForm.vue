<script setup>
import { useForm, useField } from "vee-validate";
import { toTypedSchema } from "@vee-validate/zod";
import * as z from "zod";

defineProps({
  hallsList: {
    type: Array,
    required: true,
  },
});

const validationSchema = toTypedSchema(
  z.object({
    hallId: z
      .string({
        required_error: "Please select a hall",
      })
      .uuid("Please select a valid hall"),

    rows: z
      .string({
        required_error: "Rows are required",
      })
      .min(1, "Enter at least one row"),

    seatsPerRow: z
      .number({
        required_error: "Seats per row is required",
      })
      .positive("Seat count must be greater than zero"),
  }),
);

const { handleSubmit, errors, resetForm, isSubmitting } = useForm({
  validationSchema,

  initialValues: {
    hallId: "",
    rows: "A,B,C,D",
    seatsPerRow: 10,
  },
});

const { value: hallId } = useField("hallId");
const { value: rows } = useField("rows");
const { value: seatsPerRow } = useField("seatsPerRow");

const emit = defineEmits(["submitSeats"]);

const onSubmit = handleSubmit((values) => {
  const generatedSeats = [];

  const rowList = values.rows
    .split(",")
    .map((row) => row.trim().toUpperCase())
    .filter(Boolean);

  rowList.forEach((row) => {
    for (let seat = 1; seat <= values.seatsPerRow; seat++) {
      generatedSeats.push({
        hall_id: values.hallId,
        row_name: row,
        seat_number: seat,
        type: "standard",
      });
    }
  });

  console.log("Generated Seats:", generatedSeats);

  emit("submitSeats", generatedSeats);

  resetForm();
});
</script>

<template>
  <div
    class="w-full max-w-2xl bg-gray-950 border border-gray-900 rounded-3xl p-6 shadow-2xl text-white"
  >
    <div class="mb-6">
      <h2 class="text-xl font-bold tracking-wide">Generate Hall Seats</h2>

      <p class="text-xs text-gray-500 mt-1">
        Automatically generate seats for a cinema hall.
      </p>
    </div>

    <form @submit.prevent="onSubmit" class="space-y-5">
      <!-- Hall -->
      <div class="flex flex-col gap-1.5">
        <label class="text-xs font-bold text-gray-400 uppercase tracking-wider">
          Cinema Hall
        </label>

        <select
          v-model="hallId"
          :class="[
            'w-full bg-gray-900 border rounded-xl px-4 py-3 text-sm focus:outline-none focus:border-lime-400 cursor-pointer',
            errors.hallId ? 'border-red-500' : 'border-gray-800',
          ]"
        >
          <option value="" disabled>Choose Hall...</option>

          <option v-for="hall in hallsList" :key="hall.id" :value="hall.id">
            {{ hall.name }}
          </option>
        </select>

        <span v-if="errors.hallId" class="text-xs text-red-500">
          {{ errors.hallId }}
        </span>
      </div>

      <!-- Rows -->
      <div class="flex flex-col gap-1.5">
        <label class="text-xs font-bold text-gray-400 uppercase tracking-wider">
          Row Labels
        </label>

        <input
          v-model="rows"
          type="text"
          placeholder="A,B,C,D,E"
          :class="[
            'w-full bg-gray-900 border rounded-xl px-4 py-3 text-sm focus:outline-none focus:border-lime-400',
            errors.rows ? 'border-red-500' : 'border-gray-800',
          ]"
        />

        <span v-if="errors.rows" class="text-xs text-red-500">
          {{ errors.rows }}
        </span>

        <p class="text-xs text-gray-500">
          Separate row names using commas (Example: A,B,C,D)
        </p>
      </div>

      <!-- Seats Per Row -->
      <div class="flex flex-col gap-1.5">
        <label class="text-xs font-bold text-gray-400 uppercase tracking-wider">
          Seats Per Row
        </label>

        <input
          v-model.number="seatsPerRow"
          type="number"
          min="1"
          :class="[
            'w-full bg-gray-900 border rounded-xl px-4 py-3 text-sm focus:outline-none focus:border-lime-400',
            errors.seatsPerRow ? 'border-red-500' : 'border-gray-800',
          ]"
        />

        <span v-if="errors.seatsPerRow" class="text-xs text-red-500">
          {{ errors.seatsPerRow }}
        </span>
      </div>

      <div class="pt-2">
        <button
          type="submit"
          :disabled="isSubmitting"
          class="w-full bg-lime-400 hover:bg-lime-300 disabled:bg-gray-800 text-black font-bold text-sm py-3 rounded-xl transition-all cursor-pointer"
        >
          {{ isSubmitting ? "Generating Seats..." : "Generate Seats" }}
        </button>
      </div>
    </form>
  </div>
</template>
