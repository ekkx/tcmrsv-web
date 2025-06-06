package usecase

import (
	"context"

	"github.com/ekkx/tcmrsv-web/server/internal/modules/reservation/dto/input"
	"github.com/ekkx/tcmrsv-web/server/internal/shared/actor"
	"github.com/ekkx/tcmrsv-web/server/internal/shared/errs"
)

func (uc *UsecaseImpl) DeleteReservation(ctx context.Context, params *input.DeleteReservation) error {
	rsv, err := uc.rsvRepo.GetReservationByID(ctx, params.ReservationID)
	if err != nil {
		return errs.ErrInternal.WithCause(err)
	}

	if rsv == nil {
		return errs.ErrReservationNotFound
	}

	// 予約者からのリクエストであるかを確認
	if params.Actor.Role == actor.RoleUser {
		if params.Actor.ID != rsv.UserID {
			return errs.ErrNotYourReservation
		}
	}

	// 予約を削除
	err = uc.rsvRepo.DeleteReservationByID(ctx, params.ReservationID)
	if err != nil {
		return errs.ErrInternal.WithCause(err)
	}

	return nil
}
