var ws = new WebSocket("ws://localhost:2303/ping");
//连接打开时触发
ws.onopen = function(evt) {
    console.log("Connection open ...");
    ws.send("Hello WebSockets!");
};
//接收到消息时触发
ws.onmessage = function(evt) {
    console.log("Received Message: " + evt.data);
};
//连接关闭时触发
ws.onclose = function(evt) {
    console.log("Connection closed.");
};

