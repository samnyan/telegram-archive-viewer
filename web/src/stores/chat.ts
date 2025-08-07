import { defineStore } from 'pinia'
import { computed, ref } from 'vue'
import type { TelegramResult } from '@/client/types.gen'
import { loadDataUrlAsync } from '@/utils/file.ts'

export interface ChatInfo {
  id: number
  name: string
  lastMessageTime?: Date | string
  lastMessage?: string
}

export const useChatStore = defineStore('chat', () => {
  // State
  const chats = ref<TelegramResult[]>([])
  const currentChatId = ref<number | null>(null)
  const mediaFileMap = ref<Map<string, File>>()

  // Getters
  const currentChat = computed(() => {
    if (currentChatId.value === null) return null
    return chats.value.find((chat) => chat.id === currentChatId.value) || null
  })

  const getChatById = computed(() => (id: number) => {
    return chats.value.find((chat) => chat.id === id) || null
  })

  const chatInfoList = computed<ChatInfo[]>(() => {
    return chats.value.map((i) => {
      return {
        id: i.id,
        name: i.type && i.type == 'saved_messages' ? 'Saved Messages' : i.name,
      }
    })
  })

  // Actions
  const addChat = (chat: TelegramResult) => {
    const existingIndex = chats.value.findIndex((c) => c.id === chat.id)
    if (existingIndex >= 0) {
      chats.value[existingIndex] = chat
    } else {
      chats.value.push(chat)
    }
  }

  const removeChat = (id: number) => {
    const index = chats.value.findIndex((chat) => chat.id === id)
    if (index >= 0) {
      chats.value.splice(index, 1)
      if (currentChatId.value === id) {
        currentChatId.value = null
      }
    }
  }

  const setCurrentChat = (id: number | null) => {
    currentChatId.value = id
  }

  const clearChats = () => {
    chats.value = []
    currentChatId.value = null
  }

  const loadFileMap = (m: Map<string, File>) => {
    mediaFileMap.value = m
  }

  const getMedia = async (path: string) => {
    if (!mediaFileMap.value) return null
    const f = mediaFileMap.value.get(path)
    if (!f) return null

    return await loadDataUrlAsync(f)
  }

  return {
    // State
    chats,
    currentChatId,

    // Getters
    currentChat,
    getChatById,
    chatInfoList,

    // Actions
    addChat,
    removeChat,
    setCurrentChat,
    clearChats,
    loadFileMap,
    getMedia,
  }
})
