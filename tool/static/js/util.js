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