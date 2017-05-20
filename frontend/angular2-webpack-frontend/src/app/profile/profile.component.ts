import { 
    Component,
    OnInit
} from '@angular/core';
import { ProfileService } from './profile.service';
import { ProfilePost } from './profile';


@Component ({
    selector: 'profile',
    templateUrl: './profile.component.html',
    providers: [ProfileService]
})


export class ProfileComponent implements OnInit{
    // Set our default values
    public localState = { value: '' };
    posts: ProfilePost[]
    
    constructor(
        public profileService: ProfileService
    ) { }

    errorMessage: string;

    getPosts() {
            this.profileService.getPosts()
                .subscribe(
                    posts => this.posts = posts,
                    error => this.errorMessage = <any>error
                )
    }


    createPost(post: ProfilePost) {
            if (!post.body) return;

            this.profileService.addPost(post)
                .subscribe(
                    newPost => this.posts = [newPost, ...this.posts],
                    error => this.errorMessage = <any>error
                )
                
    }

    public ngOnInit() {
        console.log('Hello `Profile` component');
        
        // this.title.getData().subscribe(data => this.data = data);
        this.getPosts();
    }

    public submitState(value: string) {
        console.log('submitState', value);
        this.localState.value = '';
    }

}