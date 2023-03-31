import { TestBed } from '@angular/core/testing';

import { GongxlsxService } from './gongxlsx.service';

describe('GongxlsxService', () => {
  let service: GongxlsxService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(GongxlsxService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
