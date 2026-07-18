<script setup>
defineProps({
  ticket: {
    type: Object,
    required: true,
  },
});

const formatDate = (date) => {
  if (!date) return "-";

  return new Date(date).toLocaleString();
};
</script>

<template>
  <NuxtLink
    :to="`/admin/tickets/${ticket.id}`"
    class="flex-1 p-5 flex flex-col justify-between hover:bg-gray-900/40 rounded-r-3xl transition-all duration-300 group block cursor-pointer"
  >
    <div>
      <div class="flex items-center justify-between mb-2">
        <div class="flex items-center gap-2">
          <span class="text-xs font-mono text-gray-500">
            #{{ ticket.booking_reference || ticket.id }}
          </span>

          <span
            class="text-[10px] bg-red-500/10 text-red-400 font-bold px-2 py-0.5 rounded uppercase tracking-wider"
          >
            High
          </span>
        </div>

        <button
          @click.stop.prevent="console.log('Ticket action menu triggered')"
          class="text-gray-500 hover:text-white transition-colors text-lg p-1 cursor-pointer"
        >
          ⋮
        </button>
      </div>

      <h3
        class="text-white text-lg font-bold leading-snug tracking-wide group-hover:text-lime-400 transition-colors"
      >
        {{
          ticket.movie_title ||
          ticket.schedule?.movie?.title ||
          "Unknown Movie"
        }}
      </h3>
    </div>

    <div class="w-full h-px bg-gray-900 my-4"></div>

    <div class="flex items-center justify-between gap-6">
      <div class="flex flex-col gap-1">
        <span
          class="text-[10px] text-gray-500 font-bold uppercase tracking-wider"
        >
          Status
        </span>

        <div class="flex items-center gap-1.5">
          <span
            class="w-2 h-2 rounded-full bg-amber-500 animate-pulse"
          ></span>

          <span class="text-amber-500 text-xs font-bold">
            {{ ticket.status }}
          </span>
        </div>
      </div>

      <div class="flex flex-col gap-1 text-right">
        <span
          class="text-[10px] text-gray-500 font-bold uppercase tracking-wider"
        >
          Show Date
        </span>

        <div
          class="flex items-center gap-1.5 justify-end text-gray-300 text-xs font-semibold"
        >
          <span>📅</span>

          <span>
            {{
              formatDate(
                ticket.schedule_time ||
                ticket.schedule?.start_time
              )
            }}
          </span>
        </div>
      </div>
    </div>
  </NuxtLink>
</template>