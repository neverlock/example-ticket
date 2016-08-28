function test1() {
        var canvas = document.getElementById('canvas1');
        var ctx = canvas.getContext('2d');
        var cw = canvas.width;
        var ch = canvas.height;
        var id = ctx.getImageData(0, 0, 1, 1);
 
        ctx.clearRect(0, 0, cw, ch);
        var t0 = new Date().getTime();
        for (var i = 0; i < 100000; ++i) {
                //var x = Math.floor(Math.random() * cw);
                //var y = Math.floor(Math.random() * ch);
                //var r = Math.floor(Math.random() * 256);
                //var g = Math.floor(Math.random() * 256);
                //var b = Math.floor(Math.random() * 256);
		var x = arguments[0]
		var y = arguments[1]
		var r = 255;
		var g = 0;
		var b = 0;
 
                ctx.fillStyle = 'rgb(' + r + ',' + g + ',' + b + ')';
                ctx.fillRect(x, y, 1, 1);
        }
        var t1 = new Date().getTime();
        console.log('fillRect() method: ' + (t1 - t0));
}
