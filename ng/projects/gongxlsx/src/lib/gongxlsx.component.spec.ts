import { ComponentFixture, TestBed } from '@angular/core/testing';

import { GongxlsxComponent } from './gongxlsx.component';

describe('GongxlsxComponent', () => {
  let component: GongxlsxComponent;
  let fixture: ComponentFixture<GongxlsxComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ GongxlsxComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(GongxlsxComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
