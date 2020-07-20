<div>
    <form method="POST">
        <h1>User: {{.FirstName}}  {{.LastName}} ({{.Email}})</h1>

        <label for="role">Role:</label><br/>
        <select id="role" name="role">
            <option value="Member">Member</option>
            <option value="Admin">Admin</option>
        </select>

        <input type="submit">
    </form>
</div>