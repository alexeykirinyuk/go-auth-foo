<script>
    function createItem() {
        window.location.href = "/bar/create";
    }

    function updateItem(id) {
        window.location.href = "/bar/update/" + id;
    }
</script>
<div>
    <button type="button" onclick="createItem()">Create</button>
    <button type="button" onclick="window.location.href = '/'">Back</button>
</div>
<div>
    <table>
        <thead>
        <th>Title</th>
        <th>Description</th>
        <th>Address</th>
        <th>Opening Date</th>
        <th></th>
        <th></th>
        </thead>
        <tbody>
        {{range .}}
            <tr>
                <td>{{.Title}}</td>
                <td>{{.Description}}</td>
                <td>{{.Address}}</td>
                <td>{{.OpeningDate}}</td>
                <td><button type="button" onclick="updateItem({{.Id}})">Update</button></td>
                <td>
                    <form method="POST" action="bar/delete/{{.Id}}">
                        <button type="submit">Delete</button>
                    </form>
                </td>
            </tr>
        {{end}}
        </tbody>
    </table>
</div>