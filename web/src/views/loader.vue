<script setup lang="ts">
import { ref } from 'vue'
import type { TelegramResult } from '@/client'

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
  <Panel header="Message Loader">
    <div>
      <input type="file" id="input" @change="onFileChange" />
    </div>
  </Panel>
</template>

<style scoped></style>
