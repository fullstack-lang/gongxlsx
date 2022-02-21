import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';

import { DisplaySelectionDB } from '../displayselection-db'
import { DisplaySelectionService } from '../displayselection.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

export interface displayselectionDummyElement {
}

const ELEMENT_DATA: displayselectionDummyElement[] = [
];

@Component({
	selector: 'app-displayselection-presentation',
	templateUrl: './displayselection-presentation.component.html',
	styleUrls: ['./displayselection-presentation.component.css'],
})
export class DisplaySelectionPresentationComponent implements OnInit {

	// insertion point for declarations

	displayedColumns: string[] = []
	dataSource = ELEMENT_DATA

	displayselection: DisplaySelectionDB = new (DisplaySelectionDB)

	// front repo
	frontRepo: FrontRepo = new (FrontRepo)
 
	constructor(
		private displayselectionService: DisplaySelectionService,
		private frontRepoService: FrontRepoService,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getDisplaySelection();

		// observable for changes in 
		this.displayselectionService.DisplaySelectionServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getDisplaySelection()
				}
			}
		)
	}

	getDisplaySelection(): void {
		const id = +this.route.snapshot.paramMap.get('id')!
		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				this.displayselection = this.frontRepo.DisplaySelections.get(id)!

				// insertion point for recovery of durations
			}
		);
	}

	// set presentation outlet
	setPresentationRouterOutlet(structName: string, ID: number) {
		this.router.navigate([{
			outlets: {
				github_com_fullstack_lang_gongxlsx_go_presentation: ["github_com_fullstack_lang_gongxlsx_go-" + structName + "-presentation", ID]
			}
		}]);
	}

	// set editor outlet
	setEditorRouterOutlet(ID: number) {
		this.router.navigate([{
			outlets: {
				github_com_fullstack_lang_gongxlsx_go_editor: ["github_com_fullstack_lang_gongxlsx_go-" + "displayselection-detail", ID]
			}
		}]);
	}
}
