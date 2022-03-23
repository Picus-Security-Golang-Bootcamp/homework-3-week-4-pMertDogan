package author

import (
	// "errors"
	"fmt"

	"gorm.io/gorm"
)

//We will use gorm
type AuthorRepository struct {
	db *gorm.DB
}
//return our repo
func NewAuthorRepository (db *gorm.DB) *AuthorRepository {
	return &AuthorRepository{db: db}
}

//Migrate curent values if exist on current DB
func (c *AuthorRepository) Migrations() {
	c.db.AutoMigrate(&Author{})
	//https://gorm.io/docs/migration.html#content-inner
	//https://gorm.io/docs/migration.html#Auto-Migration
}

//Used to Insert data to SQL 
//Json to SQL :)
func (c *AuthorRepository) InsertSampleData(authors Authors ) {
	

	for _, author := range authors {
		c.db.Where(Author{AuthorID:  author.AuthorID}).
			Attrs(Author{AuthorID: author.AuthorID, Name: author.Name}).
			FirstOrCreate(&author)
	}
	
}

//Just type full book name
func (c *AuthorRepository) FindByName(authorName string) (*Author, error) {
	var author *Author
	//lke quert
	result := c.db.First(&author, "name like ?", "%"+fmt.Sprintf("%s", authorName)+"%")
	if result.Error != nil {
		return nil, result.Error
	}

	return author, nil
}

// func (c *BookRepository) GetByID(bookID string) (*Book, error) {
// 	var book *Book
// 	result := c.db.First(&book, "ID = ?", bookID)

// 	if result.Error != nil {
// 		return nil, result.Error
// 	}

// 	return book, nil
// }
