package conference

import (
	"context"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

type service struct {
	repostory Repository
	logger log.Logger
}

func (s service) CreateConfOwnProfile(ctx context.Context, confprofile ConfOwnProfile)(string, error) {
	logger := log.With(s.logger,"method","CreateConfOwnProfile")
	conf_profile := ConfOwnProfile{
		Reviews:   confprofile.Reviews,
		Reads:     confprofile.Reads,
		Useful:    confprofile.Useful,
		Attend:    confprofile.Attend,
		Favourite: confprofile.Favourite,
		Picture:   confprofile.Picture,
	}
	if _,err := s.repostory.CreateConfOwnProfile(ctx, conf_profile); err != nil{
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

func (s service) UpdateConfOwnProfile(ctx context.Context, confprofile ConfOwnProfile) (interface{}, error) {
	logger := log.With(s.logger,"method","UpdateConfOwnProfile")
	name, err := s.repostory.UpdateConfOwnProfile(ctx, confprofile)
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

func (s service) CreateConference(ctx context.Context, conference Conference) (string, error) {
	logger := log.With(s.logger,"method","CreateConference")
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

func (s service) UpdateCreateConference(ctx context.Context, conference Conference) (interface{}, error) {
	logger := log.With(s.logger,"method","UpdateCreateConference")
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

func (s service) CreateLocation(ctx context.Context, location Location) (string, error) {
	logger := log.With(s.logger,"method","CreateLocation")
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

func (s service) UpdateCreateLocation(ctx context.Context, location Location) (interface{}, error) {
	logger := log.With(s.logger,"method","UpdateCreateLocation")
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

func (s service) CreateRating(ctx context.Context, rating Rating) (string, error) {
	logger := log.With(s.logger,"method","CreateRating")
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

func (s service) UpdateCreateRating(ctx context.Context, rating Rating) (interface{}, error) {
	logger := log.With(s.logger,"method","UpdateCreateRating")
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

func (s service) CreateReport(ctx context.Context, report Report) (string, error) {
	logger := log.With(s.logger,"method","CreateReport")
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

func (s service) UpdateCreateReport(ctx context.Context, report Report) (interface{}, error) {
	logger := log.With(s.logger,"method","UpdateCreateReport")
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

func (s service) CreateSpeaker(ctx context.Context, speaker Speaker) (string, error) {
	logger := log.With(s.logger,"method","CreateSpeaker")
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

func (s service) UpdateCreateSpeaker(ctx context.Context, speaker Speaker) (interface{}, error) {
	logger := log.With(s.logger,"method","UpdateCreateSpeaker")
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