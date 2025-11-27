#!/bin/bash

curl -X POST http://localhost:8080/books \
-H "Content-Type: application/json" \
-d '{"id": "1", "title": "The Go Programming Language", "author": "Alan Donovan"}'

curl -X POST http://localhost:8080/books \
-H "Content-Type: application/json" \
-d '{"id": "2", "title": "Clean Code", "author": "Robert Martin"}'