const express = require("express");
const router = express.Router();
const app = express();
const http = require("http");
const server = http.createServer(app);
const { Server } = require("socket.io");
const io = new Server(server);
const messages = [];

let count = 0;

io.of(/^\/\d+$/).on("connection", (socket) => {
  count++;
  console.log(count);
  const newNamespace = socket.nsp; 
  let roomChatId = socket.nsp.name.replace("/", '');
  socket.on("disconnect", () => {
    socket.disconnect();
  });

  const username = socket.handshake.query.username;
  console.log("----- Connect --------");
  console.log("username : ", username);
  console.log("----- Connect --------");
  socket.on(roomChatId, (data) => {
    const message = {
      message: data.message,
      senderUsername: username,
      sentAt: Date.now(),
    };
    messages.push(message);
    console.log("message : ", message);
    console.log("roomChatId :",roomChatId);
    newNamespace.emit(roomChatId, message);
  });
});



server.listen(3000, () => {
  console.log("listening on *:3000");
});



