
import { Injectable } from '@angular/core';
import { Http, Response, Headers } from '@angular/http';
import * as jwt_decode from 'jwt-decode';
import { environment } from '../../../environments/environment';
import { User } from '../../models/user.model';

export const TOKEN_NAME = 'jwt_token';

@Injectable()
export class AuthService {

  constructor(private http: Http) { }

  getToken(): string {
    return localStorage.getItem(TOKEN_NAME);
  }

  setToken(token: string): void {
    localStorage.setItem(TOKEN_NAME, token);
  }

  getTokenExpirationDate(token: string): Date {
    const decoded = jwt_decode(token);

    if (decoded.exp === undefined) {
      return null;
    }

    const date = new Date(0);
    date.setUTCSeconds(decoded.exp);
    return date;
  }

  isTokenExpired(token?: string): boolean {
    if (!token) {
      token = this.getToken();
    }
    if (!token) {
      return true;
    }

    const date = this.getTokenExpirationDate(token);
    if (date === undefined) {
      return false;
    }
    return !(date.valueOf() > new Date().valueOf());
  }

  login(user: User): Promise<string> {
    return this.http
      .post(`${environment.apiUrl}/login`, JSON.stringify(user))
      .toPromise()
      .then(res => res.text());
  }

  logout() {
    localStorage.removeItem(TOKEN_NAME);
  }

  register(user: User): Promise<string> {
    return this.http
      .post(`${environment.apiUrl}/user`, JSON.stringify(user))
      .toPromise()
      .then(res => res.text());
  }

}
