// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: reservation/v1/reservation.proto

package reservationv1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/ekkx/tcmrsv-web/server/internal/shared/pb/reservation/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// ReservationServiceName is the fully-qualified name of the ReservationService service.
	ReservationServiceName = "reservation.v1.ReservationService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// ReservationServiceGetReservationProcedure is the fully-qualified name of the ReservationService's
	// GetReservation RPC.
	ReservationServiceGetReservationProcedure = "/reservation.v1.ReservationService/GetReservation"
	// ReservationServiceListAvailableRoomsProcedure is the fully-qualified name of the
	// ReservationService's ListAvailableRooms RPC.
	ReservationServiceListAvailableRoomsProcedure = "/reservation.v1.ReservationService/ListAvailableRooms"
	// ReservationServiceListReservationsProcedure is the fully-qualified name of the
	// ReservationService's ListReservations RPC.
	ReservationServiceListReservationsProcedure = "/reservation.v1.ReservationService/ListReservations"
	// ReservationServiceCreateReservationProcedure is the fully-qualified name of the
	// ReservationService's CreateReservation RPC.
	ReservationServiceCreateReservationProcedure = "/reservation.v1.ReservationService/CreateReservation"
	// ReservationServiceDeleteReservationProcedure is the fully-qualified name of the
	// ReservationService's DeleteReservation RPC.
	ReservationServiceDeleteReservationProcedure = "/reservation.v1.ReservationService/DeleteReservation"
)

// ReservationServiceClient is a client for the reservation.v1.ReservationService service.
type ReservationServiceClient interface {
	GetReservation(context.Context, *connect.Request[v1.GetReservationRequest]) (*connect.Response[v1.GetReservationResponse], error)
	ListAvailableRooms(context.Context, *connect.Request[v1.ListAvailableRoomsRequest]) (*connect.Response[v1.ListAvailableRoomsResponse], error)
	ListReservations(context.Context, *connect.Request[v1.ListReservationsRequest]) (*connect.Response[v1.ListReservationsResponse], error)
	CreateReservation(context.Context, *connect.Request[v1.CreateReservationRequest]) (*connect.Response[v1.CreateReservationResponse], error)
	DeleteReservation(context.Context, *connect.Request[v1.DeleteReservationRequest]) (*connect.Response[v1.DeleteReservationResponse], error)
}

// NewReservationServiceClient constructs a client for the reservation.v1.ReservationService
// service. By default, it uses the Connect protocol with the binary Protobuf Codec, asks for
// gzipped responses, and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply
// the connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewReservationServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) ReservationServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	reservationServiceMethods := v1.File_reservation_v1_reservation_proto.Services().ByName("ReservationService").Methods()
	return &reservationServiceClient{
		getReservation: connect.NewClient[v1.GetReservationRequest, v1.GetReservationResponse](
			httpClient,
			baseURL+ReservationServiceGetReservationProcedure,
			connect.WithSchema(reservationServiceMethods.ByName("GetReservation")),
			connect.WithClientOptions(opts...),
		),
		listAvailableRooms: connect.NewClient[v1.ListAvailableRoomsRequest, v1.ListAvailableRoomsResponse](
			httpClient,
			baseURL+ReservationServiceListAvailableRoomsProcedure,
			connect.WithSchema(reservationServiceMethods.ByName("ListAvailableRooms")),
			connect.WithClientOptions(opts...),
		),
		listReservations: connect.NewClient[v1.ListReservationsRequest, v1.ListReservationsResponse](
			httpClient,
			baseURL+ReservationServiceListReservationsProcedure,
			connect.WithSchema(reservationServiceMethods.ByName("ListReservations")),
			connect.WithClientOptions(opts...),
		),
		createReservation: connect.NewClient[v1.CreateReservationRequest, v1.CreateReservationResponse](
			httpClient,
			baseURL+ReservationServiceCreateReservationProcedure,
			connect.WithSchema(reservationServiceMethods.ByName("CreateReservation")),
			connect.WithClientOptions(opts...),
		),
		deleteReservation: connect.NewClient[v1.DeleteReservationRequest, v1.DeleteReservationResponse](
			httpClient,
			baseURL+ReservationServiceDeleteReservationProcedure,
			connect.WithSchema(reservationServiceMethods.ByName("DeleteReservation")),
			connect.WithClientOptions(opts...),
		),
	}
}

// reservationServiceClient implements ReservationServiceClient.
type reservationServiceClient struct {
	getReservation     *connect.Client[v1.GetReservationRequest, v1.GetReservationResponse]
	listAvailableRooms *connect.Client[v1.ListAvailableRoomsRequest, v1.ListAvailableRoomsResponse]
	listReservations   *connect.Client[v1.ListReservationsRequest, v1.ListReservationsResponse]
	createReservation  *connect.Client[v1.CreateReservationRequest, v1.CreateReservationResponse]
	deleteReservation  *connect.Client[v1.DeleteReservationRequest, v1.DeleteReservationResponse]
}

// GetReservation calls reservation.v1.ReservationService.GetReservation.
func (c *reservationServiceClient) GetReservation(ctx context.Context, req *connect.Request[v1.GetReservationRequest]) (*connect.Response[v1.GetReservationResponse], error) {
	return c.getReservation.CallUnary(ctx, req)
}

// ListAvailableRooms calls reservation.v1.ReservationService.ListAvailableRooms.
func (c *reservationServiceClient) ListAvailableRooms(ctx context.Context, req *connect.Request[v1.ListAvailableRoomsRequest]) (*connect.Response[v1.ListAvailableRoomsResponse], error) {
	return c.listAvailableRooms.CallUnary(ctx, req)
}

// ListReservations calls reservation.v1.ReservationService.ListReservations.
func (c *reservationServiceClient) ListReservations(ctx context.Context, req *connect.Request[v1.ListReservationsRequest]) (*connect.Response[v1.ListReservationsResponse], error) {
	return c.listReservations.CallUnary(ctx, req)
}

// CreateReservation calls reservation.v1.ReservationService.CreateReservation.
func (c *reservationServiceClient) CreateReservation(ctx context.Context, req *connect.Request[v1.CreateReservationRequest]) (*connect.Response[v1.CreateReservationResponse], error) {
	return c.createReservation.CallUnary(ctx, req)
}

// DeleteReservation calls reservation.v1.ReservationService.DeleteReservation.
func (c *reservationServiceClient) DeleteReservation(ctx context.Context, req *connect.Request[v1.DeleteReservationRequest]) (*connect.Response[v1.DeleteReservationResponse], error) {
	return c.deleteReservation.CallUnary(ctx, req)
}

// ReservationServiceHandler is an implementation of the reservation.v1.ReservationService service.
type ReservationServiceHandler interface {
	GetReservation(context.Context, *connect.Request[v1.GetReservationRequest]) (*connect.Response[v1.GetReservationResponse], error)
	ListAvailableRooms(context.Context, *connect.Request[v1.ListAvailableRoomsRequest]) (*connect.Response[v1.ListAvailableRoomsResponse], error)
	ListReservations(context.Context, *connect.Request[v1.ListReservationsRequest]) (*connect.Response[v1.ListReservationsResponse], error)
	CreateReservation(context.Context, *connect.Request[v1.CreateReservationRequest]) (*connect.Response[v1.CreateReservationResponse], error)
	DeleteReservation(context.Context, *connect.Request[v1.DeleteReservationRequest]) (*connect.Response[v1.DeleteReservationResponse], error)
}

// NewReservationServiceHandler builds an HTTP handler from the service implementation. It returns
// the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewReservationServiceHandler(svc ReservationServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	reservationServiceMethods := v1.File_reservation_v1_reservation_proto.Services().ByName("ReservationService").Methods()
	reservationServiceGetReservationHandler := connect.NewUnaryHandler(
		ReservationServiceGetReservationProcedure,
		svc.GetReservation,
		connect.WithSchema(reservationServiceMethods.ByName("GetReservation")),
		connect.WithHandlerOptions(opts...),
	)
	reservationServiceListAvailableRoomsHandler := connect.NewUnaryHandler(
		ReservationServiceListAvailableRoomsProcedure,
		svc.ListAvailableRooms,
		connect.WithSchema(reservationServiceMethods.ByName("ListAvailableRooms")),
		connect.WithHandlerOptions(opts...),
	)
	reservationServiceListReservationsHandler := connect.NewUnaryHandler(
		ReservationServiceListReservationsProcedure,
		svc.ListReservations,
		connect.WithSchema(reservationServiceMethods.ByName("ListReservations")),
		connect.WithHandlerOptions(opts...),
	)
	reservationServiceCreateReservationHandler := connect.NewUnaryHandler(
		ReservationServiceCreateReservationProcedure,
		svc.CreateReservation,
		connect.WithSchema(reservationServiceMethods.ByName("CreateReservation")),
		connect.WithHandlerOptions(opts...),
	)
	reservationServiceDeleteReservationHandler := connect.NewUnaryHandler(
		ReservationServiceDeleteReservationProcedure,
		svc.DeleteReservation,
		connect.WithSchema(reservationServiceMethods.ByName("DeleteReservation")),
		connect.WithHandlerOptions(opts...),
	)
	return "/reservation.v1.ReservationService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case ReservationServiceGetReservationProcedure:
			reservationServiceGetReservationHandler.ServeHTTP(w, r)
		case ReservationServiceListAvailableRoomsProcedure:
			reservationServiceListAvailableRoomsHandler.ServeHTTP(w, r)
		case ReservationServiceListReservationsProcedure:
			reservationServiceListReservationsHandler.ServeHTTP(w, r)
		case ReservationServiceCreateReservationProcedure:
			reservationServiceCreateReservationHandler.ServeHTTP(w, r)
		case ReservationServiceDeleteReservationProcedure:
			reservationServiceDeleteReservationHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedReservationServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedReservationServiceHandler struct{}

func (UnimplementedReservationServiceHandler) GetReservation(context.Context, *connect.Request[v1.GetReservationRequest]) (*connect.Response[v1.GetReservationResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("reservation.v1.ReservationService.GetReservation is not implemented"))
}

func (UnimplementedReservationServiceHandler) ListAvailableRooms(context.Context, *connect.Request[v1.ListAvailableRoomsRequest]) (*connect.Response[v1.ListAvailableRoomsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("reservation.v1.ReservationService.ListAvailableRooms is not implemented"))
}

func (UnimplementedReservationServiceHandler) ListReservations(context.Context, *connect.Request[v1.ListReservationsRequest]) (*connect.Response[v1.ListReservationsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("reservation.v1.ReservationService.ListReservations is not implemented"))
}

func (UnimplementedReservationServiceHandler) CreateReservation(context.Context, *connect.Request[v1.CreateReservationRequest]) (*connect.Response[v1.CreateReservationResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("reservation.v1.ReservationService.CreateReservation is not implemented"))
}

func (UnimplementedReservationServiceHandler) DeleteReservation(context.Context, *connect.Request[v1.DeleteReservationRequest]) (*connect.Response[v1.DeleteReservationResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("reservation.v1.ReservationService.DeleteReservation is not implemented"))
}
