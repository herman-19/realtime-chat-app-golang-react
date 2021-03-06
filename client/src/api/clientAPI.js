const socket = new WebSocket("ws://localhost:8080/ws");

// Connects to the WebSocket endpoint and listens
// for events.
const connect = (callback) => {
    console.log("Attempting connection...");

    socket.onopen = () => {
        console.log("Successfully connected!");
    };

    socket.onmessage = msg => {
        console.log(msg);
        callback(msg);
    };

    socket.onclose = event => {
        console.log("Socket closed Connection: ", event);
    };

    socket.onerror = err => {
        console.log("Socket Error: ", err);
    };
};

// Allows client to send messages via the WebSocket connection.
const sendMessage = msg => {
    socket.send(msg);
};

export { connect, sendMessage };