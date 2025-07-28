<script setup lang="ts">
import MessageList from '@/components/chat/MessageList.vue'
import type { TelegramResult } from '@/client'
import { ref } from 'vue'

const currentChat = ref<TelegramResult>()

const onFileChange = (ev: any) => {
  const fileList: FileList = ev.target.files
  const reader = new FileReader()
  reader.onload = (ev) => {
    console.debug('File Loaded')
    const result: string = ev.target!!.result
    currentChat.value = JSON.parse(result)
  }
  reader.readAsText(fileList[0])
}
</script>

<template>
  <div>
    <!--  Chat list -->
    <div></div>
    <!--  Message list -->
    <div>
      <MessageList v-if="currentChat" :chat="currentChat" />
      <div v-else class="p-2">
        <Panel header="Message Loader">
          <div>
            <input type="file" id="input" @change="onFileChange" />
          </div>
        </Panel>
      </div>
    </div>
  </div>
</template>

<style scoped></style>
