{{define "content"}}
<div>
    <table>
        <tr valign="top">
            <td><img src="{{ .user.Avatar 128 }}"></td>
            <td>
                <h1>User: {{ .user.Username }}</h1>
                {{ if .user.AboutMe }}
                <p>{{ .user.AboutMe }}</p>
                {{ end }}
                {{ if .user.LastSeen }}
                <p>Last seen on: {{ .user.LastSeen }}</p>
                {{ end }}
                {{ if eq .current_user.Username .user.Username }}
                <p><a href="/edit_profile">Edit your profile</a></p>
                {{ end }}
            </td>
        </tr>
    </table>
    <hr>
    {{ template "posts" .posts }}
</div>
{{end}}
