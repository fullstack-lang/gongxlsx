import { NgModule } from '@angular/core';

import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { CommonModule } from '@angular/common';
import { Routes, RouterModule } from '@angular/router';

// for angular material
import { MatSliderModule } from '@angular/material/slider';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatSelectModule } from '@angular/material/select'
import { MatDatepickerModule } from '@angular/material/datepicker'
import { MatTableModule } from '@angular/material/table'
import { MatSortModule } from '@angular/material/sort'
import { MatPaginatorModule } from '@angular/material/paginator'
import { MatCheckboxModule } from '@angular/material/checkbox';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';
import { MatToolbarModule } from '@angular/material/toolbar'
import { MatListModule } from '@angular/material/list'
import { MatExpansionModule } from '@angular/material/expansion';
import { MatDialogModule, MatDialogRef } from '@angular/material/dialog';
import { MatGridListModule } from '@angular/material/grid-list';
import { MatTreeModule } from '@angular/material/tree';
import { DragDropModule } from '@angular/cdk/drag-drop';

import { AngularSplitModule, SplitComponent } from 'angular-split';

import { AppRoutingModule } from './app-routing.module';

import { SplitterComponent } from './splitter/splitter.component'
import { SidebarComponent } from './sidebar/sidebar.component';

// insertion point for imports 
import { XLFilesTableComponent } from './xlfiles-table/xlfiles-table.component'
import { XLFileSortingComponent } from './xlfile-sorting/xlfile-sorting.component'
import { XLFileDetailComponent } from './xlfile-detail/xlfile-detail.component'
import { XLFilePresentationComponent } from './xlfile-presentation/xlfile-presentation.component'

import { XLSheetsTableComponent } from './xlsheets-table/xlsheets-table.component'
import { XLSheetSortingComponent } from './xlsheet-sorting/xlsheet-sorting.component'
import { XLSheetDetailComponent } from './xlsheet-detail/xlsheet-detail.component'
import { XLSheetPresentationComponent } from './xlsheet-presentation/xlsheet-presentation.component'


@NgModule({
	declarations: [
		// insertion point for declarations 
		XLFilesTableComponent,
		XLFileSortingComponent,
		XLFileDetailComponent,
		XLFilePresentationComponent,

		XLSheetsTableComponent,
		XLSheetSortingComponent,
		XLSheetDetailComponent,
		XLSheetPresentationComponent,


		SplitterComponent,
		SidebarComponent
	],
	imports: [
		FormsModule,
		ReactiveFormsModule,
		CommonModule,
		RouterModule,

		AppRoutingModule,

		MatSliderModule,
		MatSelectModule,
		MatFormFieldModule,
		MatInputModule,
		MatDatepickerModule,
		MatTableModule,
		MatSortModule,
		MatPaginatorModule,
		MatCheckboxModule,
		MatButtonModule,
		MatIconModule,
		MatToolbarModule,
		MatExpansionModule,
		MatListModule,
		MatDialogModule,
		MatGridListModule,
		MatTreeModule,
		DragDropModule,

		AngularSplitModule,
	],
	exports: [
		// insertion point for declarations 
		XLFilesTableComponent,
		XLFileSortingComponent,
		XLFileDetailComponent,
		XLFilePresentationComponent,

		XLSheetsTableComponent,
		XLSheetSortingComponent,
		XLSheetDetailComponent,
		XLSheetPresentationComponent,


		SplitterComponent,
		SidebarComponent,

	],
	providers: [
		{
			provide: MatDialogRef,
			useValue: {}
		},
	],
})
export class GongxslxModule { }