// generated by gong
import { Component, OnInit, Inject, Optional } from '@angular/core';
import { CdkDragDrop, moveItemInArray } from '@angular/cdk/drag-drop';

import { MatDialogRef, MAT_DIALOG_DATA, MatDialog } from '@angular/material/dialog'
import { DialogData } from '../front-repo.service'
import { SelectionModel } from '@angular/cdk/collections';

import { Router, RouterState } from '@angular/router';
import { XLSheetDB } from '../xlsheet-db'
import { XLSheetService } from '../xlsheet.service'

import { FrontRepoService, FrontRepo, NullInt64 } from '../front-repo.service'
@Component({
  selector: 'lib-xlsheet-sorting',
  templateUrl: './xlsheet-sorting.component.html',
  styleUrls: ['./xlsheet-sorting.component.css']
})
export class XLSheetSortingComponent implements OnInit {

  frontRepo: FrontRepo

  // array of XLSheet instances that are in the association
  associatedXLSheets = new Array<XLSheetDB>();

  constructor(
    private xlsheetService: XLSheetService,
    private frontRepoService: FrontRepoService,

    // not null if the component is called as a selection component of xlsheet instances
    public dialogRef: MatDialogRef<XLSheetSortingComponent>,
    @Optional() @Inject(MAT_DIALOG_DATA) public dialogData: DialogData,

    private router: Router,
  ) {
    this.router.routeReuseStrategy.shouldReuseRoute = function () {
      return false;
    };
  }

  ngOnInit(): void {
    this.getXLSheets()
  }

  getXLSheets(): void {
    this.frontRepoService.pull().subscribe(
      frontRepo => {
        this.frontRepo = frontRepo

        let index = 0
        for (let xlsheet of this.frontRepo.XLSheets_array) {
          let ID = this.dialogData.ID
          let revPointerID = xlsheet[this.dialogData.ReversePointer]
          let revPointerID_Index = xlsheet[this.dialogData.ReversePointer+"_Index"]
          if (revPointerID.Int64 == ID) {
            if (revPointerID_Index == undefined) {
              revPointerID_Index = new NullInt64
              revPointerID_Index.Valid = true
              revPointerID_Index.Int64 = index++
            }
            this.associatedXLSheets.push(xlsheet)
          }
        }

        // sort associated xlsheet according to order
        this.associatedXLSheets.sort((t1, t2) => {
          let t1_revPointerID_Index = t1[this.dialogData.ReversePointer+"_Index"]
          let t2_revPointerID_Index = t2[this.dialogData.ReversePointer+"_Index"]
          if (t1_revPointerID_Index && t2_revPointerID_Index) {
            if (t1_revPointerID_Index.Int64 > t2_revPointerID_Index.Int64) {
              return 1;
            }
            if (t1_revPointerID_Index.Int64 < t2_revPointerID_Index.Int64) {
              return -1;
            }  
          }
          return 0;
        });
      }
    )
  }

  drop(event: CdkDragDrop<string[]>) {
    moveItemInArray(this.associatedXLSheets, event.previousIndex, event.currentIndex);

    // set the order of XLSheet instances
    let index = 0
    
    for (let xlsheet of this.associatedXLSheets) {
      let revPointerID_Index = xlsheet[this.dialogData.ReversePointer+"_Index"]
      revPointerID_Index.Valid = true
      revPointerID_Index.Int64 = index++
    }
  }

  save() {

    this.associatedXLSheets.forEach(
      xlsheet => {
        this.xlsheetService.updateXLSheet(xlsheet)
          .subscribe(xlsheet => {
            this.xlsheetService.XLSheetServiceChanged.next("update")
          });
      }
    )

    this.dialogRef.close('Sorting of '+ this.dialogData.ReversePointer+' done');
  }
}