var bCheckEnabled = true;
var bFinishCheck = false;

var img;
var imgArray = new Array();
var i = 0;

var myInterval = setInterval(loadImage, 1);

function loadImage() {

    while (!bFinishCheck) {
 
    if (bCheckEnabled) {

        bCheckEnabled = false;

        img = new Image();
        img.onload = fExists;
        img.onerror = fDoesntExist;
        img.src = 'images/myFolder/' + i + '.png';

    }
    }
    clearInterval(myInterval);
    alert('Loaded ' + i + ' image(s)!)');
    return;

}

function fExists() {
    imgArray.push(img);
    i++;
    bCheckEnabled = true;
}

function fDoesntExist() {
    bFinishCheck = true;
} 