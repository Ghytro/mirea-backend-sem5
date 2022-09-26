function stackElementOnMouseOver(showObject, row) {
    let descEl = document.getElementById(showObject+"-text"+row);
    let imgEl = document.getElementById(showObject+"-img"+row);
    
    descEl.style.opacity = 1;
    imgEl.style.opacity = 0;
    
    descEl.style.bottom = "300px";
    imgEl.style.top = "20px";
}

function stackElementOnMouseOut(hideObject, row) {
    let descEl = document.getElementById(hideObject+"-text"+row);
    let imgEl = document.getElementById(hideObject+"-img"+row);
    
    descEl.style.opacity = 0;
    imgEl.style.opacity = 1;
    
    descEl.style.bottom = "290px";
    imgEl.style.top = 0;
}
