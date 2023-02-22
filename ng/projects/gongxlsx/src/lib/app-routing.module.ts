import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

// insertion point for imports
import { DisplaySelectionsTableComponent } from './displayselections-table/displayselections-table.component'
import { DisplaySelectionDetailComponent } from './displayselection-detail/displayselection-detail.component'

import { XLCellsTableComponent } from './xlcells-table/xlcells-table.component'
import { XLCellDetailComponent } from './xlcell-detail/xlcell-detail.component'

import { XLFilesTableComponent } from './xlfiles-table/xlfiles-table.component'
import { XLFileDetailComponent } from './xlfile-detail/xlfile-detail.component'

import { XLRowsTableComponent } from './xlrows-table/xlrows-table.component'
import { XLRowDetailComponent } from './xlrow-detail/xlrow-detail.component'

import { XLSheetsTableComponent } from './xlsheets-table/xlsheets-table.component'
import { XLSheetDetailComponent } from './xlsheet-detail/xlsheet-detail.component'


const routes: Routes = [ // insertion point for routes declarations
	{ path: 'github_com_fullstack_lang_gongxlsx_go-displayselections/:GONG__StackPath', component: DisplaySelectionsTableComponent, outlet: 'github_com_fullstack_lang_gongxlsx_go_table' },
	{ path: 'github_com_fullstack_lang_gongxlsx_go-displayselection-adder/:GONG__StackPath', component: DisplaySelectionDetailComponent, outlet: 'github_com_fullstack_lang_gongxlsx_go_editor' },
	{ path: 'github_com_fullstack_lang_gongxlsx_go-displayselection-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath', component: DisplaySelectionDetailComponent, outlet: 'github_com_fullstack_lang_gongxlsx_go_editor' },
	{ path: 'github_com_fullstack_lang_gongxlsx_go-displayselection-detail/:id/:GONG__StackPath', component: DisplaySelectionDetailComponent, outlet: 'github_com_fullstack_lang_gongxlsx_go_editor' },

	{ path: 'github_com_fullstack_lang_gongxlsx_go-xlcells/:GONG__StackPath', component: XLCellsTableComponent, outlet: 'github_com_fullstack_lang_gongxlsx_go_table' },
	{ path: 'github_com_fullstack_lang_gongxlsx_go-xlcell-adder/:GONG__StackPath', component: XLCellDetailComponent, outlet: 'github_com_fullstack_lang_gongxlsx_go_editor' },
	{ path: 'github_com_fullstack_lang_gongxlsx_go-xlcell-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath', component: XLCellDetailComponent, outlet: 'github_com_fullstack_lang_gongxlsx_go_editor' },
	{ path: 'github_com_fullstack_lang_gongxlsx_go-xlcell-detail/:id/:GONG__StackPath', component: XLCellDetailComponent, outlet: 'github_com_fullstack_lang_gongxlsx_go_editor' },

	{ path: 'github_com_fullstack_lang_gongxlsx_go-xlfiles/:GONG__StackPath', component: XLFilesTableComponent, outlet: 'github_com_fullstack_lang_gongxlsx_go_table' },
	{ path: 'github_com_fullstack_lang_gongxlsx_go-xlfile-adder/:GONG__StackPath', component: XLFileDetailComponent, outlet: 'github_com_fullstack_lang_gongxlsx_go_editor' },
	{ path: 'github_com_fullstack_lang_gongxlsx_go-xlfile-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath', component: XLFileDetailComponent, outlet: 'github_com_fullstack_lang_gongxlsx_go_editor' },
	{ path: 'github_com_fullstack_lang_gongxlsx_go-xlfile-detail/:id/:GONG__StackPath', component: XLFileDetailComponent, outlet: 'github_com_fullstack_lang_gongxlsx_go_editor' },

	{ path: 'github_com_fullstack_lang_gongxlsx_go-xlrows/:GONG__StackPath', component: XLRowsTableComponent, outlet: 'github_com_fullstack_lang_gongxlsx_go_table' },
	{ path: 'github_com_fullstack_lang_gongxlsx_go-xlrow-adder/:GONG__StackPath', component: XLRowDetailComponent, outlet: 'github_com_fullstack_lang_gongxlsx_go_editor' },
	{ path: 'github_com_fullstack_lang_gongxlsx_go-xlrow-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath', component: XLRowDetailComponent, outlet: 'github_com_fullstack_lang_gongxlsx_go_editor' },
	{ path: 'github_com_fullstack_lang_gongxlsx_go-xlrow-detail/:id/:GONG__StackPath', component: XLRowDetailComponent, outlet: 'github_com_fullstack_lang_gongxlsx_go_editor' },

	{ path: 'github_com_fullstack_lang_gongxlsx_go-xlsheets/:GONG__StackPath', component: XLSheetsTableComponent, outlet: 'github_com_fullstack_lang_gongxlsx_go_table' },
	{ path: 'github_com_fullstack_lang_gongxlsx_go-xlsheet-adder/:GONG__StackPath', component: XLSheetDetailComponent, outlet: 'github_com_fullstack_lang_gongxlsx_go_editor' },
	{ path: 'github_com_fullstack_lang_gongxlsx_go-xlsheet-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath', component: XLSheetDetailComponent, outlet: 'github_com_fullstack_lang_gongxlsx_go_editor' },
	{ path: 'github_com_fullstack_lang_gongxlsx_go-xlsheet-detail/:id/:GONG__StackPath', component: XLSheetDetailComponent, outlet: 'github_com_fullstack_lang_gongxlsx_go_editor' },

];

@NgModule({
	imports: [RouterModule.forRoot(routes)],
	exports: [RouterModule]
})
export class AppRoutingModule { }
