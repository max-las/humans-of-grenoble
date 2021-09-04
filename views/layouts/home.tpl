<!DOCTYPE html>

<html>

  {{ template "layouts/sections/head.tpl" . }}

  <body data-barba="wrapper">

    <main data-barba="container" data-barba-namespace="home">

      {{.LayoutContent}}

    </main>

  </body>
</html>
