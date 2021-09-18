document.addEventListener("DOMContentLoaded", function(){
  if(document.querySelector("#storyForm")){
    initEdit();
  }
});

function initEdit(){
  resizedImage = null;

  mode = document.querySelector("#storyForm").dataset.mode;

  fileChange(document.getElementById("imageFile"));

  $("#imageFile").on("change", function(){
    fileChange(this);
  });
}

function fileChange(fileInput){
  var file = fileInput.files[0];

  if (file) {
    $(".file-name").html(file.name);
    $(".file").addClass("has-name");
    $("#imageField .file-name").removeClass("is-hidden");

    if(file.type != "image/jpeg" && file.type != "image/png"){
      $("#fileError .message-body").html("Format de fichier non pris en charge.<br>Les formats acceptés sont : JPG et PNG.");
      $("#fileError").removeClass("is-hidden");
    }else{
      resizeImage(file, function(resized){
        resizedImage = resized;
        $("#fileError").addClass("is-hidden");
        $("#preview img").attr("src", URL.createObjectURL(resizedImage));
        $("#preview").removeClass("is-hidden");
      });
    }

  }else{

    if(mode == "new"){
      clearFileInputStyle();
    }

  }
}

function uploadStory(event) {
  var storyForm = document.getElementById("storyForm");
  if(!storyForm.checkValidity()){
    document.querySelector("#fakeSubmit").click();
  }else{
    if($("#storyForm .message.is-danger").not(".is-hidden").length > 0){
      animateCSS("#fileError", "headShake");
    }else{
      $("#sendButton").addClass("is-hidden");
      $("#sendingButton").removeClass("is-hidden");

      let localData = new FormData();
      localData.set("text", document.querySelector("#text").value)

      if(resizedImage){

        let cloudinaryData = new FormData();
        cloudinaryData.set("folder", "humans-of-grenoble");
        fetchCloudinarySign(cloudinaryData)
        .then(function(res){
          cloudinaryData.set("api_key", res.api_key)
          cloudinaryData.set("timestamp", res.timestamp)
          cloudinaryData.set("signature", res.signature)
          cloudinaryData.set("file", resizedImage)

          cloudinaryUpload(cloudinaryData)
          .then(function(res){
            localData.set("photoUrl", res.secure_url)
            localData.set("photoPublicId", res.public_id)
            setStory(localData)
            .then(function(storyUrl){
              window.location.href = storyUrl
            })
            .catch(function(error){
              modAlert("failure", error)
            })
          })
          .catch(function(error){
            modAlert("failure", error);
          });
        })
        .catch(function(error) {
          modAlert("failure", error);
        });

      }else{

        setStory(localData)
        .then(function(storyUrl){
          window.location.href = storyUrl
        })
        .catch(function(error){
          modAlert("failure", error)
        })

      }

    }
  }
}

function cloudinaryUpload(data) {
  return new Promise((resolve, reject) => {
    var ajax = new XMLHttpRequest();
    ajax.responseType = "json";
    ajax.upload.addEventListener("progress", function(event) {
      var percent = Math.round((event.loaded / event.total) * 100);
      $("progress").val(percent);
      $("progress").html(percent + "%");
    }, false);
    ajax.addEventListener("load", function(event) {
      $("#sendButton").removeClass("is-hidden");
      $("#sendingButton").addClass("is-hidden");
      if(event.target.status == 200){
        resolve(event.target.response);
      }else{
        reject("Le serveur a répondu par une erreur " + event.target.status + ".");
      }
    }, false);
    ajax.addEventListener("error", function(event) {
      $("#sendButton").removeClass("is-hidden");
      $("#sendingButton").addClass("is-hidden");
      reject("Le serveur est injoignable.");
    }, false);
    ajax.addEventListener("abort", function(event) {
      reject("Requête annulée.");
    }, false);
    $("progress").removeClass("is-hidden");
    ajax.open("POST", "https://api.cloudinary.com/v1_1/dehn7bofz/image/upload");
    ajax.send(data);
  });
}

function fetchCloudinarySign(data) {
  return new Promise((resolve, reject) => {
    var ajax = new XMLHttpRequest();
    ajax.responseType = "json";
    ajax.addEventListener("load", function(event) {
      if(event.target.status == 200){
        resolve(event.target.response);
      }else{
        reject("Le serveur a répondu par une erreur " + event.target.status + ".");
      }
    }, false);
    ajax.addEventListener("error", function(event) {
      reject("Le serveur est injoignable.");
    }, false);
    ajax.addEventListener("abort", function(event) {
      reject("Requête annulée.");
    }, false);
    ajax.open("POST", "/admin/cloudinary");
    ajax.send(data);
  });
}

function setStory(data) {
  return new Promise((resolve, reject) => {
    var ajax = new XMLHttpRequest();
    ajax.addEventListener("load", function(event) {
      if(event.target.status == 200){
        resolve(event.target.responseText);
      }else{
        reject("Le serveur a répondu par une erreur " + event.target.status + ".");
      }
    }, false);
    ajax.addEventListener("error", function(event) {
      reject("Le serveur est injoignable.");
    }, false);
    ajax.addEventListener("abort", function(event) {
      reject("Requête annulée.");
    }, false);
    ajax.open("POST", window.location.href)
    ajax.send(data);
  });
}

function clearForm(){
  document.getElementById("storyForm").reset();
  clearFileInput();
  $("progress").addClass("is-hidden");
  $("progress").val(0);
  $("progress").html("");
}

function clearFileInputStyle(){
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
      var filename = imageFile.name.substring(0, imageFile.name.lastIndexOf(".")) + ".jpg";
      var result = new File([blob], filename, {
        type: "image/jpeg"
      });
      callback(result);
    }, "image/jpeg", 0.70);
  };
  imgBefore.src = URL.createObjectURL(imageFile);
}

function deleteStory(event){
  let id = $(event.target).closest(".box").attr("data-id");
  let selector = "[data-id='" + id + "']";

  var modalContent = $(selector + " img").prop("outerHTML");

  modalConfirm("Supprimer cette story ?", modalContent, "Supprimer", function(){
    let ajax = new XMLHttpRequest();

    ajax.addEventListener("load", function(event) {
      if(event.target.status == 200){
        if(event.target.responseText.trim() == "OK"){
          animateCSS(selector, "zoomOut").then((message) => {
            $(selector).remove();
          });
        }else{
          modAlert("failure", "Le service est buggé. Contactez l'administrateur.");
        }
      }else{
        modAlert("failure", "Le serveur a répondu par une erreur " + event.target.status + ".");
      }
    }, false);

    ajax.addEventListener("error", function(event) {
      modAlert("failure", "Le serveur est injoignable.");
    }, false);

    ajax.addEventListener("abort", function(event) {
      modAlert("failure", "Requête annulée.");
    }, false);

    ajax.open("DELETE", "/admin/edit/" + id);
    ajax.send();
  });
}

function updatePassword(event){
  var newPassForm = document.querySelector("#newPassForm");
  let data = new FormData(newPassForm);

  let ajax = new XMLHttpRequest();
  ajax.addEventListener("load", function(event) {
    if(event.target.status == 200){
      if(event.target.responseText.trim() == "OK"){
        modAlert("success", "Mis à jour");
      }else{
        modAlert("failure", "Une erreur inattendue est survenue. Contactez l'administrateur.");
      }
    }else{
      modAlert("failure", "Le serveur a répondu par une erreur " + event.target.status + ".");
    }
  }, false);

  ajax.addEventListener("error", function(event) {
    modAlert("failure", "Le serveur est injoignable.");
  }, false);

  ajax.addEventListener("abort", function(event) {
    modAlert("failure", "Requête annulée.");
  }, false);

  ajax.open("POST", "/admin/new-password");
  ajax.send(data);
}

function logout(){
  let ajax = new XMLHttpRequest();
  ajax.addEventListener("load", function(event) {
    if(event.target.status == 200){
      if(event.target.responseText.trim() == "OK"){
        window.location.href = "/admin/login";
      }else{
        modAlert("failure", "Une erreur inattendue est survenue. Contactez l'administrateur.");
      }
    }else{
      modAlert("failure", "Le serveur a répondu par une erreur " + event.target.status + ".");
    }
  }, false);

  ajax.addEventListener("error", function(event) {
    modAlert("failure", "Le serveur est injoignable.");
  }, false);

  ajax.addEventListener("abort", function(event) {
    modAlert("failure", "Requête annulée.");
  }, false);

  ajax.open("DELETE", "/admin/login");
  ajax.send();
}
