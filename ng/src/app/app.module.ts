import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { AppComponent } from './app.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';

import { GongxlsxModule } from 'gongxlsx'
import { GongxlsxspecificModule } from 'gongxlsxspecific'

// mandatory
import { HttpClientModule } from '@angular/common/http';
import { AngularSplitModule } from 'angular-split';

@NgModule({
  declarations: [
    AppComponent
  ],
  imports: [
    BrowserModule,
    BrowserAnimationsModule,

    AngularSplitModule,

    HttpClientModule,
    GongxlsxModule,
    GongxlsxspecificModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
