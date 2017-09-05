jQuery(document).ready(function($) {
    $(".clickable-row").click(function() {
        window.location = $(this).data("href");
    });
});

function allowDrop(ev) {
    alert("test");
    ev.preventDefault();
}

function drag(ev) {    
    ev.dataTransfer.setData("text", ev.target.id);
}

function drop(ev) {
    
    ev.preventDefault();
    var data = ev.dataTransfer.getData("text");
    ev.target.appendChild(document.getElementById(data));
}

var pageSaved;

function leavePage(){  
    if(pageSaved !== true){
        return "Page Not Saved!";
    }else{
        return;
    }  
}
function savePage(){
    pageSaved = true;
}

var imageId;
function setImageId(id){
    imageId = id;
}
function getImageId(){
    return imageId;
}


var tmplId;
function setTemplateId(id){
    tmplId = id;
}
function getTemplateId(){
    return tmplId;
}