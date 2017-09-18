import { Injectable } from '@angular/core';
import { Http, Headers, RequestOptions, Response } from '@angular/http';
import 'rxjs/add/operator/toPromise';
import { AppComponent } from './app.component';
import {Post} from './post';
import { Observable }     from 'rxjs/Observable';
import 'rxjs/add/operator/map';
import 'rxjs/add/operator/catch';

export type InternalStateType = {
  [key: string]: any
};


@Injectable()
export class IndexService {

  private indexURL = 'http://localhost:9000/';  // URL to web api
    constructor(private http: Http) { }


  public getPosts(): Observable<Post[]> { 
              return this.http
              .get(this.indexURL)
              .map(response => response.json().data as Post[])
              .catch(this.handleError);
  }


  public addPost (post: Post): Observable<Post> {
        let headers = new Headers({'Content-Type': 'application/json'});
        let options = new RequestOptions({headers});
    
        // add similar code here
        /*
        headers.append('Access-Control-Allow-Headers', 'Content-Type');
        headers.append('Access-Control-Allow-Methods', 'GET');
        headers.append('Access-Control-Allow-Origin', '*');
        
        might have error: XMLHttpRequest cannot load http://URI. 
        No 'Access-Control-Allow-Origin' header is present on the requested resource. 
        Origin 'http://localhost:8000' is therefore not allowed access.
        
        */

        return this.http.post(this.indexURL, { post }, options)
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
