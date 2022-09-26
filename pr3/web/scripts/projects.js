class GalleryElement {
    constructor(leftElement, rightElement) {
        this.leftElement = leftElement;
        this.rightElement = rightElement;
    }

    setOpacity(opacity) {
        this.leftElement.style.opacity = opacity;
        this.rightElement.style.opacity = opacity;
    }

    setDisplay(display) {
        this.leftElement.style.display = display;
        this.rightElement.style.display = display;
    }
}

var galleryElements = [];
var galleryIndex = 0;

function sleep(ms) {
    return new Promise(resolve => setTimeout(resolve, ms));
}

function loadGalleryElements() {
    let leftElements = document.getElementsByClassName("gallery-content-left");
    let rightElements = document.getElementsByClassName("gallery-content-right");
    for (let i = 0; i < leftElements.length; ++i) {
        galleryElements.push(new GalleryElement(leftElements[i], rightElements[i]));
    }
}

function incGalleryIndex() {
    ++galleryIndex;
    if (galleryIndex >= galleryElements.length) {
        galleryIndex = 0;
    }
}

function decGalleryIndex() {
    --galleryIndex;
    if (galleryIndex < 0) {
        galleryIndex = galleryElements.length - 1;
    }
}

async function changeElement(indexModifyCallback) {
    let el = galleryElements[galleryIndex];
    el.setOpacity(0);
    await sleep(100);
    el.setDisplay("none");
    indexModifyCallback();
    el = galleryElements[galleryIndex];
    el.setDisplay("block");
    await sleep(100);
    el.setOpacity(1);
}
