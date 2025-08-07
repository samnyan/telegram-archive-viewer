<script setup lang="ts">
import { onMounted, ref, watch } from 'vue'
import { useChatStore } from '@/stores/chat.ts'

const props = defineProps({
  photo: {
    type: String,
    default: () => undefined,
  },
})

const imgSrc = ref<string>()

const chatStore = useChatStore()

onMounted(() => {
  loadImage()
})
watch(
  () => props.photo,
  () => {
    loadImage()
  },
)
const loadImage = async () => {
  if (props.photo) {
    const resp = await chatStore.getMedia(props.photo)
    if (resp) {
      imgSrc.value = resp
    } else {
      imgSrc.value = undefined
    }
  } else {
    imgSrc.value = undefined
  }
}
</script>

<template>
  <img :src="imgSrc" class="object-contain max-w-64" />
</template>

<style scoped></style>
