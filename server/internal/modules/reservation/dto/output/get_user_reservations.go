package output

import (
	"time"

	"github.com/ekkx/tcmrsv-web/server/internal/domain/entity"
	"github.com/ekkx/tcmrsv-web/server/internal/domain/enum"
	rsv_v1 "github.com/ekkx/tcmrsv-web/server/internal/shared/api/v1/reservation"
	room_v1 "github.com/ekkx/tcmrsv-web/server/internal/shared/api/v1/room"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type GetMyReservations struct {
	Reservations []entity.Reservation
}

func NewGetMyReservations(reservations []entity.Reservation) *GetMyReservations {
	return &GetMyReservations{
		Reservations: reservations,
	}
}

func (output *GetMyReservations) ToProto() *rsv_v1.GetUserReservationsReply {
	protoRsvs := make([]*rsv_v1.Reservation, 0, len(output.Reservations))
	if len(output.Reservations) == 0 {
		return &rsv_v1.GetUserReservationsReply{
			Reservations: protoRsvs,
		}
	}

	for _, rsv := range output.Reservations {
		var campusType room_v1.CampusType
		switch rsv.CampusType {
		case enum.CampusTypeNakameguro:
			campusType = room_v1.CampusType_NAKAMEGURO
		case enum.CampusTypeIkebukuro:
			campusType = room_v1.CampusType_IKEBUKURO
		default:
			campusType = room_v1.CampusType_CAMPUS_UNSPECIFIED
		}

		// For date-only fields, we want to preserve the date as-is
		// Create a UTC time with the same date components
		dateUTC := time.Date(rsv.Date.Year(), rsv.Date.Month(), rsv.Date.Day(), 0, 0, 0, 0, time.UTC)

		protoRsvs = append(protoRsvs, &rsv_v1.Reservation{
			Id:         int64(rsv.ID),
			ExternalId: rsv.ExternalID,
			CampusType: campusType,
			Date:       timestamppb.New(dateUTC),
			RoomId:     rsv.RoomID,
			FromHour:   int32(rsv.FromHour),
			FromMinute: int32(rsv.FromMinute),
			ToHour:     int32(rsv.ToHour),
			ToMinute:   int32(rsv.ToMinute),
			BookerName: rsv.BookerName,
			CreatedAt:  timestamppb.New(rsv.CreatedAt.In(time.UTC)),
		})
	}

	return &rsv_v1.GetUserReservationsReply{
		Reservations: protoRsvs,
	}
}
