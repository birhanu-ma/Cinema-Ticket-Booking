<script setup>
const handleCreate = (resourceType) => {
  console.log(`Opening creation modal/drawer for: ${resourceType}`);
};
</script>

<template>
  <div class="p-8 bg-gray-50">
    <h1 class="text-2xl font-bold mb-6 text-gray-900">Admin Control Catalog</h1>

    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
      <AdminSharedAddCard
        subheading="Reservations"
        title="."
        button-text="Create Cinema hall"
        theme="green"
        @create="handleCreate('cinemahall')"
        :to="`/admin/tickets/form/${ticketId}`"
      />
    </div>
  </div>
</template>
