<form method="POST">
    <label for="first_name">First Name:</label><br/>
    <input id="first_name" type="text" name="first_name" required><br/>
    {{with .errors}}
        {{range .first_name}}
            <span>{{.}}</span>
            <br />
        {{end}}
    {{end -}}

    <label for="last_name">Last Name:</label><br/>
    <input id="last_name" type="text" name="last_name" required><br/>
    {{with .errors}}
        {{range .last_name}}
            <span>{{.}}</span>
            <br />
        {{end}}
    {{end -}}

    <label for="email">Email:</label><br/>
    <input id="email" type="text" name="email" required><br/>
    {{with .errors}}
        {{range .email}}
            <span>{{.}}</span>
            <br />
        {{end}}
    {{end -}}

    <label for="password">Password:</label><br/>
    <input id="password" type="password" name="password" required><br/>
    {{with .errors}}
        {{range .password}}
            <span>{{.}}</span>
            <br />
        {{end}}
    {{end -}}

    <label for="confirm_password">Confirm Password:</label><br/>
    <input id="confirm_password" type="password" name="confirm_password" required><br/>
    {{with .errors}}
        {{range .confirmed_password}}
            <span>{{.}}</span>
            <br />
        {{end}}
    {{end -}}
    <input type="submit">
</form>
