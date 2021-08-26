@echo off

REM Watches for changes in sass/bulma-custom.scss and reflect them in static/css/bulma-custom.css
REM Replace "..\dart-sass\sass" with path to sass executable

..\dart-sass\sass --watch --no-source-map sass/bulma-custom.scss:static/css/bulma-custom.css
