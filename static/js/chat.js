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
		//appendMessage($("<div>" + message + "</div>"));
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

function test1() {
        var canvas = document.getElementById('canvas');
        var ctx = canvas.getContext('2d');
        //var cw = canvas.width;
        //var ch = canvas.height;
        var id = ctx.getImageData(0, 0, 1, 1);
        var x = arguments[0]
        var y = arguments[1]

	/*
        var t0 = new Date().getTime();
                var r = 255;
                var g = 0;
                var b = 0;

                ctx.fillStyle = 'rgb(' + r + ',' + g + ',' + b + ')';
                ctx.fillRect(x, y, 1, 1);
        var t1 = new Date().getTime();
        console.log('fillRect() method: ' + (t1 - t0));
	*/


        id.data[3] = 255;
        var t0 = new Date().getTime();
                var g = 0;
                var b = 0;
 
                id.data[0] = r;
                id.data[1] = g;
                id.data[2] = b;
                ctx.putImageData(id, x, y);
        var t1 = new Date().getTime();
        console.log('putImage() method: ' + (t1 - t0));
}
