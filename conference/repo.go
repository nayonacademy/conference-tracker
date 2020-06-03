package conference

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/jinzhu/gorm"
)

var RepoErr = errors.New("unable to handle repo request")

type repo struct {
	db *gorm.DB
	logger log.Logger
}

func (r repo) CreateConfOwnProfile(ctx context.Context, confprofile ConfOwnProfile) (string, error) {
	//Create
	err := r.db.Create(&ConfOwnProfile{
		Reviews:   confprofile.Reviews,
		Reads:     confprofile.Reads,
		Useful:    confprofile.Useful,
		Attend:    confprofile.Attend,
		Favourite: confprofile.Favourite,
		Picture:   confprofile.Picture,
	}).Error
	if err != nil{
		return "",RepoErr
	}
	return "success", nil
}

func (r repo) GetConfOwnProfile(ctx context.Context, id string) (interface{}, error) {
	var profile string
	var confprofile ConfOwnProfile
	err := r.db.Find(&confprofile, "id = ?",id)
	if err != nil{
		return "", RepoErr
	}
	return profile, nil
}

func (r repo) UpdateConfOwnProfile(ctx context.Context, confprofile ConfOwnProfile) (string, error) {
	err := r.db.Where("id = ?", confprofile.ID).Find(&confprofile).Update("reviews", confprofile.Reviews,"reads",confprofile.Reads,"useful",confprofile.Useful,"attend",confprofile.Attend,"favourite",confprofile.Favourite,"picture",confprofile.Picture).Error
	if err != nil{
		return "", RepoErr
	}
	return "confprofile.ID", nil
}

func (r repo) DeleteConfOwnProfile(ctx context.Context, id string) (string, error) {
	var confprofile ConfOwnProfile
	err := r.db.Where("id = ?", id).Find(&confprofile).Delete(ConfOwnProfile{}).Error
	if err != nil{
		return "", RepoErr
	}
	return confprofile.Picture, nil
}

func (r repo) CreateConference(ctx context.Context, conference Conference) (string, error) {
	err := r.db.Create(&Conference{
		Name:        conference.Name,
		Website:     conference.Website,
		About:       conference.About,
		Phone:       conference.About,
		Email:       conference.Phone,
		Address:     conference.Address,
		City:        conference.City,
		ZipCode:     conference.ZipCode,
		Speakers:    conference.Speakers,
		Facebook:    conference.Facebook,
		Twitter:     conference.Twitter,
		Instagram:   conference.Instagram,
		OrganizerID: conference.OrganizerID,
		Locations:   conference.Locations,
		Rating:      conference.Rating,
	}).Error
	if err != nil{
		return "", RepoErr
	}
	return "success", nil
}

func (r repo) GetConference(ctx context.Context, id string) (interface{}, error) {
	var profile string
	var conference Conference
	err := r.db.Find(&conference, "id = ?",id)
	if err != nil{
		return "", RepoErr
	}
	return profile, nil
}

func (r repo) UpdateCreateConference(ctx context.Context, conference Conference) (string, error) {
	err := r.db.Where("id = ?", conference.ID).Find(&conference).Update(
		"name",conference.Name,
		"website",conference.Website,
		"about", conference.About,
		"phone",conference.Phone,
		"email",conference.Email,
		"address",conference.Address,
		"city",conference.City,
		"zip_code",conference.ZipCode,
		"speakers",conference.Speakers,
		"facebook", conference.Facebook,
		"twitter",conference.Twitter,
		"instagram",conference.Instagram,
		"organizer_id",conference.OrganizerID,
		"locations",conference.Locations,
		"rating",conference.Rating).Error
	if err != nil{
		return "", RepoErr
	}
	return "conference.ID", nil
}

func (r repo) DeleteCreateConference(ctx context.Context, id string) (string, error) {
	var conference Conference
	err := r.db.Where("id = ?", id).Find(&conference).Delete(Conference{}).Error
	if err != nil{
		return "", RepoErr
	}
	return conference.Name, nil
}

func (r repo) CreateLocation(ctx context.Context, location Location) (string, error) {
	err := r.db.Create(&Location{
		Name:  location.Name,
		Date:  location.Date,
		Time:  location.Time,
	}).Error
	if err != nil{
		return "", RepoErr
	}
	return "success", nil
}

func (r repo) GetLocation(ctx context.Context, id string) (interface{}, error) {
	var profile string
	var location Location
	err := r.db.Find(&location, "id = ?",id)
	if err != nil{
		return "", RepoErr
	}
	return profile, nil
}

func (r repo) UpdateCreateLocation(ctx context.Context, location Location) (string, error) {
	err := r.db.Where("id = ?", location.ID).Find(&location).Update(
	"name",location.Name,
	"date", location.Date,
	"time", location.Time).Error
	if err != nil{
		return "", RepoErr
	}
	return "location.ID", nil
}

func (r repo) DeleteCreateLocation(ctx context.Context, id string) (string, error) {
	var location Location
	err := r.db.Where("id = ?", id).Find(&location).Delete(Location{}).Error
	if err != nil{
		return "", RepoErr
	}
	return location.Name, nil
}

func (r repo) CreateRating(ctx context.Context, rating Rating) (string, error) {
	err := r.db.Create(&Rating{
		Rate:         rating.Rate,
		Comment:      rating.Comment,
		Caption:      rating.Caption,
		AttendQ:      rating.AttendQ,
		EnjoyQ:       rating.EnjoyQ,
		LocationQ:    rating.LocationQ,
		ConnectPeerQ: rating.ConnectPeerQ,
		Awesome:      rating.Awesome,
		Great:        rating.Great,
		Average:      rating.Average,
		Poor:         rating.Poor,
		Terrible:     rating.Terrible,
		Favorite:     rating.Favorite,
		Like:         rating.Like,
	}).Error
	if err != nil{
		return "", RepoErr
	}
	return "success", nil
}

func (r repo) GetRating(ctx context.Context, id string) (interface{}, error) {
	var profile string
	var rating Rating
	err := r.db.Find(&rating, "id = ?",id)
	if err != nil{
		return "", RepoErr
	}
	return profile, nil
}

func (r repo) UpdateCreateRating(ctx context.Context, rating Rating) (string, error) {
	err := r.db.Where("id = ?", rating.ID).Find(&rating).Update(
		"rate",rating.Rate,
		"comment", rating.Comment,
		"caption", rating.Caption,
		"attend_q", rating.AttendQ,
		"enjoy_q",rating.EnjoyQ,
		"location_q",rating.LocationQ,
		"connect_peer_q",rating.ConnectPeerQ,
		"awesome",rating.Awesome,
		"great", rating.Great,
		"average",rating.Average,
		"poor",rating.Poor,
		"terrible",rating.Terrible,
		"favorite",rating.Favorite,
		"like", rating.Like,
		).Error
	if err != nil{
		return "", RepoErr
	}
	return "rating.ID", nil
}

func (r repo) DeleteCreateRating(ctx context.Context, id string) (string, error) {
	var rating Rating
	err := r.db.Where("id = ?", id).Find(&rating).Delete(Rating{}).Error
	if err != nil{
		return "", RepoErr
	}
	return rating.Caption, nil
}

func (r repo) CreateReport(ctx context.Context, report Report) (string, error) {
	err := r.db.Create(&Report{
		Offensive:     report.Offensive,
		Violence:      report.Violence,
		Spam:          report.Spam,
		InAppropriate: report.InAppropriate,
	}).Error
	if err != nil{
		return "", RepoErr
	}
	return "success", nil
}

func (r repo) GetReport(ctx context.Context, id string) (interface{}, error) {
	var profile string
	var report Report
	err := r.db.Find(&report, "id = ?",id)
	if err != nil{
		return "", RepoErr
	}
	return profile, nil
}

func (r repo) UpdateCreateReport(ctx context.Context, report Report) (string, error) {
	err := r.db.Where("id = ?", report.ID).Find(&report).Update(
		"offensive", report.Offensive,
		"violence", report.Violence,
		"spam",report.Spam,
		"in_appropriate",report.InAppropriate,
	).Error
	if err != nil{
		return "", RepoErr
	}
	return "report.ID", nil
}

func (r repo) DeleteCreateReport(ctx context.Context, id string) (string, error) {
	var report Report
	err := r.db.Where("id = ?", id).Find(&report).Delete(Report{}).Error
	if err != nil{
		return "", RepoErr
	}
	return "delete", nil
}

func (r repo) CreateSpeaker(ctx context.Context, speaker Speaker) (string, error) {
	err := r.db.Create(&Speaker{
		Name:     speaker.Name,
		Position: speaker.Position,
	}).Error
	if err != nil{
		return "", RepoErr
	}
	return "success", nil
}

func (r repo) GetSpeaker(ctx context.Context, id string) (interface{}, error) {
	var profile string
	var speaker Speaker
	err := r.db.Find(&speaker, "id = ?",id)
	if err != nil{
		return "", RepoErr
	}
	fmt.Println(speaker)
	return profile, nil
}

func (r repo) UpdateCreateSpeaker(ctx context.Context, speaker Speaker) (string, error) {
	err := r.db.Where("id = ?", speaker.ID).Find(&speaker).Update(
		"name", speaker.Name,
		"position", speaker.Position,
	).Error
	if err != nil{
		return "", RepoErr
	}
	return "speaker.ID", nil
}

func (r repo) DeleteCreateSpeaker(ctx context.Context, id string) (string, error) {
	var speaker Speaker
	err := r.db.Where("id = ?", id).Find(&speaker).Delete(Speaker{}).Error
	if err != nil{
		return "", RepoErr
	}
	return speaker.Name, nil
}


func NewRepo(db *gorm.DB, logger log.Logger) Repository{
	return &repo{
		db: db,
		logger : log.With(logger, "repo","gorm"),
	}
}