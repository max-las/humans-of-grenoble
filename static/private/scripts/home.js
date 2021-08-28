function deleteStory(event){
  let id = $(event.target).closest(".box").attr("data-id");
  let selector = "[data-id='" + id + "']";

  var modalContent = $(selector + " img").prop("outerHTML");

  modalConfirm("Supprimer cette story ?", modalContent, "Supprimer", function(){
    let ajax = new XMLHttpRequest();

    ajax.addEventListener("load", function(event) {
      if(event.target.status == 200){
        animateCSS(selector, "zoomOut").then((message) => {
          $(selector).css("opacity", "0");
          $(selector).animate({
            height: "0px",
          }, 200, "swing", function(){
            $(selector).remove();
          });
        });
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
