/*
 * Angular 2 decorators and services
 */
import {
  Component,
  OnInit,
  ViewEncapsulation
} from '@angular/core';
import { AppState } from './app.service';

/*
 * App Component
 * Top Level Component
 */
@Component({
  selector: 'app',
  encapsulation: ViewEncapsulation.None,
  styleUrls: [
    './app.component.css'
  ],
  template: `
    <div class="nav-fixed">
      <nav class="nav purple">
        <div class="nav-wrapper">
          <a [routerLink]=" ['./'] "
            routerLinkActive="active" [routerLinkActiveOptions]= "{exact: true}" class="brand-logo white-text">HEXERENT</a>
          <a href="#" data-activates="mobile-demo" class="button-collapse"><i class="material-icons">menu</i></a>
          <ul id="nav-mobile" class="right hide-on-med-and-down">

            <li><a class="white-text" href="/download">Get the app</a></li>

            <li><a class="white-text" href="/about">About</a></li>
            <li><a class="white-text" href="/login">Login</a></li>
            <li><a class="white-text" href="/register">Register</a></li>
          </ul>
          <ul class="side-nav purple" id="mobile-demo">
            <li><a class="white-text" href="/download">Get the app</a></li>


            <li><a class="white-text" href="/about">About</a></li>
            <li><a class="white-text" href="/login">Login</a></li>
            <li><a class="white-text" href="/register">Register</a></li>
          </ul>
        </div>
      </nav>
    </div>
    

    <main>
      <router-outlet></router-outlet>
    </main>

    <footer class="page-footer white">
    <div class="footer-copyright white">
      <div class="container white"> 
        <div class="row">
          <div class="col l12 s12">
              <div class="container"> 
                <a class="black-text text-darken-5 right"
                  href="/about">About</a>
                <a class="text-black text-darken-5 right"
                  href="#!">More Links</a>
                <a class="text-black text-darken-5 right"
                  href="#!">More Links</a>  
                <h6 class="black-text">© 2014 Copyright Text</h6>
              </div>  
          </div>
        </div>
      </div>
    </div>
    </footer>  

  `
})
export class AppComponent implements OnInit {
  public angularclassLogo = 'assets/img/angularclass-avatar.png';
  public name = 'Angular 2 Webpack Starter';
  public url = 'https://twitter.com/AngularClass';

  constructor(
    public appState: AppState
  ) {}

  public ngOnInit() {
    console.log('Initial App State', this.appState.state);
  }

}

/*
 * Please review the https://github.com/AngularClass/angular2-examples/ repo for
 * more angular app examples that you may copy/paste
 * (The examples may not be updated as quickly. Please open an issue on github for us to update it)
 * For help or questions please contact us at @AngularClass on twitter
 * or our chat on Slack at https://AngularClass.com/slack-join
 */
