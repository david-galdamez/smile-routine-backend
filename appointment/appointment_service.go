package appointment

import (
	"context"
	"time"

	"github.com/david-galdamez/smile-routine-backend/utils"
)

type AppointmentService struct {
	repository *AppointmentRepository
}

func NewAppointmentService(repository *AppointmentRepository) *AppointmentService {
	return &AppointmentService{
		repository: repository,
	}
}

func (as *AppointmentService) CreateAppointment(ctx context.Context, userId int, registerAppointment *RegisterAppointmentDto) utils.ServiceResponse[AppointmentDto] {

	value, err := time.Parse("2006-01-02", registerAppointment.AppointmentDate)
	if err != nil {
		return utils.Error[AppointmentDto]("Formato de fecha incorrecto")
	}

	appointment, err := as.repository.CreateAppointment(ctx, userId, value, false)
	if err != nil {
		return utils.Error[AppointmentDto]("Error al crear la cita")
	}

	appointmentDto := AppointmentDto{}
	appointmentDto.Id = int(appointment.ID)
	appointmentDto.Date = appointment.AppointmentDate.Format("2006-01-02")
	appointmentDto.Completed = appointment.Completed

	return utils.Ok[AppointmentDto](appointmentDto)
}

func (as *AppointmentService) UpdateAppointment(ctx context.Context, completed bool, id int) utils.ServiceResponse[AppointmentDto] {
	appointment, err := as.repository.UpdateAppointment(ctx, completed, id)
	if err != nil {
		return utils.Error[AppointmentDto]("Error al actualizar la cita")
	}
	appointmentDto := AppointmentDto{}
	appointmentDto.Id = int(appointment.ID)
	appointmentDto.Date = appointment.AppointmentDate.Format("2006-01-02")
	appointmentDto.Completed = appointment.Completed

	return utils.Ok[AppointmentDto](appointmentDto)
}
