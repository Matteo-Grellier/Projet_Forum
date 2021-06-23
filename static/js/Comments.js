function createDiv() {
    if (document.body.contains(document.getElementById('notvisible'))) {
        alert("euuuhh")
        document.getElementById('notvisible').setAttribute("id", "visible")
        var del = document.getElementById('visible').setAttribute("style", "");
    } else {
        alert("yes")
        document.getElementById('visible').setAttribute("id", "notvisible")
        document.getElementById("notvisible");
        rep = document.getElementById('notvisible').setAttribute("style", "display:none;");
    }
}