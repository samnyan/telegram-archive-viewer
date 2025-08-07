<script setup lang="ts">
import { useChatStore } from '@/stores/chat.ts'
import { useToast } from 'primevue/usetoast'

const toast = useToast()

const chatStore = useChatStore()

const onFileChange = (ev: any) => {
  const fileList: FileList = ev.target.files
  if (fileList.length > 1) {
    // read dir
    const fileMap = new Map<string, File>()
    for (const f of fileList) {
      let relativePath = f.webkitRelativePath

      const paths = relativePath.split('/')
      paths.splice(0, 1)

      relativePath = paths.join('/')
      fileMap.set(relativePath, f)

      if (relativePath == 'result.json') {
        readJson(f)
      }
    }

    chatStore.loadFileMap(fileMap)
  } else {
    readJson(fileList[0])
  }
}

const readJson = (f: File) => {
  const reader = new FileReader()
  reader.onload = (ev) => {
    console.debug('File Loaded')
    const result: string = ev.target!!.result
    chatStore.addChat(JSON.parse(result))
    toast.add({ severity: 'success', summary: 'Load success!', life: 3000 })
  }
  reader.readAsText(f)
}
</script>

<template>
  <div class="p-4">
    <Panel header="Message Loader">
      <div class="flex flex-col gap-2">
        <IftaLabel>
          <InputText id="json" type="file" @change="onFileChange" />
          <label for="json">Select result.json</label>
        </IftaLabel>
        <div>Or</div>
        <IftaLabel>
          <InputText id="dir" type="file" webkitdirectory @change="onFileChange" />
          <label for="dir">ChatExport directory</label>
        </IftaLabel>
      </div>
    </Panel>
  </div>
</template>

<style scoped></style>
