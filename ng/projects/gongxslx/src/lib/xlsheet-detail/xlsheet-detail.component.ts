// generated from NgDetailTemplateTS
import { Component, OnInit } from '@angular/core';
import { FormControl } from '@angular/forms';

import { XLSheetDB } from '../xlsheet-db'
import { XLSheetService } from '../xlsheet.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'
import { MapOfComponents } from '../map-components'
import { MapOfSortingComponents } from '../map-components'

// insertion point for imports

import { Router, RouterState, ActivatedRoute } from '@angular/router';

import { MatDialog, MAT_DIALOG_DATA, MatDialogRef, MatDialogConfig } from '@angular/material/dialog';

import { NullInt64 } from '../front-repo.service'

@Component({
	selector: 'app-xlsheet-detail',
	templateUrl: './xlsheet-detail.component.html',
	styleUrls: ['./xlsheet-detail.component.css'],
})
export class XLSheetDetailComponent implements OnInit {

	// insertion point for declarations

	// the XLSheetDB of interest
	xlsheet: XLSheetDB;

	// front repo
	frontRepo: FrontRepo

	constructor(
		private xlsheetService: XLSheetService,
		private frontRepoService: FrontRepoService,
		public dialog: MatDialog,
		private route: ActivatedRoute,
		private router: Router,
	) {
		// https://stackoverflow.com/questions/54627478/angular-7-routing-to-same-component-but-different-param-not-working
		// this is for routerLink on same component when only queryParameter changes
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getXLSheet()

		// observable for changes in structs
		this.xlsheetService.XLSheetServiceChanged.subscribe(
			message => {
				if (message == "post" || message == "update" || message == "delete") {
					this.getXLSheet()
				}
			}
		)

		// insertion point for initialisation of enums list
	}

	getXLSheet(): void {
		const id = +this.route.snapshot.paramMap.get('id');
		const association = this.route.snapshot.paramMap.get('association');

		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo
				console.log("front repo XLSheetPull returned")

				if (id != 0 && association == undefined) {
					this.xlsheet = frontRepo.XLSheets.get(id)
				} else {
					this.xlsheet = new (XLSheetDB)
				}

				// insertion point for recovery of form controls value for bool fields
			}
		)


	}

	save(): void {
		const id = +this.route.snapshot.paramMap.get('id');
		const association = this.route.snapshot.paramMap.get('association');

		// some fields needs to be translated into serializable forms
		// pointers fields, after the translation, are nulled in order to perform serialization
		
		// insertion point for translation/nullation of each field
		
		// save from the front pointer space to the non pointer space for serialization
		if (association == undefined) {
			// insertion point for translation/nullation of each pointers
			if (this.xlsheet.XLFile_Sheets_reverse != undefined) {
				this.xlsheet.XLFile_SheetsDBID = new NullInt64
				this.xlsheet.XLFile_SheetsDBID.Int64 = this.xlsheet.XLFile_Sheets_reverse.ID
				this.xlsheet.XLFile_SheetsDBID.Valid = true
				this.xlsheet.XLFile_SheetsDBID_Index.Valid = true
				this.xlsheet.XLFile_Sheets_reverse = undefined // very important, otherwise, circular JSON
			}
		}

		if (id != 0 && association == undefined) {

			this.xlsheetService.updateXLSheet(this.xlsheet)
				.subscribe(xlsheet => {
					this.xlsheetService.XLSheetServiceChanged.next("update")

					console.log("xlsheet saved")
				});
		} else {
			switch (association) {
				// insertion point for saving value of ONE_MANY association reverse pointer
				case "XLFile_Sheets":
					this.xlsheet.XLFile_SheetsDBID = new NullInt64
					this.xlsheet.XLFile_SheetsDBID.Int64 = id
					this.xlsheet.XLFile_SheetsDBID.Valid = true
					this.xlsheet.XLFile_SheetsDBID_Index.Valid = true
					break
			}
			this.xlsheetService.postXLSheet(this.xlsheet).subscribe(xlsheet => {

				this.xlsheetService.XLSheetServiceChanged.next("post")

				this.xlsheet = {} // reset fields
				console.log("xlsheet added")
			});
		}
	}

	// openReverseSelection is a generic function that calls dialog for the edition of 
	// ONE-MANY association
	// It uses the MapOfComponent provided by the front repo
	openReverseSelection(AssociatedStruct: string, reverseField: string) {

		const dialogConfig = new MatDialogConfig();

		// dialogConfig.disableClose = true;
		dialogConfig.autoFocus = true;
		dialogConfig.width = "50%"
		dialogConfig.height = "50%"
		dialogConfig.data = {
			ID: this.xlsheet.ID,
			ReversePointer: reverseField,
			OrderingMode: false,
		};
		const dialogRef: MatDialogRef<string, any> = this.dialog.open(
			MapOfComponents.get(AssociatedStruct).get(
				AssociatedStruct + 'sTableComponent'
			),
			dialogConfig
		);

		dialogRef.afterClosed().subscribe(result => {
			console.log('The dialog was closed');
		});
	}

	openDragAndDropOrdering(AssociatedStruct: string, reverseField: string) {

		const dialogConfig = new MatDialogConfig();

		// dialogConfig.disableClose = true;
		dialogConfig.autoFocus = true;
		dialogConfig.data = {
			ID: this.xlsheet.ID,
			ReversePointer: reverseField,
			OrderingMode: true,
		};
		const dialogRef: MatDialogRef<string, any> = this.dialog.open(
			MapOfSortingComponents.get(AssociatedStruct).get(
				AssociatedStruct + 'SortingComponent'
			),
			dialogConfig
		);

		dialogRef.afterClosed().subscribe(result => {
			console.log('The dialog was closed');
		});
	}
}