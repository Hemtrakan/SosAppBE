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

  let channelId = socket.nsp.name.replace("/", "");
  socket.on("disconnect", () => {
    console.log("----- disconnect --------");
    socket.disconnect();
  });

  const username = socket.handshake.query.username;
  if (channelId != 0) {
    console.log("----- Connect --------");
    console.log("username : ", username);
    console.log("channelId :", channelId);
    console.log("----- Connect --------");
    socket.on(channelId, (data) => {
      const message = {
        message: data.message,
        senderUsername: username,
        sentAt: Date.now(),
      };
      messages.push(message);
      console.log("message : ", message);
      console.log("channelId :", channelId);
      newNamespace.emit(channelId, message);
    });
  } else {
    console.log("----- Connect --------");
    console.log("username : ", username);
    console.log("channelId :", channelId);
    console.log("----- Connect --------");
    socket.on(channelId, (data) => {
      const message = {
        description: data.description,
        phoneNumberCallBack: data.phoneNumberCallBack,
        latitude: data.latitude,
        longitude: data.longitude,
        username: username,
        type: data.type,
      };
      messages.push(message);
      console.log("message : ", message);
      console.log("channelId :", channelId);
      newNamespace.emit(channelId, message);
    });
  }
});

server.listen(84, () => {
  console.log("listening on *:84");
});

function getDistanceBetweenPoints(
  latitude1,
  longitude1,
  latitude2,
  longitude2,
  unit = "miles"
) {
  let theta = longitude1 - longitude2;
  let distance =
    60 *
    1.1515 *
    (180 / Math.PI) *
    Math.acos(
      Math.sin(latitude1 * (Math.PI / 180)) *
        Math.sin(latitude2 * (Math.PI / 180)) +
        Math.cos(latitude1 * (Math.PI / 180)) *
          Math.cos(latitude2 * (Math.PI / 180)) *
          Math.cos(theta * (Math.PI / 180))
    );
  if (unit == "miles") {
    return Math.round(distance, 2);
  } else if (unit == "kilometers") {
    return Math.round(distance * 1.609344, 2);
  }
}
