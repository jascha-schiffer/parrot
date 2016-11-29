import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';

import { AuthService } from './../services/auth.service';

@Component({
  selector: 'register',
  templateUrl: './register.component.html'
})
export class RegisterComponent implements OnInit {

  constructor(private auth: AuthService, private router: Router) { }

  ngOnInit() { }

  navigateToLogin() {
    this.router.navigate(['/login']);
  }

  onSubmit(email, password) {
    this.auth.register(email, password).subscribe(
      result => {
        // TODO: Redirect
        this.auth.login(email, password).subscribe(
          result => {
            this.router.navigate(['/projects']);
          },
          error => {
            // TODO: handle error
          });
      },
      error => {
        // TODO: handle error
      });
  }
}
