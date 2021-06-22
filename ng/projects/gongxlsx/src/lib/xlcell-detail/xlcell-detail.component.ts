// generated from NgDetailTemplateTS
import { Component, OnInit } from '@angular/core';
import { FormControl } from '@angular/forms';

import { XLCellDB } from '../xlcell-db'
import { XLCellService } from '../xlcell.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'
import { MapOfComponents } from '../map-components'
import { MapOfSortingComponents } from '../map-components'

// insertion point for imports

import { Router, RouterState, ActivatedRoute } from '@angular/router';

import { MatDialog, MAT_DIALOG_DATA, MatDialogRef, MatDialogConfig } from '@angular/material/dialog';

import { NullInt64 } from '../front-repo.service'

@Component({
	selector: 'app-xlcell-detail',
	templateUrl: './xlcell-detail.component.html',
	styleUrls: ['./xlcell-detail.component.css'],
})
export class XLCellDetailComponent implements OnInit {

	// insertion point for declarations

	// the XLCellDB of interest
	xlcell: XLCellDB;

	// front repo
	frontRepo: FrontRepo

	constructor(
		private xlcellService: XLCellService,
		private frontRepoService: FrontRepoService,
		public dialog: MatDialog,
		private route: ActivatedRoute,
		private router: Router,
	) {
	}

	ngOnInit(): void {
		this.getXLCell()

		// observable for changes in structs
		this.xlcellService.XLCellServiceChanged.subscribe(
			message => {
				if (message == "post" || message == "update" || message == "delete") {
					this.getXLCell()
				}
			}
		)

		// insertion point for initialisation of enums list
	}

	getXLCell(): void {
		const id = +this.route.snapshot.paramMap.get('id');
		const association = this.route.snapshot.paramMap.get('association');

		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo
				if (id != 0 && association == undefined) {
					this.xlcell = frontRepo.XLCells.get(id)
				} else {
					this.xlcell = new (XLCellDB)
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
			if (this.xlcell.XLRow_Cells_reverse != undefined) {
				if (this.xlcell.XLRow_CellsDBID == undefined) {
					this.xlcell.XLRow_CellsDBID = new NullInt64
				}
				this.xlcell.XLRow_CellsDBID.Int64 = this.xlcell.XLRow_Cells_reverse.ID
				this.xlcell.XLRow_CellsDBID.Valid = true
				if (this.xlcell.XLRow_CellsDBID_Index == undefined) {
					this.xlcell.XLRow_CellsDBID_Index = new NullInt64
				}
				this.xlcell.XLRow_CellsDBID_Index.Valid = true
				this.xlcell.XLRow_Cells_reverse = undefined // very important, otherwise, circular JSON
			}
			if (this.xlcell.XLSheet_SheetCells_reverse != undefined) {
				if (this.xlcell.XLSheet_SheetCellsDBID == undefined) {
					this.xlcell.XLSheet_SheetCellsDBID = new NullInt64
				}
				this.xlcell.XLSheet_SheetCellsDBID.Int64 = this.xlcell.XLSheet_SheetCells_reverse.ID
				this.xlcell.XLSheet_SheetCellsDBID.Valid = true
				if (this.xlcell.XLSheet_SheetCellsDBID_Index == undefined) {
					this.xlcell.XLSheet_SheetCellsDBID_Index = new NullInt64
				}
				this.xlcell.XLSheet_SheetCellsDBID_Index.Valid = true
				this.xlcell.XLSheet_SheetCells_reverse = undefined // very important, otherwise, circular JSON
			}
		}

		if (id != 0 && association == undefined) {

			this.xlcellService.updateXLCell(this.xlcell)
				.subscribe(xlcell => {
					this.xlcellService.XLCellServiceChanged.next("update")
				});
		} else {
			switch (association) {
				// insertion point for saving value of ONE_MANY association reverse pointer
				case "XLRow_Cells":
					this.xlcell.XLRow_CellsDBID = new NullInt64
					this.xlcell.XLRow_CellsDBID.Int64 = id
					this.xlcell.XLRow_CellsDBID.Valid = true
					this.xlcell.XLRow_CellsDBID_Index = new NullInt64
					this.xlcell.XLRow_CellsDBID_Index.Valid = true
					break
				case "XLSheet_SheetCells":
					this.xlcell.XLSheet_SheetCellsDBID = new NullInt64
					this.xlcell.XLSheet_SheetCellsDBID.Int64 = id
					this.xlcell.XLSheet_SheetCellsDBID.Valid = true
					this.xlcell.XLSheet_SheetCellsDBID_Index = new NullInt64
					this.xlcell.XLSheet_SheetCellsDBID_Index.Valid = true
					break
			}
			this.xlcellService.postXLCell(this.xlcell).subscribe(xlcell => {

				this.xlcellService.XLCellServiceChanged.next("post")

				this.xlcell = {} // reset fields
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
			ID: this.xlcell.ID,
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
		});
	}

	openDragAndDropOrdering(AssociatedStruct: string, reverseField: string) {

		const dialogConfig = new MatDialogConfig();

		// dialogConfig.disableClose = true;
		dialogConfig.autoFocus = true;
		dialogConfig.data = {
			ID: this.xlcell.ID,
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
		});
	}

	fillUpNameIfEmpty(event) {
		if (this.xlcell.Name == undefined) {
			this.xlcell.Name = event.value.Name		
		}
	}
}
