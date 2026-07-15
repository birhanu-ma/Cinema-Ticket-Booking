<script setup>
import { watch } from 'vue'
import { useForm, useField } from 'vee-validate'
import { toTypedSchema } from '@vee-validate/zod'
import * as z from 'zod'

const props = defineProps({
  initialData: { type: Object, required: true },
  moviesList: { type: Array, required: true },
  cinemasList: { type: Array, required: true }
})

const validationSchema = toTypedSchema(
  z.object({
    moviesId: z.string().uuid(),
    cinemaId: z.string().uuid(),
    price: z.number().positive(),
    startTime: z.string().min(1)
  })
)

const formatISOToLocalDatetime = (isoString) => {
  if (!isoString) return ''
  const d = new Date(isoString)
  return new Date(d.getTime() - d.getTimezoneOffset() * 60000).toISOString().slice(0, 16)
}

const { handleSubmit, isSubmitting, errors, resetForm } = useForm({
  validationSchema,
  initialValues: {
    moviesId: props.initialData?.movies_id || '',
    cinemaId: props.initialData?.cinema_id || '',
    price: props.initialData?.price ? Number(props.initialData.price) : 0,
    startTime: formatISOToLocalDatetime(props.initialData?.start_time)
  }
})

const { value: moviesId } = useField('moviesId')
const { value: cinemaId } = useField('cinemaId')
const { value: price } = useField('price')
const { value: startTime } = useField('startTime')

watch(() => props.initialData, (newData) => {
  if (newData) {
    resetForm({
      values: {
        moviesId: newData.movies_id,
        cinemaId: newData.cinema_id,
        price: Number(newData.price),
        startTime: formatISOToLocalDatetime(newData.start_time)
      }
    })
  }
}, { immediate: true })

const emit = defineEmits(['updateSchedule'])
const onSubmit = handleSubmit((values) => {
  emit('updateSchedule', {
    movies_id: values.moviesId,
    cinema_id: values.cinemaId,
    price: values.price,
    start_time: new Date(values.startTime).toISOString()
  })
})
</script>

<template>
  <div class="w-full max-w-lg bg-gray-950 border border-gray-900 rounded-3xl p-6 shadow-2xl text-white">
    <div class="mb-6">
      <h2 class="text-xl font-bold tracking-wide">Edit Showtime Schedule</h2>
      <p class="text-xs text-gray-500 mt-1">Re-assign room locations, pricing, or calendar timeline vectors.</p>
    </div>

    <form @submit.prevent="onSubmit" class="space-y-5">
      <div class="flex flex-col gap-1.5">
        <label class="text-xs font-bold text-gray-400 uppercase tracking-wider">Select Movie</label>
        <select v-model="moviesId" :class="['w-full bg-gray-900 border rounded-xl px-4 py-3 text-sm focus:outline-none focus:border-lime-400 text-white cursor-pointer', errors.moviesId ? 'border-red-500' : 'border-gray-800']">
          <option v-for="movie in moviesList" :key="movie.id" :value="movie.id">{{ movie.title }}</option>
        </select>
        <span v-if="errors.moviesId" class="text-xs text-red-500">{{ errors.moviesId }}</span>
      </div>

      <div class="flex flex-col gap-1.5">
        <label class="text-xs font-bold text-gray-400 uppercase tracking-wider">Cinema Hall / Room Auditorium</label>
        <select v-model="cinemaId" :class="['w-full bg-gray-900 border rounded-xl px-4 py-3 text-sm focus:outline-none focus:border-lime-400 text-white cursor-pointer', errors.cinemaId ? 'border-red-500' : 'border-gray-800']">
          <option v-for="cinema in cinemasList" :key="cinema.id" :value="cinema.id">{{ cinema.name }}</option>
        </select>
        <span v-if="errors.cinemaId" class="text-xs text-red-500">{{ errors.cinemaId }}</span>
      </div>

      <div class="flex flex-col gap-1.5">
        <label class="text-xs font-bold text-gray-400 uppercase tracking-wider">Ticket Base Price</label>
        <div class="relative">
          <input v-model.number="price" type="number" step="0.01" :class="['w-full bg-gray-900 border rounded-xl pl-8 pr-4 py-3 text-sm focus:outline-none focus:border-lime-400 text-white', errors.price ? 'border-red-500' : 'border-gray-800']" />
          <span class="absolute left-4 top-3.5 text-sm font-bold text-gray-500">$</span>
        </div>
        <span v-if="errors.price" class="text-xs text-red-500">{{ errors.price }}</span>
      </div>

      <div class="flex flex-col gap-1.5">
        <label class="text-xs font-bold text-gray-400 uppercase tracking-wider">Showtime Start Date & Time</label>
        <input v-model="startTime" type="datetime-local" :class="['w-full bg-gray-900 border rounded-xl px-4 py-3 text-sm focus:outline-none focus:border-lime-400 text-white scheme-dark cursor-pointer', errors.startTime ? 'border-red-500' : 'border-gray-800']" />
        <span v-if="errors.startTime" class="text-xs text-red-500">{{ errors.startTime }}</span>
      </div>

      <div class="pt-2">
        <button type="submit" :disabled="isSubmitting" class="w-full bg-lime-400 hover:bg-lime-300 disabled:bg-gray-800 text-black font-bold text-sm py-3 rounded-xl shadow-lg transition-all cursor-pointer">
          {{ isSubmitting ? 'Modifying show node...' : 'Update Schedule Slot' }}
        </button>
      </div>
    </form>
  </div>
</template>