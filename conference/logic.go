package conference

import (
	"context"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"time"
)

type service struct {
	repostory Repository
	logger log.Logger
}

func (s service) CreateConfOwnProfile(ctx context.Context, reviews int16,reads int16,useful int16,attend int16,favourite string,picture string)(string, error) {
	logger := log.With(s.logger,"method","CreateConfOwnProfile")
	conf_profile := ConfOwnProfile{
		Reviews:   reviews,
		Reads:     reads,
		Useful:    useful,
		Attend:    attend,
		Favourite: favourite,
		Picture:   picture,
	}
	if _,err := s.repostory.CreateConfOwnProfile(ctx,conf_profile); err != nil{
		level.Error(logger).Log("err",err)
		return "", err
	}
	logger.Log("create user", conf_profile.Picture)
	return "success", nil
}

func (s service) GetConfOwnProfile(ctx context.Context, id string) (interface{}, error) {
	logger := log.With(s.logger,"method","GetConfOwnProfile")
	name, err := s.repostory.GetConfOwnProfile(ctx, id)
	if err != nil{
		level.Error(logger).Log("err",err)
		return "", err
	}
	logger.Log("get category", name)
	return name, nil
}

func (s service) UpdateConfOwnProfile(ctx context.Context, reviews int16,reads int16,useful int16,attend int16,favourite string,picture string) (string, error) {
	logger := log.With(s.logger,"method","UpdateConfOwnProfile")
	conf_profile := ConfOwnProfile{
		Reviews:   reviews,
		Reads:     reads,
		Useful:    useful,
		Attend:    attend,
		Favourite: favourite,
		Picture:   picture,
	}
	name, err := s.repostory.UpdateConfOwnProfile(ctx, conf_profile)
	if err != nil{
		level.Error(logger).Log("err",err)
		return "", err
	}
	logger.Log("get category", name)
	return name, nil
}

func (s service) DeleteConfOwnProfile(ctx context.Context, id string) (string, error) {
	logger := log.With(s.logger,"method","DeleteConfOwnProfile")
	name, err := s.repostory.DeleteConfOwnProfile(ctx, id)
	if err != nil{
		level.Error(logger).Log("err",err)
		return "", err
	}
	logger.Log("get category", name)
	return name, nil
}

func (s service) CreateConference(ctx context.Context, name string,website string, about string,	phone string,email string,address string,city string,zipcode string,speakers []Speaker,	facebook string,twitter string,	instagram string,organizer_id int16, locations []Location, rating Rating) (string, error) {
	logger := log.With(s.logger,"method","CreateConference")
	conference := Conference{
		Name:        name,
		Website:     website,
		About:       about,
		Phone:       phone,
		Email:       email,
		Address:     address,
		City:        city,
		ZipCode:     zipcode,
		Speakers:    speakers,
		Facebook:    facebook,
		Twitter:     twitter,
		Instagram:   instagram,
		OrganizerID: organizer_id,
		Locations:   locations,
		Rating:      rating,
	}
	conf, err := s.repostory.CreateConference(ctx, conference)
	if err != nil{
		level.Error(logger).Log("err",err)
		return "", err
	}
	logger.Log("get category")
	return conf, nil
}

func (s service) GetConference(ctx context.Context, id string) (interface{}, error) {
	logger := log.With(s.logger,"method","GetConference")
	name, err := s.repostory.GetConference(ctx, id)
	if err != nil{
		level.Error(logger).Log("err",err)
		return "", err
	}
	logger.Log("get category", name)
	return name, nil
}

func (s service) UpdateCreateConference(ctx context.Context, name string,website string, about string,	phone string,email string,address string,city string,zipcode string,speakers []Speaker,	facebook string,twitter string,	instagram string,organizer_id int16, locations []Location, rating Rating) (string, error) {
	logger := log.With(s.logger,"method","UpdateCreateConference")
	conference := Conference{
		Name:        name,
		Website:     website,
		About:       about,
		Phone:       phone,
		Email:       email,
		Address:     address,
		City:        city,
		ZipCode:     zipcode,
		Speakers:    speakers,
		Facebook:    facebook,
		Twitter:     twitter,
		Instagram:   instagram,
		OrganizerID: organizer_id,
		Locations:   locations,
		Rating:      rating,
	}
	name, err := s.repostory.UpdateCreateConference(ctx, conference)
	if err != nil{
		level.Error(logger).Log("err",err)
		return "", err
	}
	logger.Log("get category", name)
	return name, nil
}

func (s service) DeleteCreateConference(ctx context.Context, id string) (string, error) {
	logger := log.With(s.logger,"method","DeleteCreateConference")
	name, err := s.repostory.DeleteCreateConference(ctx, id)
	if err != nil{
		level.Error(logger).Log("err",err)
		return "", err
	}
	logger.Log("get category", name)
	return name, nil
}

func (s service) CreateLocation(ctx context.Context, name string, date time.Month, time time.Time) (string, error) {
	logger := log.With(s.logger,"method","CreateLocation")
	location := Location{
		Name:  name,
		Date:  date,
		Time:  time,
	}
	_, err := s.repostory.CreateLocation(ctx, location)
	if err != nil{
		level.Error(logger).Log("err",err)
		return "", err
	}
	logger.Log("get category")
	return "location", nil
}

func (s service) GetLocation(ctx context.Context, id string) (interface{}, error) {
	logger := log.With(s.logger,"method","GetLocation")
	name, err := s.repostory.GetLocation(ctx, id)
	if err != nil{
		level.Error(logger).Log("err",err)
		return "", err
	}
	logger.Log("get category", name)
	return name, nil
}

func (s service) UpdateCreateLocation(ctx context.Context, name string, date time.Month, time time.Time) (string, error) {
	logger := log.With(s.logger,"method","UpdateCreateLocation")
	location := Location{
		Name:  name,
		Date:  date,
		Time:  time,
	}
	name, err := s.repostory.UpdateCreateLocation(ctx, location)
	if err != nil{
		level.Error(logger).Log("err",err)
		return "", err
	}
	logger.Log("get category", name)
	return name, nil
}

func (s service) DeleteCreateLocation(ctx context.Context, id string) (string, error) {
	logger := log.With(s.logger,"method","DeleteCreateLocation")
	name, err := s.repostory.DeleteCreateConference(ctx, id)
	if err != nil{
		level.Error(logger).Log("err",err)
		return "", err
	}
	logger.Log("get category", name)
	return name, nil
}

func (s service) CreateRating(ctx context.Context, rate int16, comment string, caption string, attend_q bool, enjoy_q bool, location_q bool, connectPeer_q bool, awesome int16, great int16, average int16, poor int16, terrible int16, favorite bool, like bool) (string, error) {
	logger := log.With(s.logger,"method","CreateRating")
	rating := Rating{
		Rate:         rate,
		Comment:      comment,
		Caption:      caption,
		AttendQ:      attend_q,
		EnjoyQ:       enjoy_q,
		LocationQ:    location_q,
		ConnectPeerQ: connectPeer_q,
		Awesome:      awesome,
		Great:        great,
		Average:      average,
		Poor:         poor,
		Terrible:     terrible,
		Favorite:     favorite,
		Like:         like,
	}
	_, err := s.repostory.CreateRating(ctx, rating)
	if err != nil{
		level.Error(logger).Log("err",err)
		return "", err
	}
	logger.Log("get category")
	return "name", nil
}

func (s service) GetRating(ctx context.Context, id string) (interface{}, error) {
	logger := log.With(s.logger,"method","GetRating")
	name, err := s.repostory.GetRating(ctx, id)
	if err != nil{
		level.Error(logger).Log("err",err)
		return "", err
	}
	logger.Log("get category", name)
	return name, nil
}

func (s service) UpdateCreateRating(ctx context.Context, rate int16, comment string, caption string, attend_q bool, enjoy_q bool, location_q bool, connectPeer_q bool, awesome int16, great int16, average int16, poor int16, terrible int16, favorite bool, like bool) (string, error) {
	logger := log.With(s.logger,"method","UpdateCreateRating")
	rating := Rating{
		Rate:         rate,
		Comment:      comment,
		Caption:      caption,
		AttendQ:      attend_q,
		EnjoyQ:       enjoy_q,
		LocationQ:    location_q,
		ConnectPeerQ: connectPeer_q,
		Awesome:      awesome,
		Great:        great,
		Average:      average,
		Poor:         poor,
		Terrible:     terrible,
		Favorite:     favorite,
		Like:         like,
	}
	name, err := s.repostory.UpdateCreateRating(ctx, rating)
	if err != nil{
		level.Error(logger).Log("err",err)
		return "", err
	}
	logger.Log("get category", name)
	return name, nil
}

func (s service) DeleteCreateRating(ctx context.Context, id string) (string, error) {
	logger := log.With(s.logger,"method","DeleteCreateRating")
	name, err := s.repostory.DeleteCreateRating(ctx, id)
	if err != nil{
		level.Error(logger).Log("err",err)
		return "", err
	}
	logger.Log("get category", name)
	return name, nil
}

func (s service) CreateReport(ctx context.Context, offensive bool,violence bool,spam bool,in_appropriate bool) (string, error) {
	logger := log.With(s.logger,"method","CreateReport")
	report := Report{

		Offensive:     offensive,
		Violence:      violence,
		Spam:          spam,
		InAppropriate: in_appropriate,
	}
	_, err := s.repostory.CreateReport(ctx, report)
	if err != nil{
		level.Error(logger).Log("err",err)
		return "", err
	}
	logger.Log("get category")
	return "name", nil
}

func (s service) GetReport(ctx context.Context, id string) (interface{}, error) {
	logger := log.With(s.logger,"method","GetReport")
	name, err := s.repostory.GetReport(ctx, id)
	if err != nil{
		level.Error(logger).Log("err",err)
		return "", err
	}
	logger.Log("get category", name)
	return name, nil
}

func (s service) UpdateCreateReport(ctx context.Context, offensive bool,violence bool,spam bool,in_appropriate bool) (string, error) {
	logger := log.With(s.logger,"method","UpdateCreateReport")
	report := Report{

		Offensive:     offensive,
		Violence:      violence,
		Spam:          spam,
		InAppropriate: in_appropriate,
	}
	name, err := s.repostory.UpdateCreateReport(ctx, report)
	if err != nil{
		level.Error(logger).Log("err",err)
		return "", err
	}
	logger.Log("get category", name)
	return name, nil
}

func (s service) DeleteCreateReport(ctx context.Context, id string) (string, error) {
	logger := log.With(s.logger,"method","GetCategory")
	name, err := s.repostory.DeleteCreateReport(ctx, id)
	if err != nil{
		level.Error(logger).Log("err",err)
		return "", err
	}
	logger.Log("get category", name)
	return name, nil
}

func (s service) CreateSpeaker(ctx context.Context, name string, position string) (string, error) {
	logger := log.With(s.logger,"method","CreateSpeaker")
	speaker := Speaker{
		Name:     name,
		Position: position,
	}
	_, err := s.repostory.CreateSpeaker(ctx, speaker)
	if err != nil{
		level.Error(logger).Log("err",err)
		return "", err
	}
	logger.Log("get category")
	return "name", nil
}

func (s service) GetSpeaker(ctx context.Context, id string) (interface{}, error) {
	logger := log.With(s.logger,"method","GetSpeaker")
	name, err := s.repostory.GetSpeaker(ctx, id)
	if err != nil{
		level.Error(logger).Log("err",err)
		return "", err
	}
	logger.Log("get category", name)
	return name, nil
}

func (s service) UpdateCreateSpeaker(ctx context.Context, name string, position string) (string, error) {
	logger := log.With(s.logger,"method","UpdateCreateSpeaker")
	speaker := Speaker{
		Name:     name,
		Position: position,
	}
	name, err := s.repostory.UpdateCreateSpeaker(ctx, speaker)
	if err != nil{
		level.Error(logger).Log("err",err)
		return "", err
	}
	logger.Log("get category", name)
	return name, nil
}

func (s service) DeleteCreateSpeaker(ctx context.Context, id string) (string, error) {
	logger := log.With(s.logger,"method","DeleteCreateSpeaker")
	name, err := s.repostory.DeleteCreateSpeaker(ctx, id)
	if err != nil{
		level.Error(logger).Log("err",err)
		return "", err
	}
	logger.Log("get category", name)
	return name, nil
}

func NewService(rep Repository, logger log.Logger) Service {
	return &service{
		repostory: rep,
		logger:    logger,
	}
}