package service

import "main/repository"

type AdminServie interface {
	Getauthor() ([]repository.RespAuthor, error)
	GenreDetail() ([]repository.RespLoan, error)
	UserDetail() ([]repository.RespUser, error)
}

type adminServie struct {
	repo repository.AdminRepository
}

func NewadminService(r repository.AdminRepository) AdminServie {
	return &adminServie{r}
}

func (s *adminServie) Getauthor() ([]repository.RespAuthor, error) {
	return s.repo.GetAuthor()
}

func (s *adminServie) GenreDetail() ([]repository.RespLoan, error) {
	return s.repo.GenreDetail()
}

func (s *adminServie) UserDetail() ([]repository.RespUser, error) {
	return s.repo.UserDetail()
}
