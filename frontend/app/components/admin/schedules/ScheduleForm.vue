<script setup>
import { useForm, useField } from 'vee-validate'
import { toTypedSchema } from '@vee-validate/zod'
import * as z from 'zod'

defineProps({
  moviesList: { type: Array, required: true },  
  cinemasList: { type: Array, required: true }  
})

const validationSchema = toTypedSchema(
  z.object({
    moviesId: z.string({ required_error: "Please select a movie" }).uuid(),
    cinemaId: z.string({ required_error: "Please assign a cinema room" }).uuid(),
    price: z
      .number({ required_error: "Ticket price is required" })
      .positive("Price must be a positive number"),
    startTime: z
      .string({ required_error: "Showtime start date and time is required" })
      .min(1, "Please pick a valid date and time")
  })
)

const { handleSubmit, isSubmitting, errors, resetForm } = useForm({
  validationSchema,
  initialValues: {
    moviesId: '',
    cinemaId: '',
    price: 0,
    startTime: ''
  }
})

const { value: moviesId } = useField('moviesId')
const { value: cinemaId } = useField('cinemaId')
const { value: price } = useField('price')
const { value: startTime } = useField('startTime')

const emit = defineEmits(['submitSchedule'])
const onSubmit = handleSubmit((values) => {
  const cleanPayload = {
    movies_id: values.moviesId,
    cinema_id: values.cinemaId,
    price: values.price,
    start_time: new Date(values.startTime).toISOString()
  }
  
  emit('submitSchedule', cleanPayload)
  resetForm()
})
</script>

<template>
  <div class="w-full max-w-lg bg-gray-950 border border-gray-900 rounded-3xl p-6 shadow-2xl text-white">
    <div class="mb-6">
      <h2 class="text-xl font-bold tracking-wide">Create Showtime Schedule</h2>
      <p class="text-xs text-gray-500 mt-1">Bind a catalog movie to a screen hall venue time slot.</p>
    </div>

    <form @submit.prevent="onSubmit" class="space-y-5">
      
      <div class="flex flex-col gap-1.5">
        <label class="text-xs font-bold text-gray-400 uppercase tracking-wider">Select Movie</label>
        <select 
          v-model="moviesId" 
          :class="['w-full bg-gray-900 border rounded-xl px-4 py-3 text-sm focus:outline-none focus:border-lime-400 cursor-pointer text-white', errors.moviesId ? 'border-red-500' : 'border-gray-800']"
        >
          <option value="" disabled selected>Choose a movie catalog item...</option>
          <option v-for="movie in moviesList" :key="movie.id" :value="movie.id">
            {{ movie.title }}
          </option>
        </select>
        <span v-if="errors.moviesId" class="text-xs font-medium text-red-500 mt-0.5">
          {{ errors.moviesId }}
        </span>
      </div>

      <div class="flex flex-col gap-1.5">
        <label class="text-xs font-bold text-gray-400 uppercase tracking-wider">Cinema Hall / Room Auditorium</label>
        <select 
          v-model="cinemaId" 
          :class="['w-full bg-gray-900 border rounded-xl px-4 py-3 text-sm focus:outline-none focus:border-lime-400 cursor-pointer text-white', errors.cinemaId ? 'border-red-500' : 'border-gray-800']"
        >
          <option value="" disabled selected>Choose a viewing hall venue...</option>
          <option v-for="cinema in cinemasList" :key="cinema.id" :value="cinema.id">
            {{ cinema.name }}
          </option>
        </select>
        <span v-if="errors.cinemaId" class="text-xs font-medium text-red-500 mt-0.5">
          {{ errors.cinemaId }}
        </span>
      </div>

      <div class="flex flex-col gap-1.5">
        <label class="text-xs font-bold text-gray-400 uppercase tracking-wider">Ticket Base Price</label>
        <div class="relative">
          <input
            v-model.number="price"
            type="number"
            step="0.01"
            placeholder="0.00"
            :class="['w-full bg-gray-900 border rounded-xl pl-8 pr-4 py-3 text-sm focus:outline-none focus:border-lime-400 text-white', errors.price ? 'border-red-500' : 'border-gray-800']"
          />
          <span class="absolute left-4 top-3.5 text-sm font-bold text-gray-500 select-none">$</span>
        </div>
        <span v-if="errors.price" class="text-xs font-medium text-red-500 mt-0.5">
          {{ errors.price }}
        </span>
      </div>

      <div class="flex flex-col gap-1.5">
        <label class="text-xs font-bold text-gray-400 uppercase tracking-wider">Showtime Start Date & Time</label>
        <input
          v-model="startTime"
          type="datetime-local"
          :class="['w-full bg-gray-900 border rounded-xl px-4 py-3 text-sm focus:outline-none focus:border-lime-400 text-white scheme-dark cursor-pointer', errors.startTime ? 'border-red-500' : 'border-gray-800']"
        />
        <span v-if="errors.startTime" class="text-xs font-medium text-red-500 mt-0.5">
          {{ errors.startTime }}
        </span>
      </div>

      <div class="pt-2">
        <button
          type="submit"
          :disabled="isSubmitting"
          class="w-full bg-lime-400 hover:bg-lime-300 disabled:bg-gray-800 disabled:text-gray-600 text-black font-bold text-sm py-3 px-4 rounded-xl shadow-lg hover:shadow-lime-400/10 transition-all duration-200 active:scale-[0.99] cursor-pointer"
        >
          {{ isSubmitting ? 'Publishing Schedule Time Node...' : 'Commit Movie Schedule' }}
        </button>
      </div>

    </form>
  </div>
</template>