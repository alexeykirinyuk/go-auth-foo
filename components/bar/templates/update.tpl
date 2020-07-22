<div>
    <form method="POST">
        <label for="title">Title:</label><br/>
        <input id="title" type="text" name="title" value="{{.Title}}" required><br/>

        <label for="description">Description:</label><br/>
        <input id="description" type="text" name="description" value="{{.Description}}"><br/>

        <label for="address">Address:</label><br/>
        <input id="address" type="text" name="address" required value="{{.Address}}"><br/>

        <label for="opening_date">Opening Date:</label><br/>
        <input id="opening_date" type="date" name="opening_date" value="{{.OpeningDate}}"><br/>

        <input type="submit">
    </form>
</div>