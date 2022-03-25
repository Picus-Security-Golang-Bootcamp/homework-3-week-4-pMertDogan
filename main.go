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

var resetAppDB = flag.Bool("init", false, "Migrate tables and save default values thats readed by json files")
var dropTable = flag.Bool("clear", false, "Drop authors and books tables for clear SQL data")





//init env and parse flags
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

	//make database connection with GORM
	db, errPG := postgres.ConnectPostgresDB()
	if errPG != nil {
		log.Fatal("Postgres cannot init: \n ", errPG)
	}
	log.Println("Postgres connected!!")

	//create repositorys for each domain
	authorRepository := author.NewAuthorRepository(db)
	bookRepository := book.NewBookRepository(db)

	//check if the user add  droptable flag
	if *dropTable {
		//drop old tables to clean all
		dropTables(authorRepository, bookRepository)
	}

	//check is user add reset  flag
	if *resetAppDB {
		//migrate and json to SQL 
		resetDB(authorRepository, bookRepository)
	}

	//migrate database struct changes.
	// postgres.MigrateDatabase(authorRepository, bookRepository)

	//save source data readed by json file to SQL
	// postgres.ReadFilesAndSaveThemToDB(authorRepository, bookRepository)

	//moved to postgres package
	// simple test functions
	// postgres.UpdateBookTest(bookRepository)
	// postgres.SoftDeleteTest(bookRepository)

	// SUPPORTED Methods example, to test just uncomment them and fix typo :)
	// b, err := bookRepository.FindByName("Lord")

	// a, _ := bookRepository.GetByID("2")
	// fmt.Println(a)
	// bookRepository.UpdateBookQuantity("2", "763")
	// a, _ = bookRepository.GetByID("2")
	// fmt.Println(a)

	// b, err := authorRepository.GetAuthorsWithBooks()
	// search, _ := authorRepository.FindByName("J.R.")
	// fmt.Print(search)
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



	// v , e := bookRepository.GetBooksWithAuthors()
	// v , e := bookRepository.GetAllBooks()
	// v , e := authorRepository.GetAuthorsWithBooks()
	v , e := authorRepository.GetAuthorWithBooks("1")
	
	if e != nil {
		log.Fatal(e)
	}
	for _,x := range v{
			fmt.Print(x)
		}

	// fmt.Print(v)

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
	fmt.Println("tip run main with -reset")
	os.Exit(0)
}
