<div>
    <form method="POST">
        <h1>User: {{.User.FirstName}}  {{.User.LastName}} ({{.User.Email}})</h1>

        <label for="role">Role:</label><br/>
        <select id="role" name="role">
            {{range $role := .Roles}}
                {{if $role.Selected}}
                    <option value="{{$role.Name}}" selected>{{$role.Name}}</option>
                {{else}}
                    <option value="{{$role.Name}}">{{$role.Name}}</option>
                {{end}}
            {{end}}
        </select>

        <input type="submit">
    </form>
</div>