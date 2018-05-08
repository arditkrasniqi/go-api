import {BrowserModule} from '@angular/platform-browser';
import {NgModule} from '@angular/core';
import {RouterModule, Routes} from '@angular/router';
import {HttpClientModule} from '@angular/common/http';
import {FormsModule} from '@angular/forms';


import {AppComponent} from './app.component';
import {UserComponent} from './user/user.component';
import {PostComponent} from './post/post.component';
import {Service} from './services/service';

const routes: Routes = [
  {
    path: '',
    component: PostComponent
  },
  {
    path: 'user',
    component: UserComponent
  },
  {
    path: 'posts',
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
    FormsModule,
    RouterModule.forRoot(routes)
  ],
  providers: [
    Service
  ],
  bootstrap: [AppComponent]
})
export class AppModule {
}
