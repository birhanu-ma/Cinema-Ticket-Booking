<script setup>
defineProps({
  seat: {
    type: Object,
    required: true,
  },
});
</script>

<template>
  <div
    class="flex bg-gray-950 border border-gray-900 rounded-3xl overflow-hidden shadow-2xl w-full group hover:border-gray-800 transition-all duration-300"
  >
    <div
      class="w-40 sm:w-44 bg-gray-900 flex items-center justify-center text-5xl"
    >
      🪑
    </div>

    <AdminSeatsSeatDetail
      :seat-id="seat.id"
      :seat-number="seat.seat_number"
      :row="seat.row_name"
      :type="seat.type"
      :hall="seat.cinema_hall?.name"
      :cinema="seat.cinema_hall?.cinema?.name"
    />
  </div>
</template>
