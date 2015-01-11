if(!window.location.origin) {
    // Some browsers (mainly IE) do not have this property, so we need to build it manually...
    window.location.origin = window.location.protocol + '//' + window.location.hostname + (window.location.port ? (':' + window.location.port) : '');
}


// Create a connection to http://localhost:9999/ws
var sock = new SockJS(window.location.origin+"/ws");

// Open the connection
sock.onopen = function() {
      console.log('open');
};

// On connection close
sock.onclose = function() {
      console.log('close');
};

// On receive message from server
sock.onmessage = function(e) {
    // Get the content
    console.log(e.data);
    var content = JSON.parse(e.data);
    if (content.username == undefined && content.message == undefined) {
        $("#chat-content").val(e.data + '\n');
    } else {

    //Append the text to textarea (using JQuery)
    $("#chat-content").val(function(i,text) {
        return text + 'User ' + content.username + ': ' + content.message + '\n';
    });

    }

};

// Function is for handling Send button click event
function sendMessage() {
    //Get the content from the textbox
    var message = $("#message").val();
    var username = $("#username").val();

    //The object to send
    var send = {
        message : message,
        username : username
    }
    sock.send(JSON.stringify(send));
    $("#message").val("");
}
