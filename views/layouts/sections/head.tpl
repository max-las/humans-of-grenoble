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
