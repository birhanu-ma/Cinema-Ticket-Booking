<script setup>
import { ref, onMounted } from "vue";
import { useRoute } from "vue-router";
import { gql } from "@apollo/client/core";

definePageMeta({
  layout: "admin",
});

const { $apollo } = useNuxtApp();

const route = useRoute();

const ticketId = route.params.id;

const isLoading = ref(false);
const ticket = ref(null);

const GET_TICKET_DETAILS = gql`
  query GetTicketDetails($id: uuid!) {
    tickets_by_pk(id: $id) {
      id
      booking_reference
      status
      created_at

      user {
        id
        name
        email
      }

      schedule_seat {
        id
        seat {
          id
          row_name
          seat_number
        }
      }

      schedule {
        id
        price
        start_time

        hall {
          id
          name

          cinema {
            id
            name
            location
          }
        }

        movie {
          id
          title

          movie_images(where: { is_featured: { _eq: true } }, limit: 1) {
            image_url
          }
        }
      }

      payments(limit: 1, order_by: { created_at: desc }) {
        id
        transaction_ref
        status
        created_at
      }
    }
  }
`;

onMounted(async () => {
  if (!ticketId) {
    console.error("Missing Ticket ID");
    return;
  }

  isLoading.value = true;

  try {
    const { data } = await $apollo.query({
      query: GET_TICKET_DETAILS,
      variables: {
        id: ticketId,
      },
      fetchPolicy: "network-only",
    });

    const result = data?.tickets_by_pk;

    if (!result) {
      ticket.value = null;
      return;
    }

    const payment = result.payments?.[0];

    const seat = result.schedule_seat?.seat;

    const seatLabel = seat
      ? `${seat.row_name ?? ""}${seat.seat_number ?? ""}`.trim() ||
        "Not Assigned"
      : "Not Assigned";

    ticket.value = {
      id: result.id,

      booking_reference: result.booking_reference || "No Reference",

      status: result.status || "Pending",

      seat_number: seatLabel,

      movie_title: result.schedule?.movie?.title || "Unknown Movie",

      movie_image: result.schedule?.movie?.movie_images?.[0]?.image_url || null,

      cinema_name: result.schedule?.hall?.cinema?.name || "Unknown Cinema",

      cinema_location: result.schedule?.hall?.cinema?.location || "",

      show_time: result.schedule?.start_time,

      ticket_price: result.schedule?.price ?? 0,

      customer_name: result.user?.name || "Guest Customer",

      customer_email: result.user?.email || "No Email",

      payment_status: payment?.status || "Unpaid",

      transaction_reference: payment?.transaction_ref || "No Transaction",

      booking_date: result.created_at,
    };

    console.log("Ticket Detail:", ticket.value);
  } catch (err) {
    console.error("Failed loading ticket details:", err);

    ticket.value = null;
  } finally {
    isLoading.value = false;
  }
});

const formatDate = (date) => {
  if (!date) return "-";

  return new Date(date).toLocaleString("en-US", {
    dateStyle: "medium",
    timeStyle: "short",
  });
};

const getStatusClasses = (status) => {
  const value = (status || "").toLowerCase();

  if (value === "confirmed") {
    return "bg-lime-400/10 text-lime-400";
  }

  if (value === "pending") {
    return "bg-yellow-500/10 text-yellow-400";
  }

  return "bg-red-500/10 text-red-400";
};
</script>

<template>
  <div
    class="min-h-screen w-full bg-linear-to-t from-[#51751f] to-transparent text-white p-8 flex flex-col items-center"
  >
    <div class="w-full max-w-2xl space-y-6">
      <div class="flex items-center justify-between">
        <NuxtLink
          to="/admin/tickets"
          class="text-xs text-lime-400 hover:underline flex items-center gap-1 w-fit"
        >
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

      <div
        v-if="isLoading"
        class="w-full text-center py-24 animate-pulse text-gray-500 text-sm"
      >
        Retrieving transactional seat records...
      </div>

      <div
        v-else-if="ticket"
        class="bg-gray-950 border border-gray-900 rounded-3xl p-6 md:p-8 shadow-2xl relative overflow-hidden"
      >
        <div class="absolute top-0 left-0 right-0 h-1 bg-lime-400"></div>

        <div class="flex flex-col gap-6">
          <div class="flex flex-wrap items-center justify-between gap-4">
            <div>
              <span
                class="text-[10px] bg-gray-800 text-gray-400 font-bold px-2.5 py-0.5 rounded-md uppercase tracking-wider block w-fit mb-1"
              >
                Ref: {{ ticket.booking_reference }}
              </span>

              <h1 class="text-2xl font-extrabold tracking-wide text-gray-100">
                Booking Details
              </h1>
            </div>

            <span
              :class="[
                'text-[10px] font-bold px-3 py-1 rounded-md uppercase tracking-wider block w-fit',
                getStatusClasses(ticket.status),
              ]"
            >
              {{ ticket.status }}
            </span>
          </div>

          <div class="w-full h-px bg-gray-900"></div>

          <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
            <div class="space-y-1">
              <span
                class="text-[10px] text-gray-500 font-bold uppercase tracking-wider block"
              >
                Movie
              </span>

              <p class="text-lime-400 text-sm font-semibold">
                {{ ticket.movie_title }}
              </p>
            </div>

            <div class="space-y-1">
              <span
                class="text-[10px] text-gray-500 font-bold uppercase tracking-wider block"
              >
                Cinema
              </span>

              <p class="text-gray-200 text-sm font-semibold">
                {{ ticket.cinema_name }}
              </p>

              <p class="text-gray-400 text-xs">
                {{ ticket.cinema_location }}
              </p>
            </div>

            <div class="space-y-1">
              <span
                class="text-[10px] text-gray-500 font-bold uppercase tracking-wider block"
              >
                Customer Account
              </span>

              <p class="text-gray-200 text-sm font-semibold">
                {{ ticket.customer_name }}
              </p>

              <p class="text-gray-400 text-xs">
                {{ ticket.customer_email }}
              </p>
            </div>

            <div class="space-y-1">
              <span
                class="text-[10px] text-gray-500 font-bold uppercase tracking-wider block"
              >
                Assigned Seat / Row
              </span>

              <p class="text-white text-base font-extrabold tracking-wider">
                {{ ticket.seat_number }}
              </p>
            </div>

            <div class="space-y-1">
              <span
                class="text-[10px] text-gray-500 font-bold uppercase tracking-wider block"
              >
                Show Date & Time
              </span>

              <p class="text-gray-300 text-xs font-semibold">
                {{ formatDate(ticket.show_time) }}
              </p>
            </div>

            <div class="space-y-1">
              <span
                class="text-[10px] text-gray-500 font-bold uppercase tracking-wider block"
              >
                Ticket Price
              </span>

              <p class="text-gray-300 text-sm font-semibold">
                {{ ticket.ticket_price }} ETB
              </p>
            </div>

            <div class="space-y-1">
              <span
                class="text-[10px] text-gray-500 font-bold uppercase tracking-wider block"
              >
                Payment Status
              </span>

              <p class="text-gray-300 text-sm font-semibold">
                {{ ticket.payment_status }}
              </p>
            </div>

            <div class="space-y-1">
              <span
                class="text-[10px] text-gray-500 font-bold uppercase tracking-wider block"
              >
                Transaction Reference
              </span>

              <p class="text-gray-300 text-xs font-mono">
                {{ ticket.transaction_reference }}
              </p>
            </div>

            <div class="space-y-1">
              <span
                class="text-[10px] text-gray-500 font-bold uppercase tracking-wider block"
              >
                Booking Date
              </span>

              <p class="text-gray-300 text-xs font-semibold">
                {{ formatDate(ticket.booking_date) }}
              </p>
            </div>
          </div>

          <div class="w-full h-px bg-gray-900"></div>

          <div
            class="flex flex-col items-center justify-center py-4 bg-gray-900/40 rounded-2xl border border-gray-900"
          >
            <div
              class="text-2xl tracking-[0.45em] font-mono text-gray-400 font-light select-none"
            >
              ||||| | | |||| |||| || |
            </div>

            <span
              class="text-[9px] text-gray-500 tracking-wider font-mono mt-1"
            >
              Ticket ID: {{ ticket.id }}
            </span>
          </div>
        </div>
      </div>

      <div v-else class="w-full text-center py-24 text-gray-600 text-sm">
        Ticket record entry not located.
      </div>
    </div>
  </div>
</template>
