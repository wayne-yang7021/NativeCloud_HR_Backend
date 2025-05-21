package model

type CheckInRequest struct {
	EmployeeID   string `json:"employee_id" binding:"required"`
	AccessTime   string `json:"access_time" binding:"required"`
	Direction    string `json:"direction" binding:"required"`
	GateType     string `json:"gate_type" binding:"required"`
	GateName     string `json:"gate_name" binding:"required"`
	AccessResult string `json:"access_result" binding:"required"`
}

func (CheckInRequest) TableName() string {
	return "access_log"
}
