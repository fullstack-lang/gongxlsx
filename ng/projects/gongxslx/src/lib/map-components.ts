// insertion point sub template for components imports 
  import { XLFilesTableComponent } from './xlfiles-table/xlfiles-table.component'
  import { XLFileSortingComponent } from './xlfile-sorting/xlfile-sorting.component'

// insertion point sub template for map of components per struct 
  export const MapOfXLFilesComponents: Map<string, any> = new Map([["XLFilesTableComponent", XLFilesTableComponent],])
  export const MapOfXLFileSortingComponents: Map<string, any> = new Map([["XLFileSortingComponent", XLFileSortingComponent],])

// map of all ng components of the stacks
export const MapOfComponents: Map<string, any> =
  new Map(
    [
      // insertion point sub template for map of components 
      ["XLFile", MapOfXLFilesComponents],
    ]
  )

// map of all ng components of the stacks
export const MapOfSortingComponents: Map<string, any> =
  new Map(
    [
    // insertion point sub template for map of sorting components 
      ["XLFile", MapOfXLFileSortingComponents],
    ]
  )
