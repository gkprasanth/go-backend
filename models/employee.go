package models

import "gorm.io/gorm"

// Employee represents the employee model
type Employee struct {
	gorm.Model

	EmpID      string `gorm:"type:varchar(100);unique" json:"emp_id"`
	EmpName    string `json:"emp_name"`
	Department string `json:"department"`
	Manager    string `json:"manager"`
	ManagerID  string `json:"manager_id"`
	HOD        string `json:"hod"`
	EmpPhone   string `json:"emp_phone"`
	Location   string `json:"location"`
}
