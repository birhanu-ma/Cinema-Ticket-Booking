<script setup>
defineProps({
  ticket: {
    type: Object,
    required: true
  }
})
</script>

<template>
  <div class="flex bg-gray-950 border border-gray-900 rounded-3xl overflow-hidden shadow-2xl  group hover:border-gray-800 transition-all duration-300">
    <div class="w-40 sm:w-48 relative flex-shrink-0 overflow-hidden bg-gray-900">
      <img
        :src="ticket.thumbnail"
        class="w-full h-full object-cover group-hover:scale-105 transition-transform duration-500"
        alt="Movie Thumbnail Poster"
      />
      <div class="absolute inset-0 bg-gradient-to-r from-transparent to-gray-950/20"></div>
    </div>

    <AdminTicketsTicketStatusDetail
      :title="ticket.title"
      :ticket-id="ticket.ticketId"
      :status="ticket.status"
      :completion-date="ticket.completionDate"
    />
  </div>
</template>