<script setup lang="ts">
import { computed, type PropType, ref } from 'vue'
import type { TelegramResult } from '@/client'
import MessageNode from '@/components/message/MessageNode.vue'

const props = defineProps({
  chat: {
    type: Object as PropType<TelegramResult>,
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
  <div v-if="chat" class="h-screen flex flex-col">
    <!--  Chat room header -->
    <div class="flex items-center p-2">
      <div>
        <Avatar :label="avatarCharacter" class="mr-2" shape="circle" />
      </div>
      <div>
        <div>
          {{ chat.name }}
        </div>
      </div>
      <div class="ml-a">
        <Button type="button" icon="pi pi-ellipsis-v" severity="secondary" />
      </div>
    </div>
    <!--  Message scroller -->
    <div class="overflow-y-auto">
      <div v-for="item of chat.messages" :key="item.id">
        <div v-if="item.type == 'service'" class="w-full flex justify-center my-2">
          <Chip label="Action" />
        </div>
        <MessageNode v-if="item.type == 'message'" :message="item" />
      </div>
    </div>
  </div>
</template>

<style scoped></style>
