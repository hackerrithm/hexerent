import {
  Component,
  OnInit
} from '@angular/core';

import { IndexService } from './index.service';
import { Title } from './title';
import { XLargeDirective } from './x-large';
import { Post } from './post';
import { ProfilePost } from './../profile/profile';

@Component({
  // The selector is what angular internally uses
  // for `document.querySelectorAll(selector)` in our index.html
  // where, in this case, selector is the string 'home'
  selector: 'index',  // <index></index>
  // We need to tell Angular's Dependency Injection which providers are in our app.
  providers: [
    Title, IndexService
  ],
  // Our list of styles in our component. We may add more to compose many styles together
  styleUrls: [ './index.component.css' ],
  // Every Angular template is first compiled by the browser before Angular runs it's compiler
  templateUrl: './index.component.html'
})
export class IndexComponent implements OnInit {
  // Set our default values
  public localState = { value: '' };

  // TypeScript public modifiers
  constructor(
    public indexService: IndexService,
    public title: Title,
  ) {}

  posts: Post[];
  errorMessage: string;

  getPosts() {
        this.indexService.getPosts()
            .subscribe(
                posts => this.posts = posts,
                error => this.errorMessage = <any>error
            )
    }


   createPost(post: Post) {
        if (!post.body) return;
        this.indexService.addPost(post)
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
    this.indexService.set('value', value);
    this.localState.value = '';
  }
}
