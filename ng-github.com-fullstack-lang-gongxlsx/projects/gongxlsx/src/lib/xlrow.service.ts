// generated by ng_file_service_ts
import { Injectable, Component, Inject } from '@angular/core';
import { HttpClientModule, HttpParams } from '@angular/common/http';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { DOCUMENT, Location } from '@angular/common'

/*
 * Behavior subject
 */
import { BehaviorSubject } from 'rxjs'
import { Observable, of } from 'rxjs'
import { catchError, map, tap } from 'rxjs/operators'

import { XLRowAPI } from './xlrow-api'
import { XLRow, CopyXLRowToXLRowAPI } from './xlrow'

import { FrontRepo, FrontRepoService } from './front-repo.service';

// insertion point for imports
import { XLCellAPI } from './xlcell-api'

@Injectable({
  providedIn: 'root'
})
export class XLRowService {

  // Kamar Raïmo: Adding a way to communicate between components that share information
  // so that they are notified of a change.
  XLRowServiceChanged: BehaviorSubject<string> = new BehaviorSubject("");

  private xlrowsUrl: string

  constructor(
    private http: HttpClient,
    @Inject(DOCUMENT) private document: Document
  ) {
    // path to the service share the same origin with the path to the document
    // get the origin in the URL to the document
    let origin = this.document.location.origin

    // if debugging with ng, replace 4200 with 8080
    origin = origin.replace("4200", "8080")

    // compute path to the service
    this.xlrowsUrl = origin + '/api/github.com/fullstack-lang/gongxlsx/go/v1/xlrows';
  }

  /** GET xlrows from the server */
  // gets is more robust to refactoring
  gets(GONG__StackPath: string, frontRepo: FrontRepo): Observable<XLRowAPI[]> {
    return this.getXLRows(GONG__StackPath, frontRepo)
  }
  getXLRows(GONG__StackPath: string, frontRepo: FrontRepo): Observable<XLRowAPI[]> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)

    return this.http.get<XLRowAPI[]>(this.xlrowsUrl, { params: params })
      .pipe(
        tap(),
        catchError(this.handleError<XLRowAPI[]>('getXLRows', []))
      );
  }

  /** GET xlrow by id. Will 404 if id not found */
  // more robust API to refactoring
  get(id: number, GONG__StackPath: string, frontRepo: FrontRepo): Observable<XLRowAPI> {
    return this.getXLRow(id, GONG__StackPath, frontRepo)
  }
  getXLRow(id: number, GONG__StackPath: string, frontRepo: FrontRepo): Observable<XLRowAPI> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)

    const url = `${this.xlrowsUrl}/${id}`;
    return this.http.get<XLRowAPI>(url, { params: params }).pipe(
      // tap(_ => this.log(`fetched xlrow id=${id}`)),
      catchError(this.handleError<XLRowAPI>(`getXLRow id=${id}`))
    );
  }

  // postFront copy xlrow to a version with encoded pointers and post to the back
  postFront(xlrow: XLRow, GONG__StackPath: string): Observable<XLRowAPI> {
    let xlrowAPI = new XLRowAPI
    CopyXLRowToXLRowAPI(xlrow, xlrowAPI)
    const id = typeof xlrowAPI === 'number' ? xlrowAPI : xlrowAPI.ID
    const url = `${this.xlrowsUrl}/${id}`;
    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.post<XLRowAPI>(url, xlrowAPI, httpOptions).pipe(
      tap(_ => {
      }),
      catchError(this.handleError<XLRowAPI>('postXLRow'))
    );
  }
  
  /** POST: add a new xlrow to the server */
  post(xlrowdb: XLRowAPI, GONG__StackPath: string, frontRepo: FrontRepo): Observable<XLRowAPI> {
    return this.postXLRow(xlrowdb, GONG__StackPath, frontRepo)
  }
  postXLRow(xlrowdb: XLRowAPI, GONG__StackPath: string, frontRepo: FrontRepo): Observable<XLRowAPI> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.post<XLRowAPI>(this.xlrowsUrl, xlrowdb, httpOptions).pipe(
      tap(_ => {
        // this.log(`posted xlrowdb id=${xlrowdb.ID}`)
      }),
      catchError(this.handleError<XLRowAPI>('postXLRow'))
    );
  }

  /** DELETE: delete the xlrowdb from the server */
  delete(xlrowdb: XLRowAPI | number, GONG__StackPath: string): Observable<XLRowAPI> {
    return this.deleteXLRow(xlrowdb, GONG__StackPath)
  }
  deleteXLRow(xlrowdb: XLRowAPI | number, GONG__StackPath: string): Observable<XLRowAPI> {
    const id = typeof xlrowdb === 'number' ? xlrowdb : xlrowdb.ID;
    const url = `${this.xlrowsUrl}/${id}`;

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    };

    return this.http.delete<XLRowAPI>(url, httpOptions).pipe(
      tap(_ => this.log(`deleted xlrowdb id=${id}`)),
      catchError(this.handleError<XLRowAPI>('deleteXLRow'))
    );
  }

  // updateFront copy xlrow to a version with encoded pointers and update to the back
  updateFront(xlrow: XLRow, GONG__StackPath: string): Observable<XLRowAPI> {
    let xlrowAPI = new XLRowAPI
    CopyXLRowToXLRowAPI(xlrow, xlrowAPI)
    const id = typeof xlrowAPI === 'number' ? xlrowAPI : xlrowAPI.ID
    const url = `${this.xlrowsUrl}/${id}`;
    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.put<XLRowAPI>(url, xlrowAPI, httpOptions).pipe(
      tap(_ => {
      }),
      catchError(this.handleError<XLRowAPI>('updateXLRow'))
    );
  }

  /** PUT: update the xlrowdb on the server */
  update(xlrowdb: XLRowAPI, GONG__StackPath: string, frontRepo: FrontRepo): Observable<XLRowAPI> {
    return this.updateXLRow(xlrowdb, GONG__StackPath, frontRepo)
  }
  updateXLRow(xlrowdb: XLRowAPI, GONG__StackPath: string, frontRepo: FrontRepo): Observable<XLRowAPI> {
    const id = typeof xlrowdb === 'number' ? xlrowdb : xlrowdb.ID;
    const url = `${this.xlrowsUrl}/${id}`;


    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    };

    return this.http.put<XLRowAPI>(url, xlrowdb, httpOptions).pipe(
      tap(_ => {
        // this.log(`updated xlrowdb id=${xlrowdb.ID}`)
      }),
      catchError(this.handleError<XLRowAPI>('updateXLRow'))
    );
  }

  /**
   * Handle Http operation that failed.
   * Let the app continue.
   * @param operation - name of the operation that failed
   * @param result - optional value to return as the observable result
   */
  private handleError<T>(operation = 'operation in XLRowService', result?: T) {
    return (error: any): Observable<T> => {

      // TODO: send the error to remote logging infrastructure
      console.error("XLRowService" + error); // log to console instead

      // TODO: better job of transforming error for user consumption
      this.log(`${operation} failed: ${error.message}`);

      // Let the app keep running by returning an empty result.
      return of(result as T);
    };
  }

  private log(message: string) {
    console.log(message)
  }
}