{{ define "menu" }}
<div>
    GoBlog
    <a href="/">Homepage</a>
    {{ if .user }}
    <a href="/user/{{ .user.Username }}">Profile</a>
    <a href="/logout">Logout</a>
    {{ else }}
    <a href="/login">Login</a>
    <a href="/register">Register</a>
    {{ end }}
</div>
{{ end }}
