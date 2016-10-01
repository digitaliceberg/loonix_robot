function httpGet(theUrl)
{
    var xmlHttp = new XMLHttpRequest();
    xmlHttp.open( "GET", theUrl, false ); // false for synchronous request
    xmlHttp.send( null );
    return xmlHttp.responseText;
}

window.addEventListener("keypress", checkKeyPressed, false);
alert("pressed.");
function checkKeyPressed(e) {
    if (e.charCode == "87") {
        alert("The 'w' key is pressed.");
    }
}
