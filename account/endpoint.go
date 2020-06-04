package account

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	CreateUser endpoint.Endpoint
	GetUser endpoint.Endpoint
	Login endpoint.Endpoint

	CreateCategory endpoint.Endpoint
	GetCategory endpoint.Endpoint
	GetCategories endpoint.Endpoint
	UpdateCategory endpoint.Endpoint

	CreateConfOwnProfile endpoint.Endpoint
	GetConfOwnProfile endpoint.Endpoint
	UpdateConfOwnProfile endpoint.Endpoint
	DeleteConfOwnProfile endpoint.Endpoint

	CreateConference endpoint.Endpoint
	GetConference endpoint.Endpoint
	UpdateCreateConference endpoint.Endpoint
	DeleteCreateConference endpoint.Endpoint

	CreateLocation endpoint.Endpoint
	GetLocation endpoint.Endpoint
	UpdateCreateLocation endpoint.Endpoint
	DeleteCreateLocation endpoint.Endpoint

	CreateRating endpoint.Endpoint
	GetRating endpoint.Endpoint
	UpdateCreateRating endpoint.Endpoint
	DeleteCreateRating endpoint.Endpoint

	CreateReport endpoint.Endpoint
	GetReport endpoint.Endpoint
	UpdateCreateReport endpoint.Endpoint
	DeleteCreateReport endpoint.Endpoint

	CreateSpeaker endpoint.Endpoint
	GetSpeaker endpoint.Endpoint
	UpdateCreateSpeaker endpoint.Endpoint
	DeleteCreateSpeaker endpoint.Endpoint
}
func MakeEndpoints(s Service) Endpoints{
	return Endpoints{
		CreateUser: makeCreateUserEndpoints(s),
		GetUser:    makeGetUserEndpoint(s),
		Login:      makeLoginEndpoint(s),
		CreateConfOwnProfile: makeCreateConfOwnProfileEndpoint(s),
		GetConfOwnProfile: makeGetConfOwnProfileEndpoint(s),
		UpdateConfOwnProfile:makeUpdateConfOwnProfileEndpoint(s),
		DeleteConfOwnProfile:makeDeleteConfOwnProfileEndpoint(s),

		CreateConference:makeCreateConferenceEndpoint(s),
		GetConference:makeGetConferenceEndpoint(s),
		UpdateCreateConference:makeUpdateCreateConferenceEndpoint(s),
		DeleteCreateConference:makeDeleteCreateConferenceEndpoint(s),

		CreateLocation:makeCreateLocationEndpoint(s),
		GetLocation:makeGetLocationEndpoint(s),
		UpdateCreateLocation:makeUpdateCreateLocationEndpoint(s),
		DeleteCreateLocation:makeDeleteCreateLocationEndpoint(s),

		CreateRating:makeCreateRatingEndpoint(s),
		GetRating:makeGetRatingEndpoint(s),
		UpdateCreateRating:makeUpdateCreateRatingEndpoint(s),
		DeleteCreateRating:makeDeleteCreateRatingEndpoint(s),

		CreateReport:makeCreateReportEndpoint(s),
		GetReport:makeGetReportEndpoint(s),
		UpdateCreateReport:makeUpdateCreateReportEndpoint(s),
		DeleteCreateReport:makeDeleteCreateReportEndpoint(s),

		CreateSpeaker:makeCreateSpeakerEndpoint(s),
		GetSpeaker:makeGetSpeakerEndpoint(s),
		UpdateCreateSpeaker:makeUpdateCreateSpeakerEndpoint(s),
		DeleteCreateSpeaker:makeDeleteCreateSpeakerEndpoint(s),

		CreateCategory: makeCreateCategoryEndpoints(s),
		GetCategory:    makeGetCategoryEndpoints(s),
		GetCategories: makeGetCategoriesEndpoints(s),
		UpdateCategory: makeUpdateCategoryEndpoints(s),
	}
}

func makeCreateUserEndpoints(s Service) endpoint.Endpoint{
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(CreateUserRequest)
		ok, err := s.CreateUser(ctx, req.Email, req.Password)
		return CreateUserResponse{Ok:ok}, err
	}
}

//func makeGetUserEndpoint(s Service) endpoint.Endpoint{
//	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
//		req := request.(GetUserRequest)
//		email, err := s.GetUser(ctx, req.Id)
//		return GetUserResponse{Email:email}, err
//	}
//}

func makeGetUserEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetUserRequest)
		user, err := s.GetUser(ctx, req.Id)

		return GetUserResponse{
			User: user,
		}, err
	}
}

func makeLoginEndpoint(s Service) endpoint.Endpoint{
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {

		req := request.(LoginRequest)
		token, err := s.Login(ctx, req.Email, req.Password)
		return LoginResponse{Token:token}, err
	}
}


func makeCreateConfOwnProfileEndpoint(s Service) endpoint.Endpoint{
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(CreateConfOwnProfileRequest)
		ok, err := s.CreateConfOwnProfile(ctx, req.Reviews, req.Reads,req.Useful,req.Attend,req.Favourite,req.Picture)
		return CreateConfOwnProfileResponse{OK:ok}, err
	}
}
func makeGetConfOwnProfileEndpoint(s Service) endpoint.Endpoint{
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(GetConfOwnProfileRequest)
		ok, err := s.GetConfOwnProfile(ctx, req.Id)
		return GetConfOwnProfileResponse{ConfOwnProfile:ok}, err
	}
}
func makeUpdateConfOwnProfileEndpoint(s Service) endpoint.Endpoint{
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(UpdateConfOwnProfileRequest)
		ok, err := s.UpdateConfOwnProfile(ctx, req.Reviews, req.Reads,req.Useful,req.Attend,req.Favourite,req.Picture)
		return UpdateConfOwnProfileResponse{OK:ok}, err
	}
}
func makeDeleteConfOwnProfileEndpoint(s Service) endpoint.Endpoint{
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(DeleteConfOwnProfileRequest)
		ok, err := s.DeleteConfOwnProfile(ctx, req.Id)
		return DeleteConfOwnProfileResponse{OK:ok}, err
	}
}

func makeCreateConferenceEndpoint(s Service) endpoint.Endpoint{
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(CreateConferenceRequest)
		ok, err := s.CreateConference(ctx, req.Name, req.Website, req.About, req.Phone, req.Email, req.Address, req.City, req.ZipCode, req.Speakers, req.Facebook, req.Twitter, req.Instagram, req.OrganizerID, req.Locations, req.Rating )
		return CreateConferenceResponse{OK:ok}, err
	}
}
func makeGetConferenceEndpoint(s Service) endpoint.Endpoint{
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(GetConferenceRequest)
		ok, err := s.GetConference(ctx, req.Id)
		return GetConferenceResponse{Conference:ok}, err
	}
}
func makeUpdateCreateConferenceEndpoint(s Service) endpoint.Endpoint{
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(UpdateCreateConferenceRequest)
		ok, err := s.UpdateCreateConference(ctx, req.Name, req.Website, req.About, req.Phone, req.Email, req.Address, req.City, req.ZipCode, req.Speakers, req.Facebook, req.Twitter, req.Instagram, req.OrganizerID, req.Locations, req.Rating )
		return UpdateCreateConferenceResponse{OK:ok}, err
	}
}
func makeDeleteCreateConferenceEndpoint(s Service) endpoint.Endpoint{
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(DeleteCreateConferenceRequest)
		ok, err := s.DeleteCreateConference(ctx, req.Id)
		return DeleteCreateConferenceResponse{OK:ok}, err
	}
}

func makeCreateLocationEndpoint(s Service) endpoint.Endpoint{
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(CreateLocationRequest)
		ok, err := s.CreateLocation(ctx, req.Name, req.Date, req.Time )
		return CreateLocationResponse{OK:ok}, err
	}
}
func makeGetLocationEndpoint(s Service) endpoint.Endpoint{
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(GetLocationRequest)
		ok, err := s.GetLocation(ctx, req.Id)
		return GetLocationResponse{Location:ok}, err
	}
}
func makeUpdateCreateLocationEndpoint(s Service) endpoint.Endpoint{
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(UpdateCreateLocationRequest)
		ok, err := s.UpdateCreateLocation(ctx, req.Name, req.Date, req.Time )
		return UpdateCreateLocationResponse{OK:ok}, err
	}
}
func makeDeleteCreateLocationEndpoint(s Service) endpoint.Endpoint{
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(DeleteCreateLocationRequest)
		ok, err := s.DeleteCreateLocation(ctx, req.Id)
		return DeleteCreateLocationResponse{OK:ok}, err
	}
}

func makeCreateRatingEndpoint(s Service) endpoint.Endpoint{
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(CreateRatingRequest)
		ok, err := s.CreateRating(ctx, req.Rate, req.Comment, req.Caption, req.AttendQ, req.EnjoyQ, req.LocationQ, req.ConnectPeerQ, req.Awesome, req.Great, req.Average, req.Poor, req.Terrible, req.Favorite, req.Like)
		return CreateRatingResponse{OK:ok}, err
	}
}
func makeGetRatingEndpoint(s Service) endpoint.Endpoint{
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(GetRatingRequest)
		ok, err := s.GetRating(ctx, req.Id)
		return GetRatingResponse{Rating:ok}, err
	}
}
func makeUpdateCreateRatingEndpoint(s Service) endpoint.Endpoint{
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(UpdateCreateRatingRequest)
		ok, err := s.UpdateCreateRating(ctx, req.Rate,req.Comment, req.Caption,req.AttendQ,req.EnjoyQ,req.LocationQ,req.ConnectPeerQ,req.Awesome,req.Great,req.Average,req.Poor,req.Terrible,req.Favorite,req.Like)
		return UpdateCreateRatingResponse{OK:ok}, err
	}
}
func makeDeleteCreateRatingEndpoint(s Service) endpoint.Endpoint{
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(DeleteCreateRatingRequest)
		ok, err := s.DeleteCreateRating(ctx, req.Id)
		return DeleteCreateRatingResponse{OK:ok}, err
	}
}

func makeCreateReportEndpoint(s Service) endpoint.Endpoint{
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(CreateReportRequest)
		ok, err := s.CreateReport(ctx, req.Offensive,req.Violence,req.Spam,req.InAppropriate)
		return CreateReportResponse{OK:ok}, err
	}
}
func makeGetReportEndpoint(s Service) endpoint.Endpoint{
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(GetReportRequest)
		ok, err := s.GetReport(ctx, req.Id)
		return GetReportResponse{Report:ok}, err
	}
}
func makeUpdateCreateReportEndpoint(s Service) endpoint.Endpoint{
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(UpdateCreateReportRequest)
		ok, err := s.UpdateCreateReport(ctx, req.Offensive,req.Violence,req.Spam,req.InAppropriate)
		return UpdateCreateReportResponse{OK:ok}, err
	}
}
func makeDeleteCreateReportEndpoint(s Service) endpoint.Endpoint{
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(DeleteCreateReportRequest)
		ok, err := s.DeleteCreateReport(ctx, req.Id)
		return DeleteCreateReportResponse{OK:ok}, err
	}
}

func makeCreateSpeakerEndpoint(s Service) endpoint.Endpoint{
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(CreateSpeakerRequest)
		ok, err := s.CreateSpeaker(ctx, req.Name,req.Position)
		return CreateSpeakerResponse{OK:ok}, err
	}
}
func makeGetSpeakerEndpoint(s Service) endpoint.Endpoint{
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(GetSpeakerRequest)
		ok, err := s.GetSpeaker(ctx, req.Id)
		return GetSpeakerResponse{Speaker:ok}, err
	}
}
func makeUpdateCreateSpeakerEndpoint(s Service) endpoint.Endpoint{
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(UpdateCreateSpeakerRequest)
		ok, err := s.UpdateCreateSpeaker(ctx, req.Name,req.Position,)
		return UpdateCreateSpeakerResponse{OK:ok}, err
	}
}
func makeDeleteCreateSpeakerEndpoint(s Service) endpoint.Endpoint{
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(DeleteCreateSpeakerRequest)
		ok, err := s.DeleteCreateSpeaker(ctx, req.Id)
		return DeleteCreateSpeakerResponse{OK:ok}, err
	}
}
func makeCreateCategoryEndpoints(s Service) endpoint.Endpoint{
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(CreateCategoryRequest)
		ok, err := s.CreateCategory(ctx, req.Name)
		return CreateCategoryResponse{OK:ok}, err
	}
}

func makeGetCategoryEndpoints(s Service) endpoint.Endpoint{
	return func(ctx context.Context, request interface{})(response interface{}, err error){
		req := request.(GetCategoryRequest)
		cat, err := s.GetCategory(ctx, req.Id)
		return GetCategoryResponse{Category:cat}, err
	}
}

func makeGetCategoriesEndpoints(s Service) endpoint.Endpoint{
	return func(ctx context.Context, request interface{})(response interface{}, err error){
		_ = request.(GetCategoriesRequest)
		cat, err := s.GetCategories(ctx)
		fmt.Println("endpoint", cat)
		return GetCategoriesResponse{Category:cat}, err
	}
}

func makeUpdateCategoryEndpoints(s Service) endpoint.Endpoint{
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(UpdateCategoryRequest)
		cat, err := s.UpdateCategory(ctx, req.Name)
		return UpdateCategoryResponse{Name:cat}, err
	}
}