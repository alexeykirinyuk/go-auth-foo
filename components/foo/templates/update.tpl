<div>
    <form method="POST">
        <label for="title">Title:</label><br/>
        <input id="title" type="text" name="title" value="{{.Title}}" required><br/>

        <label for="description">Description:</label><br/>
        <input id="description" type="text" name="description" value="{{.Description}}"><br/>

        <input type="submit">
    </form>
</div>