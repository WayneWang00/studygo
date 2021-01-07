<template>
    <div class="view">
      <input type="text" v-model="msg"><button @click="Send">发言</button>
      <div class="chat-title">聊天记录：</div>
      <div v-for="(item, index) in msgList" :key="index" class="chat-box">{{ item }}</div>
    </div>
</template>

<script>
  export default {
    data() {
      return {
        msg: '',
        ws: '',
        msgList: []
      }
    },
    methods: {
      Send() {
        this.ws.send(this.msg)
        this.msg = ''
      }
    },
    mounted() {
      this.ws = new WebSocket('ws://localhost:8000/ws')
      this.ws.onopen = function (evt) {
        console.log('connection open ...')
        this.msg = 'connection open'
        this.Send()
      }
      this.ws.onmessage = function (evt) {
        console.log('evt', evt)
        this.msgList.push(evt.data)
      }
      this.ws.onclose = (evt) => {
        console.log('connection close !!!', evt)
        // this.msg = 'connection close'
        // this.Send()
      }
    },
    beforeDestroy() {
      console.log('websocket close')
      this.ws.close()
    }
  }
</script>

<style scoped>
.view{
  width: 600px;
  margin: 0 auto;
  background-color: aliceblue;
  height: 500px;
  text-align: center;
  padding-top: 80px;
}
.chat-title{
  text-align: left;
  margin-left: 100px;
  margin-top: 50px;
}
.chat-box{
  background-color: white;
  width: 400px;
  margin: 0 auto;
}
</style>
