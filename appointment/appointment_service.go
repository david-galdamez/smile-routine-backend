package appointment

import (
	"context"
	"database/sql"
	"time"

	"github.com/david-galdamez/smile-routine-backend/database"
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

func toAppointmentDto(appointment database.Appointment) AppointmentDto {
	dto := AppointmentDto{}
	dto.Id = int(appointment.ID)
	dto.Title = appointment.Title
	if appointment.Note.Valid {
		dto.Note = &appointment.Note.String
	} else {
		dto.Note = nil
	}
	dto.Date = appointment.AppointmentDate.Format("2006-01-02")
	dto.Time = appointment.AppointmentTime.Format("15:04")
	dto.Completed = appointment.Completed
	return dto
}

func (as *AppointmentService) CreateAppointment(ctx context.Context, userId int, registerAppointment *RegisterAppointmentDto) utils.ServiceResponse[AppointmentDto] {

	value, err := time.Parse("2006-01-02", registerAppointment.AppointmentDate)
	if err != nil {
		return utils.Error[AppointmentDto]("Formato de fecha incorrecto")
	}

	timeValue, err := time.Parse("15:04", registerAppointment.AppointmentTime)
	if err != nil {
		return utils.Error[AppointmentDto]("Formato de tiempo incorrecto")
	}

	var note sql.NullString

	if registerAppointment.Note != nil {
		note.String = *registerAppointment.Note
		note.Valid = true
	} else {
		note.Valid = false
	}

	appointment, err := as.repository.CreateAppointment(ctx, userId, value, timeValue, registerAppointment.Title, note)
	if err != nil {
		return utils.Error[AppointmentDto]("Error al crear la cita")
	}

	return utils.Ok(toAppointmentDto(*appointment))
}

func (as *AppointmentService) GetAppointmentById(ctx context.Context, id int) utils.ServiceResponse[AppointmentDto] {
	appointment, err := as.repository.GetAppointmentById(ctx, id)
	if err != nil {
		return utils.Error[AppointmentDto]("Cita no encontrada")
	}

	return utils.Ok(toAppointmentDto(*appointment))
}

func (as *AppointmentService) GetAppointments(ctx context.Context, userId int) utils.ServiceResponse[[]AppointmentDto] {
	appointments, err := as.repository.GetAppointments(ctx, userId)
	if err != nil {
		return utils.Error[[]AppointmentDto]("Error al obtener las citas")
	}

	appointmentDtos := make([]AppointmentDto, 0, len(appointments))
	for _, appointment := range appointments {
		appointmentDtos = append(appointmentDtos, toAppointmentDto(appointment))
	}

	return utils.Ok(appointmentDtos)
}

func (as *AppointmentService) UpdateAppointment(ctx context.Context, completed bool, id int) utils.ServiceResponse[AppointmentDto] {
	appointment, err := as.repository.UpdateAppointment(ctx, completed, id)
	if err != nil {
		return utils.Error[AppointmentDto]("Error al actualizar la cita")
	}

	return utils.Ok(toAppointmentDto(*appointment))
}
