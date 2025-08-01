<script setup lang="ts">
import MessageList from '@/components/message/MessageList.vue'
import { ref } from 'vue'
import type { TelegramResult } from '@/client'

const currentChat = ref<TelegramResult>()
</script>

<template>
  <MessageList v-if="currentChat" :chat="currentChat" />
</template>

<style scoped></style>
