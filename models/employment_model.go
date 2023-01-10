package models

type Data struct {
	Employment []Employment
}

type Employment struct {
	Id          int64  `json:"id,omitempty"`
	ProfileCode int64  `json:"profileCode,omitempty"`
	JobTitle    string `json:"jobTitle,omitempty"  validate:"required"`
	Employer    string `json:"employer,omitempty" validate:"required"`
	StartDate   string `json:"startDate,omitempty" validate:"required"`
	EndDate     string `json:"endDate,omitempty" validate:"required"`
	City        string `json:"city,omitempty" validate:"required"`
	Description string `json:"description,omitempty" validate:"required"`
}

// {
//     "jobTitle": "CEO",
//     "employer": "Toko Lapak",
//     "startDate": "01-01-2020",
//     "endDate": "01-01-2021",
//     "city": "Jakarta",
//     "description": "CEO"
// }
