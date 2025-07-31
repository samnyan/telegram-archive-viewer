<script setup lang="ts">
import type { TelegramResult } from '@/client'
import type { PropType } from 'vue'
import TextItem from '@/components/chat/TextItem.vue'

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
    class="box-border message m-2 p-2 w-fit border-rounded-lg dark:bg-neutral-800 bg-neutral-100"
  >
    <div class="from_name">
      {{ message.from }}
    </div>
    <div v-if="message.text_entities" class="text">
      <TextItem v-for="(i, idx) in message.text_entities" :key="idx" :text="i" />
    </div>

    <div>
      {{ message.date }}
    </div>
  </div>
</template>

<style scoped></style>
