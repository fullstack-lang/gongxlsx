import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

// insertion point for imports
import { XLFilesTableComponent } from './xlfiles-table/xlfiles-table.component'
import { XLFileDetailComponent } from './xlfile-detail/xlfile-detail.component'
import { XLFilePresentationComponent } from './xlfile-presentation/xlfile-presentation.component'

import { XLSheetsTableComponent } from './xlsheets-table/xlsheets-table.component'
import { XLSheetDetailComponent } from './xlsheet-detail/xlsheet-detail.component'
import { XLSheetPresentationComponent } from './xlsheet-presentation/xlsheet-presentation.component'


const routes: Routes = [ // insertion point for routes declarations
	{ path: 'xlfiles', component: XLFilesTableComponent, outlet: 'table' },
	{ path: 'xlfile-adder', component: XLFileDetailComponent, outlet: 'editor' },
	{ path: 'xlfile-adder/:id/:association', component: XLFileDetailComponent, outlet: 'editor' },
	{ path: 'xlfile-detail/:id', component: XLFileDetailComponent, outlet: 'editor' },
	{ path: 'xlfile-presentation/:id', component: XLFilePresentationComponent, outlet: 'presentation' },
	{ path: 'xlfile-presentation-special/:id', component: XLFilePresentationComponent, outlet: 'xlfilepres' },

	{ path: 'xlsheets', component: XLSheetsTableComponent, outlet: 'table' },
	{ path: 'xlsheet-adder', component: XLSheetDetailComponent, outlet: 'editor' },
	{ path: 'xlsheet-adder/:id/:association', component: XLSheetDetailComponent, outlet: 'editor' },
	{ path: 'xlsheet-detail/:id', component: XLSheetDetailComponent, outlet: 'editor' },
	{ path: 'xlsheet-presentation/:id', component: XLSheetPresentationComponent, outlet: 'presentation' },
	{ path: 'xlsheet-presentation-special/:id', component: XLSheetPresentationComponent, outlet: 'xlsheetpres' },

];

@NgModule({
	imports: [RouterModule.forRoot(routes)],
	exports: [RouterModule]
})
export class AppRoutingModule { }