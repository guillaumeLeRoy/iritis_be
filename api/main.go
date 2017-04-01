package api

import (
	"net/http"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"errors"
	"strings"
	"golang.org/x/net/context"
	"eritis_be/firebase"
)

/* ######## HOW TO SERVE DIFFERENT ENVIRONMENTS #######

##### LIVE ENV ######
serve locally :
dev_appserver.py -A eritis-150320 dispatch.yaml default/app.yaml api/app.yaml web/app.yaml --enable_sendmail

deploy :
goapp deploy -application eritis-150320 -version 1 default/app.yaml api/app.yaml web/app.yaml
appcfg.py -A eritis-150320 update_dispatch .
appcfg.py update_indexes -A eritis-150320 ./default



##### DEV ENV ######
serve locally :
dev_appserver.py -A eritis-be-dev dispatch.yaml default/app.yaml api/app.yaml web/app.yaml --enable_sendmail

deploy :
goapp deploy -application eritis-be-dev -version 1 default/app.yaml api/app.yaml web/app.yaml
appcfg.py -A eritis-be-dev update_dispatch .
appcfg.py update_indexes -A eritis-be-dev ./default



##### GLR ENV ######
serve :
dev_appserver.py -A eritis-be-glr dispatch.yaml default/app.yaml api/app.yaml web/app.yaml --enable_sendmail

deploy :
for now there is no glr backend on app engine

*/
const LIVE_ENV_PROJECT_ID string = "eritis-150320"
const DEV_ENV_PROJECT_ID string = "eritis-be-dev"
const GLR_ENV_PROJECT_ID string = "eritis-be-glr"

func init() {

	http.HandleFunc("/api/login/", authHandler(HandleLogin))

	//meetings
	http.HandleFunc("/api/meeting/", authHandler(HandleMeeting))
	http.HandleFunc("/api/meetings/", authHandler(HandleMeeting))

	//coach
	http.HandleFunc("/api/coachs/", authHandler(HandleCoachs))

	//coachee
	http.HandleFunc("/api/coachees/", authHandler(HandleCoachees))
}

func getFirebaseJsonPath(ctx context.Context) (string, error) {
	appId := appengine.AppID(ctx)
	log.Debugf(ctx, "appId %s", appId)

	pathToJson := ""

	if strings.EqualFold(LIVE_ENV_PROJECT_ID, appId) {
		pathToJson = "firebase_keys/eritis-be-live.json"
	} else if strings.EqualFold(DEV_ENV_PROJECT_ID, appId) {
		pathToJson = "firebase_keys/eritis-be-dev.json"
	} else if strings.EqualFold(GLR_ENV_PROJECT_ID, appId) {
		pathToJson = "firebase_keys/eritis-be-glr.json"
	} else {
		return "", errors.New("AppId doesn't match any environment")
	}

	log.Debugf(ctx, "getFirebaseJsonPath path %s", pathToJson)

	return pathToJson, nil
}

func authHandler(handler func(w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		ctx := appengine.NewContext(r)

		log.Debugf(ctx, "authHandler start")

		//handle preflight in here
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")

		//check token validity

		if (r.Method == "OPTIONS") {
			log.Debugf(ctx, "authHandler, handle OPTIONS")

			w.WriteHeader(http.StatusOK)
		} else {
			log.Debugf(ctx, "authHandler, handle all requests")

			token := r.Header.Get("Authorization")
			log.Debugf(ctx, "authHandler auth token: %s", token)

			// If the token is empty...
			if token == "" {
				// If we get here, the required token is missing
				RespondErr(ctx, w, r, errors.New("invalid token"), http.StatusUnauthorized)
				return
			}

			//read the Bearer param
			if len(token) >= 1 {
				token = strings.TrimPrefix(token, "Bearer ")
			}

			log.Debugf(ctx, "authHandler token: %s", token)

			// If the token is empty...
			if token == "" {
				// If we get here, the required token is missing
				RespondErr(ctx, w, r, errors.New("invalid token"), http.StatusUnauthorized)
				return
			}

			log.Debugf(ctx, "authHandler VERIFY token")

			//init Firebase with the correct .json

			path, err := getFirebaseJsonPath(ctx)
			if err != nil {
				log.Debugf(ctx, "authHandler InitializeApp failed %s", err)
				RespondErr(ctx, w, r, err, http.StatusUnauthorized)
				return
			}

			_, err = firebase.InitializeApp(&firebase.Options{
				ServiceAccountPath: path,
			})

			if err != nil {
				log.Debugf(ctx, "authHandler InitializeApp failed %s", err)
				RespondErr(ctx, w, r, err, http.StatusUnauthorized)
				return
			}

			log.Debugf(ctx, "authHandler InitializeApp ok")

			//verify token
			auth, _ := firebase.GetAuth()
			decodedToken, err := auth.VerifyIDToken(token, ctx)
			if err != nil {
				log.Debugf(ctx, "authHandler VerifyIDToken failed %s", err)
				RespondErr(ctx, w, r, err, http.StatusUnauthorized)
				return
			}

			if err == nil {
				uid, found := decodedToken.UID()
				log.Debugf(ctx, "authHandler decodedToken uid %s, found %s", uid, found)
			}

			uid, found := decodedToken.UID()
			if !found {
				RespondErr(ctx, w, r, errors.New("UID not found"), http.StatusUnauthorized)
			}

			log.Debugf(ctx, "authHandler, UID: %s", uid)

			//auth ok, continue
			handler(w, r)
		}
	}
}

