<section class="section">
  <div class="container">
    <h1 class="title">Contact</h1>

    <form id="contactForm" onkeydown="preventEnterKey(event);" action="" method="post">

      <div class="field">
        <label class="label">Nom</label>
        <div class="control">
          <input class="input" type="text" name="name" required>
        </div>
      </div>

      <div class="field">
        <label class="label">Email</label>
        <div class="control">
          <input class="input" type="email" name="email" required>
        </div>
      </div>

      <div class="field">
        <label class="label">Sujet</label>
        <div class="control">
          <input class="input" type="text" name="subject" required>
        </div>
      </div>

      <div class="field">
        <label class="label">Message</label>
        <div class="control">
          <textarea class="textarea" name="message" required></textarea>
        </div>
      </div>

      <div class="field">
        <div class="control">
          <button type="button" id="sendButton" class="button is-outlined is-link" onclick="sendMail(event)">Envoyer</button>
          <button type="button" id="sendingButton" class="button is-outlined is-hidden">
            <span class="icon">
              <i class="fas fa-circle-notch fa-spin"></i>
            </span>
            <span>Envoi en cours...</span>
          </button>
        </div>
      </div>

      <input type="submit" class="is-hidden" id="fakeSubmit">

      <div class="notify"></div>

    </form>
  </div>
</section>
