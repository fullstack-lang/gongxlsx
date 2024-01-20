import { Injectable } from '@angular/core'
import { HttpClient, HttpHeaders } from '@angular/common/http'

import { Observable, combineLatest, BehaviorSubject, of } from 'rxjs'

// insertion point sub template for services imports
import { DisplaySelectionDB } from './displayselection-db'
import { DisplaySelection, CopyDisplaySelectionDBToDisplaySelection } from './displayselection'
import { DisplaySelectionService } from './displayselection.service'

import { XLCellDB } from './xlcell-db'
import { XLCell, CopyXLCellDBToXLCell } from './xlcell'
import { XLCellService } from './xlcell.service'

import { XLFileDB } from './xlfile-db'
import { XLFile, CopyXLFileDBToXLFile } from './xlfile'
import { XLFileService } from './xlfile.service'

import { XLRowDB } from './xlrow-db'
import { XLRow, CopyXLRowDBToXLRow } from './xlrow'
import { XLRowService } from './xlrow.service'

import { XLSheetDB } from './xlsheet-db'
import { XLSheet, CopyXLSheetDBToXLSheet } from './xlsheet'
import { XLSheetService } from './xlsheet.service'

export const StackType = "github.com/fullstack-lang/gongxlsx/go/models"

// FrontRepo stores all instances in a front repository (design pattern repository)
export class FrontRepo { // insertion point sub template
  DisplaySelections_array = new Array<DisplaySelectionDB>() // array of repo instances
  DisplaySelections = new Map<number, DisplaySelectionDB>() // map of repo instances
  DisplaySelections_batch = new Map<number, DisplaySelectionDB>() // same but only in last GET (for finding repo instances to delete)

  array_DisplaySelections = new Array<DisplaySelection>() // array of front instances
  map_ID_DisplaySelection = new Map<number, DisplaySelection>() // map of front instances

  XLCells_array = new Array<XLCellDB>() // array of repo instances
  XLCells = new Map<number, XLCellDB>() // map of repo instances
  XLCells_batch = new Map<number, XLCellDB>() // same but only in last GET (for finding repo instances to delete)

  array_XLCells = new Array<XLCell>() // array of front instances
  map_ID_XLCell = new Map<number, XLCell>() // map of front instances

  XLFiles_array = new Array<XLFileDB>() // array of repo instances
  XLFiles = new Map<number, XLFileDB>() // map of repo instances
  XLFiles_batch = new Map<number, XLFileDB>() // same but only in last GET (for finding repo instances to delete)

  array_XLFiles = new Array<XLFile>() // array of front instances
  map_ID_XLFile = new Map<number, XLFile>() // map of front instances

  XLRows_array = new Array<XLRowDB>() // array of repo instances
  XLRows = new Map<number, XLRowDB>() // map of repo instances
  XLRows_batch = new Map<number, XLRowDB>() // same but only in last GET (for finding repo instances to delete)

  array_XLRows = new Array<XLRow>() // array of front instances
  map_ID_XLRow = new Map<number, XLRow>() // map of front instances

  XLSheets_array = new Array<XLSheetDB>() // array of repo instances
  XLSheets = new Map<number, XLSheetDB>() // map of repo instances
  XLSheets_batch = new Map<number, XLSheetDB>() // same but only in last GET (for finding repo instances to delete)

  array_XLSheets = new Array<XLSheet>() // array of front instances
  map_ID_XLSheet = new Map<number, XLSheet>() // map of front instances


  // getArray allows for a get function that is robust to refactoring of the named struct name
  // for instance frontRepo.getArray<Astruct>( Astruct.GONGSTRUCT_NAME), is robust to a refactoring of Astruct identifier
  // contrary to frontRepo.Astructs_array which is not refactored when Astruct identifier is modified
  getArray<Type>(gongStructName: string): Array<Type> {
    switch (gongStructName) { // deprecated
      // insertion point
      case 'DisplaySelection':
        return this.DisplaySelections_array as unknown as Array<Type>
      case 'XLCell':
        return this.XLCells_array as unknown as Array<Type>
      case 'XLFile':
        return this.XLFiles_array as unknown as Array<Type>
      case 'XLRow':
        return this.XLRows_array as unknown as Array<Type>
      case 'XLSheet':
        return this.XLSheets_array as unknown as Array<Type>
      default:
        throw new Error("Type not recognized");
    }
  }

  getFrontArray<Type>(gongStructName: string): Array<Type> {
    switch (gongStructName) {
      // insertion point
      case 'DisplaySelection':
        return this.array_DisplaySelections as unknown as Array<Type>
      case 'XLCell':
        return this.array_XLCells as unknown as Array<Type>
      case 'XLFile':
        return this.array_XLFiles as unknown as Array<Type>
      case 'XLRow':
        return this.array_XLRows as unknown as Array<Type>
      case 'XLSheet':
        return this.array_XLSheets as unknown as Array<Type>
      default:
        throw new Error("Type not recognized");
    }
  }

  // getMap allows for a get function that is robust to refactoring of the named struct name
  getMap<Type>(gongStructName: string): Map<number, Type> { // deprecated
    switch (gongStructName) {
      // insertion point
      case 'DisplaySelection':
        return this.DisplaySelections as unknown as Map<number, Type>
      case 'XLCell':
        return this.XLCells as unknown as Map<number, Type>
      case 'XLFile':
        return this.XLFiles as unknown as Map<number, Type>
      case 'XLRow':
        return this.XLRows as unknown as Map<number, Type>
      case 'XLSheet':
        return this.XLSheets as unknown as Map<number, Type>
      default:
        throw new Error("Type not recognized");
    }
  }
  
  getFrontMap<Type>(gongStructName: string): Map<number, Type> {
    switch (gongStructName) {
      // insertion point
      case 'DisplaySelection':
        return this.map_ID_DisplaySelection as unknown as Map<number, Type>
      case 'XLCell':
        return this.map_ID_XLCell as unknown as Map<number, Type>
      case 'XLFile':
        return this.map_ID_XLFile as unknown as Map<number, Type>
      case 'XLRow':
        return this.map_ID_XLRow as unknown as Map<number, Type>
      case 'XLSheet':
        return this.map_ID_XLSheet as unknown as Map<number, Type>
      default:
        throw new Error("Type not recognized");
    }
  }
}

// the table component is called in different ways
//
// DISPLAY or ASSOCIATION MODE
//
// in ASSOCIATION MODE, it is invoked within a diaglo and a Dialog Data item is used to
// configure the component
// DialogData define the interface for information that is forwarded from the calling instance to 
// the select table
export class DialogData {
  ID: number = 0 // ID of the calling instance

  // the reverse pointer is the name of the generated field on the destination
  // struct of the ONE-MANY association
  ReversePointer: string = "" // field of {{Structname}} that serve as reverse pointer
  OrderingMode: boolean = false // if true, this is for ordering items

  // there are different selection mode : ONE_MANY or MANY_MANY
  SelectionMode: SelectionMode = SelectionMode.ONE_MANY_ASSOCIATION_MODE

  // used if SelectionMode is MANY_MANY_ASSOCIATION_MODE
  //
  // In Gong, a MANY-MANY association is implemented as a ONE-ZERO/ONE followed by a ONE_MANY association
  // 
  // in the MANY_MANY_ASSOCIATION_MODE case, we need also the Struct and the FieldName that are
  // at the end of the ONE-MANY association
  SourceStruct: string = ""  // The "Aclass"
  SourceField: string = "" // the "AnarrayofbUse"
  IntermediateStruct: string = "" // the "AclassBclassUse" 
  IntermediateStructField: string = "" // the "Bclass" as field
  NextAssociationStruct: string = "" // the "Bclass"

  GONG__StackPath: string = ""
}

export enum SelectionMode {
  ONE_MANY_ASSOCIATION_MODE = "ONE_MANY_ASSOCIATION_MODE",
  MANY_MANY_ASSOCIATION_MODE = "MANY_MANY_ASSOCIATION_MODE",
}

//
// observable that fetch all elements of the stack and store them in the FrontRepo
//
@Injectable({
  providedIn: 'root'
})
export class FrontRepoService {

  GONG__StackPath: string = ""

  httpOptions = {
    headers: new HttpHeaders({ 'Content-Type': 'application/json' })
  };

  //
  // Store of all instances of the stack
  //
  frontRepo = new (FrontRepo)

  constructor(
    private http: HttpClient, // insertion point sub template 
    private displayselectionService: DisplaySelectionService,
    private xlcellService: XLCellService,
    private xlfileService: XLFileService,
    private xlrowService: XLRowService,
    private xlsheetService: XLSheetService,
  ) { }

  // postService provides a post function for each struct name
  postService(structName: string, instanceToBePosted: any) {
    let service = this[structName.toLowerCase() + "Service" + "Service" as keyof FrontRepoService]
    let servicePostFunction = service[("post" + structName) as keyof typeof service] as (instance: typeof instanceToBePosted) => Observable<typeof instanceToBePosted>

    servicePostFunction(instanceToBePosted).subscribe(
      instance => {
        let behaviorSubject = instanceToBePosted[(structName + "ServiceChanged") as keyof typeof instanceToBePosted] as unknown as BehaviorSubject<string>
        behaviorSubject.next("post")
      }
    );
  }

  // deleteService provides a delete function for each struct name
  deleteService(structName: string, instanceToBeDeleted: any) {
    let service = this[structName.toLowerCase() + "Service" as keyof FrontRepoService]
    let serviceDeleteFunction = service["delete" + structName as keyof typeof service] as (instance: typeof instanceToBeDeleted) => Observable<typeof instanceToBeDeleted>

    serviceDeleteFunction(instanceToBeDeleted).subscribe(
      instance => {
        let behaviorSubject = instanceToBeDeleted[(structName + "ServiceChanged") as keyof typeof instanceToBeDeleted] as unknown as BehaviorSubject<string>
        behaviorSubject.next("delete")
      }
    );
  }

  // typing of observable can be messy in typescript. Therefore, one force the type
  observableFrontRepo: [
    Observable<null>, // see below for the of(null) observable
    // insertion point sub template 
    Observable<DisplaySelectionDB[]>,
    Observable<XLCellDB[]>,
    Observable<XLFileDB[]>,
    Observable<XLRowDB[]>,
    Observable<XLSheetDB[]>,
  ] = [
      // Using "combineLatest" with a placeholder observable.
      //
      // This allows the typescript compiler to pass when no GongStruct is present in the front API
      //
      // The "of(null)" is a "meaningless" observable that emits a single value (null) and completes.
      // This is used as a workaround to satisfy TypeScript requirements and the "combineLatest" 
      // expectation for a non-empty array of observables.
      of(null), // 
      // insertion point sub template
      this.displayselectionService.getDisplaySelections(this.GONG__StackPath, this.frontRepo),
      this.xlcellService.getXLCells(this.GONG__StackPath, this.frontRepo),
      this.xlfileService.getXLFiles(this.GONG__StackPath, this.frontRepo),
      this.xlrowService.getXLRows(this.GONG__StackPath, this.frontRepo),
      this.xlsheetService.getXLSheets(this.GONG__StackPath, this.frontRepo),
    ];

  //
  // pull performs a GET on all struct of the stack and redeem association pointers 
  //
  // This is an observable. Therefore, the control flow forks with
  // - pull() return immediatly the observable
  // - the observable observer, if it subscribe, is called when all GET calls are performs
  pull(GONG__StackPath: string = ""): Observable<FrontRepo> {

    this.GONG__StackPath = GONG__StackPath

    this.observableFrontRepo = [
      of(null), // see above for justification
      // insertion point sub template
      this.displayselectionService.getDisplaySelections(this.GONG__StackPath, this.frontRepo),
      this.xlcellService.getXLCells(this.GONG__StackPath, this.frontRepo),
      this.xlfileService.getXLFiles(this.GONG__StackPath, this.frontRepo),
      this.xlrowService.getXLRows(this.GONG__StackPath, this.frontRepo),
      this.xlsheetService.getXLSheets(this.GONG__StackPath, this.frontRepo),
    ]

    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest(
          this.observableFrontRepo
        ).subscribe(
          ([
            ___of_null, // see above for the explanation about of
            // insertion point sub template for declarations 
            displayselections_,
            xlcells_,
            xlfiles_,
            xlrows_,
            xlsheets_,
          ]) => {
            // Typing can be messy with many items. Therefore, type casting is necessary here
            // insertion point sub template for type casting 
            var displayselections: DisplaySelectionDB[]
            displayselections = displayselections_ as DisplaySelectionDB[]
            var xlcells: XLCellDB[]
            xlcells = xlcells_ as XLCellDB[]
            var xlfiles: XLFileDB[]
            xlfiles = xlfiles_ as XLFileDB[]
            var xlrows: XLRowDB[]
            xlrows = xlrows_ as XLRowDB[]
            var xlsheets: XLSheetDB[]
            xlsheets = xlsheets_ as XLSheetDB[]

            // 
            // First Step: init map of instances
            // insertion point sub template for init 
            // init the array
            this.frontRepo.DisplaySelections_array = displayselections

            // clear the map that counts DisplaySelection in the GET
            this.frontRepo.DisplaySelections_batch.clear()

            displayselections.forEach(
              displayselectionDB => {
                this.frontRepo.DisplaySelections.set(displayselectionDB.ID, displayselectionDB)
                this.frontRepo.DisplaySelections_batch.set(displayselectionDB.ID, displayselectionDB)
              }
            )

            // clear displayselections that are absent from the batch
            this.frontRepo.DisplaySelections.forEach(
              displayselectionDB => {
                if (this.frontRepo.DisplaySelections_batch.get(displayselectionDB.ID) == undefined) {
                  this.frontRepo.DisplaySelections.delete(displayselectionDB.ID)
                }
              }
            )

            // sort DisplaySelections_array array
            this.frontRepo.DisplaySelections_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            this.frontRepo.XLCells_array = xlcells

            // clear the map that counts XLCell in the GET
            this.frontRepo.XLCells_batch.clear()

            xlcells.forEach(
              xlcellDB => {
                this.frontRepo.XLCells.set(xlcellDB.ID, xlcellDB)
                this.frontRepo.XLCells_batch.set(xlcellDB.ID, xlcellDB)
              }
            )

            // clear xlcells that are absent from the batch
            this.frontRepo.XLCells.forEach(
              xlcellDB => {
                if (this.frontRepo.XLCells_batch.get(xlcellDB.ID) == undefined) {
                  this.frontRepo.XLCells.delete(xlcellDB.ID)
                }
              }
            )

            // sort XLCells_array array
            this.frontRepo.XLCells_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            this.frontRepo.XLFiles_array = xlfiles

            // clear the map that counts XLFile in the GET
            this.frontRepo.XLFiles_batch.clear()

            xlfiles.forEach(
              xlfileDB => {
                this.frontRepo.XLFiles.set(xlfileDB.ID, xlfileDB)
                this.frontRepo.XLFiles_batch.set(xlfileDB.ID, xlfileDB)
              }
            )

            // clear xlfiles that are absent from the batch
            this.frontRepo.XLFiles.forEach(
              xlfileDB => {
                if (this.frontRepo.XLFiles_batch.get(xlfileDB.ID) == undefined) {
                  this.frontRepo.XLFiles.delete(xlfileDB.ID)
                }
              }
            )

            // sort XLFiles_array array
            this.frontRepo.XLFiles_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            this.frontRepo.XLRows_array = xlrows

            // clear the map that counts XLRow in the GET
            this.frontRepo.XLRows_batch.clear()

            xlrows.forEach(
              xlrowDB => {
                this.frontRepo.XLRows.set(xlrowDB.ID, xlrowDB)
                this.frontRepo.XLRows_batch.set(xlrowDB.ID, xlrowDB)
              }
            )

            // clear xlrows that are absent from the batch
            this.frontRepo.XLRows.forEach(
              xlrowDB => {
                if (this.frontRepo.XLRows_batch.get(xlrowDB.ID) == undefined) {
                  this.frontRepo.XLRows.delete(xlrowDB.ID)
                }
              }
            )

            // sort XLRows_array array
            this.frontRepo.XLRows_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            this.frontRepo.XLSheets_array = xlsheets

            // clear the map that counts XLSheet in the GET
            this.frontRepo.XLSheets_batch.clear()

            xlsheets.forEach(
              xlsheetDB => {
                this.frontRepo.XLSheets.set(xlsheetDB.ID, xlsheetDB)
                this.frontRepo.XLSheets_batch.set(xlsheetDB.ID, xlsheetDB)
              }
            )

            // clear xlsheets that are absent from the batch
            this.frontRepo.XLSheets.forEach(
              xlsheetDB => {
                if (this.frontRepo.XLSheets_batch.get(xlsheetDB.ID) == undefined) {
                  this.frontRepo.XLSheets.delete(xlsheetDB.ID)
                }
              }
            )

            // sort XLSheets_array array
            this.frontRepo.XLSheets_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });


            // 
            // Second Step: reddeem slice of pointers fields
            // insertion point sub template for redeem 
            displayselections.forEach(
              displayselection => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming
                // insertion point for pointer field XLFile redeeming
                {
                  let _xlfile = this.frontRepo.XLFiles.get(displayselection.DisplaySelectionPointersEncoding.XLFileID.Int64)
                  if (_xlfile) {
                    displayselection.XLFile = _xlfile
                  }
                }
                // insertion point for pointer field XLSheet redeeming
                {
                  let _xlsheet = this.frontRepo.XLSheets.get(displayselection.DisplaySelectionPointersEncoding.XLSheetID.Int64)
                  if (_xlsheet) {
                    displayselection.XLSheet = _xlsheet
                  }
                }
                // insertion point for pointers decoding
              }
            )
            xlcells.forEach(
              xlcell => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming
                // insertion point for pointers decoding
              }
            )
            xlfiles.forEach(
              xlfile => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming
                // insertion point for pointers decoding
                xlfile.Sheets = new Array<XLSheetDB>()
                for (let _id of xlfile.XLFilePointersEncoding.Sheets) {
                  let _xlsheet = this.frontRepo.XLSheets.get(_id)
                  if (_xlsheet != undefined) {
                    xlfile.Sheets.push(_xlsheet!)
                  }
                }
              }
            )
            xlrows.forEach(
              xlrow => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming
                // insertion point for pointers decoding
                xlrow.Cells = new Array<XLCellDB>()
                for (let _id of xlrow.XLRowPointersEncoding.Cells) {
                  let _xlcell = this.frontRepo.XLCells.get(_id)
                  if (_xlcell != undefined) {
                    xlrow.Cells.push(_xlcell!)
                  }
                }
              }
            )
            xlsheets.forEach(
              xlsheet => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming
                // insertion point for pointers decoding
                xlsheet.Rows = new Array<XLRowDB>()
                for (let _id of xlsheet.XLSheetPointersEncoding.Rows) {
                  let _xlrow = this.frontRepo.XLRows.get(_id)
                  if (_xlrow != undefined) {
                    xlsheet.Rows.push(_xlrow!)
                  }
                }
                xlsheet.SheetCells = new Array<XLCellDB>()
                for (let _id of xlsheet.XLSheetPointersEncoding.SheetCells) {
                  let _xlcell = this.frontRepo.XLCells.get(_id)
                  if (_xlcell != undefined) {
                    xlsheet.SheetCells.push(_xlcell!)
                  }
                }
              }
            )

            // 
            // Third Step: reddeem front objects
            // insertion point sub template for redeem 
            
            // init front objects
            this.frontRepo.array_DisplaySelections = []
            this.frontRepo.map_ID_DisplaySelection.clear()
            this.frontRepo.DisplaySelections_array.forEach(
              displayselectionDB => {
                let displayselection = new DisplaySelection
                CopyDisplaySelectionDBToDisplaySelection(displayselectionDB, displayselection, this.frontRepo)
                this.frontRepo.array_DisplaySelections.push(displayselection)
                this.frontRepo.map_ID_DisplaySelection.set(displayselection.ID, displayselection)
              }
            )

            
            // init front objects
            this.frontRepo.array_XLCells = []
            this.frontRepo.map_ID_XLCell.clear()
            this.frontRepo.XLCells_array.forEach(
              xlcellDB => {
                let xlcell = new XLCell
                CopyXLCellDBToXLCell(xlcellDB, xlcell, this.frontRepo)
                this.frontRepo.array_XLCells.push(xlcell)
                this.frontRepo.map_ID_XLCell.set(xlcell.ID, xlcell)
              }
            )

            
            // init front objects
            this.frontRepo.array_XLFiles = []
            this.frontRepo.map_ID_XLFile.clear()
            this.frontRepo.XLFiles_array.forEach(
              xlfileDB => {
                let xlfile = new XLFile
                CopyXLFileDBToXLFile(xlfileDB, xlfile, this.frontRepo)
                this.frontRepo.array_XLFiles.push(xlfile)
                this.frontRepo.map_ID_XLFile.set(xlfile.ID, xlfile)
              }
            )

            
            // init front objects
            this.frontRepo.array_XLRows = []
            this.frontRepo.map_ID_XLRow.clear()
            this.frontRepo.XLRows_array.forEach(
              xlrowDB => {
                let xlrow = new XLRow
                CopyXLRowDBToXLRow(xlrowDB, xlrow, this.frontRepo)
                this.frontRepo.array_XLRows.push(xlrow)
                this.frontRepo.map_ID_XLRow.set(xlrow.ID, xlrow)
              }
            )

            
            // init front objects
            this.frontRepo.array_XLSheets = []
            this.frontRepo.map_ID_XLSheet.clear()
            this.frontRepo.XLSheets_array.forEach(
              xlsheetDB => {
                let xlsheet = new XLSheet
                CopyXLSheetDBToXLSheet(xlsheetDB, xlsheet, this.frontRepo)
                this.frontRepo.array_XLSheets.push(xlsheet)
                this.frontRepo.map_ID_XLSheet.set(xlsheet.ID, xlsheet)
              }
            )



            // hand over control flow to observer
            observer.next(this.frontRepo)
          }
        )
      }
    )
  }

  // insertion point for pull per struct 

  // DisplaySelectionPull performs a GET on DisplaySelection of the stack and redeem association pointers 
  DisplaySelectionPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.displayselectionService.getDisplaySelections(this.GONG__StackPath, this.frontRepo)
        ]).subscribe(
          ([ // insertion point sub template 
            displayselections,
          ]) => {
            // init the array
            this.frontRepo.DisplaySelections_array = displayselections

            // clear the map that counts DisplaySelection in the GET
            this.frontRepo.DisplaySelections_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            displayselections.forEach(
              displayselection => {
                this.frontRepo.DisplaySelections.set(displayselection.ID, displayselection)
                this.frontRepo.DisplaySelections_batch.set(displayselection.ID, displayselection)

                // insertion point for redeeming ONE/ZERO-ONE associations
                // insertion point for pointer field XLFile redeeming
                {
                  let _xlfile = this.frontRepo.XLFiles.get(displayselection.DisplaySelectionPointersEncoding.XLFileID.Int64)
                  if (_xlfile) {
                    displayselection.XLFile = _xlfile
                  }
                }
                // insertion point for pointer field XLSheet redeeming
                {
                  let _xlsheet = this.frontRepo.XLSheets.get(displayselection.DisplaySelectionPointersEncoding.XLSheetID.Int64)
                  if (_xlsheet) {
                    displayselection.XLSheet = _xlsheet
                  }
                }
              }
            )

            // clear displayselections that are absent from the GET
            this.frontRepo.DisplaySelections.forEach(
              displayselection => {
                if (this.frontRepo.DisplaySelections_batch.get(displayselection.ID) == undefined) {
                  this.frontRepo.DisplaySelections.delete(displayselection.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(this.frontRepo)
          }
        )
      }
    )
  }

  // XLCellPull performs a GET on XLCell of the stack and redeem association pointers 
  XLCellPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.xlcellService.getXLCells(this.GONG__StackPath, this.frontRepo)
        ]).subscribe(
          ([ // insertion point sub template 
            xlcells,
          ]) => {
            // init the array
            this.frontRepo.XLCells_array = xlcells

            // clear the map that counts XLCell in the GET
            this.frontRepo.XLCells_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            xlcells.forEach(
              xlcell => {
                this.frontRepo.XLCells.set(xlcell.ID, xlcell)
                this.frontRepo.XLCells_batch.set(xlcell.ID, xlcell)

                // insertion point for redeeming ONE/ZERO-ONE associations
              }
            )

            // clear xlcells that are absent from the GET
            this.frontRepo.XLCells.forEach(
              xlcell => {
                if (this.frontRepo.XLCells_batch.get(xlcell.ID) == undefined) {
                  this.frontRepo.XLCells.delete(xlcell.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(this.frontRepo)
          }
        )
      }
    )
  }

  // XLFilePull performs a GET on XLFile of the stack and redeem association pointers 
  XLFilePull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.xlfileService.getXLFiles(this.GONG__StackPath, this.frontRepo)
        ]).subscribe(
          ([ // insertion point sub template 
            xlfiles,
          ]) => {
            // init the array
            this.frontRepo.XLFiles_array = xlfiles

            // clear the map that counts XLFile in the GET
            this.frontRepo.XLFiles_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            xlfiles.forEach(
              xlfile => {
                this.frontRepo.XLFiles.set(xlfile.ID, xlfile)
                this.frontRepo.XLFiles_batch.set(xlfile.ID, xlfile)

                // insertion point for redeeming ONE/ZERO-ONE associations
              }
            )

            // clear xlfiles that are absent from the GET
            this.frontRepo.XLFiles.forEach(
              xlfile => {
                if (this.frontRepo.XLFiles_batch.get(xlfile.ID) == undefined) {
                  this.frontRepo.XLFiles.delete(xlfile.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(this.frontRepo)
          }
        )
      }
    )
  }

  // XLRowPull performs a GET on XLRow of the stack and redeem association pointers 
  XLRowPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.xlrowService.getXLRows(this.GONG__StackPath, this.frontRepo)
        ]).subscribe(
          ([ // insertion point sub template 
            xlrows,
          ]) => {
            // init the array
            this.frontRepo.XLRows_array = xlrows

            // clear the map that counts XLRow in the GET
            this.frontRepo.XLRows_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            xlrows.forEach(
              xlrow => {
                this.frontRepo.XLRows.set(xlrow.ID, xlrow)
                this.frontRepo.XLRows_batch.set(xlrow.ID, xlrow)

                // insertion point for redeeming ONE/ZERO-ONE associations
              }
            )

            // clear xlrows that are absent from the GET
            this.frontRepo.XLRows.forEach(
              xlrow => {
                if (this.frontRepo.XLRows_batch.get(xlrow.ID) == undefined) {
                  this.frontRepo.XLRows.delete(xlrow.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(this.frontRepo)
          }
        )
      }
    )
  }

  // XLSheetPull performs a GET on XLSheet of the stack and redeem association pointers 
  XLSheetPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.xlsheetService.getXLSheets(this.GONG__StackPath, this.frontRepo)
        ]).subscribe(
          ([ // insertion point sub template 
            xlsheets,
          ]) => {
            // init the array
            this.frontRepo.XLSheets_array = xlsheets

            // clear the map that counts XLSheet in the GET
            this.frontRepo.XLSheets_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            xlsheets.forEach(
              xlsheet => {
                this.frontRepo.XLSheets.set(xlsheet.ID, xlsheet)
                this.frontRepo.XLSheets_batch.set(xlsheet.ID, xlsheet)

                // insertion point for redeeming ONE/ZERO-ONE associations
              }
            )

            // clear xlsheets that are absent from the GET
            this.frontRepo.XLSheets.forEach(
              xlsheet => {
                if (this.frontRepo.XLSheets_batch.get(xlsheet.ID) == undefined) {
                  this.frontRepo.XLSheets.delete(xlsheet.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(this.frontRepo)
          }
        )
      }
    )
  }
}

// insertion point for get unique ID per struct 
export function getDisplaySelectionUniqueID(id: number): number {
  return 31 * id
}
export function getXLCellUniqueID(id: number): number {
  return 37 * id
}
export function getXLFileUniqueID(id: number): number {
  return 41 * id
}
export function getXLRowUniqueID(id: number): number {
  return 43 * id
}
export function getXLSheetUniqueID(id: number): number {
  return 47 * id
}
