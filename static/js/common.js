function mylog(msg) {
    const span = document.createElement("span");
    span.style = "display:block"
    span.innerText = new Date(Date.now()).toLocaleString() + " " + msg
    el("#log").append(span);
}

function el(id) {
    return document.querySelector(id)
}

function ajaxGet(url, endFunc) {
    let ajax = new XMLHttpRequest();
    ajax.open('get', url);
    ajax.send();
    ajax.onreadystatechange = function () {
        if (ajax.readyState == 4) {
            endFunc(ajax.responseText)
        }
    }
}