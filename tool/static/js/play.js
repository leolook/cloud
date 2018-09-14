var log = new Logger();

var type = GetBorwserType();
if (type != IPHONE_UC) {
    var log = new Logger();
    log.Info(video.style.cssText);
    // video.style.cssText += "transform: rotate(90deg); -ms-transform: rotate(90deg);-moz-transform: rotate(90deg); -webkit-transform: rotate(90deg);-o-transform: rotate(90deg);"
    log.Info(video.style.cssText);
    var h = document.documentElement.clientHeight;
    var tmp = document.getElementById("video");
    tmp.setAttribute("width", h);
    log.Info(tmp.getAttribute("width"));
}