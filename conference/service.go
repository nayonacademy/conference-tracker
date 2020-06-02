package conference

import "context"

type Service interface {
	CreateConfOwnProfile(ctx context.Context, reviews int16,reads int16,useful int16,attend int16,favourite string,picture string) (string, error)
	GetConfOwnProfile(ctx context.Context, id string)(interface{}, error)
	UpdateConfOwnProfile(ctx context.Context, confprofile ConfOwnProfile)(interface{}, error)
	DeleteConfOwnProfile(ctx context.Context, id string)(string, error)

	CreateConference(ctx context.Context, conference Conference) (string, error)
	GetConference(ctx context.Context, id string)(interface{}, error)
	UpdateCreateConference(ctx context.Context, conference Conference)(interface{}, error)
	DeleteCreateConference(ctx context.Context, id string)(string, error)

	CreateLocation(ctx context.Context, location Location) (string, error)
	GetLocation(ctx context.Context, id string)(interface{}, error)
	UpdateCreateLocation(ctx context.Context, location Location)(interface{}, error)
	DeleteCreateLocation(ctx context.Context, id string)(string, error)

	CreateRating(ctx context.Context, rating Rating) (string, error)
	GetRating(ctx context.Context, id string)(interface{}, error)
	UpdateCreateRating(ctx context.Context, rating Rating)(interface{}, error)
	DeleteCreateRating(ctx context.Context, id string)(string, error)

	CreateReport(ctx context.Context, report Report) (string, error)
	GetReport(ctx context.Context, id string)(interface{}, error)
	UpdateCreateReport(ctx context.Context, report Report)(interface{}, error)
	DeleteCreateReport(ctx context.Context, id string)(string, error)

	CreateSpeaker(ctx context.Context, speaker Speaker) (string, error)
	GetSpeaker(ctx context.Context, id string)(interface{}, error)
	UpdateCreateSpeaker(ctx context.Context, speaker Speaker)(interface{}, error)
	DeleteCreateSpeaker(ctx context.Context, id string)(string, error)
}