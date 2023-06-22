import { ComponentFixture, TestBed } from '@angular/core/testing';

import { GongxlsxdatamodelComponent } from './gongxlsxdatamodel.component';

describe('GongxlsxdatamodelComponent', () => {
  let component: GongxlsxdatamodelComponent;
  let fixture: ComponentFixture<GongxlsxdatamodelComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ GongxlsxdatamodelComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(GongxlsxdatamodelComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
