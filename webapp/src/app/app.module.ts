import { BrowserModule } from '@angular/platform-browser';
import { NgModule, ErrorHandler } from '@angular/core';
import { RouterModule, Routes, Router } from '@angular/router';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { routes } from './app.routes';


import { AppComponent } from './app.component';
import { UserListComponent } from './user-list/user-list.component';
import { RequestOptions, Http, HttpModule } from '@angular/http';
import { AuthRequestOptions } from './service/authentication/auth-request';
import { AuthErrorHandler } from './service/authentication/auth-error-handler';
import { LoginComponent } from './login/login.component';
import { AuthService } from './service/authentication/auth.service';
import { AuthGuard } from './service/authentication/auth.guard';
import { UserDetailComponent } from './user-detail/user-detail.component';
import { RegisterComponent } from './register/register.component';

@NgModule({
  declarations: [
    AppComponent,
    UserListComponent,
    LoginComponent,
    UserDetailComponent,
    RegisterComponent
  ],
  imports: [
    BrowserModule,
    FormsModule,
    ReactiveFormsModule,
    HttpModule,
    RouterModule.forRoot(routes),
  ],
  providers: [
    AuthGuard,
    AuthService,
    {
      provide: RequestOptions,
      useClass: AuthRequestOptions
    },
    {
      provide: ErrorHandler,
      useClass: AuthErrorHandler
    }
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
