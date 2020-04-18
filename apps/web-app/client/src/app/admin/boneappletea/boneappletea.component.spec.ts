import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { BoneappleteaComponent } from './boneappletea.component';

describe('BoneappleteaComponent', () => {
  let component: BoneappleteaComponent;
  let fixture: ComponentFixture<BoneappleteaComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ BoneappleteaComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(BoneappleteaComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
