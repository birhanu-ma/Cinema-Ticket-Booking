<script setup>
import { computed, ref, watch } from "vue";
import { gql } from "graphql-tag";

const props = defineProps({
  movie: {
    type: Object,
    required: true,
  },
});

const { $apollo } = useNuxtApp();
const { loggedIn } = useAuth();

const selectedHall = ref("");
const selectedSeats = ref([]);

const isBooking = ref(false);
const bookingError = ref("");

const schedule = computed(() => {
  return props.movie.schedules?.[0];
});

const halls = computed(() => {
  if (!props.movie.schedules) {
    return [];
  }

  return props.movie.schedules.map((schedule) => schedule.hall);
});

const currentSchedule = computed(() => {
  return (
    props.movie.schedules?.find(
      (schedule) => schedule.hall.id === selectedHall.value,
    ) || schedule.value
  );
});

const seats = computed(() => {
  return currentSchedule.value?.schedule_seats ?? [];
});

const toggleSeat = (seat) => {
  if (!seat.is_available) {
    return;
  }

  const exists = selectedSeats.value.find((item) => item.id === seat.id);

  if (exists) {
    selectedSeats.value = selectedSeats.value.filter(
      (item) => item.id !== seat.id,
    );
  } else {
    selectedSeats.value.push(seat);
  }
};

const totalPrice = computed(() => {
  return selectedSeats.value.length * Number(currentSchedule.value?.price ?? 0);
});

const seatsPerPage = 12;
const currentPage = ref(0);

const totalPages = computed(() => {
  return Math.ceil(seats.value.length / seatsPerPage) || 1;
});

const paginatedSeats = computed(() => {
  const start = currentPage.value * seatsPerPage;
  return seats.value.slice(start, start + seatsPerPage);
});

const nextPage = () => {
  if (currentPage.value < totalPages.value - 1) {
    currentPage.value++;
  }
};

const prevPage = () => {
  if (currentPage.value > 0) {
    currentPage.value--;
  }
};

watch(selectedHall, () => {
  selectedSeats.value = [];
  currentPage.value = 0;
});

const CREATE_BOOKING = gql`
  mutation CreateBooking($schedule_id: uuid!, $schedule_seat_ids: [uuid!]!) {
    createBooking(
      schedule_id: $schedule_id
      schedule_seat_ids: $schedule_seat_ids
    ) {
      success
      message
      booking_reference
      ticket_ids
    }
  }
`;

const INITIATE_PAYMENT = gql`
  mutation InitiatePayment($booking_reference: String!) {
    initiatePayment(booking_reference: $booking_reference) {
      checkoutUrl
      txRef
    }
  }
`;

const buyTickets = async () => {
  bookingError.value = "";

  if (!loggedIn.value) {
    await navigateTo("/auth/signup");
    return;
  }

  if (selectedSeats.value.length === 0) {
    bookingError.value = "Please select at least one seat";
    return;
  }

  if (!currentSchedule.value?.id) {
    bookingError.value = "Please select a hall";
    return;
  }

  isBooking.value = true;

  try {
    const { data: bookingData } = await $apollo.mutate({
      mutation: CREATE_BOOKING,
      variables: {
        schedule_id: currentSchedule.value.id,
        schedule_seat_ids: selectedSeats.value.map((seat) => seat.id),
      },
    });

    if (!bookingData.createBooking.success) {
      bookingError.value = bookingData.createBooking.message;
      return;
    }

    const bookingReference = bookingData.createBooking.booking_reference;

    const { data: paymentData } = await $apollo.mutate({
      mutation: INITIATE_PAYMENT,
      variables: { booking_reference: bookingReference },
    });

    window.location.href = paymentData.initiatePayment.checkoutUrl;
  } catch (err) {
    console.error(err);

    const message =
      err?.graphQLErrors?.[0]?.message ||
      err?.networkError?.result?.errors?.[0]?.message ||
      err?.message ||
      "Booking failed, please try again";

    bookingError.value = message;
  } finally {
    isBooking.value = false;
  }
};
</script>

<template>
  <div class="bg-gray-950 rounded-3xl p-7 shadow-xl border border-gray-900">
    <h2 class="text-2xl font-bold text-white">
      {{ movie.title }}
    </h2>

    <div class="mt-4 space-y-3 text-gray-300">
      <div class="flex justify-between">
        <span> Cinema </span>

        <span class="text-white">
          {{ currentSchedule?.hall?.cinema?.name }}
        </span>
      </div>

      <div class="flex justify-between">
        <span> Hall </span>

        <div class="flex items-center gap-6">
          <select
            v-model="selectedHall"
            class="bg-gray-900 text-white cusor-pointer rounded-lg px-3 py-1"
          >
            <option v-for="hall in halls" :key="hall.id" :value="hall.id">
              {{ hall.name }}
            </option>
          </select>
        </div>
      </div>
      <div class="flex justify-between">
        <span> Date </span>

        <span class="text-white">
          {{
            new Date(currentSchedule.start_time).toLocaleString([], {
              day: "numeric",
              month: "short",
              hour: "2-digit",
              minute: "2-digit",
            })
          }}
        </span>
      </div>
    </div>

    <div class="mt-6">
      <h3 class="text-white font-bold mb-3">Select Seats</h3>

      <div class="flex items-center gap-3">
        <button
          @click="prevPage"
          :disabled="currentPage === 0"
          class="w-8 h-8 flex items-center justify-center cursor-pointer rounded-full bg-gray-800 text-white disabled:opacity-30 disabled:cursor-not-allowed"
        >
          ‹
        </button>

        <div class="grid grid-cols-6 grid-rows-2 gap-3 flex-1">
          <button
            v-for="seat in paginatedSeats"
            :key="seat.id"
            @click="toggleSeat(seat)"
            class="w-10 h-10 rounded-xl cursor-pointer text-xs font-bold"
            :class="
              selectedSeats.find((item) => item.id === seat.id)
                ? 'bg-lime-400 text-black'
                : seat.is_available
                  ? 'bg-gray-800 text-white'
                  : 'bg-gray-900 text-gray-600'
            "
          >
            {{ seat.id.slice(0, 2) }}
          </button>
        </div>

        <button
          @click="nextPage"
          :disabled="currentPage === totalPages - 1"
          class="w-8 h-8 flex items-center justify-center cursor-pointer rounded-full bg-gray-800 text-white disabled:opacity-30 disabled:cursor-not-allowed"
        >
          ›
        </button>
      </div>

      <div class="flex justify-center gap-1 mt-3">
        <span
          v-for="page in totalPages"
          :key="page"
          class="w-1.5 h-1.5 rounded-full"
          :class="page - 1 === currentPage ? 'bg-lime-400' : 'bg-gray-700'"
        ></span>
      </div>
    </div>

    <div class="mt-8 flex justify-between items-center">
      <div>
        <p class="text-gray-400">Seats:</p>

        <p class="text-lime-400 font-bold">
          {{ selectedSeats.length || "None" }}
        </p>
      </div>

      <div class="text-right">
        <p class="text-gray-400">Total</p>

        <p class="text-3xl font-bold text-white">{{ totalPrice }} ETB</p>
      </div>
    </div>

    <p v-if="bookingError" class="mt-3 text-red-400 text-sm text-center">
      {{ bookingError }}
    </p>

    <button
      @click="buyTickets"
      :disabled="isBooking || selectedSeats.length === 0"
      class="mt-6 w-full cursor-pointer bg-lime-400 text-black font-bold py-3 rounded-full hover:bg-lime-500 disabled:opacity-50 disabled:cursor-not-allowed"
    >
      {{ isBooking ? "Processing..." : "Buy Tickets" }}
    </button>

    <div class="flex justify-between pt-8">
      <span> Book Mark </span>

      <div class="flex items-center gap-6">
        <UiBookMarkButton
          v-if="currentSchedule?.id"
          :schedule-id="currentSchedule.id"
        />
      </div>
    </div>
  </div>
</template>