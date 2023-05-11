const express = require("express");
const router = express.Router();
const app = express();
const http = require("http");
const server = http.createServer(app);
const { Server } = require("socket.io");
const io = new Server(server);
const messages = [];

// let count = 0;

io.of(/^\/\d+$/).on("connection", (socket) => {
  // count++;
  const newNamespace = socket.nsp; 

  // let emergency = socket.nsp.name.replace("/", '');
  
  console.log(newNamespace.name);

  let roomChatId = socket.nsp.name.replace("/", '');
  socket.on("disconnect", () => {
    console.log("----- disconnect --------");
    socket.disconnect();
  });

  const username = socket.handshake.query.username;
  console.log("----- Connect --------");
  console.log("username : ", username);
  console.log("roomChatId :",roomChatId);
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

io.on("connection", (socket) => {
  const username = socket.handshake.query.username;

  socket.on("disconnect", () => {
    console.log("----- disconnect --------");
  console.log("username : ", username);
    console.log("----- disconnect --------");
    socket.disconnect();
  });
const eme = 'emergency';
  console.log("----- Connect "+ eme+ " --------");
  console.log("username : ", username);
  console.log("----- Connect "+ eme+ " --------");

  socket.on(eme, (data) => {
    const message = {
      message: data.message,
      senderUsername: username,
      sentAt: Date.now(),
    };
    messages.push(message);
    console.log("message : ", message);
    socket.emit(eme, message);
  });
});

server.listen(3000, () => {
  console.log("listening on *:3000");
});



