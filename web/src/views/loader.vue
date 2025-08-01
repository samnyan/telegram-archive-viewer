<script setup lang="ts">
import { useChatStore } from '@/stores/chat.ts'
import { useToast } from 'primevue/usetoast'

const toast = useToast()

const chatStore = useChatStore()

const onFileChange = (ev: any) => {
  const fileList: FileList = ev.target.files
  const reader = new FileReader()
  reader.onload = (ev) => {
    console.debug('File Loaded')
    const result: string = ev.target!!.result
    chatStore.addChat(JSON.parse(result))
    toast.add({ severity: 'success', summary: 'Load success!', life: 3000 })
  }
  reader.readAsText(fileList[0])
}
</script>

<template>
  <div class="p-4">
    <Panel header="Message Loader">
      <div>
        <input type="file" id="input" @change="onFileChange" />
      </div>
    </Panel>
  </div>
</template>

<style scoped></style>
