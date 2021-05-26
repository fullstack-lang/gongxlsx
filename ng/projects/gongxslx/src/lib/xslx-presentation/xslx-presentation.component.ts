import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup} from '@angular/forms';

import { XslxDB } from '../xslx-db'
import { XslxService } from '../xslx.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

export interface xslxDummyElement {
}

const ELEMENT_DATA: xslxDummyElement[] = [
];

@Component({
	selector: 'app-xslx-presentation',
	templateUrl: './xslx-presentation.component.html',
	styleUrls: ['./xslx-presentation.component.css'],
})
export class XslxPresentationComponent implements OnInit {

	// insertion point for declarations

	displayedColumns: string[] = [];
	dataSource = ELEMENT_DATA;

	xslx: XslxDB;
 
	constructor(
		private xslxService: XslxService,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getXslx();

		// observable for changes in 
		this.xslxService.XslxServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getXslx()
				}
			}
		)
	}

	getXslx(): void {
		const id = +this.route.snapshot.paramMap.get('id');
		this.xslxService.getXslx(id)
			.subscribe(
				xslx => {
					this.xslx = xslx

					// insertion point for recovery of durations

				}
			);
	}

	// set presentation outlet
	setPresentationRouterOutlet(structName: string, ID: number) {
		this.router.navigate([{
			outlets: {
				presentation: [structName + "-presentation", ID]
			}
		}]);
	}

	// set editor outlet
	setEditorRouterOutlet(ID: number) {
		this.router.navigate([{
			outlets: {
				editor: ["xslx-detail", ID]
			}
		}]);
	}
}
