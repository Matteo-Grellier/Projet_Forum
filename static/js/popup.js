function closes() {
    var box = document.getElementById("box")
    box.style.display = "none"
    var h1 = document.getElementById("erreur")

    if (h1 == null) {
        box.style.display = "none"
    }
}