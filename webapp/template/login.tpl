{{define "content"}}
<div>
    <h1>Sign In</h1>
    <form method="POST" action="/login">
        <label>Username:</label>
        <input type="text" name="username">
        <br>
        <label>Password:</label>
        <input type="password" name="password">
        <br>
        <button type="submit">Sign In</button>
    </form>
</div>
{{end}}
