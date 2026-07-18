<script setup>
import { gql } from "@apollo/client/core";
import { useRouter } from "vue-router";

definePageMeta({
  layout: "admin",
});

const { $apollo } = useNuxtApp();

const router = useRouter();

const INSERT_TICKET = gql`
  mutation CreateTicket($ticket: Object!, $payment: Object!) {
    insert_tickets_one(object: $ticket) {
      id
    }
  }
`;

const handleTicketSubmit = async (cleanPayload) => {
  try {
    await $apollo.mutate({
      mutation: INSERT_TICKET,

      variables: {
        ticket: {
          schedule_id: cleanPayload.schedule_id,

          user_id: cleanPayload.user_id,

          status: "pending",
        },

        payment: {
          status: "pending",
        },
      },
    });

    router.push("/admin/tickets");
  } catch (err) {
    console.error("Ticket creation failed", err);
  }
};
</script>

<template>
  <div
    class="h-screen w-full flex flex-col items-center justify-center p-8 gap-6 bg-linear-to-t from-[#51751f] to-transparent text-white"
  >
    <div class="w-full max-w-lg flex flex-col gap-1 text-left">
      <NuxtLink
        to="/admin/tickets"
        class="text-xs text-lime-400 hover:underline w-fit"
      >
        ← Back to Ticket Logs
      </NuxtLink>

      <h1 class="text-2xl font-bold tracking-wide mt-2">Add New Ticket</h1>
    </div>

    <AdminTicketsTicketForm @submit-ticket="handleTicketSubmit" />
  </div>
</template>
