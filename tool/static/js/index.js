//db info update event
$("#db_update_btn").click(function() {
    var obj = GetForm($(this));

    if (obj.addr == "") {
        alert("addr is empty");
        return
    }

    db_info = obj;

    var ajax = new Ajax("connectDb");
    ajax.Post(obj, function(result) {
        console.log(db_info);
        window.localStorage.setItem(db_info.name, result.url);
        //初始化table select
        var tmp = ""
        for (var i = 0; i < result.list.length; i++) {
            tmp += "<option>" + result.list[i] + "</option>"
        }
        $("#table_select").children().remove();
        $("#table_select").append(tmp);
    });
});

//db run
$("#db_run_btn").click(function() {
    var obj = GetForm($(this));
    console.log(obj);
    if (obj.name == "") {
        alert("table is empty");
        return;
    }

    var ajax = new Ajax("createModel");
    ajax.Post(obj, function(result) {
        var txt = $("#data_show").text();
        result = txt + result;
        $("#data_show").text(result);
    });

});

//db clear
$("#db_clear_btn").click(function() {

    $("#data_show").text("");
});