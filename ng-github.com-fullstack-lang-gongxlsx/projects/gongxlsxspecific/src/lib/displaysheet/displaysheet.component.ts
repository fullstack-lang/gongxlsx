import { Component, Input, OnInit } from '@angular/core';

import * as gongxlsx from '../../../../gongxlsx/src/public-api'

import { MatRadioModule } from '@angular/material/radio';
import { FormsModule } from '@angular/forms';
import { CommonModule } from '@angular/common';
import { AngularSplitModule } from 'angular-split';
import { MatTableModule } from '@angular/material/table';
import { MatTabsModule } from '@angular/material/tabs';

import { MatTableDataSource } from '@angular/material/table';
import { MatSort } from '@angular/material/sort';
import { Observable, timer } from 'rxjs';
import { UntypedFormControl } from '@angular/forms';
import { MatTabChangeEvent } from '@angular/material/tabs';

type Constructor = new () => Object;

const json2Instance = (source: string, destinationConstructor: Constructor) =>
  Object.assign(new destinationConstructor(), JSON.parse(source));

@Component({
  selector: 'lib-displaysheet',
  templateUrl: './displaysheet.component.html',
  standalone: true,
  styleUrls: ['./displaysheet.component.css'],
  imports: [
    MatTableModule,
    MatTabsModule,
    MatRadioModule,
    FormsModule,
    CommonModule,
  ]
})
export class DisplaysheetComponent implements OnInit {

  @Input() GONG__StackPath: string = ""

  // tabs to select the xl file
  tabsForFile = new Array<string>();
  selectedFile = new UntypedFormControl(0);
  selectedFileIndex = 0

  // tabs to select the sheet within the xl file
  tabsForSheetsOfFile = new Array<Array<string>>();
  selectedSheet = new UntypedFormControl(0);
  selectedSheetIndex = 0

  columns = [] as any
  displayedColumns: string[] = []

  sort: MatSort | undefined
  matTableDataSource: MatTableDataSource<string> = new (MatTableDataSource)

  public gongxlsxFrontRepo?: gongxlsx.FrontRepo

  // autonmatic refresh of sheet display
  obsTimer: Observable<number> = timer(1000, 500) // due time 1', period 0.5'

  // commitNb stores the number of commit on the backend
  commitFromBackNb: number = 0

  // commitNb stores the number of commit on the frontend
  commitFromFrontNb: number = 0

  constructor(
    private gongxlsxFrontRepoService: gongxlsx.FrontRepoService,
    private commitNbService: gongxlsx.CommitNbFromBackService,
    private pushFromFrontService: gongxlsx.PushFromFrontNbService,
  ) { }

  ngOnInit(): void {
    this.displaySelectedSheet()

    this.commitNbService.getCommitNbFromBack(500, this.GONG__StackPath).subscribe(
      commitFromBackNb => {
        // console.log("commit nb in the back " + commitFromBackNb + " local commit nb " + this.commitFromBackNb)
        if (commitFromBackNb > this.commitFromBackNb) {
          this.displaySelectedSheet()
          this.commitFromBackNb = commitFromBackNb
        }
      }
    )
  }

  displaySelectedSheet() {
    this.gongxlsxFrontRepoService.pull(this.GONG__StackPath).subscribe(
      gongxlsxsFrontRepo => {

        // reset tabs and create one tab per XL file

        this.tabsForFile = new Array<string>()
        for (let fileId = 0; fileId < gongxlsxsFrontRepo.getFrontArray<gongxlsx.XLFile>(gongxlsx.XLFile.GONGSTRUCT_NAME)!.length; fileId++) {
          let xlFile = gongxlsxsFrontRepo.getFrontArray<gongxlsx.XLFile>(gongxlsx.XLFile.GONGSTRUCT_NAME)[fileId]

          this.tabsForFile.push(xlFile.Name);

          // create on array of the sheets of the file
          let tabForSheet = new Array<string>()
          for (let sheetId = 0; sheetId < xlFile.Sheets!.length; sheetId++) {
            let gongXLSheet = xlFile.Sheets![sheetId]
            tabForSheet.push(gongXLSheet.Name)
          }
          this.tabsForSheetsOfFile.push(tabForSheet)
        }

        // reset the data to display
        this.matTableDataSource = new (MatTableDataSource)

        this.gongxlsxFrontRepo = gongxlsxsFrontRepo

        if (this.gongxlsxFrontRepo.getFrontArray<gongxlsx.XLFile>(gongxlsx.XLFile.GONGSTRUCT_NAME).length == 0) {
          console.error("cannot deal with 0 file")
        }

        // file selected by the end user
        let gongXLFile = this.gongxlsxFrontRepo.getFrontArray<gongxlsx.XLFile>(gongxlsx.XLFile.GONGSTRUCT_NAME)[this.selectedFileIndex]

        let displayselection = new gongxlsx.DisplaySelection
        if (this.gongxlsxFrontRepo.getFrontArray<gongxlsx.DisplaySelection>(gongxlsx.DisplaySelection.GONGSTRUCT_NAME).length > 0) {
          // default file to display is rank 0
          displayselection = this.gongxlsxFrontRepo.getFrontArray<gongxlsx.DisplaySelection>(gongxlsx.DisplaySelection.GONGSTRUCT_NAME)[this.selectedFileIndex]
        }


        // is selection is activated
        if (displayselection.XLFile != null) {
          gongXLFile = displayselection.XLFile
        }

        if (gongXLFile.Sheets!.length == 0) {
          console.error("cannot deal with 0 sheets")
        }

        // default display choice
        let gongXLSheet = gongXLFile.Sheets![0]

        // if the selected sheet is among the sheets of the xl file, select it for display
        if (displayselection.XLSheet != null) {
          for (let sheetId = 0; sheetId < gongXLFile.Sheets!.length; sheetId++) {
            let gongXLSheet_ = gongXLFile.Sheets![sheetId]
            if (displayselection.XLSheet.ID == gongXLSheet_.ID) {
              gongXLSheet = gongXLSheet_
            }
          }
        } else {
          gongXLSheet = gongXLFile.Sheets![this.selectedSheetIndex]
        }

        // cells are provided in random order. on need to order them in the correct order
        if (gongXLSheet == undefined) {
          console.log("No selected sheet")
          gongXLSheet = gongXLFile.Sheets![0]
        }

        let contentArray = new Array<string[]>(gongXLSheet.NbRows)

        for (let rowNb = 0; rowNb < gongXLSheet.NbRows; rowNb++) {
          contentArray[rowNb] = new Array<string>(gongXLSheet.MaxCol)
        }

        for (let gongXLCell of gongXLSheet.SheetCells!) {
          contentArray[gongXLCell.Y][gongXLCell.X] = gongXLCell.Name
        }

        // now one need to fill up element data with synthetic object
        for (let rowNb = 1; rowNb < gongXLSheet.NbRows; rowNb++) {
          let oneRow = [] as any
          for (let columnNb = 0; columnNb < gongXLSheet.MaxCol; columnNb++) {
            oneRow[contentArray[0][columnNb]] = contentArray[rowNb][columnNb]
          }
          this.matTableDataSource.data.push(oneRow)
        }

        this.columns = []
        for (let columnNb = 0; columnNb < gongXLSheet.MaxCol; columnNb++) {

          this.columns.push(
            {
              columnDef: contentArray[0][columnNb],
              header: contentArray[0][columnNb],
              cell: (element: any) => `${element[contentArray[0][columnNb]]}`
            }
          )
        }

        // make up displayed columns
        this.displayedColumns = this.columns.map((c: { columnDef: any; }) => c.columnDef);

        console.log(this.displayedColumns)

        this.postProcessing()
      }

    )
  }

  ngAfterViewInit() {
    // enable sorting on all fields (including pointers and reverse pointer)
    this.matTableDataSource.sortingDataAccessor = (cell: any, property: string) => {
      switch (property) {
        case 'ID':
          return cell

        // insertion point for specific sorting accessor
        case 'Name':
          return cell

        default:
          return cell
      };
    }
    this.matTableDataSource.sort = this.sort!
  }

  postProcessing() {
    // enable sorting on all fields (including pointers and reverse pointer)
    this.matTableDataSource.sortingDataAccessor = (cell: any, property: string) => {
      switch (property) {
        case 'ID':
          return cell

        // insertion point for specific sorting accessor
        case 'Name':
          return cell

        default:
          return cell
      };
    }
    this.matTableDataSource.sort = this.sort!
  }

  onTableSelection(event: MatTabChangeEvent) {
    console.log(event)
    this.selectedFileIndex = event.index
    this.displaySelectedSheet()
  }

  onSheetSelection(event: MatTabChangeEvent) {
    console.log(event)
    this.selectedSheetIndex = event.index
    this.displaySelectedSheet()
  }
}