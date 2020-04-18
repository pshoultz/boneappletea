import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

import { BoneappleteaComponent } from './admin/boneappletea/boneappletea.component';

const routes: Routes = [
    { path: 'admin', component: BoneappleteaComponent },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
