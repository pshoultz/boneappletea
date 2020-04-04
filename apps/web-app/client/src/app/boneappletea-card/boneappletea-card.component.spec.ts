import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { BoneappleteaCardComponent } from './boneappletea-card.component';

describe('BoneappleteaCardComponent', () => {
  let component: BoneappleteaCardComponent;
  let fixture: ComponentFixture<BoneappleteaCardComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ BoneappleteaCardComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(BoneappleteaCardComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
