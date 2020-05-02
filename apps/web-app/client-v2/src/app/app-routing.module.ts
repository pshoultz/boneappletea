import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

import { MakeComponent } from './public/make/make.component';
import { AddComponent } from './public/add/add.component';

import { LoginComponent } from './admin/login/login.component';
import { ApproveComponent } from './admin/approve/approve.component';


const routes: Routes = [
    {path:'', redirectTo:'', pathMatch:'full'},
    {path:'make', component:MakeComponent}, 
    {path:'add', component:AddComponent}, 
    //{path:'admin', component:LoginComponent}, //NOTE: admins log in?
    {path:'admin', component:ApproveComponent}, //NOTE: once approved login, admins can approve of bats?
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
