// generated by gong
import { Component, OnInit, AfterViewInit, ViewChild, Inject, Optional, Input } from '@angular/core';
import { BehaviorSubject } from 'rxjs'
import { MatSort } from '@angular/material/sort';
import { MatPaginator } from '@angular/material/paginator';
import { MatTableDataSource } from '@angular/material/table';
import { MatButton } from '@angular/material/button'

import { MatDialogRef, MAT_DIALOG_DATA, MatDialog } from '@angular/material/dialog'
import { DialogData, FrontRepoService, FrontRepo, SelectionMode } from '../front-repo.service'
import { NullInt64 } from '../null-int64'
import { SelectionModel } from '@angular/cdk/collections';

const allowMultiSelect = true;

import { ActivatedRoute, Router, RouterState } from '@angular/router';
import { LinkDB } from '../link-db'
import { LinkService } from '../link.service'

// insertion point for additional imports

import { RouteService } from '../route-service';

// TableComponent is initilizaed from different routes
// TableComponentMode detail different cases 
enum TableComponentMode {
  DISPLAY_MODE,
  ONE_MANY_ASSOCIATION_MODE,
  MANY_MANY_ASSOCIATION_MODE,
}

// generated table component
@Component({
  selector: 'app-linkstable',
  templateUrl: './links-table.component.html',
  styleUrls: ['./links-table.component.css'],
})
export class LinksTableComponent implements OnInit {

  @Input() GONG__StackPath: string = ""

  // mode at invocation
  mode: TableComponentMode = TableComponentMode.DISPLAY_MODE

  // used if the component is called as a selection component of Link instances
  selection: SelectionModel<LinkDB> = new (SelectionModel)
  initialSelection = new Array<LinkDB>()

  // the data source for the table
  links: LinkDB[] = []
  matTableDataSource: MatTableDataSource<LinkDB> = new (MatTableDataSource)

  // front repo, that will be referenced by this.links
  frontRepo: FrontRepo = new (FrontRepo)

  // displayedColumns is referenced by the MatTable component for specify what columns
  // have to be displayed and in what order
  displayedColumns: string[];

  // for sorting & pagination
  @ViewChild(MatSort)
  sort: MatSort | undefined
  @ViewChild(MatPaginator)
  paginator: MatPaginator | undefined;

  ngAfterViewInit() {

    // enable sorting on all fields (including pointers and reverse pointer)
    this.matTableDataSource.sortingDataAccessor = (linkDB: LinkDB, property: string) => {
      switch (property) {
        case 'ID':
          return linkDB.ID

        // insertion point for specific sorting accessor
        case 'Name':
          return linkDB.Name;

        case 'Identifier':
          return linkDB.Identifier;

        case 'Fieldtypename':
          return linkDB.Fieldtypename;

        case 'TargetMultiplicity':
          return linkDB.TargetMultiplicity;

        case 'SourceMultiplicity':
          return linkDB.SourceMultiplicity;

        case 'Middlevertice':
          return (linkDB.Middlevertice ? linkDB.Middlevertice.Name : '');

        case 'GongStructShape_Links':
          if (this.frontRepo.GongStructShapes.get(linkDB.GongStructShape_LinksDBID.Int64) != undefined) {
            return this.frontRepo.GongStructShapes.get(linkDB.GongStructShape_LinksDBID.Int64)!.Name
          } else {
            return ""
          }

        default:
          console.assert(false, "Unknown field")
          return "";
      }
    };

    // enable filtering on all fields (including pointers and reverse pointer, which is not done by default)
    this.matTableDataSource.filterPredicate = (linkDB: LinkDB, filter: string) => {

      // filtering is based on finding a lower case filter into a concatenated string
      // the linkDB properties
      let mergedContent = ""

      // insertion point for merging of fields
      mergedContent += linkDB.Name.toLowerCase()
      mergedContent += linkDB.Identifier.toLowerCase()
      mergedContent += linkDB.Fieldtypename.toLowerCase()
      mergedContent += linkDB.TargetMultiplicity.toLowerCase()
      mergedContent += linkDB.SourceMultiplicity.toLowerCase()
      if (linkDB.Middlevertice) {
        mergedContent += linkDB.Middlevertice.Name.toLowerCase()
      }
      if (linkDB.GongStructShape_LinksDBID.Int64 != 0) {
        mergedContent += this.frontRepo.GongStructShapes.get(linkDB.GongStructShape_LinksDBID.Int64)!.Name.toLowerCase()
      }


      let isSelected = mergedContent.includes(filter.toLowerCase())
      return isSelected
    };

    this.matTableDataSource.sort = this.sort!
    this.matTableDataSource.paginator = this.paginator!
  }

  applyFilter(event: Event) {
    const filterValue = (event.target as HTMLInputElement).value;
    this.matTableDataSource.filter = filterValue.trim().toLowerCase();
  }

  constructor(
    private linkService: LinkService,
    private frontRepoService: FrontRepoService,

    // not null if the component is called as a selection component of link instances
    public dialogRef: MatDialogRef<LinksTableComponent>,
    @Optional() @Inject(MAT_DIALOG_DATA) public dialogData: DialogData,

    private router: Router,
    private activatedRoute: ActivatedRoute,

    private routeService: RouteService,
  ) {

    // compute mode
    if (dialogData == undefined) {
      this.mode = TableComponentMode.DISPLAY_MODE
    } else {
      this.GONG__StackPath = dialogData.GONG__StackPath
      switch (dialogData.SelectionMode) {
        case SelectionMode.ONE_MANY_ASSOCIATION_MODE:
          this.mode = TableComponentMode.ONE_MANY_ASSOCIATION_MODE
          break
        case SelectionMode.MANY_MANY_ASSOCIATION_MODE:
          this.mode = TableComponentMode.MANY_MANY_ASSOCIATION_MODE
          break
        default:
      }
    }

    // observable for changes in structs
    this.linkService.LinkServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.getLinks()
        }
      }
    )
    if (this.mode == TableComponentMode.DISPLAY_MODE) {
      this.displayedColumns = ['ID', 'Delete', // insertion point for columns to display
        "Name",
        "Identifier",
        "Fieldtypename",
        "TargetMultiplicity",
        "SourceMultiplicity",
        "Middlevertice",
        "GongStructShape_Links",
      ]
    } else {
      this.displayedColumns = ['select', 'ID', // insertion point for columns to display
        "Name",
        "Identifier",
        "Fieldtypename",
        "TargetMultiplicity",
        "SourceMultiplicity",
        "Middlevertice",
        "GongStructShape_Links",
      ]
      this.selection = new SelectionModel<LinkDB>(allowMultiSelect, this.initialSelection);
    }

  }

  ngOnInit(): void {
    let stackPath = this.activatedRoute.snapshot.paramMap.get('GONG__StackPath')
    if (stackPath != undefined) {
      this.GONG__StackPath = stackPath
    }

    this.getLinks()

    this.matTableDataSource = new MatTableDataSource(this.links)
  }

  getLinks(): void {
    this.frontRepoService.pull(this.GONG__StackPath).subscribe(
      frontRepo => {
        this.frontRepo = frontRepo

        this.links = this.frontRepo.Links_array;

        // insertion point for time duration Recoveries
        // insertion point for enum int Recoveries

        // in case the component is called as a selection component
        if (this.mode == TableComponentMode.ONE_MANY_ASSOCIATION_MODE) {
          for (let link of this.links) {
            let ID = this.dialogData.ID
            let revPointer = link[this.dialogData.ReversePointer as keyof LinkDB] as unknown as NullInt64
            if (revPointer.Int64 == ID) {
              this.initialSelection.push(link)
            }
            this.selection = new SelectionModel<LinkDB>(allowMultiSelect, this.initialSelection);
          }
        }

        if (this.mode == TableComponentMode.MANY_MANY_ASSOCIATION_MODE) {

          let mapOfSourceInstances = this.frontRepo[this.dialogData.SourceStruct + "s" as keyof FrontRepo] as Map<number, LinkDB>
          let sourceInstance = mapOfSourceInstances.get(this.dialogData.ID)!

          // we associates on sourceInstance of type SourceStruct with a MANY MANY associations to LinkDB
          // the field name is sourceField
          let sourceFieldArray = sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance]! as unknown as LinkDB[]
          if (sourceFieldArray != null) {
            for (let associationInstance of sourceFieldArray) {
              let link = associationInstance[this.dialogData.IntermediateStructField as keyof typeof associationInstance] as unknown as LinkDB
              this.initialSelection.push(link)
            }
          }

          this.selection = new SelectionModel<LinkDB>(allowMultiSelect, this.initialSelection);
        }

        // update the mat table data source
        this.matTableDataSource.data = this.links
      }
    )
  }

  // newLink initiate a new link
  // create a new Link objet
  newLink() {
  }

  deleteLink(linkID: number, link: LinkDB) {
    // list of links is truncated of link before the delete
    this.links = this.links.filter(h => h !== link);

    this.linkService.deleteLink(linkID, this.GONG__StackPath).subscribe(
      link => {
        this.linkService.LinkServiceChanged.next("delete")
      }
    );
  }

  editLink(linkID: number, link: LinkDB) {

  }

  // set editor outlet
  setEditorRouterOutlet(linkID: number) {
    let outletName = this.routeService.getEditorOutlet(this.GONG__StackPath)
    let fullPath = this.routeService.getPathRoot() + "-" + "link" + "-detail"

    let outletConf: any = {}
    outletConf[outletName] = [fullPath, linkID, this.GONG__StackPath]

    this.router.navigate([{ outlets: outletConf }])
  }

  /** Whether the number of selected elements matches the total number of rows. */
  isAllSelected() {
    const numSelected = this.selection.selected.length;
    const numRows = this.links.length;
    return numSelected === numRows;
  }

  /** Selects all rows if they are not all selected; otherwise clear selection. */
  masterToggle() {
    this.isAllSelected() ?
      this.selection.clear() :
      this.links.forEach(row => this.selection.select(row));
  }

  save() {

    if (this.mode == TableComponentMode.ONE_MANY_ASSOCIATION_MODE) {

      let toUpdate = new Set<LinkDB>()

      // reset all initial selection of link that belong to link
      for (let link of this.initialSelection) {
        let index = link[this.dialogData.ReversePointer as keyof LinkDB] as unknown as NullInt64
        index.Int64 = 0
        index.Valid = true
        toUpdate.add(link)

      }

      // from selection, set link that belong to link
      for (let link of this.selection.selected) {
        let ID = this.dialogData.ID as number
        let reversePointer = link[this.dialogData.ReversePointer as keyof LinkDB] as unknown as NullInt64
        reversePointer.Int64 = ID
        reversePointer.Valid = true
        toUpdate.add(link)
      }


      // update all link (only update selection & initial selection)
      for (let link of toUpdate) {
        this.linkService.updateLink(link, this.GONG__StackPath)
          .subscribe(link => {
            this.linkService.LinkServiceChanged.next("update")
          });
      }
    }

    if (this.mode == TableComponentMode.MANY_MANY_ASSOCIATION_MODE) {

      // get the source instance via the map of instances in the front repo
      let mapOfSourceInstances = this.frontRepo[this.dialogData.SourceStruct + "s" as keyof FrontRepo] as Map<number, LinkDB>
      let sourceInstance = mapOfSourceInstances.get(this.dialogData.ID)!

      // First, parse all instance of the association struct and remove the instance
      // that have unselect
      let unselectedLink = new Set<number>()
      for (let link of this.initialSelection) {
        if (this.selection.selected.includes(link)) {
          // console.log("link " + link.Name + " is still selected")
        } else {
          console.log("link " + link.Name + " has been unselected")
          unselectedLink.add(link.ID)
          console.log("is unselected " + unselectedLink.has(link.ID))
        }
      }

      // delete the association instance
      let associationInstance = sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance]
      let link = associationInstance![this.dialogData.IntermediateStructField as keyof typeof associationInstance] as unknown as LinkDB
      if (unselectedLink.has(link.ID)) {
        this.frontRepoService.deleteService(this.dialogData.IntermediateStruct, associationInstance)


      }

      // is the source array is empty create it
      if (sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance] == undefined) {
        (sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance] as unknown as Array<LinkDB>) = new Array<LinkDB>()
      }

      // second, parse all instance of the selected
      if (sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance]) {
        this.selection.selected.forEach(
          link => {
            if (!this.initialSelection.includes(link)) {
              // console.log("link " + link.Name + " has been added to the selection")

              let associationInstance = {
                Name: sourceInstance["Name"] + "-" + link.Name,
              }

              let index = associationInstance[this.dialogData.IntermediateStructField + "ID" as keyof typeof associationInstance] as unknown as NullInt64
              index.Int64 = link.ID
              index.Valid = true

              let indexDB = associationInstance[this.dialogData.IntermediateStructField + "DBID" as keyof typeof associationInstance] as unknown as NullInt64
              indexDB.Int64 = link.ID
              index.Valid = true

              this.frontRepoService.postService(this.dialogData.IntermediateStruct, associationInstance)

            } else {
              // console.log("link " + link.Name + " is still selected")
            }
          }
        )
      }

      // this.selection = new SelectionModel<LinkDB>(allowMultiSelect, this.initialSelection);
    }

    // why pizza ?
    this.dialogRef.close('Pizza!');
  }
}