<section class="section">
  <div class="container">
    <h1 class="title">Changement de mot de passe</h1>

    <form id="newPassForm" action="" method="post" enctype="multipart/form-data">

      <div class="field">
        <label class="label">Nouveau mot de passe</label>
        <div class="control">
          <input class="input" type="password" name="password" required>
        </div>
      </div>

      <div class="field">
        <label class="label">Confirmation nouveau mot de passe</label>
        <div class="control">
          <input class="input" type="password" name="password-confirm" required>
        </div>
      </div>

      <div class="field">
        <div class="control">
          <button type="button" id="sendButton" class="button is-outlined is-link" onclick="updatePassword(event)">Valider</button>
        </div>
      </div>

      <input type="submit" class="is-hidden" id="fakeSubmit">

    </form>
  </div>
</section>
