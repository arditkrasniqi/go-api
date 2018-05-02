import { Injectable } from '@angular/core';
import {HttpClient, HttpHeaders} from '@angular/common/http';

@Injectable()
export class UserService {

  constructor(private http: HttpClient) {}

  getUsers() {
    return this.http.get('/api/users');
  }

  deleteUser(id: number) {
    return this.http.delete(`/api/delete-user/${id}`);
  }
}