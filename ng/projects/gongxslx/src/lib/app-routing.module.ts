import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

// insertion point for imports
import { XslxsTableComponent } from './xslxs-table/xslxs-table.component'
import { XslxDetailComponent } from './xslx-detail/xslx-detail.component'
import { XslxPresentationComponent } from './xslx-presentation/xslx-presentation.component'


const routes: Routes = [ // insertion point for routes declarations
	{ path: 'xslxs', component: XslxsTableComponent, outlet: 'table' },
	{ path: 'xslx-adder', component: XslxDetailComponent, outlet: 'editor' },
	{ path: 'xslx-adder/:id/:association', component: XslxDetailComponent, outlet: 'editor' },
	{ path: 'xslx-detail/:id', component: XslxDetailComponent, outlet: 'editor' },
	{ path: 'xslx-presentation/:id', component: XslxPresentationComponent, outlet: 'presentation' },
	{ path: 'xslx-presentation-special/:id', component: XslxPresentationComponent, outlet: 'xslxpres' },

];

@NgModule({
	imports: [RouterModule.forRoot(routes)],
	exports: [RouterModule]
})
export class AppRoutingModule { }
