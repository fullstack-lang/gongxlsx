import { Component } from '@angular/core';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {

  view = 'Default view'
  default = 'Default view'
  diagrams = 'Diagrams view'
  meta = 'Meta view'
  xlsx = 'xlsx view'

  views: string[] = [this.default, this.diagrams, this.meta, this.xlsx];
}
