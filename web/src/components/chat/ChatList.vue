<script setup lang="ts">
import ChatToolbar from '@/components/chat/ChatToolbar.vue'
import ChatItem from '@/components/chat/ChatItem.vue'
import { type ChatInfo, useChatStore } from '@/stores/chat.ts'
import { storeToRefs } from 'pinia'
import { useRouter } from 'vue-router'

const router = useRouter()

const chatStore = useChatStore()
const { chatInfoList } = storeToRefs(chatStore)

const onChatClick = (c: ChatInfo) => {
  router.push({
    name: 'message',
  })
  chatStore.setCurrentChat(c.id)
}
</script>

<template>
  <div class="w-full h-full flex flex-col">
    <ChatToolbar />
    <div class="overflow-y-auto h-full">
      <ChatItem v-for="c of chatInfoList" :key="c.id" :chat="c" @click="onChatClick(c)" />
    </div>
  </div>
</template>

<style scoped></style>
