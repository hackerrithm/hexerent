import {
  Component,
  OnInit
} from '@angular/core';

import { HomeService } from './home.service';
import { Title } from './title';
import { XLargeDirective } from './x-large';
import { Post } from "./post";

@Component({
  // The selector is what angular internally uses
  // for `document.querySelectorAll(selector)` in our index.html
  // where, in this case, selector is the string 'home'
  selector: 'home',  // <home></home>
  // We need to tell Angular's Dependency Injection which providers are in our app.
  providers: [
    Title, HomeService
  ],
  // Our list of styles in our component. We may add more to compose many styles together
  styleUrls: [ './home.component.css' ],
  // Every Angular template is first compiled by the browser before Angular runs it's compiler
  templateUrl: './home.component.html'
})
export class HomeComponent implements OnInit {
  // Set our default values
  public localState = { value: '' };
  // TypeScript public modifiers
  constructor(
    public homeService: HomeService,
    public title: Title
  ) {}

  posts: Post[];
  errorMessage: string;

  getPosts() {
        this.homeService.getPosts()
            .subscribe(
                posts => this.posts = posts,
                error => this.errorMessage = <any>error
            )
    }


   createPost(post: Post) {
        if (!post.body) return;
        this.homeService.addPost(post)
            .subscribe(
                newPost => this.posts = [newPost, ...this.posts],
                error => this.errorMessage = <any>error
            )
            
    }

  public ngOnInit() {
    console.log('Hello Kemz `Home` component');
    console.log("Kemar is cool");
    
    // this.title.getData().subscribe(data => this.data = data);
    this.getPosts();
  }

  public submitState(value: string) {
    console.log('submitState', value);
    this.homeService.set('value', value);
    this.localState.value = '';
  }
}
