package service

import (
	"errors"
	"golang-practice-ma/internal/model"
	"golang-practice-ma/internal/store"
)

type Logger interface {
	Log(msg, errStr string)
}

type Service struct {
	store  store.Store
	logger Logger
}

func New(s store.Store, l Logger) *Service {
	return &Service{
		store:  s,
		logger: l,
	}
}

func (s *Service) ObtenTodosLosLibros() ([]*model.Libro, error) {
	if s.logger != nil {
		s.logger.Log("Estamos obteniendo los libros", "")
	}

	libros, err := s.store.GetAll()
	if err != nil {
		if s.logger != nil {
			s.logger.Log("Error al obtener libros", err.Error())
		}
		return nil, err
	}

	return libros, nil
}

func (s *Service) ObtenLibroPorID(id int) (*model.Libro, error) {
	return s.store.GetBYID(id)
}

func (s *Service) CrearLibro(libro model.Libro) (*model.Libro, error) {
	if libro.Titulo == "" || libro.Autor == "" {
		return nil, errors.New("no se puede añadir libro o autor sin nombre")
	}

	return s.store.Create(&libro)
}

func (s *Service) ActualizarLibro(id int, libro model.Libro) (*model.Libro, error) {
	if libro.Titulo == "" || libro.Autor == "" {
		return nil, errors.New("no se puede añadir libro o autor sin nombre")
	}

	return s.store.Update(id, &libro)
}

func (s *Service) EliminarLibro(id int) error {
	return s.store.Delete(id)
}
