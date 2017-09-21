package handler

import (
	"net/http"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"eritis_be/pkg/response"
	"google.golang.org/appengine/user"
	"errors"
	"strings"
)

func HandleAdmin(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	log.Debugf(ctx, "handle admin")

	switch r.Method {
	case "GET":

		if strings.Contains(r.URL.Path, "user") {
			handleGetConnectedAdminUser(w, r) // GET /api/v1/admins/user
		} else if strings.Contains(r.URL.Path, "possible_coachs") {

			// detect an id
			params := response.PathParams(ctx, r, "/api/v1/admins/possible_coachs/:id")
			userId, ok := params[":id"]
			if ok {
				getPossibleCoach(w, r, userId) // GET /api/v1/admins/possible_coachs/:id
				return
			}
			// return all possible coachs
			getAllPossibleCoachs(w, r) // GET /api/v1/admins/possible_coachs
		} else if strings.Contains(r.URL.Path, "admins/coachs") {

			params := response.PathParams(ctx, r, "/api/v1/admins/coachs/:id")
			userId, ok := params[":id"]
			if ok {
				handleGetCoachForId(w, r, userId) // GET /api/v1/admins/coachs/ID
				return
			}

			handleAdminGetCoachs(w, r) // GET /api/v1/admins/coachs
		} else if strings.Contains(r.URL.Path, "admins/coachees") {

			/**
		 		GET a specific coachee
		 	*/
			params := response.PathParams(ctx, r, "/api/v1/admins/coachees/:uid")
			userId, ok := params[":uid"]
			if ok {
				handleGetCoacheeForId(w, r, userId) // GET /api/v1/admins/coachees/:uid
				return
			}

			// get ALL coachees
			handleAdminGetCoachees(w, r) // GET /api/v1/admins/coachees
		} else if strings.Contains(r.URL.Path, "admins/rhs") {

			/**
			* Get HR for uid
			 */
			params := response.PathParams(ctx, r, "/api/v1/admins/rhs/:id")
			userId, ok := params[":id"]
			if ok {
				handleGetHrForId(w, r, userId) // GET /api/v1/admins/rhs/ID
				return
			}

			handleAdminGetRhs(w, r) // GET /api/v1/admins/rhs
		} else if strings.Contains(r.URL.Path, "admins/meetings/coachees") {

			/**
			* Get meetings for specific Coachee
			 */
			params := response.PathParams(ctx, r, "/api/v1/admins/meetings/coachees/:uid")
			//get uid param
			uid, ok := params[":uid"]
			if ok {
				getAllMeetingsForCoachee(w, r, uid) // GET /api/v1/admins/meetings/coachees/:uid
				return
			}
		}else if strings.Contains(r.URL.Path, "admins/meetings/coachs") {

			/**
			* Get meetings for specific Coachee
			 */
			params := response.PathParams(ctx, r, "/api/v1/admins/meetings/coachs/:uid")
			//get uid param
			uid, ok := params[":uid"]
			if ok {
				getAllMeetingsForCoach(w, r, uid) // GET /api/v1/admins/meetings/coachs/:uid
				return
			}
		}


	case "PUT":

		// upload picture
		contains := strings.Contains(r.URL.Path, "profile_picture")
		if contains {
			params := response.PathParams(ctx, r, "/api/v1/admins/coachs/:uid/profile_picture")
			uid, ok := params[":uid"]
			if ok {
				uploadCoachProfilePicture(w, r, uid)
				return
			}
		}
	default:
		http.NotFound(w, r)
	}
}

func handleGetConnectedAdminUser(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	log.Debugf(ctx, "handleGetConnectedAdminUser")

	u := user.Current(ctx)

	log.Debugf(ctx, "handleGetConnectedAdminUser, %s", u)

	if u != nil && u.Admin {

		var admin struct {
			Email string `json:"email"`
		}

		admin.Email = u.Email

		response.Respond(ctx, w, r, &admin, http.StatusOK)
		return
	}
	response.RespondErr(ctx, w, r, errors.New("No user or not an admin"), http.StatusBadRequest)
}

func handleAdminGetCoachs(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	log.Debugf(ctx, "handleAdminGetCoachs")

	handleGetAllCoachs(w, r)
}

func handleAdminGetCoachees(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	log.Debugf(ctx, "handleAdminGetCoachees")

	handleGetAllCoachees(w, r)
}

func handleAdminGetRhs(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	log.Debugf(ctx, "handleAdminGetRhs")

	handleGetAllRHs(w, r)
}
