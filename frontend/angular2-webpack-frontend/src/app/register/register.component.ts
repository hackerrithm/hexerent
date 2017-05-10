import { 
    Component,
    OnInit
} from '@angular/core';
import { RegisterService } from './register.service';
import { RegisterPost } from "./register";

@Component ({
    selector: 'register',
    templateUrl: './register.component.html',
    providers: [RegisterService]
})


export class RegisterComponent implements OnInit{
    // Set our default values
    public localState = { value: '' };
    
    constructor(
        public registerService: RegisterService,
    ) { }


    posts: RegisterPost[];
    errorMessage: string;

    getPosts() {
            this.registerService.getPosts()
                .subscribe(
                    posts => this.posts = posts,
                    error => this.errorMessage = <any>error
                )
        }


    createPost(post: RegisterPost) {
            if (!post.firstname) return;
            if (!post.lastname) return;
            if (!post.username) return;
            if (!post.password) return;            
            if (!post.email) return;
            if (!post.subscribe) return;

            this.registerService.addPost(post)
                .subscribe(
                    newPost => this.posts = [newPost, ...this.posts],
                    error => this.errorMessage = <any>error
                )
                
        }

    public ngOnInit() {
        console.log('Hello `Register` component');
        
        // this.title.getData().subscribe(data => this.data = data);
        this.getPosts();
    }

    public submitState(value: string) {
        console.log('submitState', value);
        this.registerService.set('value', value);
        this.localState.value = '';
    }

}