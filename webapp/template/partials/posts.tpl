{{ define "posts" }}
    {{ range $post := . }}
    <table>
        <tr valign="top">
            <td><img src="{{ $post.Author.Avatar 36 }}"></td>
            <td>{{ $post.Author.Username }} says:<br>{{ $post.Body }}</td>
        </tr>
    </table>
    {{ end }}
{{ end }}
