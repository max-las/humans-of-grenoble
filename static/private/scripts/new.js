var resizedImage = null;

$(function() {
  $("#imageFile").change(function(){
    var file = this.files[0];

    if (file) {
      $(".file-name").html(file.name);
      $(".file").addClass("has-name");
      $("#imageField .file-name").removeClass("is-hidden");

      if(file.type != "image/jpeg" && file.type != "image/png" && file.type != "image/gif"){
        $("#fileError .message-body").html("Format de fichier non pris en charge.<br>Les formats acceptés sont : JPG, PNG et GIF.");
        $("#fileError").removeClass("is-hidden");
      }else if(file.size > (50 << (10 * 2))){ // > 50MB
        $("#fileError .message-body").html("Ce fichier dépasse la taille maximum de 50 Mo.");
        $("#fileError").removeClass("is-hidden");
      }else{
        resizeImage(file, function(){
          $("#fileError").addClass("is-hidden");
          $("#preview img").attr("src", URL.createObjectURL(resizedImage));
          $("#preview").removeClass("is-hidden");
        });
      }

    }else{

      clearFileInput();

    }
  });
});

function upload(event) {
  var storyForm = document.getElementById("storyForm");
  if(!storyForm.checkValidity()){
    $("#fakeSubmit").click();
  }else{
    if($(".message.is-danger:visible").length > 0){
      animateCSS(".message.is-danger", "headShake");
    }else{
      $("#sendButton").addClass("is-hidden");
      $("#sendingButton").removeClass("is-hidden");
      var data = new FormData(storyForm);
      var ajax = new XMLHttpRequest();
      ajax.upload.addEventListener("progress", function(event) {
        var percent = Math.round((event.loaded / event.total) * 100);
        $("progress").val(percent);
        $("progress").html(percent + "%");
      }, false);
      ajax.addEventListener("load", function(event) {
        $("#sendButton").removeClass("is-hidden");
        $("#sendingButton").addClass("is-hidden");
        if(event.target.status == 200){
          clearForm();
          modAlert("success", "La story a été publiée !");
        }else{
          modAlert("failure", "Le serveur a répondu par une erreur " + event.target.status + ".");
        }
      }, false);
      ajax.addEventListener("error", function(event) {
        $("#sendButton").removeClass("is-hidden");
        $("#sendingButton").addClass("is-hidden");
        modAlert("failure", "Le serveur est injoignable.");
      }, false);
      ajax.addEventListener("abort", function(event) {
        modAlert("failure", "Requête annulée.");
      }, false);
      $("progress").removeClass("is-hidden");
      ajax.open("POST", "/admin/new");
      ajax.send(data);
    }
  }
}

function clearForm(){
  document.getElementById("storyForm").reset();
  clearFileInput();
  $("progress").addClass("is-hidden");
  $("progress").val(0);
  $("progress").html("");
}

function clearFileInput(){
  $("#preview, #imageField .file-name").addClass("is-hidden");
  $("#imageField .file").removeClass("has-name");
  $("#imageField .file-name").html("");
}

function resizeImage(imageFile, callback){
  let canvas = document.createElement("canvas");
  let ctx = canvas.getContext("2d");
  var imgBefore = new Image();
  imgBefore.onload = function(){
    var ratio = this.width / this.height;
    var heightAfter, widthAfter;
    if(this.height < this.width){
      heightAfter = 1500;
      widthAfter = 1500 * ratio;
    }else{
      widthAfter = 1500;
      heightAfter = 1500 / ratio;
    }
    canvas.width = widthAfter;
    canvas.height = heightAfter;
    ctx.drawImage(imgBefore, 0, 0, widthAfter, heightAfter);
    canvas.toBlob(function(blob){
      resizedImage = new File([blob], imageFile.name, {
        type: "image/png"
      });
      callback();
    }, "image/png");
  };
  imgBefore.src = URL.createObjectURL(imageFile);
}
