package handler

import (
	"context"

	"connectrpc.com/connect"
	reservationv1 "github.com/ekkx/tcmrsv-web/server/internal/shared/pb/reservation/v1"
)

func (h *HandlerImpl) DeleteReservation(ctx context.Context, req *connect.Request[reservationv1.DeleteReservationRequest]) (*connect.Response[reservationv1.DeleteReservationResponse], error) {
    return nil, nil
}

