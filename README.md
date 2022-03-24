# Homework 3 Week 4

Simple application written in Go. Main focus is using GORM as ORM. Postgres is used as background

Its about bookLiblary that store books and author information.

Please edit .env file with your connection information.Iss filled with example  values

For the first usage please use -init flag to fill SQL with dumy values

> go run .\main.go -init

You can clear SQL after the test operations with -clear flag

>go run .\main.go -clear

> go run .\main.go -init

Then use -init again to repair & fill SQL with data

There is no any other flag available. Just edit comments inside main.go and have fun!

#Example functions for each domain (Author , book)

  - GetByID
  - FindByName
  - GetBooksWithAuthor
  - GetAuthorWithBooks






