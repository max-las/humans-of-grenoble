@echo off

REM Watches for changes in sass/bulma-custom.scss and reflect them in static/css/bulma-custom.css

sass --watch --no-source-map --style=compressed sass/bulma-custom.scss:static/css/bulma-custom.css
