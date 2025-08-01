<script setup lang="ts">
import { type PropType, ref } from 'vue'
import type { TelegramResult } from '@/client'
import MessageNode from '@/components/message/MessageNode.vue'

const props = defineProps({
  chat: {
    type: Object as PropType<TelegramResult>,
    default: () => undefined,
  },
})
</script>

<template>
  <div class="w-full h-screen">
    <!--  Chat room header -->
    <div class="w-full flex items-center px-2" style="height: 32px">
      <div>{{ chat?.name }}</div>
    </div>
    <!--  Message scroller -->
    <div>
      <div v-if="chat" class="flex justify-center">
        <VirtualScroller
          :items="chat.messages"
          :itemSize="50"
          class="border border-surface-200 dark:border-surface-700 rounded"
          style="width: 100%; height: calc(100vh - 32px)"
        >
          <template v-slot:item="{ item }">
            <div v-if="item.type == 'service'" class="w-full flex justify-center my-2">
              <Chip label="Action" />
            </div>
            <MessageNode v-if="item.type == 'message'" :message="item" />
          </template>
        </VirtualScroller>
      </div>
    </div>
  </div>
</template>

<style scoped></style>
