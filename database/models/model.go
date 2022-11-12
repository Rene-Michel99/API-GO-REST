package models

import (
    "gorm.io/gorm"
    "time"
)


type Book struct{
    gorm.Model
    ID int                  `json:"id" gorm:primaryKey;autoIncrement`
    Name string             `json:"name" gorm:text;not null;default:null`
    Author string           `json:"author" gorm:text;not null;default:null`
    Synopsis string         `json:"synopsis" gorm:text;not null;default:null`
    LaunchDate time.Time    `json:"launchDate" gorm:date;not null;default:null`
    CopyQnt int32           `json:"copyQnt" gorm:int;not null;default:null`
}