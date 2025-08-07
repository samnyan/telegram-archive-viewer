<script setup lang="ts">
import { onMounted, ref, watch } from 'vue'
import { useChatStore } from '@/stores/chat.ts'

const props = defineProps({
  mediaType: {
    type: String,
    default: () => undefined,
  },
  thumbnail: {
    type: String,
    default: () => undefined,
  },
  file: {
    type: String,
    default: () => undefined,
  },
  stickerEmoji: {
    type: String,
    default: () => undefined,
  },
  mimeType: {
    type: String,
    default: () => undefined,
  },
})

const src = ref<string>()
const thumbSrc = ref<string>()
const type = ref<'video' | 'image' | 'audio' | string>('video')

const chatStore = useChatStore()

onMounted(() => {
  loadFile()
  loadThumbnail()
  updateMimeType()
})
watch(
  () => props.file,
  () => {
    loadFile()
  },
)
watch(
  () => props.thumbnail,
  () => {
    loadThumbnail()
  },
)
const loadFile = async () => {
  if (props.file) {
    const resp = await chatStore.getMedia(props.file)
    if (resp) {
      src.value = resp
    } else {
      src.value = undefined
    }
  } else {
    src.value = undefined
  }
}
const loadThumbnail = async () => {
  if (props.thumbnail) {
    const resp = await chatStore.getMedia(props.thumbnail)
    if (resp) {
      thumbSrc.value = resp
    } else {
      thumbSrc.value = undefined
    }
  } else {
    thumbSrc.value = undefined
  }
}
const updateMimeType = () => {
  if (props.mimeType) {
    if (props.mimeType.startsWith('video')) {
      type.value = 'video'
    } else if (props.mimeType.startsWith('image')) {
      type.value = 'image'
    } else if (props.mimeType.startsWith('audio')) {
      type.value = 'audio'
    } else {
      type.value = 'Unknown media type: ' + props.mimeType
    }
  } else {
    type.value = 'No media type'
  }
}
</script>

<template>
  <div v-if="src" :title="stickerEmoji" class="max-w-64">
    <video
      v-if="type == 'video'"
      :src="src"
      :poster="thumbSrc"
      class="object-contain"
      :controls="mediaType != 'animation' && mediaType != 'sticker'"
      :autoplay="mediaType == 'animation' || mediaType == 'sticker'"
      :loop="mediaType == 'animation' || mediaType == 'sticker'"
    />
    <div v-else-if="type == 'audio'">
      <audio controls>
        <source :src="src" :type="mediaType" />
        Your browser does not support the audio tag.
      </audio>
      <span>[Voice message]</span>
    </div>
    <img v-else-if="type == 'image'" :src="src" class="object-contain" />
    <span v-else>{{ type }}</span>
  </div>
</template>

<style scoped></style>
