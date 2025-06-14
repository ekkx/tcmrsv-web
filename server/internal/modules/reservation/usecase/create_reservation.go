package usecase

import (
	"context"
	"time"

	"github.com/ekkx/tcmrsv-web/server/internal/domain/entity"
	"github.com/ekkx/tcmrsv-web/server/internal/modules/reservation/dto/input"
	"github.com/ekkx/tcmrsv-web/server/internal/modules/reservation/dto/output"
	rsv_repo "github.com/ekkx/tcmrsv-web/server/internal/modules/reservation/repository"
	room_repo "github.com/ekkx/tcmrsv-web/server/internal/modules/room/repository"
	"github.com/ekkx/tcmrsv-web/server/internal/shared/errs"
	"github.com/ekkx/tcmrsv-web/server/pkg/utils"
)

func (u *UsecaseImpl) CreateReservation(ctx context.Context, params *input.CreateReservation) (*output.CreateReservation, error) {
	if err := params.Validate(); err != nil {
		return nil, errs.ErrInvalidArgument.WithCause(err)
	}

	// Convert the received date to JST first, then extract date components
	dateInJST := params.Date.In(utils.JST())
	var date = time.Date(dateInJST.Year(), dateInJST.Month(), dateInJST.Day(), 0, 0, 0, 0, utils.JST())

	// 練習室の存在チェック
	rooms := u.roomRepo.SearchRooms(ctx, &room_repo.SearchRoomsArgs{
		ID: &params.RoomID,
	})

	if len(rooms) == 0 {
		return nil, errs.ErrRoomNotFound
	}

	// 予約時間と練習室が被っていないか確認
	hasConflict, err := u.rsvRepo.CheckReservationConflict(ctx, &rsv_repo.CheckReservationConflictArgs{
		RoomID:     params.RoomID,
		Date:       date,
		FromHour:   params.FromHour,
		FromMinute: params.FromMinute,
		ToHour:     params.ToHour,
		ToMinute:   params.ToMinute,
	})
	if err != nil {
		return nil, errs.ErrInternal.WithCause(err)
	}

	if hasConflict {
		return nil, errs.ErrReservationConflict
	}

	rsv, err := u.rsvRepo.CreateReservation(ctx, &rsv_repo.CreateReservationArgs{
		UserID:     params.UserID,
		CampusType: params.CampusType,
		RoomID:     params.RoomID,
		Date:       date,
		FromHour:   params.FromHour,
		FromMinute: params.FromMinute,
		ToHour:     params.ToHour,
		ToMinute:   params.ToMinute,
		BookerName: params.BookerName,
	})
	if err != nil {
		return nil, errs.ErrInternal.WithCause(err)
	}

	return output.NewCreateReservation([]entity.Reservation{*rsv}), nil
}
