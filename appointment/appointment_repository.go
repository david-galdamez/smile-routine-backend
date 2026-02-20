package appointment

import (
	"context"
	"time"

	"github.com/david-galdamez/smile-routine-backend/database"
)

type AppointmentRepository struct {
	db *database.Queries
}

func NewAppointmentRepository(db *database.Queries) *AppointmentRepository {
	return &AppointmentRepository{db: db}
}

func (ar *AppointmentRepository) CreateAppointment(ctx context.Context, userId int, appointmentDate time.Time, completed bool) (*database.Appointment, error) {
	appointment, err := ar.db.CreateAppointment(ctx, database.CreateAppointmentParams{
		UserID:          int32(userId),
		AppointmentDate: appointmentDate,
		Completed:       completed,
	})
	if err != nil {
		return nil, err
	}
	return &appointment, nil
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
