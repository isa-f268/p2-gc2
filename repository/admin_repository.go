package repository

import (
	"fmt"

	"gorm.io/gorm"
)

type RespAuthor struct {
	Name      string
	Published int
}

type RespLoan struct {
	GenreName string `gorm:"column:genre_name"`
	Loaner    int
}

type RespUser struct {
	Name string
	Loan int
}

type AdminRepository interface {
	GetAuthor() ([]RespAuthor, error)
	GenreDetail() ([]RespLoan, error)
	UserDetail() ([]RespUser, error)
}

type adminRepository struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) AdminRepository {
	return adminRepository{db}
}

func (r adminRepository) GetAuthor() ([]RespAuthor, error) {
	var author []RespAuthor

	resp := r.db.Raw(`
	SELECT CONCAT(a.first_name,' ', a.last_name)as Name, COUNT(*) as Published
	FROM authors a
	JOIN books b ON a.id = b.author_id
	GROUP BY a.id
`).Scan(&author)

	if resp.Error != nil {
		return []RespAuthor{}, fmt.Errorf("get author list error")
	}

	return author, nil
}

func (r adminRepository) GenreDetail() ([]RespLoan, error) {
	var genre []RespLoan
	resp := r.db.Raw(`
	select genre_name,count(*) as Loaner from loans a
	join books b on a.book_id=b.id
	join genres c on b.genre_id=c.id
	group by genre_id,genre_name
	`).Scan(&genre)

	if resp.Error != nil {
		return []RespLoan{}, fmt.Errorf("genre detail error")
	}

	return genre, nil
}

func (r adminRepository) UserDetail() ([]RespUser, error) {
	var users []RespUser
	resp := r.db.Raw(`
	select concat(first_name,' ',last_name)as Name, count(*) as Loan 
	from users a
	join loans b on a.id= b.user_id
	group by Name
	order by Loan DESC
	`).Scan(&users)
	if resp.Error != nil {
		return []RespUser{}, fmt.Errorf("user detail error")
	}
	return users, nil

}
