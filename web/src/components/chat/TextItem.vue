<script setup lang="ts">
import type { PropType } from 'vue'
import type { TelegramTextEntity } from '@/client'

const props = defineProps({
  text: {
    type: Object as PropType<TelegramTextEntity>,
    default: () => undefined,
  },
})
</script>

<template>
  <div
    v-if="text"
    class="inline-block whitespace-pre-wrap"
    :class="{
      'font-bold': text.type == 'bold',
      'font-italic': text.type == 'italic',
      underline: text.type == 'underline',
      'line-through': text.type == 'strikethrough',
    }"
  >
    {{ text.text }}
  </div>
</template>

<style scoped></style>
