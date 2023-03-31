import { NgModule } from '@angular/core';

import { GongxlsxspecificComponent } from './gongxlsxspecific.component';
import { DataModelPanelComponent } from './data-model-panel/data-model-panel.component';
import { DisplaysheetComponent } from './displaysheet/displaysheet.component';

import { GongdocModule } from 'gongdoc'
import { GongdocdiagramsModule } from 'gongdocdiagrams'

import { GongxlsxModule } from 'gongxlsx'

import { MatRadioModule } from '@angular/material/radio';
import { FormsModule } from '@angular/forms';
import { CommonModule } from '@angular/common';
import { AngularSplitModule } from 'angular-split';
import { MatTableModule } from '@angular/material/table';
import { MatTabsModule } from '@angular/material/tabs';

@NgModule({
  declarations: [
    GongxlsxspecificComponent,
    DataModelPanelComponent,
    DisplaysheetComponent
  ],
  imports: [
    GongdocModule,
    GongdocdiagramsModule,

    MatTableModule,
    MatTabsModule,
    MatRadioModule,
    FormsModule,
    CommonModule,

    AngularSplitModule,

    GongxlsxModule,
  ],
  exports: [
    GongxlsxspecificComponent,
    DataModelPanelComponent,
    DisplaysheetComponent
  ]
})
export class GongxlsxspecificModule { }
