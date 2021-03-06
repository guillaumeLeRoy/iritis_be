/**
 * Created by guillaume on 01/02/2017.
 */
var Coach = (function () {
    function Coach(id) {
        this.id = id;
    }
    Coach.parseCoach = function (json) {
        console.log("parseCoach, json : ", json);
        var coach = new Coach(json.id);
        coach.email = json.email;
        coach.first_name = json.first_name;
        coach.last_name = json.last_name;
        coach.avatar_url = json.avatar_url;
        coach.start_date = json.start_date;
        coach.score = json.score;
        coach.sessionsCount = json.sessions_count;
        coach.description = json.description;
        coach.chat_room_url = json.chat_room_url;
        coach.linkedin_url = json.linkedin_url;
        coach.training = json.training;
        coach.degree = json.degree;
        coach.extraActivities = json.extraActivities;
        coach.coachForYears = json.coachForYears;
        coach.coachingExperience = json.coachingExperience;
        coach.coachingHours = json.coachingHours;
        coach.supervisor = json.supervisor;
        coach.favoriteCoachingSituation = json.favoriteCoachingSituation;
        coach.status = json.status;
        coach.revenues = json.revenue;
        coach.insurance_url = json.insurance_url;
        coach.invoice_address = json.invoice_address;
        coach.invoice_city = json.invoice_city;
        coach.invoice_entity = json.invoice_entity;
        coach.invoice_postcode = json.invoice_postcode;
        coach.languages = json.languages;
        coach.experienceShortSession = json.experienceShortSession;
        coach.coachingSpecifics = json.coachingSpecifics;
        coach.therapyElements = json.therapyElements;
        return coach;
    };
    return Coach;
}());
export { Coach };
//# sourceMappingURL=/Users/guillaume/angular/eritis_fe/src/app/model/Coach.js.map