package account

import (
	"context"
	"errors"
	"github.com/davecgh/go-spew/spew"
	"github.com/go-kit/kit/log"
	"github.com/jinzhu/gorm"
)

var RepoErr = errors.New("unable to handle repo request")

type repo struct {
	db *gorm.DB
	logger log.Logger
}

func (r *repo) CreateUser(ctx context.Context, user User) error {
	var usr User
	err := r.db.First(&user, "email = ?", user.Email).Scan(&usr).Error
	if usr.Email == user.Email{
		return errors.New("user already exists")
	}
	// Create
	err = r.db.Create(&User{Email: user.Email, Password: user.Password}).Error
	if err != nil{
		return RepoErr
	}
	return nil
}

//func (r *repo) GetUser(ctx context.Context, id string) (string, error) {
//	var user User
//	err := r.db.First(&user, "id = ?", id)
//	if err != nil{
//		return "", RepoErr
//	}
//	spew.Dump(user)
//	spew.Dump(id)
//	return "email", nil
//}

func (r *repo) GetUser(ctx context.Context, id string) (User, error) {
	var user User
	result := r.db.Select("email").Where("id = ?", id).First(&user).Scan(&user)
	//db.Select("name, age").Find(&users)
	if result.Error != nil{
		return User{}, RepoErr
	}
	spew.Dump(user.Email)
	return user, nil
}

func (r *repo) Login(ctx context.Context, email string, password string) (string, error) {
	var user User
	var token string
	err := r.db.Where("email = ? and password = ?", email, password).Find(&user).Error
	if err != nil{
		return "", RepoErr
	}
	token , errs := Sign(email, password)
	if errs != nil{
		return "", RepoErr
	}
	return token, nil
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

func (r repo) GetLocation(ctx context.Context, id string) (Location, error) {
	var location Location
	result := r.db.Where("id = ?", id).First(&location).Scan(&location)
	//result := r.db.First(&category).Scan(&category)
	if result.Error != nil{
		return Location{}, RepoErr
	}
	return location, nil
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

func (r repo) GetSpeaker(ctx context.Context, id string) (Speaker, error) {
	var speaker Speaker
	result := r.db.Where("id = ?", id).First(&speaker).Scan(&speaker)
	//result := r.db.First(&category).Scan(&category)
	if result.Error != nil{
		return Speaker{}, RepoErr
	}
	return speaker, nil
}

func (r repo) GetAllSpeaker(ctx context.Context)([]Speaker, error) {
	var speaker []Speaker
	result := r.db.Find(&speaker).Scan(&speaker)
	if result.Error != nil{
		return []Speaker{}, RepoErr
	}
	return speaker, nil
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
func (r repo) CreateCategory(ctx context.Context, category Category) error {
	var cat Category
	err := r.db.Find(&category, "name = ?", category.Name).Scan(&cat).Error
	if cat.Name == category.Name{
		return errors.New("category already exists")
	}
	//Create
	err = r.db.Create(&Category{Name:category.Name}).Error
	if err != nil{
		return RepoErr
	}
	return nil
}

func (r repo) GetCategory(ctx context.Context, id string) ([]Category, error) {

	var category []Category
	result := r.db.Where("id = ?", id).First(&category).Scan(&category)
	//result := r.db.First(&category).Scan(&category)
	if result.Error != nil{
		return []Category{}, RepoErr
	}
	return category, nil
}
func (r repo) GetCategories(ctx context.Context) ([]Category, error) {
	var category []Category
	//result := r.db.Where("id = ?", id).First(&category).Scan(&category)
	result := r.db.Find(&category).Scan(&category)
	if result.Error != nil{
		return []Category{}, RepoErr
	}
	return category, nil
}
func (r repo) UpdateCategory(ctx context.Context, name string) (string, error) {
	var category Category
	err := r.db.Where("name = ?", name).Find(&category).Update("name", name).Error
	if err != nil{
		return "", RepoErr
	}
	return category.Name, nil
}
func NewRepo(db *gorm.DB, logger log.Logger) Repository{
	return &repo{
		db:    db,
		logger: log.With(logger,"repo","gorm"),
	}
}