package conference

import "context"

type Service interface {
	CreateConfOwnProfile(ctx context.Context, confprofile ConfOwnProfile) error
	GetConfOwnProfile(ctx context.Context, id string)(interface{}, error)
	UpdateConfOwnProfile(ctx context.Context, confprofile ConfOwnProfile)(interface{}, error)
	DeleteConfOwnProfile(ctx context.Context, id string)(string, error)

	CreateConference(ctx context.Context, conference Conference) error
	GetConference(ctx context.Context, id string)(interface{}, error)
	UpdateCreateConference(ctx context.Context, confprofile ConfOwnProfile)(interface{}, error)
	DeleteCreateConference(ctx context.Context, id string)(string, error)

	CreateLocation(ctx context.Context, location Location) error
	GetLocation(ctx context.Context, id string)(interface{}, error)
	UpdateCreateLocation(ctx context.Context, confprofile ConfOwnProfile)(interface{}, error)
	DeleteCreateLocation(ctx context.Context, id string)(string, error)

	CreateRating(ctx context.Context, rating Rating) error
	GetRating(ctx context.Context, id string)(interface{}, error)
	UpdateCreateRating(ctx context.Context, confprofile ConfOwnProfile)(interface{}, error)
	DeleteCreateRating(ctx context.Context, id string)(string, error)

	CreateReport(ctx context.Context, report Report) error
	GetReport(ctx context.Context, id string)(interface{}, error)
	UpdateCreateReport(ctx context.Context, confprofile ConfOwnProfile)(interface{}, error)
	DeleteCreateReport(ctx context.Context, id string)(string, error)

	CreateSpeaker(ctx context.Context, speaker Speaker) error
	GetSpeaker(ctx context.Context, id string)(interface{}, error)
	UpdateCreateSpeaker(ctx context.Context, confprofile ConfOwnProfile)(interface{}, error)
	DeleteCreateSpeaker(ctx context.Context, id string)(string, error)
}