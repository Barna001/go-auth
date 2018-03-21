import { Component, OnInit } from '@angular/core';
import { AuthService } from '../service/authentication/auth.service';
import { NgForm, FormBuilder, FormGroup, AbstractControl, FormControl, Validators } from '@angular/forms';
import { Router } from '@angular/router';

@Component({
  selector: 'app-register',
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.css']
})
export class RegisterComponent implements OnInit {

  formGroup: FormGroup;

  constructor(private authService: AuthService, private router: Router, private fb: FormBuilder) {

  }

  ngOnInit() {
    this.formGroup = this.fb.group({
      name: [''],
      email: ['', [Validators.required, Validators.email]],
      password: ['', [Validators.required]],
      password2: ['', [Validators.required]],
    }, {
      validator: this.matchPass('password', 'password2')
    });
  }

  matchPass(firstControlName, secondControlName) {
    return (AC: AbstractControl) => {
      const firstControlValue = AC.get(firstControlName).value;
      const secondControlValue = AC.get(secondControlName).value;
      if (firstControlValue !== secondControlValue) {
        AC.get(secondControlName).setErrors({ match: true });
      } else {
        return null;
      }
    };
  }

  get password2(): AbstractControl {
    return this.formGroup.get('password2');
  }

  get email(): AbstractControl {
    return this.formGroup.get('email');
  }

  onSubmit() {
    this.authService.register({
      email: this.email.value,
      password: this.password2.value,
      name: this.formGroup.get('name').value,
    }).then(() => {
      this.router.navigate(['/login']);
    });
  }

}
