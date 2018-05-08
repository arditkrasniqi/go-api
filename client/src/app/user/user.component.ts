import {Component, OnInit} from '@angular/core';
import {Service} from '../services/service';
import {Observable} from 'rxjs/Observable';

@Component({
  selector: 'app-user',
  templateUrl: './user.component.html',
  styleUrls: ['./user.component.css']
})
export class UserComponent implements OnInit {
  users: Observable<any>;
  username: string;
  password: string;
  userId: number;
  modalData: string;

  constructor(private service: Service) {
  }

  ngOnInit() {
    this.users = this.service.getUsers();
  }

  deleteUser(id: number) {
    this.users = this.service.deleteUser(id);
  }

  updateUser() {
    const user = {
      username: this.username,
      password: this.password
    };
    this.users = this.service.updateUser(this.userId, user);
  }

  activeUser(user: any) {
    this.setModalData('editUser');
    this.userId = user.id;
    this.username = user.username;
    this.password = user.password;
  }

  newUser() {
    const user = {
      username: this.username,
      password: this.password
    };
    this.users = this.service.newUser(user);
  }

  setModalData(data: string) {
    if (data === 'newUser') {
      this.username = this.password = '';
    }
    this.modalData = data;
  }
}
