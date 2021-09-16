package book

type Book struct {
	Id          int `gorm:"primary_key, AUTO_INCREMENT"`
	Name        string
	Price       string
	Description string
}

func (book *Book) TableName() string {
	return "book"
}
