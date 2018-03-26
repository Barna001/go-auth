import { Component, OnInit } from '@angular/core';
import { Http, RequestOptions } from '@angular/http';
import { User } from '../models/user.model';
import { environment } from '../../environments/environment';
import { Router } from '@angular/router';

@Component({
  selector: 'app-user-detail',
  templateUrl: './user-detail.component.html',
  styleUrls: ['./user-detail.component.css']
})
export class UserDetailComponent implements OnInit {

  user: User;

  constructor(private http: Http, private router: Router) {
    this.user = {
      name: '',
      email: '',
      password: '',
    };
  }

  ngOnInit() {
    this.loadUser();
  }

  loadUser() {
    this.http.get(
      `${environment.apiUrl}/user`,
      {params: { email: this.router.url.replace('/user/', '')}}
    ).subscribe(res => {
      this.user = res.json();
    });
  }

}
