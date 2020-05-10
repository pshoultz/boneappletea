import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { LoginComponent } from './login/login.component';
import { ApproveComponent } from './approve/approve.component';
import { MaterialModule } from '../material/material.module';

@NgModule({
  declarations: [LoginComponent, ApproveComponent],
  imports: [
    CommonModule,
    MaterialModule,
  ]
})
export class AdminModule { }
