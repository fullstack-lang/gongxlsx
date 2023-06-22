import { TestBed } from '@angular/core/testing';

import { GongxlsxdatamodelService } from './gongxlsxdatamodel.service';

describe('GongxlsxdatamodelService', () => {
  let service: GongxlsxdatamodelService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(GongxlsxdatamodelService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
