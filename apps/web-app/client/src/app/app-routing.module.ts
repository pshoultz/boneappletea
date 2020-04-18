import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

import { BoneappleteaComponent } from './admin/boneappletea/boneappletea.component';
import { FrontComponent } from './public/front/front.component';

const routes: Routes = [
    { path: '', component: FrontComponent, pathMatch: 'full' },
    { path: 'admin', component: BoneappleteaComponent },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
