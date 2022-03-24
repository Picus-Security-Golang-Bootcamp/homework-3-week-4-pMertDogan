package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	postgres "github.com/pMertDogan/picusWeek4/common/db"
	"github.com/pMertDogan/picusWeek4/domain/author"
	"github.com/pMertDogan/picusWeek4/domain/book"
)

var resetAppDB = flag.Bool("reset", false, "Migrate tables and save default values thats readed by json files")
var dropTable = flag.Bool("dropTable", false, "Drop authors and books tables for clear SQL data")

//init app config
func init() {
	//Load env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file. \n ERROR : " + err.Error())
	}
	//parse reset flag or any other
	flag.Parse()

}

func main() {

	//make database connection
	db, errPG := postgres.ConnectPostgresDB()
	if errPG != nil {
		log.Fatal("Postgres cannot init: \n ", errPG)
	}
	log.Println("Postgres connected!!")

	//create repositorys for each domain
	authorRepository := author.NewAuthorRepository(db)
	bookRepository := book.NewBookRepository(db)

	//check if the user request it drop table
	//drop old tables to clean all

	if *dropTable {
		dropTables(authorRepository, bookRepository)
	}

	//check is user request reset using flag
	if *resetAppDB {

		resetDB(authorRepository, bookRepository)
	}

	//migrate database struct changes.
	// migrateDatabase(authorRepository, bookRepository)

	//save source data readed by json file to SQL
	// readFilesAndSaveThemToDB(authorRepository, bookRepository)

	//moved to postgres package
	// simple test functions
	// postgres.UpdateBookTest(bookRepository)
	// postgres.SoftDeleteTest(bookRepository)

	// SUPPORTED Methods example
	// b, err := bookRepository.FindByName("Lord")

	a, _ := bookRepository.GetByID("2")
	fmt.Println(a)
	bookRepository.UpdateBookQuantity("2", "21")
	a, _ = bookRepository.GetByID("2")
	fmt.Println(a)

	// softDeleteTest(bookRepository)
	// b, err := authorRepository.GetAuthorsWithBooks()
	// a, err := authorRepository.FindByName("J.R.R")
	// a, err := authorRepository.FindByName("Gogo")
	// a, err2 := authorRepository.GetByID("2")

	// bookRepository.UpdateBookQuantity("2","763")
	// if err2 != nil {
	// 	log.Fatal("Unable read all data " + err2.Error())
	// }
	//To test String override
	//if its return BOOKS not Book

	// b, err := bookRepository.GetBooksWithAuthors()
	// if err != nil {
	// 	log.Fatal("Unable read all data " + err.Error())
	// }
	// for _,v := range b{
	// 	fmt.Print(v)
	// }
	//fmt.Println(b)

	//Print uses String interface thats we owerwrite

}

//drop old tables to clean all
func dropTables(authorRepository *author.AuthorRepository, bookRepository *book.BookRepository) {
	err := postgres.DropTables(authorRepository, bookRepository)
	if err != nil {
		log.Fatal("Unable drop table" + err.Error())
	}
	os.Exit(1)
}

// migrate and save json data to sql
func resetDB(authorRepository *author.AuthorRepository, bookRepository *book.BookRepository) {
	fmt.Println("reset start")
	//create our DB struct on SQL
	postgres.MigrateDatabase(authorRepository, bookRepository)

	//store local data to sql
	postgres.ReadFilesAndSaveThemToDB(authorRepository, bookRepository)
	fmt.Println("reset end")
	os.Exit(0)
}
