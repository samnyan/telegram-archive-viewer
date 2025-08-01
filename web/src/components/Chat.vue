<script setup lang="ts">
import ChatList from '@/components/chat/ChatList.vue'
</script>

<template>
  <div class="h-screen flex">
    <!--  Chat list -->
    <div class="w-sm">
      <ChatList />
    </div>
    <!--  Message list -->
    <div class="h-full w-full">
      <RouterView />
    </div>
  </div>
</template>

<style scoped></style>
