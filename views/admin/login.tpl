<section class="section">
  <div class="container">
    <h1 class="title">Identification</h1>

    <form action="" method="post">
      <div class="field">
        <label class="label">Identifiant</label>
        <div class="control">
          <input class="input" type="text" name="username" required>
        </div>
      </div>

      <div class="field">
        <label class="label">Mot de passe</label>
        <div class="control">
          <input class="input" type="password" name="password" required>
        </div>
      </div>

      {{if .Success}}
        <div class="notification is-success">
          <button type="button" class="delete closeNotification"></button>
          {{.Success}}
        </div>
      {{end}}

      {{if .Error}}
        <div class="notification is-danger">
          <button type="button" class="delete closeNotification"></button>
          {{.Error}}
        </div>
      {{end}}

      <div class="field">
        <div class="control">
          <button class="button is-outlined is-link">Connexion</button>
        </div>
      </div>
    </form>
  </div>
</section>
