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

// type NewStudentBean struct {
// 	Name        string `json:"name"`
// 	YearOfStudy int    `json:"year"`
// 	Department  string `json:"dept"`
// 	BloodGroup  string `json:"bloodgroup"`
// 	PhoneNo     int    `json:"phno"`
// 	Email       string `json:"email"`
// 	Location    string `json:"loc"`
// }
