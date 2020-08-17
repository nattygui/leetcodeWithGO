package _690EmployeeImportance

type Employee struct {
	Id           int
	Importance   int
	Subordinates []int
}

func getImportance(employees []*Employee, id int) int {
	res := 0
	getImportanceByemployee(id, employees, &res)
	return res
}

func getImportanceByemployee(Id int, employees []*Employee, res *int) {
	for _, employee := range employees {
		if employee.Id == Id {
			*res += employee.Importance
			if len(employee.Subordinates) != 0 {
				for _, id := range employee.Subordinates {
					getImportanceByemployee(id, employees, res)
				}
			}
		}
	}
}
