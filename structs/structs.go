package structs

type Abiturient struct {
	ID          string `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	BirthDate   string `json:"birth_date"`
	BirthPlace  string `json:"birth_place"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phone_number"`
	MiddleName  string `json:"middle_name"`
}

type Login struct {
	Login        string `json:"login"`
	Password     string `json:"password"`
	Role         string `json:"role"`
	AbiturientID int64  `json:"abiturient_id"`
}

type UserCheck struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

type UserCheckResult struct {
	Role         string `json:"role"`
	AbiturientID int64  `json:"abiturient_id"`
}

type Subject struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Specialty struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type University struct {
	Code      string `json:"code"`
	Name      string `json:"name"`
	Active    bool   `json:"active"`
	URL       string `json:"url"`
	ShortName string `json:"short_name"`
	City      string `json:"city"`
	Rate      int    `json:"rate"`
}

type AbiturientScore struct {
	AbiturientID int64  `json:"abiturient_id"`
	SubjectID    string `json:"subject_id"`
	SubjectName  string `json:"subject_name"`
	Score        int    `json:"score"`
}

type Faculty struct {
	ID             int    `json:"id"`
	UniversityCode string `json:"university_code"`
	UniversityName string `json:"university_name"`
	Name           string `json:"name"`
}
