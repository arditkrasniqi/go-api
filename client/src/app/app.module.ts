import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import {RouterModule, Routes} from '@angular/router';


import { AppComponent } from './app.component';
import { UserComponent } from './user/user.component';
import { PostComponent } from './post/post.component';

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
    RouterModule.forRoot(routes)
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
