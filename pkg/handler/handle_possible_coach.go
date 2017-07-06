package handler

import (
	"net/http"
	"strings"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"eritis_be/pkg/response"
	"eritis_be/pkg/model"
	"eritis_be/pkg/utils"
	"fmt"
	"errors"
	"google.golang.org/appengine/datastore"
	"golang.org/x/net/context"
)

func HandlePossibleCoach(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	log.Debugf(ctx, "handle possible coach")

	switch r.Method {
	case "PUT":

		// upload picture
		if ok := strings.Contains(r.URL.Path, "profile_picture"); ok {
			uploadPossibleCoachProfilePicture(w, r)
			return
		}

		// upload assurance
		if ok := strings.Contains(r.URL.Path, "assurance"); ok {
			uploadPossibleCoachAssurance(w, r)
			return
		}

		//try to detect a coach
		if ok := strings.Contains(r.URL.Path, "possible_coachs"); ok {
			handleCreatePossibleCoach(w, r)
			return
		}

		http.NotFound(w, r)

	default:
		http.NotFound(w, r)
	}
}

func handleCreatePossibleCoach(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	log.Debugf(ctx, "handleCreatePossibleCoach")

	var possibleCoach struct {
		Email                     string `json:"email"`
		FirstName                 string `json:"firstName"`
		LastName                  string `json:"lastName"`
		LinkedinUrl               string `json:"linkedin_url"`
		Description               string `json:"description"`
		Training                  string
		Degree                    string
		ExtraActivities           string //ActivitiesOutOfCoaching
		CoachForYears             string // been a coach xx years
		CoachingExperience        string // coaching experience
		CoachingHours             string // number of coaching hours
		Supervisor                string
		FavoriteCoachingSituation string
		Status                    string
		Revenue                   string //revenues for last 3 years
	}

	err := response.Decode(r, &possibleCoach)
	if err != nil {
		response.RespondErr(ctx, w, r, err, http.StatusBadRequest)
		return
	}

	log.Debugf(ctx, "handleCreatePossibleCoach, create possible coach")

	var possibleCoachToUpdate *model.PossibleCoach
	err = datastore.RunInTransaction(ctx, func(ctx context.Context) error {
		log.Debugf(ctx, "SearchForPossibleCoach, RunInTransaction")

		possibleCoachToUpdate, err = model.FindPossibleCoachByEmail(ctx, possibleCoach.Email)

		if err != nil && err != model.ErrNoPossibleCoach {
			// response.RespondErr(ctx, w, r, err, http.StatusInternalServerError)
			return err
		}

		log.Debugf(ctx, "handleCreatePossibleCoach, RunInTransaction, no error")

		if err == model.ErrNoPossibleCoach {
			log.Debugf(ctx, "handleCreatePossibleCoach, RunInTransaction, no coach found")

			// create new possible coach
			newPossibleCoach, err := model.CreatePossibleCoach(ctx, possibleCoach.Email)
			if err != nil {
				// response.RespondErr(ctx, w, r, err, http.StatusInternalServerError)
				return err
			}
			possibleCoachToUpdate = newPossibleCoach
		} else {
			log.Debugf(ctx, "handleCreatePossibleCoach, RunInTransaction, already have a coach, %s", possibleCoachToUpdate)
		}

		return nil
	}, &datastore.TransactionOptions{XG: true})
	if err != nil {
		response.RespondErr(ctx, w, r, err, http.StatusInternalServerError)
		return
	}

	log.Debugf(ctx, "handleCreatePossibleCoach, update values")

	possibleCoachToUpdate.FirstName = possibleCoach.FirstName
	possibleCoachToUpdate.LastName = possibleCoach.LastName
	possibleCoachToUpdate.Description = possibleCoach.Description
	possibleCoachToUpdate.LinkedinUrl = possibleCoach.LinkedinUrl
	possibleCoachToUpdate.Training = possibleCoach.Training
	possibleCoachToUpdate.Degree = possibleCoach.Degree
	possibleCoachToUpdate.ExtraActivities = possibleCoach.ExtraActivities
	possibleCoachToUpdate.CoachForYears = possibleCoach.CoachForYears
	possibleCoachToUpdate.CoachingExperience = possibleCoach.CoachingExperience
	possibleCoachToUpdate.CoachingHours = possibleCoach.CoachingHours
	possibleCoachToUpdate.Supervisor = possibleCoach.Supervisor
	possibleCoachToUpdate.FavoriteCoachingSituation = possibleCoach.FavoriteCoachingSituation
	possibleCoachToUpdate.Status = possibleCoach.Status
	possibleCoachToUpdate.Revenue = possibleCoach.Revenue

	err = possibleCoachToUpdate.Update(ctx)
	if err != nil {
		response.RespondErr(ctx, w, r, err, http.StatusInternalServerError)
		return
	}

	log.Debugf(ctx, "handleCreatePossibleCoach, DONE")

	response.Respond(ctx, w, r, &possibleCoachToUpdate, http.StatusCreated)
}

func uploadPossibleCoachProfilePicture(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	log.Debugf(ctx, "uploadPossibleCoachProfilePicture")

	// get email
	emailSender := r.FormValue("email")
	if emailSender == "" {
		response.RespondErr(ctx, w, r, errors.New("Empty email"), http.StatusBadRequest)
		return
	}

	var coach *model.PossibleCoach
	err := datastore.RunInTransaction(ctx, func(ctx context.Context) error {
		log.Debugf(ctx, "uploadPossibleCoachProfilePicture, RunInTransaction")

		var erro error
		coach, erro = model.FindPossibleCoachByEmail(ctx, emailSender)
		if erro != nil {
			return erro
		}

		return nil
	}, &datastore.TransactionOptions{XG: true})
	if err != nil {
		response.RespondErr(ctx, w, r, err, http.StatusInternalServerError)
		return
	}

	log.Debugf(ctx, "uploadPossibleCoachProfilePicture, coach ok, %s", coach)

	hash, err := utils.GetEmailHash(ctx, coach.Email)
	if err != nil {
		response.RespondErr(ctx, w, r, err, http.StatusInternalServerError)
		return
	}

	log.Debugf(ctx, "uploadPossibleCoachProfilePicture, email hash %s", hash)

	fileName, err := utils.UploadPictureProfile(r, hash, "possible_profile_pict")
	if err != nil {
		response.RespondErr(ctx, w, r, err, http.StatusInternalServerError)
		return
	}

	// save new picture url
	storage, err := utils.GetStorageUrl(ctx)
	if err != nil {
		response.RespondErr(ctx, w, r, err, http.StatusInternalServerError)
		return
	}

	fmt.Sprintf("%s/%s", storage, fileName)
	avatarUrl := fmt.Sprintf("%s/%s", storage, fileName)
	coach.AvatarURL = avatarUrl
	coach.Update(ctx)

	log.Debugf(ctx, "uploadPossibleCoachProfilePicture, avatar url updated")

	response.Respond(ctx, w, r, nil, http.StatusOK)
}

func uploadPossibleCoachAssurance(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	log.Debugf(ctx, "uploadPossibleCoachAssurance")

	// get email
	emailSender := r.FormValue("email")
	if emailSender == "" {
		response.RespondErr(ctx, w, r, errors.New("Empty email"), http.StatusBadRequest)
		return
	}

	var coach *model.PossibleCoach
	err := datastore.RunInTransaction(ctx, func(ctx context.Context) error {
		log.Debugf(ctx, "uploadPossibleCoachAssurance, RunInTransaction")

		var erro error
		coach, erro = model.FindPossibleCoachByEmail(ctx, emailSender)
		if erro != nil {
			//response.RespondErr(ctx, w, r, err, http.StatusInternalServerError)
			return erro
		}

		return nil
	}, &datastore.TransactionOptions{XG: true})
	if err != nil {
		response.RespondErr(ctx, w, r, err, http.StatusInternalServerError)
		return
	}

	log.Debugf(ctx, "uploadPossibleCoachAssurance, coach ok")

	hash, err := utils.GetEmailHash(ctx, coach.Email)
	if err != nil {
		response.RespondErr(ctx, w, r, err, http.StatusInternalServerError)
		return
	}

	log.Debugf(ctx, "uploadPossibleCoachAssurance, hash %s", hash)

	fileName, err := utils.UploadPossibleCoachAssurance(r, hash, "possible_assurance")
	if err != nil {
		response.RespondErr(ctx, w, r, err, http.StatusInternalServerError)
		return
	}

	// save new assurance url
	storage, err := utils.GetStorageUrl(ctx)
	if err != nil {
		response.RespondErr(ctx, w, r, err, http.StatusInternalServerError)
		return
	}

	assuranceUrl := fmt.Sprintf("%s/%s", storage, fileName)
	coach.AssuranceUrl = assuranceUrl
	coach.Update(ctx)

	log.Debugf(ctx, "uploadPossibleCoachAssurance, url updated")

	response.Respond(ctx, w, r, nil, http.StatusOK)
}
