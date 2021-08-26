function deleteStory(event){
  let id = $(event.target).closest(".box").attr("data-id");

  modalConfirm("Supprimer cette story ?", "danger", "Supprimer", function(){
    let ajax = new XMLHttpRequest();

    ajax.addEventListener("load", function(event) {
      if(event.target.status == 200){
        let selector = "[data-id='" + id + "']";
        animateCSS(selector, "zoomOut").then((message) => {
          $(selector).remove();
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
