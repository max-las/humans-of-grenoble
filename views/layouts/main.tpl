<!DOCTYPE html>

<html>
  <head>
    <title>{{.PageTitle}}</title>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/animate.css/4.1.1/animate.min.css">
    <link rel="stylesheet" href="/static/css/bulma-custom.css">
    <link rel="stylesheet" href="/static/css/fa-all.min.css">
    <link rel="stylesheet" href="/static/css/fonts.css">
    <link rel="stylesheet" href="/static/css/style.css">
    {{ range .AdditionnalStyles }}
      <link rel="stylesheet" href="{{.}}">
    {{ end }}
    <script src="https://code.jquery.com/jquery-3.6.0.min.js"></script>
    <script src="/static/js/common.js"></script>
    {{ range .AdditionnalScripts }}
      <script src="{{.}}"></script>
    {{ end }}
  </head>
  <body>
    <nav class="navbar" role="navigation" aria-label="main navigation">

      <div class="navbar-brand is-family-playfair-display">
        <a class="navbar-item" href="/">
          <div class="block">
            <p class="is-size-3"><span class="has-text-primary">H</span>umans <span class="has-text-primary">O</span>f <span class="has-text-primary">G</span>renoble</p>
          </div>
        </a>

        <a role="button" class="navbar-burger" aria-label="menu" aria-expanded="false" data-target="navbarBasicExample">
          <span aria-hidden="true"></span>
          <span aria-hidden="true"></span>
          <span aria-hidden="true"></span>
        </a>
      </div>

      <div id="navbarBasicExample" class="navbar-menu">
        <div class="navbar-start">
          <a class="navbar-item is-family-sans-serif" href="/stories">
            <span>Stories</span>
          </a>

          <div class="navbar-item">
            <div class="dot"></div>
          </div>

          <a class="navbar-item is-family-sans-serif" href="/projet">
            <span>Projet</span>
          </a>

          <div class="navbar-item">
            <div class="dot"></div>
          </div>

          <a class="navbar-item is-family-sans-serif" href="/auteure">
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
      <div class="modal-card">
        <header class="modal-card-head">
          <p class="modal-card-title">Confirmation requise</p>
          <button class="delete closeModal" aria-label="close"></button>
        </header>
        <section class="modal-card-body"></section>
        <footer class="modal-card-foot">
          <button class="button confirmButton"></button>
          <button class="button closeModal">Annuler</button>
        </footer>
      </div>
    </div>

    {{.LayoutContent}}

  </body>
</html>
