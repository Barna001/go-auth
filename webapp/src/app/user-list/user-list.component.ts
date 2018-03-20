import { Component, OnInit } from '@angular/core';
import { Http } from '@angular/http';
import { environment } from '../../environments/environment';
import { User } from '../models/user.model';

@Component({
  selector: 'app-user-list',
  templateUrl: './user-list.component.html',
  styleUrls: ['./user-list.component.css']
})
export class UserListComponent implements OnInit {

  users: User[];

  constructor(private http: Http) {}

  ngOnInit() {
    this.loadUsers();
  }

  loadUsers() {
    this.http.get(`${environment.apiUrl}/user`).subscribe(res => this.users = res.json());
  }
}
