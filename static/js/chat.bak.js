var messageTxt;
var messages;

$(function () {

	messageTxt = $("#messageTxt");
	messages = $("#messages");


	w = new Ws("ws://" + HOST + "/my_endpoint");
	w.OnConnect(function () {
		console.log("Websocket connection established");
	});

	w.OnDisconnect(function () {
		appendMessage($("<div><center><h3>Disconnected</h3></center></div>"));
	});

	w.On("chat", function (message) {
		appendMessage($("<div>" + message + "</div>"));
		//hack render img
		var x = Math.floor(Math.random() * 300);
		var y = Math.floor(Math.random() * 100);
		test1(x,y);
	});

	$("#sendBtn").click(function () {
		w.Emit("chat", messageTxt.val().toString());
		messageTxt.val("");
	});

})


function appendMessage(messageDiv) {
    var theDiv = messages[0];
    var doScroll = theDiv.scrollTop == theDiv.scrollHeight - theDiv.clientHeight;
    messageDiv.appendTo(messages);
    if (doScroll) {
        theDiv.scrollTop = theDiv.scrollHeight - theDiv.clientHeight;
    }
}

