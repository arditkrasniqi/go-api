import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import {RouterModule, Routes} from '@angular/router';
import {HttpClientModule} from '@angular/common/http';


import { AppComponent } from './app.component';
import { UserComponent } from './user/user.component';
import { PostComponent } from './post/post.component';
import {UserService} from './services/user-service';

const routes: Routes = [
  {
    path: 'user',
    component: UserComponent
  },
  {
    path: 'post',
    component: PostComponent
  }
];

@NgModule({
  declarations: [
    AppComponent,
    UserComponent,
    PostComponent
  ],
  imports: [
    BrowserModule,
    HttpClientModule,
    RouterModule.forRoot(routes)
  ],
  providers: [
    UserService
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
