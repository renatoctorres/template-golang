package employee

type Employee struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Position     string `json:"position"`
	DepartmentId string `json:"department:id"`
}
