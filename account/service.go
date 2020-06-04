package account

import (
	"context"
	"time"
)

type Service interface {
	CreateUser(ctx context.Context, email string, password string) (string, error)
	GetUser(ctx context.Context, id string) (User, error)
	Login(ctx context.Context, email string, password string)(string, error)

	CreateCategory(ctx context.Context, name string) (string, error)
	GetCategories(ctx context.Context) ([]Category, error)
	GetCategory(ctx context.Context, id string)(Category, error)
	UpdateCategory(ctx context.Context, name string)(string, error)

	CreateConfOwnProfile(ctx context.Context, reviews int16,reads int16,useful int16,attend int16,favourite string,picture string) (string, error)
	GetConfOwnProfile(ctx context.Context, id string)(interface{}, error)
	UpdateConfOwnProfile(ctx context.Context, reviews int16,reads int16,useful int16,attend int16,favourite string,picture string)(string, error)
	DeleteConfOwnProfile(ctx context.Context, id string)(string, error)

	CreateConference(ctx context.Context, name string,website string, about string,	phone string,email string,address string,city string,zipcode string,speakers []Speaker,	facebook string,twitter string,	instagram string,organizer_id int16, locations []Location, rating Rating) (string, error)
	GetConference(ctx context.Context, id string)(interface{}, error)
	UpdateCreateConference(ctx context.Context, name string,website string, about string,	phone string,email string,address string,city string,zipcode string,speakers []Speaker,	facebook string,twitter string,	instagram string,organizer_id int16, locations []Location, rating Rating)(string, error)
	DeleteCreateConference(ctx context.Context, id string)(string, error)

	CreateLocation(ctx context.Context, name string, date time.Month, time time.Time) (string, error)
	GetLocation(ctx context.Context, id string)(interface{}, error)
	UpdateCreateLocation(ctx context.Context, name string, date time.Month, time time.Time)(string, error)
	DeleteCreateLocation(ctx context.Context, id string)(string, error)

	CreateRating(ctx context.Context, rate int16, comment string, caption string, attend_q bool, enjoy_q bool, location_q bool, connectPeer_q bool, awesome int16, great int16, average int16, poor int16, terrible int16, favorite bool, like bool ) (string, error)
	GetRating(ctx context.Context, id string)(interface{}, error)
	UpdateCreateRating(ctx context.Context, rate int16, comment string, caption string, attend_q bool, enjoy_q bool, location_q bool, connectPeer_q bool, awesome int16, great int16, average int16, poor int16, terrible int16, favorite bool, like bool )(string, error)
	DeleteCreateRating(ctx context.Context, id string)(string, error)

	CreateReport(ctx context.Context, offensive bool,violence bool,spam bool,in_appropriate bool) (string, error)
	GetReport(ctx context.Context, id string)(interface{}, error)
	UpdateCreateReport(ctx context.Context, offensive bool,violence bool,spam bool,in_appropriate bool)(string, error)
	DeleteCreateReport(ctx context.Context, id string)(string, error)

	CreateSpeaker(ctx context.Context, name string, position string) (string, error)
	GetSpeaker(ctx context.Context, id string)(interface{}, error)
	UpdateCreateSpeaker(ctx context.Context, name string, position string)(string, error)
	DeleteCreateSpeaker(ctx context.Context, id string)(string, error)
}