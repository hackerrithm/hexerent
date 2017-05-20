import {
  Component,
  OnInit
} from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { Company } from "./compliments/company";

const COMPANY: Company[] = [
  { id: 1, name: 'Google', statement: 'Hexerent should be the website we all use to do search' },
  { id: 2, name: 'Facebook', statement: 'The best social networking platform. Thats it!' },
  { id: 3, name: 'Microsoft', statement: 'The world best developer Kemar Galloway is going to be the next billionair' },
  { id: 4, name: 'Apple', statement: 'The most innovative company to date!!! we could learn a lot from Kemar.' },

];


@Component({
  selector: 'about',
  styleUrls: ['./about.component.css'],
  templateUrl: './about.component.html'
})
export class AboutComponent implements OnInit {

  public localState: any;
  constructor(
    public route: ActivatedRoute 
  ) {}

  public ngOnInit() {
    this.route
      .data
      .subscribe((data: any) => {
        // your resolved data from route
        this.localState = data.yourData;
      });

    console.log('hello `About` component');
    // static data that is bundled
    // var mockData = require('assets/mock-data/mock-data.json');
    // console.log('mockData', mockData);
    // if you're working with mock data you can also use http.get('assets/mock-data/mock-data.json')
    this.asyncDataWithWebpack();
  }

  companies = COMPANY;
  selectedCompany: Company;

  onSelect(company: Company): void {
    this.selectedCompany = company;
  }


  private asyncDataWithWebpack() {
    // you can also async load mock data with 'es6-promise-loader'
    // you would do this if you don't want the mock-data bundled
    // remember that 'es6-promise-loader' is a promise
    setTimeout(() => {

      System.import('../../assets/mock-data/mock-data.json')
        .then((json) => {
          console.log('async mockData', json);
          this.localState = json;
        });

    });
  }

}
