<script>
    function updateItem(id) {
        window.location.href = "/user/" + id;
    }
</script>
<div>
    <button type="button" onclick="window.location.href = '/'">Back</button>
</div>
<div>
    <table>
        <thead>
        <th>Full Name</th>
        <th>Email</th>
        <th>Role</th>
        <th></th>
        </thead>
        <tbody>
        {{range .}}
            <tr>
                <td>{{.FirstName}} {{.LastName}}</td>
                <td>{{.Email}}</td>
                <td>{{.Role}}</td>
                <td><button type="button" onclick="updateItem({{.Id}})">Change Role</button></td>
            </tr>
        {{end}}
        </tbody>
    </table>
</div>