<nav class="navbar" role="navigation" aria-label="main navigation">

  <div class="navbar-brand is-family-playfair-display">
    <a class="navbar-item" href="/">
      <div class="block">
        <p class="is-size-3"><span class="has-text-primary">H</span>umans <span class="has-text-primary">O</span>f <span class="has-text-primary">G</span>renoble</p>
      </div>
    </a>

    <a role="button" class="navbar-burger" aria-label="menu" aria-expanded="false" data-target="mainNavbar">
      <span aria-hidden="true"></span>
      <span aria-hidden="true"></span>
      <span aria-hidden="true"></span>
    </a>
  </div>

  <div id="mainNavbar" class="navbar-menu">
    <div class="navbar-start">
      <a class="navbar-item" href="/stories">
        <span>Stories</span>
      </a>

      <a class="navbar-item" href="/projet">
        <span>Projet</span>
      </a>

      <a class="navbar-item" href="/auteure">
        <span>Auteure</span>
      </a>
    </div>

    <div class="navbar-end">
      <div class="navbar-item">
        <div class="buttons">
          <a class="button is-white">
            <span class="icon">
              <i class="fab fa-twitter"></i>
            </span>
          </a>
          <a class="button is-white">
            <span class="icon">
              <i class="fab fa-instagram"></i>
            </span>
          </a>
          <a class="button is-white">
            <span class="icon">
              <i class="fab fa-pinterest"></i>
            </span>
          </a>
        </div>
      </div>
    </div>

  </div>
</nav>

<div class="modal" id="successModal">
  <div class="modal-background"></div>
  <div class="modal-content">
    <article class="message is-success">
      <div class="message-header">
        <p>Succès</p>
        <button class="delete closeModal" aria-label="delete"></button>
      </div>
      <div class="message-body"></div>
    </article>
  </div>
</div>

<div class="modal" id="failureModal">
  <div class="modal-background"></div>
  <div class="modal-content">
    <article class="message is-danger">
      <div class="message-header">
        <p>Erreur</p>
        <button class="delete closeModal" aria-label="delete"></button>
      </div>
      <div class="message-body"></div>
    </article>
  </div>
</div>

<div class="modal" id="confirmModal">
  <div class="modal-background"></div>
  <div class="modal-card px-3">
    <header class="modal-card-head">
      <p class="modal-card-title"></p>
      <button class="delete closeModal" aria-label="close"></button>
    </header>
    <section class="modal-card-body"></section>
    <footer class="modal-card-foot">
      <button class="button confirmButton is-danger"></button>
      <button class="button closeModal">Annuler</button>
    </footer>
  </div>
</div>
