package services

// Person represents a person in the system
type Person struct {
	Id           string  `json:"id"`
	FirstName    string  `json:"firstName"`
	LastName     string  `json:"lastName"`
	Prefix       *string `json:"prefix"`
	Room         string  `json:"room"`
	Phone        string  `json:"phone"`
	Floor        int     `json:"floor"`
	DepartmentId string  `json:"departmentId"`
	Title        string  `json:"title"`
}

// Department represents a department in the system
type Department struct {
	Id                 string  `json:"id"`
	Name               string  `json:"name"`
	Phone              string  `json:"phone"`
	ManagerId          string  `json:"managerId"`
	ParentDepartmentId *string `json:"parentDepartmentId"`
}
