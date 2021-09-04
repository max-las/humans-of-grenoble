document.addEventListener("DOMContentLoaded", function(){
  barba.init({
    transitions: [{
      name: 'main-to-main',
      to: {
        custom: function(data){
          return data.current.container.querySelector("#mainNavbar") !== null && data.next.container.querySelector("#mainNavbar") !== null;
        }
      },
      leave(data) {
        return gsap.to(data.current.container.querySelector(".barba-content"), {
          opacity: 0
        });
      },
      enter(data) {
        return gsap.from(data.next.container.querySelector(".barba-content"), {
          opacity: 0
        });
      }
    }, {
      name: 'home-main',
      to: {
        custom: function(data){
          var navBefore = data.current.container.querySelector("#mainNavbar") !== null;
          var navAfter = data.next.container.querySelector("#mainNavbar") !== null;
          return !(navBefore && navAfter);
        }
      },
      leave(data) {
        return gsap.to(data.current.container, {
          opacity: 0
        });
      },
      enter(data) {
        return gsap.from(data.next.container, {
          opacity: 0
        });
      }
    }],
    views: [{
      namespace: 'main',
      beforeEnter(data) {
        if(data.next.container.querySelector("#mainNavbar") !== null){
          initNavAndModals();
        }
        if(data.next.container.querySelector("#storyForm") !== null){
          initEdit();
        }
      }
    }]
  });

  initNavAndModals();
});

function initNavAndModals(){
  $(".navbar-burger").on("click", function() {
      $(".navbar-burger").toggleClass("is-active");
      $(".navbar-menu").toggleClass("is-active");
  });

  $(".closeNotification").on("click", function(){
    $(this).closest(".notification").hide();
  });

  $(".closeModal, .modal-background, #confirmModal .confirmButton").on("click", function(){
    $("#confirmModal .confirmButton").off("click.tmp");
    $(this).closest(".modal").removeClass("is-active");
  });
}

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

function modalConfirm(title, content, confirmText, callback){
  $("#confirmModal .modal-card-title").text(title);
  $("#confirmModal .modal-card-body").html(content);

  var confirmButton = $("#confirmModal .confirmButton");
  confirmButton.html(confirmText);

  confirmButton.one("click.tmp", callback);

  $("#confirmModal").addClass("is-active");
}
