import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { MakeComponent } from './make.component';

describe('MakeComponent', () => {
  let component: MakeComponent;
  let fixture: ComponentFixture<MakeComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ MakeComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(MakeComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
