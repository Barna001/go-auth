import { Component, OnInit } from '@angular/core';
import { AuthService } from '../service/authentication/auth.service';
import { NgForm } from '@angular/forms';
import { Router } from '@angular/router';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent {

  constructor(private authService: AuthService, private router: Router) {}

  onSubmit(form: NgForm) {
    this.authService.login({
      email: form.value.email,
      password: form.value.password,
    }).then((jwtToken: string) => {
      this.authService.setToken(jwtToken);
      this.router.navigate(['/dashboard']);
    });
  }

}