<section class="section">
  <div class="container">
    <h1 class="title">Nouvelle Story</h1>

    <form id="storyForm" action="" method="post" enctype="multipart/form-data">

      <div class="field" id="imageField">
        <div class="control">
          <div class="file">
            <label class="file-label">
              <input class="file-input" id="imageFile" name="imageFile" type="file" accept="image/jpeg, image/png, image/gif" required>
              <span class="file-cta">
                <span class="file-icon">
                  <i class="fas fa-upload"></i>
                </span>
                <span class="file-label">
                  Choisir une image…
                </span>
              </span>
              <span class="file-name is-hidden"></span>
            </label>
          </div>
        </div>
      </div>

      <article class="message is-danger is-hidden" id="fileError">
        <div class="message-body"></div>
      </article>

      <div class="field is-hidden" id="preview">
        <div class="control">
          <figure class="image">
            <img src="about:blank" style="max-height: 500px; width: auto;">
          </figure>
        </div>
      </div>

      <div class="field">
        <label class="label">Texte</label>
        <div class="control">
          <textarea class="textarea" name="text" placeholder="Un beau jour..." required></textarea>
        </div>
      </div>

      <div class="field">
        <div class="control">
          <progress class="progress is-success is-hidden" value="0" max="100">0%</progress>
        </div>
      </div>

      <div class="field">
        <div class="control">
          <button type="button" id="sendButton" class="button is-outlined is-link" onclick="upload(event)">Envoyer</button>
          <button type="button" id="sendingButton" class="button is-outlined is-hidden">
            <span class="icon">
              <i class="fas fa-circle-notch fa-spin"></i>
            </span>
            <span>Envoi en cours...</span>
          </button>
        </div>
      </div>

      <input type="submit" class="is-hidden" id="fakeSubmit">

    </form>
  </div>
</section>
