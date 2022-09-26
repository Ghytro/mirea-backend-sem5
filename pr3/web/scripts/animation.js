function loadAnimation() {
    for (let el of document.getElementsByClassName("left-to-right-animation")) {
        if (el.className.includes("animation-order-")) {
            el.style.transitionDelay = "calc(var(--animation-transition-time) * "+
            (parseInt(el.className.match(/animation-order-\d+/)[0].slice(-1)) - 1)+")";
        }
        el.style.right = 0;
        el.style.opacity = 1;
    }
    for (let el of document.getElementsByClassName("right-to-left-animation")) {
        if (el.className.includes("animation-order-")) {
            el.style.transitionDelay = "calc(var(--animation-transition-time) * "+
            (parseInt(el.className.match(/animation-order-\d+/)[0].slice(-1)) - 1)+")";
        }
        el.style.left = 0;
        el.style.opacity = 1;
    }
    for (let el of document.getElementsByClassName("bottom-to-top-animation")) {
        if (el.className.includes("animation-order-")) {
            el.style.transitionDelay = "calc(var(--animation-transition-time) * "+
            (parseInt(el.className.match(/animation-order-\d+/)[0].slice(-1)) - 1)+")";
        }
        el.style.top = 0;
        el.style.opacity = 1;
    }
    for (let el of document.getElementsByClassName("top-to-bottom-animation")) {
        if (el.className.includes("animation-order-")) {
            el.style.transitionDelay = "calc(var(--animation-transition-time) * "+
            (parseInt(el.className.match(/animation-order-\d+/)[0].slice(-1)) - 1)+")";
        }
        el.style.bottom = 0;
        el.style.opacity = 1;
    }
    for (let el of document.getElementsByClassName("plain-animation")) {
        if (el.className.includes("animation-order-")) {
            el.style.transitionDelay = "calc(var(--animation-transition-time) * "+
            (parseInt(el.className.match(/animation-order-\d+/)[0].slice(-1)) - 1)+")";
        }
        el.style.opacity = 1;
    }
}
