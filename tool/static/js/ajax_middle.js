const http_head = "http://127.0.0.1:12020/tool/";
var db_info = new Object();

class Ajax {
    constructor(url) {
        this.url = url;
    }

    Post(obj, callback) {
        // console.log(this.url, this.header, obj);
        this.url = http_head + this.url
        const data = JSON.stringify(obj);
        $.ajax({
            "url": this.url,
            "type": "post",
            "contentType": "application/json;utf-8",
            "data": data,
            "dataType": "json",
            "beforeSend": function(request) {
                var head = window.localStorage.getItem(db_info.name);
                if (head != "") {
                    console.log(head);
                    request.setRequestHeader("url", head);
                }
            },
            "success": function(result) {
                if (result.code == 200) {
                    callback(result.data);
                } else {
                    alert(result.message);
                    return;
                }
            },
            "error": function(result) {
                console.log(result);
            },
        })
    }
}