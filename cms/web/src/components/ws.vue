<script>
  var webSocket = new WebSocket('ws://localhost:8000/websocket')
  webSocket.onopen = function (evt) {
    console.log('connection open ...')
    webSocket.send("connection open")
  };
  webSocket.onmessage = function (evt) {
    console.log('Received Message:', evt)
  };
  webSocket.onclose = function (evt) {
    console.log('Connection Close !!!')
    webSocket.send('Connection Close !!!')
  }
</script>

