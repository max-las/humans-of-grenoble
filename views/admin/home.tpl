<section class="section">
  <div class="container">
    <h1 class="title">Administration</h1>

    <div class="block">
      <a class="button is-success" href="/admin/new">
        <span class="icon">
          <i class="fas fa-plus"></i>
        </span>
        <span>Nouvelle story</span>
      </a>
      <a class="button is-warning" href="/admin/new-password">
        <span class="icon">
          <i class="fas fa-lock"></i>
        </span>
        <span>Changer de mot de passe</span>
      </a>
      <button class="button is-danger" onclick="logout()">
        <span class="icon">
          <i class="fas fa-power-off"></i>
        </span>
        <span>Déconnexion</span>
      </button>
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
                <a class="button is-info" href="/admin/edit/{{.Id}}">
                  <span class="icon">
                    <i class="fas fa-edit"></i>
                  </span>
                  <span>Éditer</span>
                </a>
                <button class="button is-danger" onclick="deleteStory(event)">
                  <span class="icon">
                    <i class="fas fa-trash-alt"></i>
                  </span>
                  <span>Supprimer</span>
                </button>
              </div>
            </div>
          </div>
        </div>
      {{ end }}
    </div>

  </div>
</section>
