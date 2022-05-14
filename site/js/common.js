let socketUrl = "ws://192.168.44.244:7000/ws";
let apiUrl = "http://192.168.44.244:7070";

function getLocalStorage(name) {
    let value = localStorage.getItem(name);
    if (value == null) {
        return "";
    }
    return value;
}

function setLocalStorage(name, value) {
    localStorage.setItem(name, value)
}
