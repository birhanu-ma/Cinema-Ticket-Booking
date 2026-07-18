<script setup>
import { ref, onMounted, onUnmounted, computed } from "vue";
import { gql } from "graphql-tag";
import { useRoute } from "#app";

const route = useRoute();
const { $apollo } = useNuxtApp();

const status = ref("checking");
const payment = ref(null);
const tickets = ref([]);

let pollTimer = null;
let attempts = 0;
const maxAttempts = 15;

const txRef = route.query.tx_ref || route.query.trx_ref || "";

const bookingRefFromRoute = route.query.booking_reference || "";

const GET_PAYMENT_STATUS = gql`
  query GetPaymentStatus($tx_ref: String!) {
    payments(where: { transaction_ref: { _eq: $tx_ref } }) {
      id
      transaction_ref
      status
      booking_reference
      amount
      payment_method
      currency
      created_at

      user {
        name
      }
    }
  }
`;

const GET_TICKETS_BY_BOOKING = gql`
  query GetTickets($booking_reference: String!) {
    tickets(where: { booking_reference: { _eq: $booking_reference } }) {
      id
      booking_reference
      status

      user {
        name
      }

      schedule_seat {
        seat {
          row_name
          seat_number
        }
      }

      schedule {
        id
        start_time

        movie {
          id
          title
          duration

          movie_images(order_by: { is_featured: desc }, limit: 1) {
            image_url
          }
        }

        hall {
          id
          name

          cinema {
            id
            name
          }
        }
      }
    }
  }
`;

const stopPolling = () => {
  if (pollTimer) {
    clearInterval(pollTimer);
    pollTimer = null;
  }
};

const fetchTickets = async (bookingReference) => {
  const { data: ticketData } = await $apollo.query({
    query: GET_TICKETS_BY_BOOKING,
    variables: {
      booking_reference: bookingReference,
    },
    fetchPolicy: "network-only",
  });

  tickets.value = ticketData.tickets;
};

const checkStatus = async () => {
  if (!txRef) {
    status.value = "failed";

    stopPolling();

    return;
  }

  try {
    const { data } = await $apollo.query({
      query: GET_PAYMENT_STATUS,

      variables: {
        tx_ref: txRef,
      },

      fetchPolicy: "network-only",
    });

    const row = data.payments?.[0];

    if (!row) {
      console.warn("Payment row not found.");

      attempts++;

      if (attempts >= maxAttempts) {
        status.value = "failed";

        stopPolling();
      }

      return;
    }

    payment.value = row;

    if (row.status === "success") {
      console.log("Payment is SUCCESS.");

      await fetchTickets(row.booking_reference);

      status.value = "confirmed";

      stopPolling();

      return;
    }

    if (row.status === "failed") {
      status.value = "failed";

      stopPolling();

      return;
    }

    attempts++;

    if (attempts >= maxAttempts) {
      status.value = "timeout";

      stopPolling();
    }
  } catch (err) {
    attempts++;

    if (attempts >= maxAttempts) {
      status.value = "timeout";

      stopPolling();
    }
  }
};

const formattedShowtime = (isoString) => {
  if (!isoString) return "";

  return new Date(isoString).toLocaleString(undefined, {
    weekday: "short",

    month: "short",

    day: "numeric",

    hour: "numeric",

    minute: "2-digit",
  });
};

const formattedAmount = computed(() => {
  if (!payment.value) return "";

  return `${payment.value.amount} ${payment.value.currency}`;
});

const seatLabel = (ticket) => {
  const seat = ticket.schedule_seat?.seat;

  if (!seat) return null;

  return `${seat.row_name}${seat.seat_number}`;
};

const ticketStatusClasses = (ticketStatus) => {
  if (ticketStatus === "confirmed") {
    return "bg-lime-400 text-black";
  }

  if (ticketStatus === "failed") {
    return "bg-red-500 text-white";
  }

  return "bg-gray-700 text-lime-300";
};

onMounted(() => {
  console.log("Success page mounted.");

  if (bookingRefFromRoute) {
    fetchTickets(bookingRefFromRoute).catch((err) => {
      console.log("Initial ticket prefetch failed:", err);
    });
  }

  checkStatus();

  pollTimer = setInterval(checkStatus, 2000);
});

onUnmounted(() => {
  stopPolling();
});
</script>
<template>
  <div class="min-h-screen bg-lime-50 px-4 py-12">
    <div class="mw-250 h-150">
      <div v-if="status === 'checking'" class="text-center py-12">
        <div
          class="w-9 h-9 mx-auto mb-4 rounded-full border-4 border-lime-200 border-t-lime-500 animate-spin"
        />

        <h2 class="text-xl font-semibold text-lime-900">
          Confirming your payment&hellip;
        </h2>

        <p class="text-lime-700 mt-2">
          This usually takes a few seconds. Please don't close this page.
        </p>
      </div>

      <div v-else-if="status === 'failed'" class="text-center py-12">
        <div
          class="w-16 h-16 bg-red-500 rounded-full flex items-center justify-center mx-auto mb-5"
        >
          <span class="text-white text-3xl font-bold"> ✕ </span>
        </div>

        <h2 class="text-xl font-semibold text-lime-900">Payment failed</h2>

        <p class="text-lime-700 mt-2">
          We couldn't confirm this payment. If you were charged, please contact
          support with this reference:
        </p>

        <code
          v-if="txRef"
          class="inline-block mt-3 px-3 py-1.5 bg-lime-100 border border-lime-200 rounded-md text-sm text-lime-800"
        >
          {{ txRef }}
        </code>

        <div class="mt-6">
          <NuxtLink
            to="/"
            class="inline-block px-5 py-2.5 bg-lime-500 text-white font-bold rounded-full text-sm hover:bg-lime-600"
          >
            Back to home
          </NuxtLink>
        </div>
      </div>

      <div v-else-if="status === 'timeout'" class="text-center py-12">
        <h2 class="text-xl font-semibold text-lime-900">Still processing</h2>

        <p class="text-lime-700 mt-2">
          Your payment is taking longer than expected to confirm. It may still
          go through.
        </p>

        <code
          v-if="txRef"
          class="inline-block mt-3 px-3 py-1.5 bg-lime-100 border border-lime-200 rounded-md text-sm text-lime-800"
        >
          {{ txRef }}
        </code>

        <div class="mt-6">
          <NuxtLink
            to="/"
            class="inline-block px-5 py-2.5 bg-lime-500 text-white font-bold rounded-full text-sm hover:bg-lime-600"
          >
            Back to home
          </NuxtLink>
        </div>
      </div>

      <div v-else-if="status === 'confirmed'">
        <div class="text-center mb-10">
          <div
            class="w-16 h-16 bg-lime-500 rounded-full flex items-center justify-center mx-auto mb-4"
          >
            <span class="text-white text-3xl font-bold"> ✓ </span>
          </div>

          <h2 class="text-2xl font-bold text-lime-900">Booking confirmed</h2>

          <p v-if="payment" class="text-lime-700 mt-2">
            Paid

            <span class="font-semibold text-lime-600">
              {{ formattedAmount }}
            </span>
          </p>
        </div>

        <div class="flex justify-center items-center mb-10">
          <div
            v-for="ticket in tickets"
            :key="ticket.id"
            class="relative flex w-full max-w-3xl bg-white rounded-2xl shadow-sm overflow-hidden border border-lime-100"
          >
            <div class="w-28 shrink-0 bg-lime-100 relative">
              <img
                v-if="ticket.schedule?.movie?.movie_images?.[0]?.image_url"
                :src="ticket.schedule.movie.movie_images[0].image_url"
                :alt="ticket.schedule.movie.title"
                class="w-full h-full object-cover opacity-90"
              />

              <div
                v-else
                class="w-full h-full flex items-center justify-center text-lime-400 text-xs px-2 text-center"
              >
                No poster
              </div>
            </div>

            <div class="relative w-0 border-l-2 border-dashed border-lime-200">
              <span
                class="absolute -top-3 -left-3 w-6 h-6 rounded-full bg-lime-50"
              />

              <span
                class="absolute -bottom-3 -left-3 w-6 h-6 rounded-full bg-lime-50"
              />
            </div>

            <div class="flex-1 p-4 flex flex-col justify-between">
              <div>
                <div class="flex items-start justify-between gap-2">
                  <h3
                    class="text-base font-semibold leading-tight text-lime-950"
                  >
                    {{ ticket.schedule?.movie?.title || "Movie" }}
                  </h3>

                  <span
                    class="shrink-0 px-2.5 py-1 rounded-full text-xs font-bold capitalize"
                    :class="ticketStatusClasses(ticket.status)"
                  >
                    {{ ticket.status }}
                  </span>
                </div>

                <p class="text-sm font-medium text-lime-800 mt-2">
                  Issued for:
                  {{ ticket.user?.name || "Guest" }}
                </p>

                <p class="text-sm text-lime-700 mt-1">
                  {{ ticket.schedule?.hall?.cinema?.name }}
                  &middot;
                  {{ ticket.schedule?.hall?.name }}
                </p>

                <p class="text-sm text-lime-700">
                  {{ formattedShowtime(ticket.schedule?.start_time) }}
                </p>
              </div>

              <div class="mt-4 pt-3 border-t border-dashed border-lime-200">
                <div class="flex items-end justify-between">
                  <div>
                    <p
                      class="text-[10px] uppercase tracking-wide text-lime-500"
                    >
                      Seat
                    </p>

                    <p class="text-lg font-bold font-mono text-lime-600">
                      {{ seatLabel(ticket) || "—" }}
                    </p>
                  </div>

                  <div class="text-right">
                    <p
                      class="text-[10px] uppercase tracking-wide text-lime-500"
                    >
                      Price
                    </p>

                    <p class="text-lg font-bold text-lime-700">
                      {{ formattedAmount }}
                    </p>
                  </div>
                </div>

                <div class="mt-3">
                  <p class="text-[10px] uppercase tracking-wide text-lime-500">
                    Reference
                  </p>

                  <p class="text-xs font-mono text-lime-700 break-all">
                    {{ payment?.transaction_ref }}
                  </p>
                </div>
              </div>
            </div>
          </div>
        </div>

        <div class="text-center">
          <NuxtLink
            to="/"
            class="inline-block px-5 py-2.5 bg-lime-500 text-white font-bold rounded-full text-sm hover:bg-lime-600"
          >
            Back to home
          </NuxtLink>
        </div>
      </div>
    </div>
  </div>
</template>
