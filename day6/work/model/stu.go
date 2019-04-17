package model

import "github.com/pkg/errors"

var (
	ErrNotFoundBook = errors.New("not fount book")
)

type Student struct {
	Name string
	Grade string
	Id string
	Sex string
	books []*BorrowItem
}

type BorrowItem struct {
	book *Book
	num int
}

func CreateStudent(name, grade, id, sex string) *Student {
	stu := &Student{
		Name:name,
		Grade:grade,
		Id:id,
		Sex:sex,
	}
	return stu
}

func (s *Student) AddBook(b *BorrowItem) {
	s.books = append(s.books, b)
}

func (s *Student) returnBook(b *BorrowItem) (err error) {
	for i, v := range s.books {
		if v.book.Name == b.book.Name {
			if v.num == b.num {
				l := s.books[0:i]
				r := s.books[i+1:]
				l = append(l, r...)
				s.books = l
				return
			}
			s.books[i].num -= b.num
			return
		}
	}
	err = ErrNotFoundBook
	return
}

func (s *Student) GetBookList() []*BorrowItem {
	//todo
	return nil
}