var __decorate = (this && this.__decorate) || function (decorators, target, key, desc) {
    var c = arguments.length, r = c < 3 ? target : desc === null ? desc = Object.getOwnPropertyDescriptor(target, key) : desc, d;
    if (typeof Reflect === "object" && typeof Reflect.decorate === "function") r = Reflect.decorate(decorators, target, key, desc);
    else for (var i = decorators.length - 1; i >= 0; i--) if (d = decorators[i]) r = (c < 3 ? d(r) : c > 3 ? d(target, key, r) : d(target, key)) || r;
    return c > 3 && r && Object.defineProperty(target, key, r), r;
};
var __metadata = (this && this.__metadata) || function (k, v) {
    if (typeof Reflect === "object" && typeof Reflect.metadata === "function") return Reflect.metadata(k, v);
};
import { Component } from "@angular/core";
import { FormBuilder, Validators } from "@angular/forms";
import { AuthService } from "../../../service/auth.service";
import { Router } from "@angular/router";
import { Observable } from "rxjs";
import { AdminAPIService } from "../../../service/adminAPI.service";
var SignUpType;
(function (SignUpType) {
    SignUpType[SignUpType["COACH"] = 0] = "COACH";
    SignUpType[SignUpType["COACHEE"] = 1] = "COACHEE";
    SignUpType[SignUpType["RH"] = 2] = "RH";
    SignUpType[SignUpType["NULL"] = 3] = "NULL";
})(SignUpType || (SignUpType = {}));
var SignupAdminComponent = (function () {
    /* ----- END Contract Plan ----*/
    function SignupAdminComponent(formBuilder, authService, adminAPIService, router) {
        this.formBuilder = formBuilder;
        this.authService = authService;
        this.adminAPIService = adminAPIService;
        this.router = router;
        this.signUpSelectedType = SignUpType.NULL;
        this.error = false;
        this.errorMessage = "";
        console.log("constructor");
    }
    SignupAdminComponent.prototype.ngOnInit = function () {
        console.log("ngOnInit");
        // this.signUpTypes = [SignUpType.COACH, SignUpType.COACHEE, SignUpType.RH];
        this.signUpTypes = [SignUpType.COACH, SignUpType.RH];
        this.signUpForm = this.formBuilder.group({
            email: ['', Validators.compose([
                    Validators.required,
                    this.isEmail
                ])]
        });
        this.getListOfContractPlans();
    };
    SignupAdminComponent.prototype.onSelectPlan = function (plan) {
        console.log("onSelectPlan, plan ", plan);
        this.mSelectedPlan = plan;
    };
    SignupAdminComponent.prototype.onSignUpSubmitted = function () {
        console.log("onSignUp");
        //reset errors
        this.error = false;
        this.errorMessage = '';
        if (this.signUpSelectedType == SignUpType.COACH) {
            console.log("onSignUp, coach");
            this.createPotentialCoach(this.signUpForm.value.email);
        }
        else if (this.signUpSelectedType == SignUpType.RH) {
            this.createPotentialRh(this.signUpForm.value.email);
        }
        else {
            Materialize.toast('Vous devez sélectionner un type', 3000, 'rounded');
        }
    };
    SignupAdminComponent.prototype.createPotentialRh = function (email) {
        console.log('createPotentialRh');
        var body = {
            "email": email,
        };
        this.adminAPIService.createPotentialRh(body).subscribe(function (res) {
            console.log('createPotentialRh, res', res);
            Materialize.toast('Collaborateur RH ajouté !', 3000, 'rounded');
        }, function (error) {
            console.log('createPotentialRh, error', error);
            Materialize.toast("Impossible d'ajouter le RH", 3000, 'rounded');
        });
    };
    SignupAdminComponent.prototype.createPotentialCoach = function (email) {
        console.log('createPotentialCoach');
        var body = {
            "email": email,
        };
        this.adminAPIService.createPotentialCoach(body).subscribe(function (res) {
            console.log('createPotentialCoach, res', res);
            Materialize.toast('Collaborateur Coach ajouté !', 3000, 'rounded');
        }, function (error) {
            console.log('createPotentialCoach, error', error);
            Materialize.toast("Impossible d'ajouter le Coach", 3000, 'rounded');
        });
    };
    SignupAdminComponent.prototype.getListOfContractPlans = function () {
        var _this = this;
        this.authService.getNotAuth(AuthService.GET_CONTRACT_PLANS, null).subscribe(function (response) {
            var json = response.json();
            console.log("getListOfContractPlans, response json : ", json);
            _this.plans = Observable.of(json);
            // this.cd.detectChanges();
        });
    };
    SignupAdminComponent.prototype.isEmail = function (control) {
        if (!control.value.match("[a-z0-9!#$%&'*+/=?^_`{|}~-]+(?:\.[a-z0-9!#$%&'*+/=?^_`{|}~-]+)*@(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?")) {
            console.log("email NOT ok");
            // this.test = false
            return { noEmail: true };
        }
        // this.test = true
        console.log("email ok");
    };
    SignupAdminComponent.prototype.isEqualPassword = function (control) {
        if (!this.signUpForm) {
            return { passwordNoMatch: true };
        }
        if (control.value !== this.signUpForm.controls["password"].value) {
            console.log("isEqualPassword, NO");
            return { passwordNoMatch: true };
        }
    };
    SignupAdminComponent.prototype.getSignUpTypeName = function (type) {
        switch (type) {
            case SignUpType.COACH:
                return "Coach";
            case SignUpType.COACHEE:
                return "Coaché";
            case SignUpType.RH:
                return "RH";
        }
    };
    return SignupAdminComponent;
}());
SignupAdminComponent = __decorate([
    Component({
        selector: 'rb-signup',
        templateUrl: './signup-admin.component.html',
        styleUrls: ['./signup-admin.component.scss']
    }),
    __metadata("design:paramtypes", [FormBuilder, AuthService, AdminAPIService, Router])
], SignupAdminComponent);
export { SignupAdminComponent };
//# sourceMappingURL=/Users/guillaume/angular/eritis_fe/src/app/login/signup/signup-admin/signup-admin.component.js.map