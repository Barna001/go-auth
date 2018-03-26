import { Component, OnInit } from '@angular/core';
import { AuthService } from '../service/authentication/auth.service';
import { NgForm } from '@angular/forms';
import { Router } from '@angular/router';
import { RequestOptions } from '@angular/http';
import { AuthRequestOptions } from '../service/authentication/auth-request';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent {

  constructor(private authService: AuthService, private router: Router, private requestOptions: RequestOptions) {}

  onSubmit(form: NgForm) {
    this.authService.login({
      email: form.value.email,
      password: form.value.password,
      name: '',
    }).then((jwtToken: string) => {
      this.authService.setToken(jwtToken);
      (this.requestOptions as AuthRequestOptions).refreshToken();
      this.router.navigate(['/dashboard']);
    });
  }

}
