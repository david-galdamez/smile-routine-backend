package appointment

type RegisterAppointmentDto struct {
	AppointmentDate string  `json:"appointment_date" validate:"required,datetime=2006-01-02"`
	AppointmentTime string  `json:"appointment_time" validate:"required"`
	Title           string  `json:"title" validate:"required"`
	Note            *string `json:"note,omitempty"`
}

type AppointmentListDto struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Date  string `json:"appointment_date" validate:"datetime=2006-01-02"`
	Time  string `json:"appointment_time"`
}

type AppointmentDto struct {
	Id        int     `json:"id"`
	Title     string  `json:"title"`
	Note      *string `json:"note,omitempty"`
	Date      string  `json:"appointment_date" validate:"datetime=2006-01-02"`
	Time      string  `json:"appointment_time"`
	Completed bool    `json:"completed"`
}
