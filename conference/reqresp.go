package conference

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

type(
	CreateConfOwnProfileRequest struct {
		Reviews int16	`json:"reviews"`
		Reads int16	`json:"reads"`
		Useful int16 `json:"useful"`
		Attend int16	`json:"attend"`
		Favourite string	`json:"favourite"`
		Picture string	`json:"picture"`
	}
	CreateConfOwnProfileResponse struct {
		OK string	`json:"ok"`
	}
	GetConfOwnProfileRequest struct {
		Id string	`json:"id"`
	}
	GetConfOwnProfileResponse struct {
		ConfOwnProfile interface{} `json:"conf_own_profile,omitempty"`
	}
	UpdateConfOwnProfileRequest struct {
		Reviews int16	`json:"reviews"`
		Reads int16	`json:"reads"`
		Useful int16 `json:"useful"`
		Attend int16	`json:"attend"`
		Favourite string	`json:"favourite"`
		Picture string	`json:"picture"`
	}
	UpdateConfOwnProfileResponse struct {
		OK string	`json:"ok"`
	}
	DeleteConfOwnProfileRequest struct {
		Id string	`json:"id"`
	}
	DeleteConfOwnProfileResponse struct {
		OK string	`json:"ok"`
	}
	CreateConferenceRequest struct {
		Name string `json:"name"`
		Website string	`json:"website"`
		About string	`json:"about"`
		Phone string	`json:"phone"`
		Email string	`json:"email"`
		Address string	`json:"address"`
		City string	`json:"city"`
		ZipCode string	`json:"zip_code"`
		Speakers []Speaker	`json:"speakers"`
		Facebook string	`json:"facebook"`
		Twitter string	`json:"twitter"`
		Instagram string	`json:"instagram"`
		OrganizerID int16	`json:"organizer_id"`
		Locations []Location `json:"locations"`
		Rating Rating	`json:"rating"`
	}
	CreateConferenceResponse struct {
		OK string	`json:"ok"`
	}
	GetConferenceRequest struct {
		Id string	`json:"id"`
	}
	GetConferenceResponse struct {
		Conference interface{}	`json:"conference"`
	}
	UpdateCreateConferenceRequest struct {
		Name string `json:"name"`
		Website string	`json:"website"`
		About string	`json:"about"`
		Phone string	`json:"phone"`
		Email string	`json:"email"`
		Address string	`json:"address"`
		City string	`json:"city"`
		ZipCode string	`json:"zip_code"`
		Speakers []Speaker	`json:"speakers"`
		Facebook string	`json:"facebook"`
		Twitter string	`json:"twitter"`
		Instagram string	`json:"instagram"`
		OrganizerID int16	`json:"organizer_id"`
		Locations []Location `json:"locations"`
		Rating Rating	`json:"rating"`
	}
	UpdateCreateConferenceResponse struct {
		OK string	`json:"ok"`
	}
	DeleteCreateConferenceRequest struct {
		Id string	`json:"id"`
	}
	DeleteCreateConferenceResponse struct {
		OK string	`json:"ok"`
	}
	CreateLocationRequest struct {
		Name string	`json:"name"`
		Date time.Month	`json:"date"`
		Time time.Time	`json:"time"`
	}
	CreateLocationResponse struct {
		OK string	`json:"ok"`
	}
	GetLocationRequest struct {
		Id string	`json:"id"`
	}
	GetLocationResponse struct {
		Location interface{} `json:"location"`
	}
	UpdateCreateLocationRequest struct {
		Name string	`json:"name"`
		Date time.Month	`json:"date"`
		Time time.Time	`json:"time"`
	}
	UpdateCreateLocationResponse struct {
		OK string	`json:"ok"`
	}
	DeleteCreateLocationRequest struct {
		Id string	`json:"id"`
	}
	DeleteCreateLocationResponse struct {
		OK string	`json:"ok"`
	}
	CreateRatingRequest struct {
		Rate int16 `json:"rate"`
		Comment string `json:"comment"`
		Caption string `json:"caption"`
		AttendQ bool `json:"attend_q"`
		EnjoyQ bool `json:"enjoy_q"`
		LocationQ bool `json:"location_q"`
		ConnectPeerQ bool `json:"connect_peer_q"`
		Awesome int16 `json:"awesome"`
		Great int16 `json:"great"`
		Average int16 `json:"average"`
		Poor int16 `json:"poor"`
		Terrible int16 `json:"terrible"`
		Favorite bool `json:"favorite"`
		Like bool `json:"like"`
	}
	CreateRatingResponse struct {
		OK string	`json:"ok"`
	}
	GetRatingRequest struct {
		Id string	`json:"id"`
	}
	GetRatingResponse struct {
		Rating interface{} `json:"rating"`
	}
	UpdateCreateRatingRequest struct {
		Rate int16 `json:"rate"`
		Comment string `json:"comment"`
		Caption string `json:"caption"`
		AttendQ bool `json:"attend_q"`
		EnjoyQ bool `json:"enjoy_q"`
		LocationQ bool `json:"location_q"`
		ConnectPeerQ bool `json:"connect_peer_q"`
		Awesome int16 `json:"awesome"`
		Great int16 `json:"great"`
		Average int16 `json:"average"`
		Poor int16 `json:"poor"`
		Terrible int16 `json:"terrible"`
		Favorite bool `json:"favorite"`
		Like bool `json:"like"`
	}
	UpdateCreateRatingResponse struct {
		OK string	`json:"ok"`
	}
	DeleteCreateRatingRequest struct {
		Id string	`json:"id"`
	}
	DeleteCreateRatingResponse struct {
		OK string	`json:"ok"`
	}
	CreateReportRequest struct {
		Offensive bool	`json:"offensive"`
		Violence bool	`json:"violence"`
		Spam bool	`json:"spam"`
		InAppropriate bool	`json:"in_appropriate"`
	}
	CreateReportResponse struct {
		OK string	`json:"ok"`
	}
	GetReportRequest struct {
		Id string	`json:"id"`
	}
	GetReportResponse struct {
		Report interface{} `json:"report"`
	}
	UpdateCreateReportRequest struct {
		Offensive bool	`json:"offensive"`
		Violence bool	`json:"violence"`
		Spam bool	`json:"spam"`
		InAppropriate bool	`json:"in_appropriate"`
	}
	UpdateCreateReportResponse struct {
		OK string	`json:"ok"`
	}
	DeleteCreateReportRequest struct {
		Id string	`json:"id"`
	}
	DeleteCreateReportResponse struct {
		OK string	`json:"ok"`
	}
	CreateSpeakerRequest struct {
		Name string `json:"name"`
		Position string `json:"position"`
	}
	CreateSpeakerResponse struct {
		OK string	`json:"ok"`
	}
	GetSpeakerRequest struct {
		Id string	`json:"id"`
	}
	GetSpeakerResponse struct {
		Speaker interface{} `json:"speaker"`
	}
	UpdateCreateSpeakerRequest struct {
		Name string `json:"name"`
		Position string `json:"position"`
	}
	UpdateCreateSpeakerResponse struct {
		OK string	`json:"ok"`
	}
	DeleteCreateSpeakerRequest struct {
		Id string	`json:"id"`
	}
	DeleteCreateSpeakerResponse struct {
		OK string	`json:"ok"`
	}
)

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error{
	return json.NewEncoder(w).Encode(response)
}

func decodeCreateConfOwnProfileRequest(ctx context.Context, r *http.Request)(interface{}, error) {
	var req CreateConfOwnProfileRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil{
		return nil, err
	}
	return req, nil
	}

func decodeGetConfOwnProfileRequest(ctx context.Context, r *http.Request)(interface{}, error) {
	var req GetConfOwnProfileRequest
	vars := mux.Vars(r)
	req = GetConfOwnProfileRequest{Id:vars["id"]}
	return req, nil
	}

func decodeUpdateConfOwnProfileRequest(ctx context.Context, r *http.Request)(interface{}, error) {
	var req UpdateConfOwnProfileRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil{
		return nil, err
	}
	return req, nil
	}

func decodeDeleteConfOwnProfileRequest(ctx context.Context, r *http.Request)(interface{}, error) {
	var req DeleteConfOwnProfileRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil{
		return nil, err
	}
	return req, nil
	}

func decodeCreateConferenceRequest(ctx context.Context, r *http.Request)(interface{}, error) {
	var req CreateConferenceRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil{
		return nil, err
	}
	return req, nil
	}

func decodeGetConferenceRequest(ctx context.Context, r *http.Request)(interface{}, error) {
	var req GetConferenceRequest
	vars := mux.Vars(r)
	req = GetConferenceRequest{Id:vars["id"]}
	return req, nil
	}

func decodeUpdateCreateConferenceRequest(ctx context.Context, r *http.Request)(interface{}, error) {
	var req UpdateCreateConferenceRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil{
		return nil, err
	}
	return req, nil
	}

func decodeDeleteCreateConferenceRequest(ctx context.Context, r *http.Request)(interface{}, error) {
	var req DeleteCreateConferenceRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil{
		return nil, err
	}
	return req, nil
	}

func decodeCreateLocationRequest(ctx context.Context, r *http.Request)(interface{}, error) {
	var req CreateLocationRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil{
		return nil, err
	}
	return req, nil
	}

func decodeGetLocationRequest(ctx context.Context, r *http.Request)(interface{}, error) {
	var req GetLocationRequest
	vars := mux.Vars(r)
	req = GetLocationRequest{Id:vars["id"]}
	return req, nil
	}

func decodeUpdateCreateLocationRequest(ctx context.Context, r *http.Request)(interface{}, error) {
	var req UpdateCreateLocationRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil{
		return nil, err
	}
	return req, nil
	}

func decodeDeleteCreateLocationRequest(ctx context.Context, r *http.Request)(interface{}, error) {
	var req DeleteCreateLocationRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil{
		return nil, err
	}
	return req, nil
	}

func decodeCreateRatingRequest(ctx context.Context, r *http.Request)(interface{}, error) {
	var req CreateRatingRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil{
		return nil, err
	}
	return req, nil
	}

func decodeGetRatingRequest(ctx context.Context, r *http.Request)(interface{}, error) {
	var req GetRatingRequest
	vars := mux.Vars(r)
	req = GetRatingRequest{Id:vars["id"]}
	return req, nil
	}

func decodeUpdateCreateRatingRequest(ctx context.Context, r *http.Request)(interface{}, error) {
	var req UpdateCreateRatingRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil{
		return nil, err
	}
	return req, nil
	}

func decodeDeleteCreateRatingRequest(ctx context.Context, r *http.Request)(interface{}, error) {
	var req DeleteCreateRatingRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil{
		return nil, err
	}
	return req, nil
	}

func decodeCreateReportRequest(ctx context.Context, r *http.Request)(interface{}, error) {
	var req CreateReportRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil{
		return nil, err
	}
	return req, nil
	}

func decodeGetReportRequest(ctx context.Context, r *http.Request)(interface{}, error) {
	var req GetReportRequest
	vars := mux.Vars(r)
	req = GetReportRequest{Id:vars["id"]}
	return req, nil
	}

func decodeUpdateCreateReportRequest(ctx context.Context, r *http.Request)(interface{}, error) {
	var req UpdateCreateReportRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil{
		return nil, err
	}
	return req, nil
	}

func decodeDeleteCreateReportRequest(ctx context.Context, r *http.Request)(interface{}, error) {
	var req DeleteCreateReportRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil{
		return nil, err
	}
	return req, nil
	}

func decodeCreateSpeakerRequest(ctx context.Context, r *http.Request)(interface{}, error) {
	var req CreateSpeakerRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil{
		return nil, err
	}
	return req, nil
	}

func decodeGetSpeakerRequest(ctx context.Context, r *http.Request)(interface{}, error) {
	var req GetSpeakerRequest
	vars := mux.Vars(r)
	req = GetSpeakerRequest{Id:vars["id"]}
	return req, nil
	}

func decodeUpdateCreateSpeakerRequest(ctx context.Context, r *http.Request)(interface{}, error) {
	var req UpdateCreateSpeakerRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil{
		return nil, err
	}
	return req, nil
	}

func decodeDeleteCreateSpeakerRequest(ctx context.Context, r *http.Request)(interface{}, error) {
	var req DeleteCreateSpeakerRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil{
		return nil, err
	}
	return req, nil
	}

