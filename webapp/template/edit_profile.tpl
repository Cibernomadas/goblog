{{define "content"}}
<h1>Edit Profile</h1>
<form action="/edit_profile" method="post">
    <p>
        <label>Username</label><br>
        <input type="text" name="username"><br>
    </p>
    <p>
        <label>About Me</label><br>
        <textarea cols="50" rows="5" name="about_me"></textarea><br>
    </p>
    <p><button type="submit">Update</button></p>
</form>
{{end}}
