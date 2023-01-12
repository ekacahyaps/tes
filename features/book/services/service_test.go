package services

import (
	"api/features/book"
	"api/helper"
	"api/mocks"
	"errors"
	"testing"

	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	repo := mocks.NewBookData(t)

	t.Run("berhasil add buku", func(t *testing.T) {
		userID := 2
		inputData := book.Core{Judul: "kalkulus 1", Penulis: "Alfian", TahunTerbit: 2019}
		resData := book.Core{ID: 3, Judul: "kalkulus 1", Penulis: "Alfian", TahunTerbit: 2019}

		repo.On("Add", userID, inputData).Return(resData, nil)

		srv := New(repo)

		_, token := helper.GenerateJWT(2)

		pToken := token.(*jwt.Token)
		pToken.Valid = true

		res, err := srv.Add(pToken, inputData)
		assert.Nil(t, err)
		assert.Equal(t, resData.ID, res.ID)
		repo.AssertExpectations(t)
	})

	// t.Run("buku tidak ditemukan", func(t *testing.T) {
	// 	userID := 2
	// 	inputData := book.Core{Judul: "kalkulus 1", Penulis: "Alfian", TahunTerbit: 2019}
	// 	// resData := book.Core{ID: 3, Judul: "kalkulus 1", Penulis: "Alfian", TahunTerbit: 2019}

	// 	repo.On("Add", userID, inputData).Return(book.Core{}, err)

	// 	srv := New(repo)

	// 	_, token := helper.GenerateJWT(2)

	// 	pToken := token.(*jwt.Token)
	// 	pToken.Valid = true

	// 	res, err := srv.Add(pToken, inputData)
	// 	assert.Nil(t, err)
	// 	assert.Equal(t, resData.ID, res.ID)
	// 	repo.AssertExpectations(t)
	// })

	// t.Run("terjadi kesalahan pada server", func(t *testing.T) {
	// 	userID := 2
	// 	inputData := book.Core{Judul: "kalkulus 1", Penulis: "Alfian", TahunTerbit: 2019}
	// 	// resData := book.Core{ID: 3, Judul: "kalkulus 1", Penulis: "Alfian", TahunTerbit: 2019}

	// 	repo.On("Add", userID, inputData).Return(book.Core{}, errors.New("terjadi kesalahan pada server"))

	// 	srv := New(repo)

	// 	_, token := helper.GenerateJWT(2)

	// 	pToken := token.(*jwt.Token)
	// 	pToken.Valid = true

	// 	res, err := srv.Add(pToken, inputData)
	// 	assert.NotNil(t, err)
	// 	// assert.ErrorContains(t, err, "query error")
	// 	assert.Equal(t, res.UserID, uint(0))
	// 	repo.AssertExpectations(t)
	// })

	// t.Run("user tidak ditemukan", func(t *testing.T) {

	// })
}

func TestUpdate(t *testing.T) {
	repo := mocks.NewBookData(t)

	t.Run("sukses update data buku", func(t *testing.T) {
		inputData := book.Core{Judul: "Kalkulus 2", Penulis: "Putra", TahunTerbit: 2015}
		resData := book.Core{ID: 3, Judul: "Kalkulus 2", Penulis: "Putra", TahunTerbit: 2015}
		userID := 2
		bookId := 3

		repo.On("Update", bookId, userID, inputData).Return(resData, nil).Once()

		srv := New(repo)

		_, token := helper.GenerateJWT(2)

		pToken := token.(*jwt.Token)
		pToken.Valid = true

		res, err := srv.Update(pToken, bookId, inputData)
		assert.Nil(t, err)
		assert.Equal(t, res.ID, uint(bookId))
		repo.AssertExpectations(t)
	})

	t.Run("data buku tidak ditemukan", func(t *testing.T) {
		inputData := book.Core{Judul: "Kalkulus 2", Penulis: "Putra", TahunTerbit: 2015}
		userID := 2
		bookId := 3

		repo.On("Update", bookId, userID, inputData).Return(book.Core{}, errors.New("data not found")).Once()

		srv := New(repo)

		_, token := helper.GenerateJWT(2)

		pToken := token.(*jwt.Token)
		pToken.Valid = true

		res, err := srv.Update(pToken, bookId, inputData)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "tidak ditemukan")
		assert.Equal(t, uint(0), res.ID)
		repo.AssertExpectations(t)

	})

	t.Run("terdapat masalah pada server", func(t *testing.T) {
		inputData := book.Core{Judul: "Kalkulus 2", Penulis: "Putra", TahunTerbit: 2015}
		userID := 2
		bookId := 3

		repo.On("Update", bookId, userID, inputData).Return(book.Core{}, errors.New("terdapat masalah pada server")).Once()

		srv := New(repo)

		_, token := helper.GenerateJWT(2)

		pToken := token.(*jwt.Token)
		pToken.Valid = true

		res, err := srv.Update(pToken, bookId, inputData)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server")
		assert.Equal(t, uint(0), res.ID)
		repo.AssertExpectations(t)
	})

	t.Run("jwt tidak valid", func(t *testing.T) {
		inputData := book.Core{Judul: "Kalkulus 2", Penulis: "Putra", TahunTerbit: 2015}
		bookId := 3

		srv := New(repo)

		_, token := helper.GenerateJWT(0)

		pToken := token.(*jwt.Token)
		pToken.Valid = true

		res, err := srv.Update(pToken, bookId, inputData)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "tidak ditemukan")
		assert.Equal(t, uint(0), res.UserID)
		repo.AssertExpectations(t)
	})
}

func TestDelete(t *testing.T) {
	repo := mocks.NewBookData(t)

	t.Run("sukses delete buku", func(t *testing.T) {
		bookID := 3
		userID := 2
		repo.On("Delete", bookID, userID).Return(nil)

		srv := New(repo)

		_, token := helper.GenerateJWT(2)

		pToken := token.(*jwt.Token)
		pToken.Valid = true

		err := srv.Delete(pToken, bookID)
		assert.Nil(t, err)
		repo.AssertExpectations(t)

	})

	// t.Run("data tidak ditemukan", func(t *testing.T) {
	// 	bookID := 3
	// 	userID := 2
	// 	repo.On("Delete", bookID, userID).Return(errors.New("not found")).Once()

	// 	srv := New(repo)

	// 	_, token := helper.GenerateJWT(2)

	// 	pToken := token.(*jwt.Token)
	// 	pToken.Valid = true

	// 	err := srv.Delete(pToken, bookID)
	// 	assert.NotNil(t, err)
	// 	assert.ErrorContains(t, err, "tidak ditemukan")
	// 	repo.AssertExpectations(t)
	// })

	// t.Run("terdapat masalah pada server", func(t *testing.T) {
	// 	bookID := 3
	// 	userID := 2
	// 	repo.On("Delete", bookID, userID).Return(errors.New("terdapat masalah pada server")).Once()

	// 	srv := New(repo)

	// 	_, token := helper.GenerateJWT(2)

	// 	pToken := token.(*jwt.Token)
	// 	pToken.Valid = true

	// 	err := srv.Delete(pToken, bookID)
	// 	assert.NotNil(t, err)
	// 	assert.ErrorContains(t, err, "server")
	// 	repo.AssertExpectations(t)
	// })
}

func TestMyBook(t *testing.T) {
	repo := mocks.NewBookData(t)

	t.Run("sukses lihat koleksi buku", func(t *testing.T) {
		userID := 2
		resData := []book.Core{{ID: 3, Judul: "Kalkulus 2", Penulis: "Putra", TahunTerbit: 2015}, {ID: 4, Judul: "Kimia 2", Penulis: "Alfian", TahunTerbit: 2015}}
		repo.On("MyBook", userID).Return(resData, nil)

		srv := New(repo)

		_, token := helper.GenerateJWT(2)

		pToken := token.(*jwt.Token)
		pToken.Valid = true

		res, err := srv.MyBook(pToken)
		assert.Nil(t, err)
		assert.Equal(t, len(res), len(resData))
		repo.AssertExpectations(t)
	})

	// t.Run("data tidak ditemukan", func(t *testing.T) {
	// 	userID := 2
	// 	repo.On("MyBook", userID).Return([]book.Core{}, errors.New("data not found"))

	// 	srv := New(repo)

	// 	_, token := helper.GenerateJWT(2)

	// 	pToken := token.(*jwt.Token)
	// 	pToken.Valid = true

	// 	_, err := srv.MyBook(pToken)
	// 	assert.NotNil(t, err)
	// 	assert.ErrorContains(t, err, "tidak ditemukan")
	// 	// assert.Equal(t, len(res), int(0))
	// 	repo.AssertExpectations(t)
	// })

	// t.Run("terdapat masalah pada server", func(t *testing.T) {
	// 	userID := 2
	// 	repo.On("MyBook", userID).Return([]book.Core{}, errors.New("terdapat masalah pada server"))

	// 	srv := New(repo)

	// 	_, token := helper.GenerateJWT(2)

	// 	pToken := token.(*jwt.Token)
	// 	pToken.Valid = true

	// 	_, err := srv.MyBook(pToken)
	// 	assert.NotNil(t, err)
	// 	assert.ErrorContains(t, err, "server")
	// 	// assert.Equal(t, len(res), int(0))
	// 	repo.AssertExpectations(t)
	// })
}

func TestAllBooks(t *testing.T) {
	// repo := mocks.NewBookData(t)

	t.Run("sukses lihat koleksi semua buku", func(t *testing.T) {

	})

	t.Run("data tidak ditemukan", func(t *testing.T) {

	})

	t.Run("terdapat masalah pada server", func(t *testing.T) {

	})
}
