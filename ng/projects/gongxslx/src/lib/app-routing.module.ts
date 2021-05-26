import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

// insertion point for imports
import { XLFilesTableComponent } from './xlfiles-table/xlfiles-table.component'
import { XLFileDetailComponent } from './xlfile-detail/xlfile-detail.component'
import { XLFilePresentationComponent } from './xlfile-presentation/xlfile-presentation.component'


const routes: Routes = [ // insertion point for routes declarations
	{ path: 'xlfiles', component: XLFilesTableComponent, outlet: 'table' },
	{ path: 'xlfile-adder', component: XLFileDetailComponent, outlet: 'editor' },
	{ path: 'xlfile-adder/:id/:association', component: XLFileDetailComponent, outlet: 'editor' },
	{ path: 'xlfile-detail/:id', component: XLFileDetailComponent, outlet: 'editor' },
	{ path: 'xlfile-presentation/:id', component: XLFilePresentationComponent, outlet: 'presentation' },
	{ path: 'xlfile-presentation-special/:id', component: XLFilePresentationComponent, outlet: 'xlfilepres' },

];

@NgModule({
	imports: [RouterModule.forRoot(routes)],
	exports: [RouterModule]
})
export class AppRoutingModule { }
