<section class="section">
  <div class="container">
    <h1 class="title">Administration</h1>

    <div class="block">
      <a class="button is-success" href="/admin/new">Nouvelle story</a>
      <a class="button is-danger" href="/admin/logout">Déconnexion</a>
    </div>

    <div class="block">
      {{ range .Stories }}
        <div class="box" data-id="{{.Id}}">
          <div class="columns">
            <div class="column is-one-quarter">
              <img src="{{.PhotoUrl}}">
            </div>
            <div class="column is-flex is-align-items-center">
              <p class="ellipsised display-new-lines">{{.Text}}</p>
            </div>
            <div class="column is-one-quarter is-flex is-justify-content-center is-align-items-center">
              <div class="buttons">
                <a class="button is-info" href="/admin/edit/{{.Id}}">Éditer</a>
                <button class="button is-danger" onclick="deleteStory(event)">Supprimer</button>
              </div>
            </div>
          </div>
        </div>
      {{ end }}
    </div>

  </div>
</section>
