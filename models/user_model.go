package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id             primitive.ObjectID `json:"id,omitempty"`
	ProfileCode    int64              `json:"profileCode,omitempty"`
	WantedJobTitle string             `json:"wantedJobTitle,omitempty"  validate:"required"`
	FirstName      string             `json:"firstName,omitempty" validate:"required"`
	LastName       string             `json:"lastName,omitempty" validate:"required"`
	Email          string             `json:"email,omitempty" validate:"email"`
	Phone          string             `json:"phone,omitempty" validate:"required"`
	Country        string             `json:"country,omitempty" validate:"required"`
	City           string             `json:"city,omitempty" validate:"required"`
	Address        string             `json:"address,omitempty" validate:"required"`
	PostalCode     int32              `json:"postalCode,omitempty" validate:"required"`
	DrivingLicense string             `json:"drivingLicense,omitempty" validate:"required"`
	Nationality    string             `json:"nationality,omitempty" validate:"required"`
	PlaceOfBirth   string             `json:"placeOfBirth,omitempty" validate:"required"`
	DateOfBirth    string             `json:"dateOfBirth,omitempty" validate:"required"`
}

// {
//     "wantedJobTitle": "Software Engineer",
//     "firstName": "Namaku",
//     "lastName": "Ukaman",
//     "email": "ukaman.namaku@gmail.com",
//     "phone": "08008880000",
//     "country": "Indonesia",
//     "city": "Jakarta",
//     "address": "Jl. Gatot Subroto",
//     "postalCode": 200001,
//     "drivingLicense": "1234567890123456",
//     "nationality": "Indonesia",
//     "placeOfBirth": "Maluku",
//     "dateOfBirth": "07-12-1988"
// }
