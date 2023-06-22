import { Component, OnInit } from '@angular/core';

import { Observable, combineLatest, timer } from 'rxjs'

import * as gongdoc from 'gongdoc'
import * as gongxlsx from 'gongxlsx'

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
})
export class AppComponent implements OnInit {

  xl_display = 'xl_display'
  default = 'Gongxlsx Data/Model'
  view = this.xl_display

  views: string[] = [this.xl_display, this.default];

  DataStack = "gongxlsx"
  ModelStacks = "github.com/fullstack-lang/gongxlsx/go/models"

  constructor(
  ) {

  }

  ngOnInit(): void {
  }
}
