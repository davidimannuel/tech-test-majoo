package main

import (
	"context"
	"errors"

	"gorm.io/gorm"
)

/*
	Jawaban No 2:
	apa saja yang harus diperbaiki ? kita bisa menggunakan interface
	untuk mengabstraksi area sehingga tidak butuh type, berikut contoh code saya berikan di bawah
*/

type (
	// interface area
	IBangunDatar interface {
		GetArea() int64
		GetType() string
	}
	// model
	Area struct {
		ID        int64  `gorm:"column:id;primaryKey;"`
		AreaValue int64  `gorm:"column:area_value"`
		AreaType  string `gorm:"column:type"`
	}
	// repository
	AreaRepository struct {
		DB gorm.DB
	}
)

// simplify insert Area Function
func (_r *AreaRepository) InsertArea(ctx context.Context, bd IBangunDatar) (err error) {
	ar := Area{
		AreaValue: bd.GetArea(),
		AreaType:  bd.GetType(),
	}
	err = _r.DB.WithContext(ctx).Create(&ar).Error
	if err != nil {
		return err
	}
	return nil
}

// implement bangun datar interface example
type (
	Persegi struct {
		Sisi int64
	}

	PersegiPanjang struct {
		Panjang int64
		Lebar   int64
	}

	Segitiga struct {
		Alas   int64
		Tinggi int64
	}
)

func (p Persegi) GetArea() int64 {
	return p.Sisi * p.Sisi
}
func (p Persegi) GetType() string {
	return "persegi"
}

func (pj PersegiPanjang) GetArea() int64 {
	return pj.Panjang * pj.Lebar
}
func (pj PersegiPanjang) GetType() string {
	return "persegi_panjang"
}

func (s Segitiga) GetArea() int64 {
	return (s.Alas * s.Tinggi) / 2
}
func (s Segitiga) GetType() string {
	return "segitiga"
}

// usecase
type AreaUsecase struct {
	repo *AreaRepository
}

// call repository
func (_u AreaUsecase) InsertPersegi(ctx context.Context, p *Persegi) error {
	err := _u.repo.InsertArea(ctx, p)
	if err != nil {
		// log.Error().Msg(err.Error())
		err = errors.New("ERROR_DATABASE")
		return err
	}

	return nil
}

func (_u AreaUsecase) InsertPersegiPanjang(ctx context.Context, p *PersegiPanjang) error {
	err := _u.repo.InsertArea(ctx, p)
	if err != nil {
		// log.Error().Msg(err.Error())
		err = errors.New("ERROR_DATABASE")
		return err
	}

	return nil
}

func (_u AreaUsecase) InsertSegitiga(ctx context.Context, s *Segitiga) error {
	err := _u.repo.InsertArea(ctx, s)
	if err != nil {
		// log.Error().Msg(err.Error())
		err = errors.New("ERROR_DATABASE")
		return err
	}

	return nil
}
