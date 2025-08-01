<script setup lang="ts">
import { computed, type PropType } from 'vue'
import type { ChatInfo } from '@/stores/chat.ts'

const props = defineProps({
  chat: {
    type: Object as PropType<ChatInfo>,
    default: () => undefined,
  },
})

const avatarCharacter = computed<string>(() => {
  if (props.chat?.name && props.chat.name.length > 0) {
    return props.chat.name.substring(0, 1)
  }
  return '-'
})
</script>

<template>
  <div
    v-if="chat"
    class="flex items-center gap-2 p-4 cursor-pointer hover:bg-neutral transition-all"
  >
    <div>
      <Avatar :label="avatarCharacter" class="mr-2" size="large" shape="circle" />
    </div>
    <div class="w-full">
      <div class="flex justify-between items-center">
        <div class="font-medium">{{ chat.name }}</div>
        <div class="text-sm">{{ chat.lastMessageTime || '' }}</div>
      </div>
      <div class="flex justify-between items-center">
        <div class="text-sm">{{ chat.lastMessage || '-' }}</div>
      </div>
    </div>
  </div>
</template>

<style scoped></style>
