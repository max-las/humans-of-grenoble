$(function() {
  $(".navbar-burger").click(function() {
      $(".navbar-burger").toggleClass("is-active");
      $(".navbar-menu").toggleClass("is-active");
  });

  $(".closeNotification").click(function(){
    $(this).closest(".notification").hide();
  });

  $(".closeModal, .modal-background, #confirmModal .confirmButton").click(function(){
    $("#confirmModal .confirmButton").off("click.tmp");
    $(this).closest(".modal").removeClass("is-active");
  });
});

const animateCSS = (element, animation, prefix = 'animate__') => new Promise((resolve, reject) => {
  const animationName = `${prefix}${animation}`;
  const node = $(element);

  node.addClass([`${prefix}animated`, animationName]);

  function handleAnimationEnd(event) {
    event.stopPropagation();
    node.removeClass([`${prefix}animated`, animationName]);
    resolve('Animation ended');
  }

  node.one('animationend', handleAnimationEnd);
});

function modAlert(status, message){
  if(status == "success"){
    $("#successModal .message-body").html(message);
    $("#successModal").addClass("is-active");
  }
  if(status == "failure"){
    $("#failureModal .message-body").html(message);
    $("#failureModal").addClass("is-active");
  }
}

function modalConfirm(message, status, confirmText, callback){
  $("#confirmModal .modal-card-body").html(message);

  var confirmButton = $("#confirmModal .confirmButton");
  confirmButton.html(confirmText);
  confirmButton.attr("class", "button confirmButton");
  if(status == "info"){
    confirmButton.addClass("is-info");
  }
  if(status == "danger"){
    confirmButton.addClass("is-danger");
  }

  confirmButton.one("click.tmp", callback);

  $("#confirmModal").addClass("is-active");
}
