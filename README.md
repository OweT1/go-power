# go-power

Just for fun :|

The projects can be found in the `projects/` folder.

To run the scripts of the different projects, we can simply do so by changing directory to the respective projects and running the `.go` files. This command has been encapsulated in the respective scripts in the `scripts/` folder.

As an example, if we want to run the project script for the url_checker (under `projects/url_checker/`), we can simply run:

```bash
scripts/run_url_checker.sh
```

## Library

For this project, we will be running an API server, hence running the script will simply start up the server at port 8080.
To view the Library we have currently, you can go to `http://localhost:8080/books` to check out the books in your Library.
We have implemented 2 functions/endpoints - Adding a Book and Deleting a Book.

To add a book, you can do something like:

```bash
curl -X POST http://localhost:8080/books \
-H "Content-Type: application/json" \
-d '{"id": "3", "title": "The Hobbit", "author": "Tolkien"}'
```

To delete a book, you can do something like:

```bash
curl -X DELETE http://localhost:8080/books/1
```

where the book(s) with `id` of value `1` will be deleted from our Library.
