<template>
  <div id="app">
    <img src="./assets/logo.png">
    <router-view/>
  </div>
</template>

<script>
import Chat from './components/chat'
// import Ws from './components/ws'
export default {
  // name: 'App'

  name: 'Chat',
  components: {
    Chat
  }

  // name: 'Ws',
  // components: {
  //   Ws
  // }
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
