import { Injectable } from '@angular/core';
import { Route, Router, Routes } from '@angular/router';

// insertion point for imports
import { DisplaySelectionsTableComponent } from './displayselections-table/displayselections-table.component'
import { DisplaySelectionDetailComponent } from './displayselection-detail/displayselection-detail.component'

import { XLCellsTableComponent } from './xlcells-table/xlcells-table.component'
import { XLCellDetailComponent } from './xlcell-detail/xlcell-detail.component'

import { XLFilesTableComponent } from './xlfiles-table/xlfiles-table.component'
import { XLFileDetailComponent } from './xlfile-detail/xlfile-detail.component'

import { XLRowsTableComponent } from './xlrows-table/xlrows-table.component'
import { XLRowDetailComponent } from './xlrow-detail/xlrow-detail.component'

import { XLSheetsTableComponent } from './xlsheets-table/xlsheets-table.component'
import { XLSheetDetailComponent } from './xlsheet-detail/xlsheet-detail.component'


@Injectable({
    providedIn: 'root'
})
export class RouteService {
    private routes: Routes = []

    constructor(private router: Router) { }

    public addRoutes(newRoutes: Routes): void {
        const existingRoutes = this.router.config
        this.routes = this.router.config

        for (let newRoute of newRoutes) {
            if (!existingRoutes.includes(newRoute)) {
                this.routes.push(newRoute)
            }
        }
        this.router.resetConfig(this.routes)
    }

    getPathRoot(): string {
        return 'github_com_fullstack_lang_gongxlsx_go'
    }
    getTableOutlet(stackPath: string): string {
        return this.getPathRoot() + '_table' + '_' + stackPath
    }
    getEditorOutlet(stackPath: string): string {
        return this.getPathRoot() + '_editor' + '_' + stackPath
    }
    // insertion point for per gongstruct route/path getters
    getDisplaySelectionTablePath(): string {
        return this.getPathRoot() + '-displayselections/:GONG__StackPath'
    }
    getDisplaySelectionTableRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getDisplaySelectionTablePath(), component: DisplaySelectionsTableComponent, outlet: this.getTableOutlet(stackPath) }
        return route
    }
    getDisplaySelectionAdderPath(): string {
        return this.getPathRoot() + '-displayselection-adder/:GONG__StackPath'
    }
    getDisplaySelectionAdderRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getDisplaySelectionAdderPath(), component: DisplaySelectionDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getDisplaySelectionAdderForUsePath(): string {
        return this.getPathRoot() + '-displayselection-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath'
    }
    getDisplaySelectionAdderForUseRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getDisplaySelectionAdderForUsePath(), component: DisplaySelectionDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getDisplaySelectionDetailPath(): string {
        return this.getPathRoot() + '-displayselection-detail/:id/:GONG__StackPath'
    }
    getDisplaySelectionDetailRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getDisplaySelectionDetailPath(), component: DisplaySelectionDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }

    getXLCellTablePath(): string {
        return this.getPathRoot() + '-xlcells/:GONG__StackPath'
    }
    getXLCellTableRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getXLCellTablePath(), component: XLCellsTableComponent, outlet: this.getTableOutlet(stackPath) }
        return route
    }
    getXLCellAdderPath(): string {
        return this.getPathRoot() + '-xlcell-adder/:GONG__StackPath'
    }
    getXLCellAdderRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getXLCellAdderPath(), component: XLCellDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getXLCellAdderForUsePath(): string {
        return this.getPathRoot() + '-xlcell-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath'
    }
    getXLCellAdderForUseRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getXLCellAdderForUsePath(), component: XLCellDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getXLCellDetailPath(): string {
        return this.getPathRoot() + '-xlcell-detail/:id/:GONG__StackPath'
    }
    getXLCellDetailRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getXLCellDetailPath(), component: XLCellDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }

    getXLFileTablePath(): string {
        return this.getPathRoot() + '-xlfiles/:GONG__StackPath'
    }
    getXLFileTableRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getXLFileTablePath(), component: XLFilesTableComponent, outlet: this.getTableOutlet(stackPath) }
        return route
    }
    getXLFileAdderPath(): string {
        return this.getPathRoot() + '-xlfile-adder/:GONG__StackPath'
    }
    getXLFileAdderRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getXLFileAdderPath(), component: XLFileDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getXLFileAdderForUsePath(): string {
        return this.getPathRoot() + '-xlfile-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath'
    }
    getXLFileAdderForUseRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getXLFileAdderForUsePath(), component: XLFileDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getXLFileDetailPath(): string {
        return this.getPathRoot() + '-xlfile-detail/:id/:GONG__StackPath'
    }
    getXLFileDetailRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getXLFileDetailPath(), component: XLFileDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }

    getXLRowTablePath(): string {
        return this.getPathRoot() + '-xlrows/:GONG__StackPath'
    }
    getXLRowTableRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getXLRowTablePath(), component: XLRowsTableComponent, outlet: this.getTableOutlet(stackPath) }
        return route
    }
    getXLRowAdderPath(): string {
        return this.getPathRoot() + '-xlrow-adder/:GONG__StackPath'
    }
    getXLRowAdderRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getXLRowAdderPath(), component: XLRowDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getXLRowAdderForUsePath(): string {
        return this.getPathRoot() + '-xlrow-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath'
    }
    getXLRowAdderForUseRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getXLRowAdderForUsePath(), component: XLRowDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getXLRowDetailPath(): string {
        return this.getPathRoot() + '-xlrow-detail/:id/:GONG__StackPath'
    }
    getXLRowDetailRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getXLRowDetailPath(), component: XLRowDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }

    getXLSheetTablePath(): string {
        return this.getPathRoot() + '-xlsheets/:GONG__StackPath'
    }
    getXLSheetTableRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getXLSheetTablePath(), component: XLSheetsTableComponent, outlet: this.getTableOutlet(stackPath) }
        return route
    }
    getXLSheetAdderPath(): string {
        return this.getPathRoot() + '-xlsheet-adder/:GONG__StackPath'
    }
    getXLSheetAdderRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getXLSheetAdderPath(), component: XLSheetDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getXLSheetAdderForUsePath(): string {
        return this.getPathRoot() + '-xlsheet-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath'
    }
    getXLSheetAdderForUseRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getXLSheetAdderForUsePath(), component: XLSheetDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getXLSheetDetailPath(): string {
        return this.getPathRoot() + '-xlsheet-detail/:id/:GONG__StackPath'
    }
    getXLSheetDetailRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getXLSheetDetailPath(), component: XLSheetDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }



    addDataPanelRoutes(stackPath: string) {

        this.addRoutes([
            // insertion point for all routes getter
            this.getDisplaySelectionTableRoute(stackPath),
            this.getDisplaySelectionAdderRoute(stackPath),
            this.getDisplaySelectionAdderForUseRoute(stackPath),
            this.getDisplaySelectionDetailRoute(stackPath),

            this.getXLCellTableRoute(stackPath),
            this.getXLCellAdderRoute(stackPath),
            this.getXLCellAdderForUseRoute(stackPath),
            this.getXLCellDetailRoute(stackPath),

            this.getXLFileTableRoute(stackPath),
            this.getXLFileAdderRoute(stackPath),
            this.getXLFileAdderForUseRoute(stackPath),
            this.getXLFileDetailRoute(stackPath),

            this.getXLRowTableRoute(stackPath),
            this.getXLRowAdderRoute(stackPath),
            this.getXLRowAdderForUseRoute(stackPath),
            this.getXLRowDetailRoute(stackPath),

            this.getXLSheetTableRoute(stackPath),
            this.getXLSheetAdderRoute(stackPath),
            this.getXLSheetAdderForUseRoute(stackPath),
            this.getXLSheetDetailRoute(stackPath),

        ])
    }
}
