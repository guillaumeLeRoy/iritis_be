package api

import (
	"time"
	"google.golang.org/appengine/datastore"
	"golang.org/x/net/context"
	"google.golang.org/appengine/log"
	"strings"
	"errors"
)

/*
Origin : from a Coach or a Coachee
*/
type MeetingReview struct {
	Key     *datastore.Key `json:"id" datastore:"-"`
	Type    ReviewType `json:"type"`
	Comment string `json:"comment"`
	Date    time.Time `json:"date"`
}

type ReviewType string

const (
	SESSION_CONTEXT ReviewType = "SESSION_CONTEXT"
	SESSION_GOAL ReviewType = "SESSION_GOAL"
	SESSION_VALUE ReviewType = "SESSION_VALUE"
	SESSION_NEXT_STEP ReviewType = "SESSION_NEXT_STEP"
)

func convertToReviewType(reviewType string) (ReviewType, error) {
	if strings.Compare(reviewType, string(SESSION_VALUE)) == 0 {
		return SESSION_VALUE, nil
	} else if strings.Compare(reviewType, string(SESSION_NEXT_STEP)) == 0 {
		return SESSION_NEXT_STEP, nil
	} else if strings.Compare(reviewType, string(SESSION_CONTEXT)) == 0 {
		return SESSION_CONTEXT, nil
	}else if strings.Compare(reviewType, string(SESSION_GOAL)) == 0 {
		return SESSION_GOAL, nil
	}

	return "", errors.New("can't convert reviewType")
}

func createReview(ctx context.Context, meetingKey *datastore.Key, comment string, reviewType ReviewType) (*MeetingReview, error) {
	log.Debugf(ctx, "Create createReview")

	var review MeetingReview
	review.Type = reviewType
	review.Comment = comment
	review.Date = time.Now()
	review.Key = datastore.NewIncompleteKey(ctx, "MeetingReview", meetingKey)

	key, err := datastore.Put(ctx, review.Key, &review)
	if err != nil {
		return nil, err
	}
	review.Key = key

	return &review, nil
}

func (r *MeetingReview)updateReview(ctx context.Context, reviewKey *datastore.Key, comment string) (*MeetingReview, error) {
	log.Debugf(ctx, "Create createReview")

	r.Comment = comment
	r.Date = time.Now()
	key, err := datastore.Put(ctx, r.Key, r)

	if err != nil {
		return nil, err
	}
	r.Key = key
	return r, nil
}

func getAllReviewsForMeeting(ctx context.Context, meetingKey *datastore.Key) ([]*MeetingReview, error) {
	log.Debugf(ctx, "getReviewsForMeeting")

	var reviews []*MeetingReview
	keys, err := datastore.NewQuery("MeetingReview").Ancestor(meetingKey).GetAll(ctx, &reviews)

	for i, review := range reviews {
		review.Key = keys[i]
	}
	if err != nil {
		return nil, err
	}

	log.Debugf(ctx, "getReviewsForMeeting", reviews)

	return reviews, nil
}

func getReviewsForMeetingAndForType(ctx context.Context, meetingKey *datastore.Key, reviewType string) ([]*MeetingReview, error) {
	log.Debugf(ctx, "getReviewsForMeetingAndForType, reviewType %s", reviewType)

	var reviews []*MeetingReview
	keys, err := datastore.NewQuery("MeetingReview").Ancestor(meetingKey).Filter("Type =", reviewType).GetAll(ctx, &reviews)

	for i, review := range reviews {
		review.Key = keys[i]
	}
	if err != nil {
		return nil, err
	}

	log.Debugf(ctx, "getReviewsForMeeting", reviews)

	return reviews, nil
}

// delete review
func deleteMeetingReview(ctx context.Context, reviewKey *datastore.Key) (error) {
	log.Debugf(ctx, "deleteMeetingReview, reviewKey %s", reviewKey)

	err := datastore.Delete(ctx, reviewKey)

	return err

}