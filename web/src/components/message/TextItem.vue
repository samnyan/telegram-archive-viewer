<script setup lang="ts">
import { type PropType, ref } from 'vue'
import type { TelegramTextEntity } from '@/client'

const props = defineProps({
  text: {
    type: Object as PropType<TelegramTextEntity>,
    default: () => undefined,
  },
})

const hidden = ref(true)
</script>

<template>
  <span
    v-if="text"
    class="whitespace-pre-line"
    :class="{
      'font-bold': text.type == 'bold',
      'font-italic': text.type == 'italic',
      underline: text.type == 'underline',
      'line-through': text.type == 'strikethrough',
    }"
  >
    <span
      v-if="text.type == 'spoiler'"
      class="spoiler"
      :class="{
        'spoiler-hidden': hidden,
      }"
      @click="() => (hidden = false)"
    >
      <span :class="{ 'opacity-0': hidden }">
        {{ text.text }}
      </span>
    </span>
    <span v-else-if="text.type == 'pre'" class="pre">
      {{ text.text }}
    </span>
    <span v-else-if="text.type == 'blockquote'" class="blockquote">
      {{ text.text }}
    </span>
    <template v-else>
      {{ text.text }}
    </template>
  </span>
</template>

<style scoped></style>
