package output

import (
	"github.com/ekkx/tcmrsv-web/server/internal/domain/entity"
	"github.com/ekkx/tcmrsv-web/server/internal/domain/enum"
	rsv_v1 "github.com/ekkx/tcmrsv-web/server/internal/shared/api/v1/reservation"
	room_v1 "github.com/ekkx/tcmrsv-web/server/internal/shared/api/v1/room"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type UpdateReservation struct {
	Reservation entity.Reservation
}

func NewUpdateReservation(reservation entity.Reservation) *UpdateReservation {
	return &UpdateReservation{
		Reservation: reservation,
	}
}

func (output *UpdateReservation) ToProto() *rsv_v1.UpdateReservationReply {
	var campusType room_v1.CampusType
	switch output.Reservation.CampusType {
	case enum.CampusTypeNakameguro:
		campusType = room_v1.CampusType_NAKAMEGURO
	case enum.CampusTypeIkebukuro:
		campusType = room_v1.CampusType_IKEBUKURO
	default:
		campusType = room_v1.CampusType_CAMPUS_UNSPECIFIED
	}

	return &rsv_v1.UpdateReservationReply{
		Reservation: &rsv_v1.Reservation{
			Id:         int64(output.Reservation.ID),
			ExternalId: output.Reservation.ExternalID,
			CampusType: campusType,
			Date:       timestamppb.New(output.Reservation.Date),
			RoomId:     output.Reservation.RoomID,
			FromHour:   int32(output.Reservation.FromHour),
			FromMinute: int32(output.Reservation.FromMinute),
			ToHour:     int32(output.Reservation.ToHour),
			ToMinute:   int32(output.Reservation.ToMinute),
			BookerName: output.Reservation.BookerName,
			CreatedAt:  timestamppb.New(output.Reservation.CreatedAt),
		},
	}
}
