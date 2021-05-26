// insertion point sub template for components imports 
  import { XslxsTableComponent } from './xslxs-table/xslxs-table.component'
  import { XslxSortingComponent } from './xslx-sorting/xslx-sorting.component'

// insertion point sub template for map of components per struct 
  export const MapOfXslxsComponents: Map<string, any> = new Map([["XslxsTableComponent", XslxsTableComponent],])
  export const MapOfXslxSortingComponents: Map<string, any> = new Map([["XslxSortingComponent", XslxSortingComponent],])

// map of all ng components of the stacks
export const MapOfComponents: Map<string, any> =
  new Map(
    [
      // insertion point sub template for map of components 
      ["Xslx", MapOfXslxsComponents],
    ]
  )

// map of all ng components of the stacks
export const MapOfSortingComponents: Map<string, any> =
  new Map(
    [
    // insertion point sub template for map of sorting components 
      ["Xslx", MapOfXslxSortingComponents],
    ]
  )
