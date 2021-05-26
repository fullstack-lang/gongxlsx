import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';

import { Observable, combineLatest } from 'rxjs';

// insertion point sub template for services imports 
import { XslxDB } from './xslx-db'
import { XslxService } from './xslx.service'


// FrontRepo stores all instances in a front repository (design pattern repository)
export class FrontRepo { // insertion point sub template 
  Xslxs_array = new Array<XslxDB>(); // array of repo instances
  Xslxs = new Map<number, XslxDB>(); // map of repo instances
  Xslxs_batch = new Map<number, XslxDB>(); // same but only in last GET (for finding repo instances to delete)
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
    private xslxService: XslxService,
  ) { }

  // typing of observable can be messy in typescript. Therefore, one force the type
  observableFrontRepo: [ // insertion point sub template 
    Observable<XslxDB[]>,
  ] = [ // insertion point sub template 
      this.xslxService.getXslxs(),
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
            xslxs_,
          ]) => {
            // Typing can be messy with many items. Therefore, type casting is necessary here
            // insertion point sub template for type casting 
            var xslxs: XslxDB[]
            xslxs = xslxs_

            // 
            // First Step: init map of instances
            // insertion point sub template for init 
            // init the array
            FrontRepoSingloton.Xslxs_array = xslxs

            // clear the map that counts Xslx in the GET
            FrontRepoSingloton.Xslxs_batch.clear()
            
            xslxs.forEach(
              xslx => {
                FrontRepoSingloton.Xslxs.set(xslx.ID, xslx)
                FrontRepoSingloton.Xslxs_batch.set(xslx.ID, xslx)
              }
            )
            
            // clear xslxs that are absent from the batch
            FrontRepoSingloton.Xslxs.forEach(
              xslx => {
                if (FrontRepoSingloton.Xslxs_batch.get(xslx.ID) == undefined) {
                  FrontRepoSingloton.Xslxs.delete(xslx.ID)
                }
              }
            )
            
            // sort Xslxs_array array
            FrontRepoSingloton.Xslxs_array.sort((t1, t2) => {
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
            xslxs.forEach(
              xslx => {
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

  // XslxPull performs a GET on Xslx of the stack and redeem association pointers 
  XslxPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.xslxService.getXslxs()
        ]).subscribe(
          ([ // insertion point sub template 
            xslxs,
          ]) => {
            // init the array
            FrontRepoSingloton.Xslxs_array = xslxs

            // clear the map that counts Xslx in the GET
            FrontRepoSingloton.Xslxs_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            xslxs.forEach(
              xslx => {
                FrontRepoSingloton.Xslxs.set(xslx.ID, xslx)
                FrontRepoSingloton.Xslxs_batch.set(xslx.ID, xslx)

                // insertion point for redeeming ONE/ZERO-ONE associations 

                // insertion point for redeeming ONE-MANY associations 
              }
            )

            // clear xslxs that are absent from the GET
            FrontRepoSingloton.Xslxs.forEach(
              xslx => {
                if (FrontRepoSingloton.Xslxs_batch.get(xslx.ID) == undefined) {
                  FrontRepoSingloton.Xslxs.delete(xslx.ID)
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
export function getXslxUniqueID(id: number): number {
  return 31 * id
}
