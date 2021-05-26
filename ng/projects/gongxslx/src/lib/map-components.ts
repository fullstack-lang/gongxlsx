// insertion point sub template for components imports 
  import { XLFilesTableComponent } from './xlfiles-table/xlfiles-table.component'
  import { XLFileSortingComponent } from './xlfile-sorting/xlfile-sorting.component'
  import { XLSheetsTableComponent } from './xlsheets-table/xlsheets-table.component'
  import { XLSheetSortingComponent } from './xlsheet-sorting/xlsheet-sorting.component'

// insertion point sub template for map of components per struct 
  export const MapOfXLFilesComponents: Map<string, any> = new Map([["XLFilesTableComponent", XLFilesTableComponent],])
  export const MapOfXLFileSortingComponents: Map<string, any> = new Map([["XLFileSortingComponent", XLFileSortingComponent],])
  export const MapOfXLSheetsComponents: Map<string, any> = new Map([["XLSheetsTableComponent", XLSheetsTableComponent],])
  export const MapOfXLSheetSortingComponents: Map<string, any> = new Map([["XLSheetSortingComponent", XLSheetSortingComponent],])

// map of all ng components of the stacks
export const MapOfComponents: Map<string, any> =
  new Map(
    [
      // insertion point sub template for map of components 
      ["XLFile", MapOfXLFilesComponents],
      ["XLSheet", MapOfXLSheetsComponents],
    ]
  )

// map of all ng components of the stacks
export const MapOfSortingComponents: Map<string, any> =
  new Map(
    [
    // insertion point sub template for map of sorting components 
      ["XLFile", MapOfXLFileSortingComponents],
      ["XLSheet", MapOfXLSheetSortingComponents],
    ]
  )
