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
func NewAuthorRepository(db *gorm.DB) *AuthorRepository {
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
func (c *AuthorRepository) InsertSampleData(authors Authors) {

	for _, author := range authors {
		c.db.Where(Author{AuthorID: author.AuthorID}).
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

func (c *AuthorRepository) GetByID(authorID string) (*Author, error) {
	var author *Author
	result := c.db.First(&author, "Author_ID = ?", authorID)

	if result.Error != nil {
		return nil, result.Error
	}

	return author, nil
}

func (b *AuthorRepository) GetAuthorsWithBooks() (Authors, error) {
	var authors Authors
	//get author books with join query
	result := b.db.Joins("JOIN books ON books.author_ID = authors.Author_ID ").Find(&authors)

	if result.Error != nil {
		return nil, result.Error
	}
	return authors, nil
}

func (b *AuthorRepository) GetAuthorWithBooks(authorID string) (Authors, error) {
	var authors Authors

	//AND emails.email = ?", "jinzhu@example.org
	//get author books with join query
	result := b.db.Joins("JOIN books ON books.author_ID = authors.Author_ID AND authors.author_id = ?", authorID).Find(&authors)
	// result := b.db.Raw("select * from authors JOIN books ON books.author_ID = authors.Author_ID and authors.author_id = ?", authorID).Scan(&authors)
	fmt.Println(result)
	if result.Error != nil {
		return nil, result.Error
	}
	return authors, nil
}

//Drop author with raw sql
func (b *AuthorRepository) DropTable() error {
	result := b.db.Exec("DROP TABLE authors")
	// b.db.dr
	if result.Error != nil {
		return result.Error
	}

	return nil

}
