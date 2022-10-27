package models

import (
    "gorm.io/gorm"
    "time"
)


type Book struct{
    gorm.Model
    id int                  `json:"id" gorm:primaryKey;autoIncrement`
    name string             `json:"name"`
    author string           `json:"author"`
    synopsis string         `json:"synopsis"`
    launchDate time.Time    `json:"launchDate"`
    copyQnt int             `json:"copyQnt"`
}