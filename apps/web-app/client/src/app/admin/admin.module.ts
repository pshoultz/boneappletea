import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { BoneappleteaComponent } from './boneappletea/boneappletea.component';
import { BoneappleteaCardComponent } from './boneappletea-card/boneappletea-card.component';
import { LoginComponent } from './login/login.component';

@NgModule({
  declarations: [
      BoneappleteaComponent,
      BoneappleteaCardComponent,
      LoginComponent
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
