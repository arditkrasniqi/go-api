import { Component, OnInit } from '@angular/core';
import {UserService} from '../services/user-service';
import {Observable} from 'rxjs/Observable';

@Component({
  selector: 'app-user',
  templateUrl: './user.component.html',
  styleUrls: ['./user.component.css']
})
export class UserComponent implements OnInit {
  users: Observable<any>;
  constructor(private userService: UserService) { }

  ngOnInit() {
    this.users = this.userService.getUsers();
  }

  deleteUser(id: number) {
    this.users = this.userService.deleteUser(id);
  }
}
