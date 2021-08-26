<section class="section">

    <h1 class="title is-1 has-text-centered is-family-secondary"><span class="has-text-primary">S</span>tories</h1>

    <div class="columns">
      {{range $column := .Columns}}
        <div class="column">
          {{range $story := $column}}
            <div class="box has-background-light storyPreview" data-id="{{$story.Id}}">
              <img src="{{$story.PhotoUrl}}">
              <div class="content has-text-dark is-family-secondary">
                <p class="ellipsised is-italic display-new-lines">“{{$story.Text}}”</p>
              </div>
            </div>
          {{end}}
        </div>
      {{end}}
    </div>

</section>
