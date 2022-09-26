var defaultContactMeButton = {
    border: "",
    borderRadius: "",
    background: "",
    color: ""
};

var defaultContactMeButtonBoxShadow = {
    left: "",
    top: "",
    borderRadius: ""
};

function onMouseOverContactMeButton() {
    let contactMeButton = document.getElementById("contact-me-button");
    defaultContactMeButton.border = contactMeButton.style.border;
    defaultContactMeButton.borderRadius = contactMeButton.style.borderRadius;
    defaultContactMeButton.background = contactMeButton.style.background;
    defaultContactMeButton.color = contactMeButton.style.color;

    let contactMeButtonBoxShadow = document.getElementById("contact-me-button-box-shadow");
    defaultContactMeButtonBoxShadow.left = contactMeButtonBoxShadow.style.left;
    defaultContactMeButtonBoxShadow.top = contactMeButtonBoxShadow.style.top;
    defaultContactMeButtonBoxShadow.borderRadius = contactMeButtonBoxShadow.style.borderRadius;
    
    contactMeButton.style.border = "solid 4px #000";
    contactMeButton.style.borderRadius = "30px";
    contactMeButton.style.background = "transparent";
    contactMeButton.style.color = "#000";

    contactMeButtonBoxShadow.style.left = "0";
    contactMeButtonBoxShadow.style.top = "-50px";
    contactMeButtonBoxShadow.style.borderRadius = "30px";
}

function onMouseOutContactMeButton() {
    let contactMeButton = document.getElementById("contact-me-button");
    let contactMeButtonBoxShadow = document.getElementById("contact-me-button-box-shadow");

    contactMeButton.style.border = defaultContactMeButton.border;
    contactMeButton.style.borderRadius = defaultContactMeButton.borderRadius;
    contactMeButton.style.background = defaultContactMeButton.background;
    contactMeButton.style.color = defaultContactMeButton.color;

    contactMeButtonBoxShadow.style.left = defaultContactMeButtonBoxShadow.left;
    contactMeButtonBoxShadow.style.top = defaultContactMeButtonBoxShadow.top;
    contactMeButtonBoxShadow.style.borderRadius = defaultContactMeButtonBoxShadow.borderRadius; 
}

function onMouseOverCVLink() {
    let arrow = document.getElementById("cv-arrow");
    arrow.style.left = "5px";
    arrow.style.opacity = "1";
}

function onMouseOutCVLink() {
    let arrow = document.getElementById("cv-arrow");
    arrow.style.left = "10px";
    arrow.style.opacity = "0";
}

function toggleHeader() {
    let header = document.querySelector("header");
    header.classList.toggle("active");
}
