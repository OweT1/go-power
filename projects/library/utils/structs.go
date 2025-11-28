package utils

import "gorm.io/gorm"

type Book struct {
	// gorm.Model adds: ID (uint), CreatedAt, UpdatedAt, DeletedAt
	gorm.Model

	Title  string `json:"title" gorm:"not null"`
	Author string `json:"author" gorm:"not null"`
}