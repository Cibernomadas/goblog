<!doctype html>
<html>
    <head>
        {{ include "layouts/header" }}
    </head>
    <body>
        {{ template "menu" . }}
        <hr>
        {{ if .error }}
            {{ template "error" . }}
        {{ end }}
        {{ template "content" . }}
        <hr>
        {{ include "layouts/footer" }}
    </body>
</html>
