import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { BoneappleteaComponent } from './boneappletea/boneappletea.component';
import { BoneappleteaCardComponent } from './boneappletea-card/boneappletea-card.component';

@NgModule({
  declarations: [
      BoneappleteaComponent,
      BoneappleteaCardComponent
  ],
  imports: [
    CommonModule,
  ],
  exports:[
      BoneappleteaComponent,
      BoneappleteaCardComponent
  ],
})
export class AdminModule { }
