import { Injectable } from '@angular/core';
import { Http, Headers, RequestOptions, Response } from '@angular/http';
import 'rxjs/add/operator/toPromise';
import {LoginPost} from './login';
import { Observable }     from 'rxjs/Observable';
import 'rxjs/add/operator/map';
import 'rxjs/add/operator/catch';

export type InternalStateType = {
  [key: string]: any
};


@Injectable()
export class LoginService {

  private loginURL = 'http://localhost:9000/login';  // URL to web api
    constructor(private http: Http) { }


  public getPosts(): Observable<LoginPost[]> { 
              return this.http
              .get(this.loginURL)
              .map(response => response.json().data as LoginPost[])
              .catch(this.handleError);
  }


  public addPost (post: LoginPost): Observable<LoginPost> {
        let headers = new Headers({'Content-Type': 'application/json'});
        let options = new RequestOptions({headers});

        return this.http.post(this.loginURL, { post }, options)
            .map(this.parseData)
            .catch(this.handleError);
    }

    private parseData(res: Response)  {
        let body = res.json();

        if (body instanceof Array) {
            return body || [];
        }

        else return body.post || {};
    }

    // Prases error based on the format
  private handleError(error: Response | any) {
        let errorMessage: string;

        errorMessage = error.message ? error.message : error.toString();

        // In real world application, call to log error to remote server
        // logError(error);

        return Observable.throw(errorMessage);
  }
    

/*
  private handleError(error: any): Promise<any> {
    console.error('An error occurred', error); // for demo purposes only
    return Promise.reject(error.message || error);
  }  */

  public _state: InternalStateType = { };

  // already return a clone of the current state
  public get state() {
    return this._state = this._clone(this._state);
  }
  // never allow mutation
  public set state(value) {
    throw new Error('do not mutate the `.state` directly');
  }

  public get(prop?: any) {
    // use our state getter for the clone
    const state = this.state;
    return state.hasOwnProperty(prop) ? state[prop] : state;
  }

  public set(prop: string, value: any) {
    // internally mutate our state
    return this._state[prop] = value;
  }

  private _clone(object: InternalStateType) {
    // simple object clone
    return JSON.parse(JSON.stringify( object ));
  }
}
