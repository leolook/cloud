function GetForm(elem) {
    var form = $(elem).parent().parent().parent();
    var child = form.children();

    var data = '{'

    for (var i = 0; i < child.length; i++) {
        var tmp = $(child[i]).children();
        tmp = $(tmp)[1];
        tmp = $(tmp).children();
        if (tmp.length > 0) {
            var name = $(tmp)[0].name;
            var value = $(tmp)[0].value;

            if ((typeof value) == "number") {
                data += '"' + name + '":' + value + ','
            } else {
                data += '"' + name + '":"' + value + '",'
            }
        }
    }
    if (data.length > 1) {
        data = data.substring(0, data.length - 1);
    }
    data = data + "}"
    return JSON.parse(data);
}

function getNowFormatDate() {
    var date = new Date();
    var seperator1 = "-";
    var seperator2 = ":";
    var month = date.getMonth() + 1;
    var strDate = date.getDate();
    if (month >= 1 && month <= 9) {
        month = "0" + month;
    }
    if (strDate >= 0 && strDate <= 9) {
        strDate = "0" + strDate;
    }
    var currentdate = date.getFullYear() + seperator1 + month + seperator1 + strDate +
        " " + date.getHours() + seperator2 + date.getMinutes() +
        seperator2 + date.getSeconds();
    return currentdate;
}

const IPHONE_OTHER = 0;
const IPHONE_UC = 1;
const IPHONE_WE_CHAT = 2;
const IPHONE_SAFARI = 3;

//获取浏览器类型
function GetBorwserType() {
    var type = navigator.userAgent;
    if (type.indexOf("UCBrowser") != -1) { //UC
        return IPHONE_UC;
    } else if (type.indexOf("MicroMessenger") != -1) { //wechat
        return IPHONE_WE_CHAT;
    } else if (type.indexOf("Safari") != -1) { //safari
        return IPHONE_SAFARI;
    } else {
        return IPHONE_OTHER;
    }
}