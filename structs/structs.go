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
