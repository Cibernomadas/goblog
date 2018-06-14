{{ define "content" }}
<form method="POST" action="/register">
    <label>Username</label>
    <input type="text" name="username">
    <br>
    <label>Email</label>
    <input type="email" name="email">
    <br>
    <label>Password</label>
    <input type="password" name="passworda">
    <br>
    <label>Repeat Password</label>
    <input type="password" name="passwordb">
    <br>
    <button type="submit">Register</button>
</form>
{{ end }}
