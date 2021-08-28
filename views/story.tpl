<section class="section">
  <div class="container">
    {{ if .JustAdded }}
      <div class="notification is-success">
        <button type="button" class="delete closeNotification"></button>
        Story publiée avec succès !
      </div>
    {{ end }}
    <div class="tile is-ancestor">
      <div class="tile is-vertical is-parent px-5">
        <div class="tile is-child is-flex is-justify-content-center is-align-items-center">
          <img src="/static/img/ornament.svg" style="max-height: 30px;">
        </div>
        <div class="tile is-child">
          <figure class="image">
            <img src="{{.PhotoUrl}}">
          </figure>
        </div>
        <div class="tile is-child is-flex is-justify-content-center is-align-items-center">
          <img src="/static/img/ornament.svg" style="max-height: 30px; transform: scale(-1, -1);">
        </div>
      </div>
      <div class="tile is-parent px-5">
        <div class="tile is-child is-flex is-align-items-center">
          <p class="display-new-lines">“{{.Text}}”</p>
        </div>
      </div>
    </div>
  </div>
</section>
