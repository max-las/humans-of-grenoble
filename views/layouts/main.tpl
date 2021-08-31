<!DOCTYPE html>

<html>

  {{ template "layouts/sections/head.tpl" . }}

  <body>

    {{ template "layouts/sections/navAndModals.tpl" . }}

    {{.LayoutContent}}

  </body>
</html>
