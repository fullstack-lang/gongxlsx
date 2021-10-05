import {Component, OnInit} from '@angular/core';

import * as gongxlsx from 'gongxlsx'

export interface PeriodicElement {
  name: string;
  position: number;
  weight: number;
  symbol: string;
}

const ELEMENT_DATA: PeriodicElement[] = [
  {position: 1, name: 'Hydrogen', weight: 1.0079, symbol: 'H'},
  {position: 2, name: 'Helium', weight: 4.0026, symbol: 'He'},
  {position: 3, name: 'Lithium', weight: 6.941, symbol: 'Li'},
  {position: 4, name: 'Beryllium', weight: 9.0122, symbol: 'Be'},
  {position: 5, name: 'Boron', weight: 10.811, symbol: 'B'},
  {position: 6, name: 'Carbon', weight: 12.0107, symbol: 'C'},
  {position: 7, name: 'Nitrogen', weight: 14.0067, symbol: 'N'},
  {position: 8, name: 'Oxygen', weight: 15.9994, symbol: 'O'},
  {position: 9, name: 'Fluorine', weight: 18.9984, symbol: 'F'},
  {position: 10, name: 'Neon', weight: 20.1797, symbol: 'Ne'},
];
@Component({
  selector: 'lib-displaysheet',
  templateUrl: './displaysheet.component.html',
  styleUrls: ['./displaysheet.component.css']
})
export class DisplaysheetComponent implements OnInit {

  columns = [
    {
      columnDef: 'position',
      header: 'No.',
      cell: (element: PeriodicElement) => `${element.position}`
    },
    {
      columnDef: 'name',
      header: 'Name',
      cell: (element: PeriodicElement) => `${element.name}`
    },
    {
      columnDef: 'weight',
      header: 'Weight',
      cell: (element: PeriodicElement) => `${element.weight}`
    },
    {
      columnDef: 'symbol',
      header: 'Symbol',
      cell: (element: PeriodicElement) => `${element.symbol}`
    }
  ];
  dataSource = ELEMENT_DATA;
  displayedColumns = this.columns.map(c => c.columnDef);

  public gongxlsxFrontRepo: gongxlsx.FrontRepo

  constructor(
    private gongxlsxFrontRepoService: gongxlsx.FrontRepoService,
  ) { }

  ngOnInit(): void {
    // this.gongxlsxFrontRepoService.pull().subscribe(
    //   gongxlsxsFrontRepo => {
    //     this.gongxlsxFrontRepo = gongxlsxsFrontRepo

    //     if (this.gongxlsxFrontRepo.XLFiles.size == 0) {
    //       console.error("cannot deal with 0 file")
    //     }

    //     let gongXLFile = this.gongxlsxFrontRepo.XLFiles_array[0];

    //     if (gongXLFile.Sheets.length == 0) {
    //       console.error("cannot deal with 0 sheets")
    //     }

    //     let gongXLSheet = gongXLFile.Sheets[0]

    //     for (let gongXLCell of gongXLSheet.SheetCells) {

    //       // only display cells of first line
    //       if (gongXLCell.Y == 0) {
    //         console.log(gongXLCell.Name)

    //         // this.columns.push(
    //         //   {
    //         //     columnDef: gongXLCell.Name,
    //         //     header: gongXLCell.Name,
    //         //     cell: (element: PeriodicElement) => `${element.symbol}`
    //         //   }
    //         // )
    //       }
    //     }

    //     // make up displayed columns
    //     // this.displayedColumns = this.columns.map(c => c.columnDef);

    //     console.log(this.displayedColumns)
    //   }
    // )
  }

}
