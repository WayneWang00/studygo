import Vue from 'vue'
import Router from 'vue-router'
import HelloWorld from '@/components/HelloWorld'
import Chat from "@/components/chat"
// import Ws from '@/components/ws'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'HelloWorld',
      component: HelloWorld
    },
    {
      path: '/ws',
      name: 'Chat',
      component: Chat
    // },
    // {
    //   path: '/websocket',
    //   name: 'Ws',
    //   component:Ws
    }
  ]
})
