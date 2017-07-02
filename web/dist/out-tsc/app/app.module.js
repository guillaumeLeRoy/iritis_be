var __decorate = (this && this.__decorate) || function (decorators, target, key, desc) {
    var c = arguments.length, r = c < 3 ? target : desc === null ? desc = Object.getOwnPropertyDescriptor(target, key) : desc, d;
    if (typeof Reflect === "object" && typeof Reflect.decorate === "function") r = Reflect.decorate(decorators, target, key, desc);
    else for (var i = decorators.length - 1; i >= 0; i--) if (d = decorators[i]) r = (c < 3 ? d(r) : c > 3 ? d(target, key, r) : d(target, key)) || r;
    return c > 3 && r && Object.defineProperty(target, key, r), r;
};
import { BrowserModule } from "@angular/platform-browser";
import { NgModule } from "@angular/core";
import { FormsModule, ReactiveFormsModule } from "@angular/forms";
import { HttpModule } from "@angular/http";
import { AppComponent } from "./app.component";
import { HeaderComponent } from "./header/header.component";
import { DataService } from "./service/data.service";
import { LogService } from "./service/log.service";
import { routing } from "./app.routing";
import { SignupAdminComponent } from "./login/signup/signup-admin/signup-admin.component";
import { SigninComponent } from "./login/signin/signin.component";
import { AuthService } from "./service/auth.service";
import { AuthGuard } from "./login/auth.guard";
import { WelcomeComponent } from "./welcome/welcome.component";
import { ChatComponent } from "./chat/chat.component";
import { ChatItemComponent } from "./chat/chat-item.component";
import { CoachItemComponent } from "./user/coach-item/coach-item.component";
import { CoachCoacheeService } from "./service/coach_coachee.service";
import { CalendarModule } from "angular-calendar"; // lib
import { BrowserAnimationsModule } from "@angular/platform-browser/animations";
import { NgbModule } from "@ng-bootstrap/ng-bootstrap";
import { MeetingsService } from "./service/meetings.service";
import { MeetingListComponent } from "./meeting/meeting-list/meeting-list.component";
import { MeetingItemCoacheeComponent } from "./meeting/meeting-list/coachee/meeting-item-coachee/meeting-item-coachee.component";
import { PreMeetingComponent } from "./meeting/meeting-list/coachee/pre-meeting/pre-meeting.component";
import { ProfileCoachComponent } from "./user/profile/coach/profile-coach.component";
import { ProfileCoacheeComponent } from "./user/profile/coachee/profile-coachee.component";
import { MeetingItemCoachComponent } from "./meeting/meeting-list/coach/meeting-item-coach/meeting-item-coach.component";
import { FirebaseService } from "./service/firebase.service";
import { MeetingDateComponent } from "./meeting/meeting-date/meeting-date.component";
import { SliderModule } from "primeng/components/slider/slider";
import { Ng2PageScrollModule } from "ng2-page-scroll";
import { AdminAPIService } from "./service/adminAPI.service";
import { AdminComponent } from "./admin/admin.component";
import { MeetingItemRhComponent } from "./meeting/meeting-list/rh/meeting-item-rh/meeting-item-rh.component";
import { ProfileRhComponent } from "./user/profile/rh/profile-rh.component";
import { SignupCoacheeComponent } from "./login/signup/signup-coachee/signup-coachee.component";
import { SignupCoachComponent } from "./login/signup/signup-coach/signup-coach.component";
import { SignupRhComponent } from "./login/signup/signup-rh/signup-rh.component";
import { AvailableMeetingsComponent } from "./meeting/meeting-list/coach/available-meetings/available-meetings.component";
import { AdminCoachsListComponent } from "./admin/coachs-list/admin-coachs-list.component";
import { CoacheesListComponent } from './admin/coachees-list/coachees-list.component';
import { RhsListComponent } from './admin/rhs-list/rhs-list.component';
import { MeetingListCoachComponent } from './meeting/meeting-list/coach/meeting-list-coach/meeting-list-coach.component';
import { MeetingListCoacheeComponent } from './meeting/meeting-list/coachee/meeting-list-coachee/meeting-list-coachee.component';
import { MeetingListRhComponent } from './meeting/meeting-list/rh/meeting-list-rh/meeting-list-rh.component';
import { CoacheeItemComponent } from './user/coachee-item/coachee-item.component';
import { RhItemComponent } from './user/rh-item/rh-item.component';
import { RegisterCoachComponent } from './login/register/register-coach/register-coach.component';
import { CodeDeontologieComponent } from './login/register/register-coach/code-deontologie/code-deontologie.component';
var AppModule = (function () {
    function AppModule() {
    }
    return AppModule;
}());
AppModule = __decorate([
    NgModule({
        declarations: [
            AppComponent,
            HeaderComponent,
            SignupAdminComponent,
            SigninComponent,
            ProfileCoachComponent,
            ProfileCoacheeComponent,
            ProfileRhComponent,
            WelcomeComponent,
            ChatComponent,
            ChatItemComponent,
            CoachItemComponent,
            MeetingListComponent,
            MeetingItemCoacheeComponent,
            MeetingItemCoachComponent,
            MeetingDateComponent,
            PreMeetingComponent,
            AdminCoachsListComponent,
            AdminComponent,
            MeetingItemRhComponent,
            SignupCoacheeComponent,
            SignupCoachComponent,
            SignupRhComponent,
            AvailableMeetingsComponent,
            CoacheesListComponent,
            RhsListComponent,
            MeetingListCoachComponent,
            MeetingListCoacheeComponent,
            MeetingListRhComponent,
            CoacheeItemComponent,
            RhItemComponent,
            RegisterCoachComponent,
            CodeDeontologieComponent
        ],
        imports: [
            BrowserAnimationsModule,
            BrowserModule,
            FormsModule,
            HttpModule,
            routing,
            ReactiveFormsModule,
            CalendarModule.forRoot(),
            NgbModule.forRoot(),
            SliderModule,
            Ng2PageScrollModule.forRoot()
        ],
        providers: [DataService, LogService, AuthService, AuthGuard, CoachCoacheeService, MeetingsService, FirebaseService, AdminAPIService],
        bootstrap: [AppComponent]
    })
], AppModule);
export { AppModule };
//# sourceMappingURL=/Users/guillaume/angular/eritis_fe/src/app/app.module.js.map