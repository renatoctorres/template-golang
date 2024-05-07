package employee

type EmployeeRoutes struct {
	Base         string
	ByID         string
	ByDepartment string
}

var Employees = EmployeeRoutes{
	Base:         "/employees",
	ByID:         "/employees/{id}",
	ByDepartment: "/employees/department/{deptId}",
}
