package book

type Service interface {
	FindAll() ([]Book, error)
	FindById(id int) (Book, error)
	Create(bookInput BookRequest) (Book, error)
	Update(ID int, bookInput BookRequest) (Book, error)
	Delete(ID int) (Book, error)
}
type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]Book, error) {
	return s.repository.FindAll()
}

func (s *service) FindById(id int) (Book, error) {
	return s.repository.FindById(id)
}

func (s *service) Create(bookInput BookRequest) (Book, error) {
	price, _ := bookInput.Price.Int64()

	book := Book{
		Title:       bookInput.Title,
		Price:       int(price),
		Description: bookInput.Description,
		Rating:      bookInput.Rating,
		Discount:    bookInput.Discount,
	}

	return s.repository.Create(book)
}

func (s *service) Update(ID int, bookInput BookRequest) (Book, error) {
	book, _ := s.repository.FindById(ID)
	price, _ := bookInput.Price.Int64()

	book.Title = bookInput.Title
	book.Price = int(price)
	book.Description = bookInput.Description
	book.Rating = bookInput.Rating
	book.Discount = bookInput.Discount

	return s.repository.Update(book)
}

func (s *service) Delete(ID int) (Book, error) {
	book, _ := s.repository.FindById(ID)

	return s.repository.Delete(book)
}
