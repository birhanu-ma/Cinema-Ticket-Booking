<script setup>
import { useForm, useField } from 'vee-validate'
import { toTypedSchema } from '@vee-validate/zod'
import * as z from 'zod'

const validationSchema = toTypedSchema(
  z.object({
    name: z
      .string({ required_error: "Cinema name is required" })
      .min(2, "Name must be at least 2 characters long"),
    location: z
      .string({ required_error: "Location is required" })
      .min(2, "Please provide a valid location description")
  })
)

const { handleSubmit, isSubmitting, errors, resetForm } = useForm({
  validationSchema,
  initialValues: {
    name: '',
    location: ''
  }
})

const { value: name } = useField('name')
const { value: location } = useField('location')

const emit = defineEmits(['submitCinema'])
const onSubmit = handleSubmit((values) => {
  const cleanPayload = {
    name: values.name,
    location: values.location
  }
  
  emit('submitCinema', cleanPayload)
  resetForm()
})
</script>

<template>
  <div class="w-full max-w-lg bg-gray-950 border border-gray-900 rounded-3xl p-6 shadow-2xl text-white">
    <div class="mb-6">
      <h2 class="text-xl font-bold tracking-wide">Register New Cinema Room</h2>
      <p class="text-xs text-gray-500 mt-1">Add a physical auditorium or screen theater layout log to the network.</p>
    </div>

    <form @submit.prevent="onSubmit" class="space-y-5">
      
      <div class="flex flex-col gap-1.5">
        <label class="text-xs font-bold text-gray-400 uppercase tracking-wider">
          Cinema Hall Name <span class="text-red-500">*</span>
        </label>
        <input
          v-model="name"
          type="text"
          placeholder="e.g., Ultra VIP Room 4, IMAX Screen"
          :class="[
            'w-full bg-gray-900 border rounded-xl px-4 py-3 text-sm focus:outline-none transition-colors placeholder:text-gray-600',
            errors.name ? 'border-red-500 focus:border-red-500' : 'border-gray-800 focus:border-lime-400'
          ]"
        />
        <span v-if="errors.name" class="text-xs font-medium text-red-500 mt-0.5">
          {{ errors.name }}
        </span>
      </div>

      <div class="flex flex-col gap-1.5">
        <label class="text-xs font-bold text-gray-400 uppercase tracking-wider">
          Physical Location / Venue <span class="text-red-500">*</span>
        </label>
        <input
          v-model="location"
          type="text"
          placeholder="e.g., 3rd Floor, Downtown Mall Arcade"
          :class="[
            'w-full bg-gray-900 border rounded-xl px-4 py-3 text-sm focus:outline-none transition-colors placeholder:text-gray-600',
            errors.location ? 'border-red-500 focus:border-red-500' : 'border-gray-800 focus:border-lime-400'
          ]"
        />
        <span v-if="errors.location" class="text-xs font-medium text-red-500 mt-0.5">
          {{ errors.location }}
        </span>
      </div>

      <div class="pt-2">
        <button
          type="submit"
          :disabled="isSubmitting"
          class="w-full bg-lime-400 hover:bg-lime-300 disabled:bg-gray-800 disabled:text-gray-600 text-black font-bold text-sm py-3 px-4 rounded-xl shadow-lg hover:shadow-lime-400/10 transition-all duration-200 active:scale-[0.99] cursor-pointer"
        >
          {{ isSubmitting ? 'Inserting Room Node...' : 'Save Cinema Layout' }}
        </button>
      </div>

    </form>
  </div>
</template>