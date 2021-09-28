<!DOCTYPE html>

<html>

  {{ template "layouts/sections/head.tpl" . }}

  <body data-barba="wrapper">

    <main data-barba="container" data-barba-namespace="main">

      {{ template "layouts/sections/navAndModals.tpl" . }}

      <div class="barba-content">

        <div class="minHeightContent">

          {{.LayoutContent}}

        </div>

      </div>

      {{ template "layouts/sections/footer.tpl" . }}

    </main>

  </body>
</html>
