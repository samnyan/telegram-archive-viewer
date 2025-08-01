<script setup lang="ts">
import MessageList from '@/components/message/MessageList.vue'
import { useChatStore } from '@/stores/chat.ts'

const chatStore = useChatStore()
</script>

<template>
  <MessageList v-if="chatStore.currentChat" :chat="chatStore.currentChat" />
</template>

<style scoped></style>
