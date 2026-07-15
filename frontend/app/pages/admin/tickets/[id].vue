<script setup>
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'

definePageMeta({
  layout: "admin",
});

const route = useRoute()
const ticketId = route.params.id

const isLoading = ref(false)
const ticket = ref(null)

onMounted(async () => {
  isLoading.value = true
  try {
    await new Promise(resolve => setTimeout(resolve, 450))

    ticket.value = {
      id: ticketId,
      user_id: "77777777-8888-9999-0000-aaaaaaaaaaaa",
      user_name: "Alex Mercer",
      user_email: "alex.mercer@gmail.com",
      schedule_id: "3fa85f64-5717-4562-b3fc-2c963f66afa6",
      movie_title: "Shazam! Fury of the Gods",
      status: "Confirmed",
      seat_number: "G-14",
      booking_reference: "TX-99081",
      purchased_at: "July 14, 2026, 3:30 PM"
    }
  } catch (err) {
    console.error("Failed downloading ticket profile:", err)
  } finally {
    isLoading.value = false
  }
})

const getStatusClasses = (status) => {
  const norm = (status || '').toLowerCase()
  if (norm === 'confirmed') return 'bg-lime-400/10 text-lime-400'
  if (norm === 'pending') return 'bg-yellow-500/10 text-yellow-400'
  return 'bg-red-500/10 text-red-400'
}
</script>

<template>
  <div class="min-h-screen w-full bg-gray-900 text-white p-8 flex flex-col items-center">
    <div class="w-full max-w-2xl space-y-6">
      
      <div class="flex items-center justify-between">
        <NuxtLink to="/admin/tickets" class="text-xs text-lime-400 hover:underline flex items-center gap-1 w-fit">
          ← Back to Ticket Matrix
        </NuxtLink>
        
        <NuxtLink 
          v-if="ticket"
          :to="`/admin/tickets/form/${ticket.id}`" 
          class="bg-gray-800 hover:bg-gray-750 text-xs font-semibold px-4 py-2 rounded-xl border border-gray-700 transition-colors cursor-pointer"
        >
          Edit Reservation
        </NuxtLink>
      </div>

      <div v-if="isLoading" class="w-full text-center py-24 animate-pulse text-gray-500 text-sm">
        Retrieving transactional seat records...
      </div>

      <div v-else-if="ticket" class="bg-gray-950 border border-gray-900 rounded-3xl p-6 md:p-8 shadow-2xl relative overflow-hidden">
        <div class="absolute top-0 left-0 right-0 h-1 bg-lime-400"></div>

        <div class="flex flex-col gap-6">
          <div class="flex flex-wrap items-center justify-between gap-4">
            <div>
              <span class="text-[10px] bg-gray-800 text-gray-400 font-bold px-2.5 py-0.5 rounded-md uppercase tracking-wider block w-fit mb-1">
                Ref: {{ ticket.booking_reference }}
              </span>
              <h1 class="text-2xl font-extrabold tracking-wide text-gray-100">Booking Details</h1>
            </div>
            <span :class="['text-[10px] font-bold px-3 py-1 rounded-md uppercase tracking-wider block w-fit', getStatusClasses(ticket.status)]">
              {{ ticket.status }}
            </span>
          </div>

          <div class="w-full h-px bg-gray-900"></div>

          <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
            <div class="space-y-1">
              <span class="text-[10px] text-gray-500 font-bold uppercase tracking-wider block">Assigned Film</span>
              <p class="text-lime-400 text-sm font-semibold">{{ ticket.movie_title }}</p>
              <span class="text-[10px] text-gray-600 font-mono block">ID: {{ ticket.schedule_id }}</span>
            </div>

            <div class="space-y-1">
              <span class="text-[10px] text-gray-500 font-bold uppercase tracking-wider block">Customer Account</span>
              <p class="text-gray-200 text-sm font-semibold">{{ ticket.user_name }}</p>
              <p class="text-gray-400 text-xs">{{ ticket.user_email }}</p>
            </div>

            <div class="space-y-1">
              <span class="text-[10px] text-gray-500 font-bold uppercase tracking-wider block">Assigned Seat / Row</span>
              <p class="text-white text-base font-extrabold tracking-wider">{{ ticket.seat_number }}</p>
            </div>

            <div class="space-y-1">
              <span class="text-[10px] text-gray-500 font-bold uppercase tracking-wider block">Transaction Registered</span>
              <p class="text-gray-300 text-xs font-semibold">{{ ticket.purchased_at }}</p>
            </div>
          </div>

          <div class="w-full h-px bg-gray-900"></div>

          <div class="flex flex-col items-center justify-center py-4 bg-gray-900/40 rounded-2xl border border-gray-900">
            <div class="text-2xl tracking-[0.45em] font-mono text-gray-400 font-light select-none">
              ||||| | | |||| |||| || |
            </div>
            <span class="text-[9px] text-gray-500 tracking-wider font-mono mt-1">UUID: {{ ticket.id }}</span>
          </div>
        </div>
      </div>

      <div v-else class="w-full text-center py-24 text-gray-600 text-sm">
        Ticket record entry not located.
      </div>

    </div>
  </div>
</template>