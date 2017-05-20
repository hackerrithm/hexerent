import { Routes } from '@angular/router';
import { IndexComponent } from './home';
import { AboutComponent } from './about';
import { LoginComponent } from './login';
import { RegisterComponent } from './register'
import { NoContentComponent } from './no-content';

import { DataResolver } from './app.resolver';

export const ROUTES: Routes = [
  { path: '',      component: IndexComponent },
  { path: 'about', component: AboutComponent },
  { path: 'login', component: LoginComponent },
  { path: 'register', component: RegisterComponent },
  { path: 'detail', loadChildren: './+detail#DetailModule'},
  { path: 'barrel', loadChildren: './+barrel#BarrelModule'},
  { path: '**',    component: NoContentComponent },
];
