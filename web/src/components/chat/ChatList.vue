<script setup lang="ts">
import ChatToolbar from '@/components/chat/ChatToolbar.vue'
import ChatItem from '@/components/chat/ChatItem.vue'
</script>

<template>
  <div class="w-full h-full flex flex-col">
    <ChatToolbar />
    <div class="overflow-y-scroll">
      <ChatItem v-for="idx in 10" :key="idx" />
    </div>
  </div>
</template>

<style scoped></style>
