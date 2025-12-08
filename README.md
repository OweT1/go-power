# go-power

Just for fun :|

The projects can be found in the `projects/` folder.

To run the scripts of the different projects, we can simply do so by changing directory to the respective projects and running the `.go` files. This command has been encapsulated in the respective scripts in the `scripts/` folder.

As an example, if we want to run the project script for the url_checker (under `projects/url_checker/`), we can simply run:

```bash
scripts/run_url_checker.sh
```

## Environmental Variables

We will require the use of environmental variables for some of our projects.
To get a copy of the `.env` file, just simply run:

```powershell
cp .env.example .env
```

## Project - Library

### Set-up Instructions & App Details

For this project, we will be running an API server, hence running the `run_library.sh` script will simply start up the server at port 8080.
To view the Library we have currently, you can go to `http://localhost:8080/books` to check out the books in your Library.

However, you may realise that your library currently does not have any books.
Hence, in order add some books to your library, you will need to run an ingestion script.
On another bash terminal, run the script:

```bash
scripts/populate_library.sh
```

This should populate your library with 2 books.

> [!NOTE]
> In order to run the above script, you must be logged in.
> Details of registering and logging in can be found in the following section.

### Endpoint Details

#### Authentication

For authentication purposes, we have 2 endpoints - Registering and Logging In.

To register a user, you can do:

```bash
curl -X POST http://localhost:8080/register \
-H "Content-Type: application/json" \
-d '{"username": "test_user", "password": "test_password"}'
```

To login, you can then do:

```bash
curl -X POST http://localhost:8080/login \
-H "Content-Type: application/json" \
-d '{"username": "test_user", "password": "test_password"}'
```

By doing the above, you will then find a file called `jwt_token.txt` created.
Inside contains the `jwt_token` required to add and delete books.

#### Books

We have implemented 2 API endpoints - Adding and Deleting a Book.

To add a book, you can do something like:

```bash
curl -X POST http://localhost:8080/books \
-H "Authorization: Bearer $(cat "projects/library/jwt_token.txt")" \
-H "Content-Type: application/json" \
-d '{"id": "3", "title": "The Hobbit", "author": "Tolkien"}'
```

To delete a book, you can do something like:

```bash
curl -X DELETE http://localhost:8080/books/1 \
-H "Authorization: Bearer $(cat "projects/library/jwt_token.txt")"
```

> [!NOTE]
> To run the above commands, you must be logged in and can only run the commands from the main directory (`go-power/`)
> If not, there will be directory issues with finding the `jwt_token.txt` file.

where the book(s) with `id` of value `1` will be deleted from our Library.
