import { 
    Component,
    OnInit
} from '@angular/core';
import { LoginService } from './login.service';
import { LoginPost } from "./login";

@Component ({
    selector: 'login',
    templateUrl: './login.component.html',
    providers: [LoginService]
})


export class LoginComponent implements OnInit{
    // Set our default values
    public localState = { value: '' };
    model: any = {};
    loading = false;
    
    constructor(
        public loginService: LoginService,
    ) { }

    posts: LoginPost[];
    errorMessage: string;

    getPosts() {
            this.loginService.getPosts()
                .subscribe(
                    posts => this.posts = posts,
                    error => this.errorMessage = <any>error
                )
        }


    createPost(post: LoginPost) {
            if (!post.username) return;
            if (!post.password) return;

            this.loginService.addPost(post)
                .subscribe(
                    newPost => this.posts = [newPost, ...this.posts],
                    error => this.errorMessage = <any>error
                )
                
        }

    public ngOnInit() {
        console.log('Hello `Login` component');
        
        // this.title.getData().subscribe(data => this.data = data);
        this.getPosts();
    }

    public submitState(value: string) {
        console.log('submitState', value);
        this.loginService.set('value', value);
        this.localState.value = '';
    }

}