import { Component, Input } from '@angular/core';

import { Company } from './company';
@Component({
  selector: 'company-detail',
  template: `
    <div *ngIf="company">
      <div>
        <label><h5>what {{company.name}} has to say about Hexerent: </h5></label>
        <h6>{{company.statement}}</h6>
      </div>
    </div>
  `
})
export class CompanyDetailComponent {
  @Input() public company: Company;
}