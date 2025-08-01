<script setup lang="ts">
import type { TelegramResult } from '@/client'
import type { PropType } from 'vue'
import TextItem from '@/components/message/TextItem.vue'
import { $filters } from '@/utils/filters.ts'

type TelegramMessageType = TelegramResult['messages'][number]

const props = defineProps({
  message: {
    type: Object as PropType<TelegramMessageType>,
    default: () => undefined,
  },
})
</script>

<template>
  <div
    class="message box-border m-2 py-2 px-3 w-fit flex flex-col border-rounded-lg dark:bg-neutral-800 bg-neutral-100"
  >
    <div class="from_name">
      {{ message.from }}
    </div>
    <div v-if="message.text_entities" class="text">
      <TextItem v-for="(i, idx) in message.text_entities" :key="idx" :text="i" />
    </div>

    <div class="ml-a opacity-50 text-sm" :title="$filters.formatDate(message.date)">
      {{ $filters.formatTimeOnly(message.date) }}
    </div>
  </div>
</template>

<style scoped></style>
