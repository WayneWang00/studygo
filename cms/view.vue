<template>
    <div class="view">
        <input type="text" v-model="msg"><button @click="send">发言</button>
        <div class="chat-title">聊天记录：</div>
        <div v-for="(item,index) in msgList" :key="index" class="chat-box">{{item}}</div>
    </div>
</template>

<script>
export default {
   data(){
       return {
           msg: '',
           ws: '',
           msgList: []
       }
   },
   created() {
       this.ws = new WebSocket('127.0.0.1:3000')
       this.ws.onopen = () => {
           console.log('connection open')
       }
       this.ws.onmessage = (evt) => {
           console.log('evt')
           this.msgList.push(evt.data)
       }
       this.ws.onclose = () => {
           console.log('connection close')
       }
   },
   methods: {
       send() {
           this.ws.send(this.msg)
           this.msg = ''
       }
   },
   mounted() {
       this.ws = new WebSocket('127.0.0.1:3000/ping')
       this.ws.onopen = () => {
           console.log('connection open')
       }
       this.ws.onmessage = (evt) => {
           console.log('evt')
           this.msgList.push(evt.data)
       }
       this.ws.onclose = () => {
           console.log('connection close')
       }
   },
   beforeDestroy() {
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
