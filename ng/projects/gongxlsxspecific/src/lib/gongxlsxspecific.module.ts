import { NgModule } from '@angular/core';
import { GongxlsxspecificComponent } from './gongxlsxspecific.component';
import { DisplaysheetComponent } from './displaysheet/displaysheet.component';
import {MatTableModule} from '@angular/material/table';




@NgModule({
  declarations: [
    GongxlsxspecificComponent,
    DisplaysheetComponent
  ],
  imports: [
    MatTableModule,
  ],
  exports: [
    GongxlsxspecificComponent,
    DisplaysheetComponent
  ]
})
export class GongxlsxspecificModule { }
