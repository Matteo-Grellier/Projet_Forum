function showPassword() {
    var input = document.getElementById("myInput");
    var eye = document.getElementById("eye");
    var slash = document.getElementById("slash");

    if (input.type === 'password') {
        input.type = 'text';
        eye.style.display = "initial";
        slash.style.display = "none";
    } else {
        input.type = "password";
        eye.style.display = "none";
        slash.style.display = "initial"
    }
}