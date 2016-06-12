var serversocket = new WebSocket("wss://lab.owulveryck.info:443/echo");

serversocket.onopen = function() {
    var msg = {
        topic:  "null",
        sender: "null",
        message: "Connection init",
        date: Date.now()
    };
     
    serversocket.send(JSON.stringify(msg));
}

// Write message on receive
serversocket.onmessage = function(e) {
    document.getElementById('output').innerHTML += "Received: " + e.data + "<br>";
};

function senddata() {
    // Construct a msg object containing the data the server needs to process the message from the chat client.
    var msg = {
        topic:  document.getElementById("topic").value,
        sender: document.getElementById("sender").value,
        message: document.getElementById("message").value,
        date: Date.now()
    };
     
    serversocket.send(JSON.stringify(msg));
    document.getElementById('output').innerHTML += "Sent: " + JSON.stringify(msg) + "<br>";
}

