// generated from NgDetailTemplateTS
import { Component, OnInit } from '@angular/core';
import { FormControl } from '@angular/forms';

import { XLRowDB } from '../xlrow-db'
import { XLRowService } from '../xlrow.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'
import { MapOfComponents } from '../map-components'
import { MapOfSortingComponents } from '../map-components'

// insertion point for imports

import { Router, RouterState, ActivatedRoute } from '@angular/router';

import { MatDialog, MAT_DIALOG_DATA, MatDialogRef, MatDialogConfig } from '@angular/material/dialog';

import { NullInt64 } from '../front-repo.service'

@Component({
	selector: 'app-xlrow-detail',
	templateUrl: './xlrow-detail.component.html',
	styleUrls: ['./xlrow-detail.component.css'],
})
export class XLRowDetailComponent implements OnInit {

	// insertion point for declarations

	// the XLRowDB of interest
	xlrow: XLRowDB;

	// front repo
	frontRepo: FrontRepo

	constructor(
		private xlrowService: XLRowService,
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
		this.getXLRow()

		// observable for changes in structs
		this.xlrowService.XLRowServiceChanged.subscribe(
			message => {
				if (message == "post" || message == "update" || message == "delete") {
					this.getXLRow()
				}
			}
		)

		// insertion point for initialisation of enums list
	}

	getXLRow(): void {
		const id = +this.route.snapshot.paramMap.get('id');
		const association = this.route.snapshot.paramMap.get('association');

		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo
				console.log("front repo XLRowPull returned")

				if (id != 0 && association == undefined) {
					this.xlrow = frontRepo.XLRows.get(id)
				} else {
					this.xlrow = new (XLRowDB)
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
			if (this.xlrow.XLSheet_Rows_reverse != undefined) {
				this.xlrow.XLSheet_RowsDBID = new NullInt64
				this.xlrow.XLSheet_RowsDBID.Int64 = this.xlrow.XLSheet_Rows_reverse.ID
				this.xlrow.XLSheet_RowsDBID.Valid = true
				this.xlrow.XLSheet_RowsDBID_Index.Valid = true
				this.xlrow.XLSheet_Rows_reverse = undefined // very important, otherwise, circular JSON
			}
		}

		if (id != 0 && association == undefined) {

			this.xlrowService.updateXLRow(this.xlrow)
				.subscribe(xlrow => {
					this.xlrowService.XLRowServiceChanged.next("update")

					console.log("xlrow saved")
				});
		} else {
			switch (association) {
				// insertion point for saving value of ONE_MANY association reverse pointer
				case "XLSheet_Rows":
					this.xlrow.XLSheet_RowsDBID = new NullInt64
					this.xlrow.XLSheet_RowsDBID.Int64 = id
					this.xlrow.XLSheet_RowsDBID.Valid = true
					this.xlrow.XLSheet_RowsDBID_Index.Valid = true
					break
			}
			this.xlrowService.postXLRow(this.xlrow).subscribe(xlrow => {

				this.xlrowService.XLRowServiceChanged.next("post")

				this.xlrow = {} // reset fields
				console.log("xlrow added")
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
			ID: this.xlrow.ID,
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
			ID: this.xlrow.ID,
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
