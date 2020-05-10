import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

import { MakeComponent } from './public/make/make.component';
import { AddComponent } from './public/add/add.component';
import { HomeComponent } from './public/home/home.component';
import { WordsComponent } from './public/words/words.component';
import { ContactComponent } from './public/contact/contact.component';
import { DonateComponent } from './public/donate/donate.component';

import { LoginComponent } from './admin/login/login.component';
import { ApproveComponent } from './admin/approve/approve.component';

const routes: Routes = [
    //NOTE: default route
    {path:'', redirectTo:'', component:HomeComponent, pathMatch:'full'},

    //NOTE: public components
    {path:'make', component:MakeComponent}, 
    {path:'add', component:AddComponent}, 
    {path:'words', component:WordsComponent}, 
    {path:'contact', component:ContactComponent}, 
    {path:'donate', component:DonateComponent}, 

    //NOTE: admin components
    //{path:'admin', component:LoginComponent}, //NOTE: admins log in?
    {path:'admin', component:ApproveComponent}, //NOTE: once approved login, admins can approve of bats?
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
