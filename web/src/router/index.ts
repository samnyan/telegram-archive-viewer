import { createRouter, createWebHistory } from 'vue-router'
import Home from '@/views/home.vue'
import Loader from '@/views/loader.vue'
import Message from '@/views/message.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: Home,
    },
    {
      path: '/loader',
      name: 'loader',
      component: Loader,
    },
    {
      path: '/message',
      name: 'message',
      component: Message,
    },
  ],
})

export default router
