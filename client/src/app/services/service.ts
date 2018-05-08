import {Injectable} from '@angular/core';
import {HttpClient} from '@angular/common/http';

@Injectable()
export class Service {

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
    return this.http.patch(`/api/update-user/${id}`, Service.mapUserFields(user));
  }

  newUser(user: object) {
    return this.http.post('/api/new-user', user);
  }

  createPost(post: object) {
    return this.http.post('/api/new-post', post);
  }
}
