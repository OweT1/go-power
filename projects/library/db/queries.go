package db

const createTableQuery string = `
	CREATE TABLE IF NOT EXISTS books (
		id TEXT PRIMARY KEY,
		title TEXT,
		author TEXT
	);
`

const GetBooksQuery string = `
	SELECT
		id,
		title,
		author
	FROM books
`

const AddBookQuery string = `
	INSERT INTO books (id, title, author)
	VALUES (?, ?, ?)
`

const DeleteBookQuery string = `
	DELETE FROM books
	WHERE id = ?
`