package beans

type StudentBean struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	YearOfStudy int    `json:"year"`
	Department  string `json:"dept"`
	BloodGroup  string `json:"bloodgroup"`
	PhoneNo     int    `json:"phno"`
	Email       string `json:"email"`
	Location    string `json:"loc"`
}

type Student struct {
	ID          int `gorm:"primary_key"`
	Name        string
	YearOfStudy int
	Department  string
	BloodGroup  string
	PhoneNo     int
	Email       string
	Location    string
}
