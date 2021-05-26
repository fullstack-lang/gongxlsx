import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';

import { Observable, combineLatest } from 'rxjs';

// insertion point sub template for services imports 
import { XLFileDB } from './xlfile-db'
import { XLFileService } from './xlfile.service'


// FrontRepo stores all instances in a front repository (design pattern repository)
export class FrontRepo { // insertion point sub template 
  XLFiles_array = new Array<XLFileDB>(); // array of repo instances
  XLFiles = new Map<number, XLFileDB>(); // map of repo instances
  XLFiles_batch = new Map<number, XLFileDB>(); // same but only in last GET (for finding repo instances to delete)
}

//
// Store of all instances of the stack
//
export const FrontRepoSingloton = new (FrontRepo)

// define the type of nullable Int64 in order to support back pointers IDs
export class NullInt64 {
    Int64: number
    Valid: boolean
}

// define the interface for information that is forwarded from the calling instance to 
// the select table
export interface DialogData {
  ID: number; // ID of the calling instance
  ReversePointer: string; // field of {{Structname}} that serve as reverse pointer
  OrderingMode: boolean; // if true, this is for ordering items
}

//
// observable that fetch all elements of the stack and store them in the FrontRepo
//
@Injectable({
  providedIn: 'root'
})
export class FrontRepoService {

  httpOptions = {
    headers: new HttpHeaders({ 'Content-Type': 'application/json' })
  };

  constructor(
    private http: HttpClient, // insertion point sub template 
    private xlfileService: XLFileService,
  ) { }

  // typing of observable can be messy in typescript. Therefore, one force the type
  observableFrontRepo: [ // insertion point sub template 
    Observable<XLFileDB[]>,
  ] = [ // insertion point sub template 
      this.xlfileService.getXLFiles(),
    ];

  //
  // pull performs a GET on all struct of the stack and redeem association pointers 
  //
  // This is an observable. Therefore, the control flow forks with
  // - pull() return immediatly the observable
  // - the observable observer, if it subscribe, is called when all GET calls are performs
  pull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest(
          this.observableFrontRepo
        ).subscribe(
          ([ // insertion point sub template for declarations 
            xlfiles_,
          ]) => {
            // Typing can be messy with many items. Therefore, type casting is necessary here
            // insertion point sub template for type casting 
            var xlfiles: XLFileDB[]
            xlfiles = xlfiles_

            // 
            // First Step: init map of instances
            // insertion point sub template for init 
            // init the array
            FrontRepoSingloton.XLFiles_array = xlfiles

            // clear the map that counts XLFile in the GET
            FrontRepoSingloton.XLFiles_batch.clear()
            
            xlfiles.forEach(
              xlfile => {
                FrontRepoSingloton.XLFiles.set(xlfile.ID, xlfile)
                FrontRepoSingloton.XLFiles_batch.set(xlfile.ID, xlfile)
              }
            )
            
            // clear xlfiles that are absent from the batch
            FrontRepoSingloton.XLFiles.forEach(
              xlfile => {
                if (FrontRepoSingloton.XLFiles_batch.get(xlfile.ID) == undefined) {
                  FrontRepoSingloton.XLFiles.delete(xlfile.ID)
                }
              }
            )
            
            // sort XLFiles_array array
            FrontRepoSingloton.XLFiles_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });
            

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template for redeem 
            xlfiles.forEach(
              xlfile => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
              }
            )

            // hand over control flow to observer
            observer.next(FrontRepoSingloton)
          }
        )
      }
    )
  }

  // insertion point for pull per struct 

  // XLFilePull performs a GET on XLFile of the stack and redeem association pointers 
  XLFilePull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.xlfileService.getXLFiles()
        ]).subscribe(
          ([ // insertion point sub template 
            xlfiles,
          ]) => {
            // init the array
            FrontRepoSingloton.XLFiles_array = xlfiles

            // clear the map that counts XLFile in the GET
            FrontRepoSingloton.XLFiles_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            xlfiles.forEach(
              xlfile => {
                FrontRepoSingloton.XLFiles.set(xlfile.ID, xlfile)
                FrontRepoSingloton.XLFiles_batch.set(xlfile.ID, xlfile)

                // insertion point for redeeming ONE/ZERO-ONE associations 

                // insertion point for redeeming ONE-MANY associations 
              }
            )

            // clear xlfiles that are absent from the GET
            FrontRepoSingloton.XLFiles.forEach(
              xlfile => {
                if (FrontRepoSingloton.XLFiles_batch.get(xlfile.ID) == undefined) {
                  FrontRepoSingloton.XLFiles.delete(xlfile.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(FrontRepoSingloton)
          }
        )
      }
    )
  }
}

// insertion point for get unique ID per struct 
export function getXLFileUniqueID(id: number): number {
  return 31 * id
}
