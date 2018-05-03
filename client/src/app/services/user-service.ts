import {Injectable} from '@angular/core';
import {HttpClient} from '@angular/common/http';

@Injectable()
export class UserService {

  constructor(private http: HttpClient) {
  }

  static mapUserFields(user: any) {
    return {
      id: user.id,
      username: user.username,
      password: user.password
    };
  }

  getUsers() {
    return this.http.get('/api/users');
  }

  deleteUser(id: number) {
    return this.http.delete(`/api/delete-user/${id}`);
  }

  updateUser(id: number, user: object) {
    return this.http.patch(`/api/update-user/${id}`, UserService.mapUserFields(user));
  }
}
