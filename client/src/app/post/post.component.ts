import {Component, OnInit} from '@angular/core';
import {Observable} from 'rxjs/Observable';
import {Service} from '../services/service';

@Component({
  selector: 'app-post',
  templateUrl: './post.component.html',
  styleUrls: ['./post.component.css']
})
export class PostComponent implements OnInit {
  showNewUserForm: boolean;
  posts: Observable<any>;
  postTitle: string;
  postContent: string;
  postedBy: number;

  constructor(private service: Service) {
    this.postedBy = 1;
  }

  ngOnInit() {
  }

  newPost() {
    const post = {
      title: this.postTitle,
      content: this.postContent,
      postedBy: this.postedBy
    };

    this.service.createPost(post);
  }

  toggleNewUserForm() {
    this.showNewUserForm = !this.showNewUserForm;
  }
}
