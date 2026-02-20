package appointment

type RegisterAppointmentDto struct {
	AppointmentDate string `json:"appointment_date" validate:"required,datetime=2006-01-02"`
}

type AppointmentDto struct {
	Id        int    `json:"id"`
	Date      string `json:"appointment_date" validate:"datetime=2006-01-02"`
	Completed bool   `json:"completed"`
}
