class Logger {
    constructor() {

    }

    Info() {
        var now = getNowFormatDate();
        var data = now + " [INFO] [";
        for (var i = 0; i < arguments.length; i++) {
            data += arguments[i] + " ";
        }
        data += "]"
        console.log(data);

        var ajax = new Ajax("mobileLog");
        var obj = new Object();
        obj.data = data;
        ajax.Post(obj, function(result) {

        })
    }
}