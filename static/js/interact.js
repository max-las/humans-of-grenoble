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
      if(file.size < 800000){
        resizedImage = file;
        showPreview();
      }else{
        resizeImage(file, function(resized){
          resizedImage = resized;
          showPreview();
        });
      }
    }

  }else{

    if(mode == "new"){
      clearFileInputStyle();
    }

  }
}

function showPreview(){
  $("#fileError").addClass("is-hidden");
  $("#preview img").attr("src", URL.createObjectURL(resizedImage));
  $("#preview").removeClass("is-hidden");
}

function uploadStory(event) {
  let storyForm = document.getElementById("storyForm");
  if(!storyForm.checkValidity()){
    document.querySelector("#fakeSubmit").click();
  }else{
    if($("#storyForm .message.is-danger").not(".is-hidden").length > 0){
      animateCSS("#fileError", "headShake");
    }else{
      setSubmitButton("sending");

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
              setSubmitButton("send");
              if(mode == "new"){
                clearForm();
                notify("success", "Story publiée avec succès.")
              }
              if(mode == "edit"){
                clearProgress();
                notify("success", "Story éditée avec succès.")
              }
            })
            .catch(function(error){
              setSubmitButton("send");
              notify("failure", error)
            })
          })
          .catch(function(error){
            setSubmitButton("send");
            notify("failure", error);
          });
        })
        .catch(function(error) {
          setSubmitButton("send");
          notify("failure", error);
        });

      }else{

        setStory(localData)
        .then(function(storyUrl){
          clearProgress();
          setSubmitButton("send");
          notify("success", "Story éditée avec succès.")
        })
        .catch(function(error){
          setSubmitButton("send");
          notify("failure", error)
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
    });
    ajax.addEventListener("load", function(event) {
      if(event.target.status == 200){
        resolve(event.target.response);
      }else{
        reject("Le serveur a répondu par une erreur " + event.target.status + ".");
      }
    });
    ajax.addEventListener("error", function(event) {
      reject("Le serveur est injoignable.");
    });
    ajax.addEventListener("abort", function(event) {
      reject("Requête annulée.");
    });
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
    });
    ajax.addEventListener("error", function(event) {
      reject("Le serveur est injoignable.");
    });
    ajax.addEventListener("abort", function(event) {
      reject("Requête annulée.");
    });
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
    });
    ajax.addEventListener("error", function(event) {
      reject("Le serveur est injoignable.");
    });
    ajax.addEventListener("abort", function(event) {
      reject("Requête annulée.");
    });
    ajax.open("POST", window.location.href)
    ajax.send(data);
  });
}

function clearForm(){
  document.getElementById("storyForm").reset();
  clearFileInputStyle();
  clearProgress();
}

function clearFileInputStyle(){
  $("#preview, #imageField .file-name").addClass("is-hidden");
  $("#imageField .file").removeClass("has-name");
  $("#imageField .file-name").html("");
}

function clearProgress(){
  $("progress").addClass("is-hidden");
  $("progress").val(0);
  $("progress").html("");
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
    });

    ajax.addEventListener("error", function(event) {
      modAlert("failure", "Le serveur est injoignable.");
    });

    ajax.addEventListener("abort", function(event) {
      modAlert("failure", "Requête annulée.");
    });

    ajax.open("DELETE", "/admin/edit/" + id);
    ajax.send();
  });
}

function updatePassword(event){
  document.querySelector(".notify").innerHtml = "";

  let newPassForm = document.querySelector("#newPassForm");

  if(!newPassForm.checkValidity()){
    document.querySelector("#fakeSubmit").click();
  }else{
    let data = new FormData(newPassForm);

    var password = data.get("password");
    var passwordConfirm = data.get("password-confirm");
    if(password != passwordConfirm){
      notify("failure", "Le mot de passe et sa confirmation sont différents.")
    }else{
      setSubmitButton("sending");

      let ajax = new XMLHttpRequest();
      ajax.addEventListener("load", function(event) {
        if(event.target.status == 200){
          if(event.target.responseText.trim() == "OK"){
            newPassForm.reset();
            notify("success", "Le mot de passe a été mis à jour.");
          }else{
            notify("failure", "Une erreur inattendue est survenue. Contactez l'administrateur.");
          }
        }else{
          notify("failure", "Le serveur a répondu par une erreur " + event.target.status + ".");
        }
      });

      ajax.addEventListener("error", function(event) {
        notify("failure", "Le serveur est injoignable.");
      });

      ajax.addEventListener("loadend", function(event) {
        setSubmitButton("send");
      });

      ajax.addEventListener("abort", function(event) {
        setSubmitButton("send");
        notify("failure", "Requête annulée.");
      });

      ajax.open("POST", "/admin/new-password");
      ajax.send(data);
    }
  }
}

function setSubmitButton(status){
  if(status == "send"){
    document.querySelector("#sendButton").classList.remove("is-hidden");
    document.querySelector("#sendingButton").classList.add("is-hidden");
  }
  if(status == "sending"){
    document.querySelector("#sendButton").classList.add("is-hidden");
    document.querySelector("#sendingButton").classList.remove("is-hidden");
  }
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
  });

  ajax.addEventListener("error", function(event) {
    modAlert("failure", "Le serveur est injoignable.");
  });

  ajax.addEventListener("abort", function(event) {
    modAlert("failure", "Requête annulée.");
  });

  ajax.open("DELETE", "/admin/login");
  ajax.send();
}

function notify(status, message){
  var notif = document.createElement("div");
  notif.classList.add("notification");
  if(status == "success"){
    notif.classList.add("is-success");
  }
  if(status == "failure"){
    notif.classList.add("is-danger");
  }

  var closeButton = document.createElement("button");
  closeButton.setAttribute("type", "button");
  closeButton.classList.add("delete");
  closeButton.addEventListener("click", function(){
    let notif = $(this).closest(".notification");
    animateCSS(notif, "zoomOut")
    .then(function(){
      notif.remove();
    })
  });

  var text = document.createTextNode(message);

  notif.appendChild(closeButton);
  notif.appendChild(text);

  document.querySelector(".notify").appendChild(notif);

  animateCSS(notif, "slideInUp");
}

function sendMail(token){
  let contactForm = document.getElementById("contactForm");
  if(!contactForm.checkValidity()){
    document.querySelector("#fakeSubmit").click();
  }else{
    setSubmitButton("sending");

    let data = new FormData(contactForm);
    data.append("g-recaptcha-response", token);

    let ajax = new XMLHttpRequest();
    ajax.addEventListener("load", function(event) {
      if(event.target.status == 200){
        if(event.target.responseText.trim() == "OK"){
          contactForm.reset();
          setSubmitButton("send");
          notify("success", "Votre message a bien été envoyé. Merci !");
        }else{
          setSubmitButton("send");
          notify("failure", "Une erreur inattendue est survenue. Contactez l'administrateur.");
        }
      }else{
        setSubmitButton("send");
        notify("failure", "Le serveur a répondu par une erreur " + event.target.status + ".");
      }
    });

    ajax.addEventListener("error", function(event) {
      setSubmitButton("send");
      notify("failure", "Le serveur est injoignable.");
    });

    ajax.addEventListener("abort", function(event) {
      setSubmitButton("send");
      notify("failure", "Requête annulée.");
    });

    ajax.open("POST", "/contact");
    ajax.send(data);
  }
}

function preventEnterKey(event){
  if(event.keyCode == 13 && document.activeElement.tagName != "TEXTAREA"){
    event.preventDefault();
    document.querySelector("#sendButton").click();
  }
}

function recaptchaHandler(action, callback){
  grecaptcha.ready(function() {
    grecaptcha.execute("6LeSGroeAAAAAAbqnDMMmYNVeAlWX_W4H-Eaobwj", {action: action}).then(callback);
  });
}