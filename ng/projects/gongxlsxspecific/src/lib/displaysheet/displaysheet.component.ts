import { Component, OnInit } from '@angular/core';

import * as gongxlsx from 'gongxlsx'

import { MatTableDataSource } from '@angular/material/table';
import { MatSort } from '@angular/material/sort';

type Constructor = new () => Object;

const json2Instance = (source: string, destinationConstructor: Constructor) =>
  Object.assign(new destinationConstructor(), JSON.parse(source));

@Component({
  selector: 'lib-displaysheet',
  templateUrl: './displaysheet.component.html',
  styleUrls: ['./displaysheet.component.css']
})
export class DisplaysheetComponent implements OnInit {

  columns = [] as any
  displayedColumns: string[] = []

  sort: MatSort | undefined
  matTableDataSource: MatTableDataSource<string> = new (MatTableDataSource)

  public gongxlsxFrontRepo?: gongxlsx.FrontRepo

  constructor(
    private gongxlsxFrontRepoService: gongxlsx.FrontRepoService,
  ) { }

  ngOnInit(): void {
    this.gongxlsxFrontRepoService.pull().subscribe(
      gongxlsxsFrontRepo => {
        this.gongxlsxFrontRepo = gongxlsxsFrontRepo

        if (this.gongxlsxFrontRepo.XLFiles.size == 0) {
          console.error("cannot deal with 0 file")
        }
        let gongXLFile = this.gongxlsxFrontRepo.XLFiles_array[0]

        if (this.gongxlsxFrontRepo.DisplaySelections.size == 0) {
          console.error("cannot deal with 0 display selections")
        }
        // default file to display is rank 0
        let displayselection = this.gongxlsxFrontRepo.DisplaySelections_array[0]

        // is selection is activated
        if (displayselection.XLFile != null) {
          gongXLFile = displayselection.XLFile
        }

        if (gongXLFile.Sheets!.length == 0) {
          console.error("cannot deal with 0 sheets")
        }
        // default display choice
        let gongXLSheet = gongXLFile.Sheets![0]

        // if the selected sheet is among the sheets of the xl file, select it to display
        for (let sheetId = 0; sheetId < gongXLFile.Sheets!.length; sheetId++) {
          let gongXLSheet_ = gongXLFile.Sheets![sheetId]
          if (displayselection.XLSheet != null) {

            if (displayselection.XLSheet.ID == gongXLSheet_.ID) {
              gongXLSheet = gongXLSheet_
            }
          }
        }

        // cells are provided in random order. on need to order them in the correct order
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
}
