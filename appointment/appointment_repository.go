package appointment

import (
	"context"
	"database/sql"
	"time"

	"github.com/david-galdamez/smile-routine-backend/database"
)

type AppointmentRepository struct {
	db *database.Queries
}

func NewAppointmentRepository(db *database.Queries) *AppointmentRepository {
	return &AppointmentRepository{db: db}
}

func (ar *AppointmentRepository) CreateAppointment(ctx context.Context, userId int, appointmentDate time.Time, appointmentTime time.Time, title string, note sql.NullString) (*database.Appointment, error) {
	appointment, err := ar.db.CreateAppointment(ctx, database.CreateAppointmentParams{
		UserID:          int32(userId),
		AppointmentDate: appointmentDate,
		AppointmentTime: appointmentTime,
		Title:           title,
		Note:            note,
	})
	if err != nil {
		return nil, err
	}
	return &appointment, nil
}

func (ar *AppointmentRepository) GetAppointmentById(ctx context.Context, id int) (*database.Appointment, error) {
	appointment, err := ar.db.GetAppointmentById(ctx, int64(id))
	if err != nil {
		return nil, err
	}

	return &appointment, nil
}

func (ar *AppointmentRepository) GetAppointments(ctx context.Context, userId int) ([]database.Appointment, error) {
	appointments, err := ar.db.GetAppointments(ctx, int32(userId))
	if err != nil {
		return nil, err
	}

	return appointments, nil
}

func (ar *AppointmentRepository) UpdateAppointment(ctx context.Context, completed bool, id int) (*database.Appointment, error) {
	appointment, err := ar.db.UpdateAppointment(ctx, database.UpdateAppointmentParams{
		Completed: completed,
		ID:        int64(id),
	})
	if err != nil {
		return nil, err
	}
	return &appointment, nil
}
