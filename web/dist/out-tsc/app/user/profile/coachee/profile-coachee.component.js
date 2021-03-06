var __decorate = (this && this.__decorate) || function (decorators, target, key, desc) {
    var c = arguments.length, r = c < 3 ? target : desc === null ? desc = Object.getOwnPropertyDescriptor(target, key) : desc, d;
    if (typeof Reflect === "object" && typeof Reflect.decorate === "function") r = Reflect.decorate(decorators, target, key, desc);
    else for (var i = decorators.length - 1; i >= 0; i--) if (d = decorators[i]) r = (c < 3 ? d(r) : c > 3 ? d(target, key, r) : d(target, key)) || r;
    return c > 3 && r && Object.defineProperty(target, key, r), r;
};
var __metadata = (this && this.__metadata) || function (k, v) {
    if (typeof Reflect === "object" && typeof Reflect.metadata === "function") return Reflect.metadata(k, v);
};
import { ChangeDetectorRef, Component } from "@angular/core";
import { Observable } from "rxjs";
import { Coachee } from "../../../model/Coachee";
import { AuthService } from "../../../service/auth.service";
import { FormBuilder, Validators } from "@angular/forms";
import { ActivatedRoute } from "@angular/router";
import { CoachCoacheeService } from "../../../service/coach_coachee.service";
import { Headers } from "@angular/http";
var ProfileCoacheeComponent = (function () {
    function ProfileCoacheeComponent(authService, cd, formBuilder, coachService, route) {
        this.authService = authService;
        this.cd = cd;
        this.formBuilder = formBuilder;
        this.coachService = coachService;
        this.route = route;
        this.isOwner = false;
        this.updateUserLoading = false;
        this.loading = true;
    }
    ProfileCoacheeComponent.prototype.ngOnInit = function () {
        console.log("ngOnInit");
        window.scrollTo(0, 0);
        this.loading = true;
        this.formCoachee = this.formBuilder.group({
            firstName: ['', Validators.required],
            lastName: ['', Validators.required],
            avatar: ['', Validators.required]
        });
        this.getCoacheeAndUser();
    };
    ProfileCoacheeComponent.prototype.ngOnDestroy = function () {
        console.log("ngOnDestroy");
        if (this.subscriptionGetCoachee) {
            console.log("Unsubscribe coach");
            this.subscriptionGetCoachee.unsubscribe();
        }
        if (this.subscriptionGetRoute) {
            console.log("Unsubscribe subscriptionGetRoute");
            this.subscriptionGetRoute.unsubscribe();
        }
    };
    ProfileCoacheeComponent.prototype.getCoacheeAndUser = function () {
        var _this = this;
        this.subscriptionGetRoute = this.route.params.subscribe(function (params) {
            var coacheeId = params['id'];
            _this.subscriptionGetCoachee = _this.coachService.getCoacheeForId(coacheeId).subscribe(function (coachee) {
                console.log("gotCoachee", coachee);
                _this.setFormValues(coachee);
                _this.coachee = Observable.of(coachee);
                console.log("getUser");
                var user = _this.authService.getConnectedUser();
                _this.isOwner = (user instanceof Coachee) && (coachee.email === user.email);
                _this.cd.detectChanges();
                _this.loading = false;
            });
        });
    };
    ProfileCoacheeComponent.prototype.setFormValues = function (coachee) {
        this.formCoachee.setValue({
            firstName: coachee.first_name,
            lastName: coachee.last_name,
            avatar: coachee.avatar_url
        });
    };
    ProfileCoacheeComponent.prototype.submitCoacheeProfilUpdate = function () {
        var _this = this;
        console.log("submitCoacheeProfilUpdate");
        this.updateUserLoading = true;
        this.coachee.last().flatMap(function (coachee) {
            console.log("submitCoacheeProfilUpdate, coachee obtained");
            return _this.authService.updateCoacheeForId(coachee.id, _this.formCoachee.value.firstName, _this.formCoachee.value.lastName, _this.formCoachee.value.avatar);
        }).flatMap(function (coachee) {
            console.log('Upload user success', coachee);
            if (_this.avatarUrl != null && _this.avatarUrl !== undefined) {
                console.log("Upload avatar");
                var params = [coachee.id];
                var formData = new FormData();
                formData.append('uploadFile', _this.avatarUrl, _this.avatarUrl.name);
                var headers = new Headers();
                headers.append('Accept', 'application/json');
                return _this.authService.put(AuthService.PUT_COACHEE_PROFILE_PICT, params, formData, { headers: headers })
                    .map(function (res) { return res.json(); })
                    .catch(function (error) { return Observable.throw(error); });
            }
            else {
                return Observable.of(coachee);
            }
        }).subscribe(function (coachee) {
            console.log('Upload avatar success', coachee);
            _this.updateUserLoading = false;
            Materialize.toast('Votre profil a été modifié !', 3000, 'rounded');
            //refresh page
            setTimeout('', 1000);
            // window.location.reload();
        }, function (error) {
            console.log('Upload avatar error', error);
            _this.updateUserLoading = false;
            Materialize.toast('Impossible de modifier votre profil', 3000, 'rounded');
        });
    };
    ProfileCoacheeComponent.prototype.filePreview = function (event) {
        if (event.target.files && event.target.files[0]) {
            this.avatarUrl = event.target.files[0];
            console.log("filePreview", this.avatarUrl);
            var reader = new FileReader();
            reader.onload = function (e) {
                // $('#avatar-preview').attr('src', e.target.result);
                $('#avatar-preview').css('background-image', 'url(' + e.target.result + ')');
            };
            reader.readAsDataURL(event.target.files[0]);
        }
    };
    ProfileCoacheeComponent = __decorate([
        Component({
            selector: 'er-profile-coachee',
            templateUrl: 'profile-coachee.component.html',
            styleUrls: ['./profile-coachee.component.scss']
        }),
        __metadata("design:paramtypes", [AuthService, ChangeDetectorRef, FormBuilder, CoachCoacheeService, ActivatedRoute])
    ], ProfileCoacheeComponent);
    return ProfileCoacheeComponent;
}());
export { ProfileCoacheeComponent };
//# sourceMappingURL=/Users/guillaume/angular/eritis_fe/src/app/user/profile/coachee/profile-coachee.component.js.map