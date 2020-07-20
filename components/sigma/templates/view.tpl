<script>
    function createItem() {
        window.location.href = "/sigma/create";
    }

    function updateItem(id) {
        window.location.href = "/sigma/update/" + id;
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
        <th></th>
        <th></th>
        </thead>
        <tbody>
        {{range .}}
            <tr>
                <td>{{.Info}}</td>
                <td><button type="button" onclick="updateItem({{.Id}})">Update</button></td>
                <td>
                    <form method="POST" action="sigma/delete/{{.Id}}">
                        <button type="submit">Delete</button>
                    </form>
                </td>
            </tr>
        {{end}}
        </tbody>
    </table>
</div>