//读取窗口大小
var winwigh, winheight;
if (window.innerWidth) {
    winwigh = window.innerWidth;
    winheight = window.innerHeight;
}
else {
    winwigh = document.documentElement.clientWidth;
    winheight = document.documentElement.clientHeight;
}

//随窗口变化
function refresh() {
    if (winwigh != window.innerWidth) {
        debounce(location.reload(), 2000);
    }
}
//消抖函数
function debounce(fun, wait) {
    var time;
    return function () {
        clearTimeout(time);
        time = setTimeout(fun, wait);
    }
}
//sleep函数
function sleep(ms) {
    var start = Date.now(), end = start + ms;
    while (Date.now() < end);
    return;
}