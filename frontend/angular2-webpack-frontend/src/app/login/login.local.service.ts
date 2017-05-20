import { Injectable, OnInit } from '@angular/core';
import { Subject }    from 'rxjs/Subject';

@Injectable()
export class LocalStorageService {

  private missionAnnouncedSource = new Subject<string>();
  private logoutAnnoucedSource = new Subject<string>();

  loginAnnounced$ = this.missionAnnouncedSource.asObservable();
  logoutAnnounced$ = this.logoutAnnoucedSource.asObservable();

  announceLogin(mission: string) {
    this.missionAnnouncedSource.next(mission);
  }    

  announceLogout(){
    this.logoutAnnoucedSource.next(null);
  }
}