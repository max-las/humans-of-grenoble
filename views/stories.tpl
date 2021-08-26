<section class="section">

    <h1 class="title is-1 has-text-centered is-family-playfair-display"><span class="has-text-primary">S</span>tories</h1>

    <div class="columns">
      {{range $column := .Columns}}
        <div class="column">
          {{range $story := $column}}
            <div class="box has-background-light storyPreview">
              <a href="/story/{{$story.Id}}">
                <img src="{{$story.PhotoUrl}}">
                <div class="content has-text-dark is-family-lora">
                  <p class="ellipsised display-new-lines">“{{$story.Text}}”</p>
                </div>
              </a>
            </div>
          {{end}}
        </div>
      {{end}}
    </div>

</section>
