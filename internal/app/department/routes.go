package department

type DepartmentRoutes struct {
	Base         string
	ByID         string
	ByDepartment string
}

var Departments = DepartmentRoutes{
	Base: "/departments",
	ByID: "/departments/{id}",
}
