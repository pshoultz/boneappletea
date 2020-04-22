import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { LoginComponent } from './login/login.component';
import { ApproveComponent } from './approve/approve.component';



@NgModule({
  declarations: [LoginComponent, ApproveComponent],
  imports: [
    CommonModule
  ]
})
export class AdminModule { }
