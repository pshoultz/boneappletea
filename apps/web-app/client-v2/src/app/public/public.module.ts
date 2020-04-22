import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { MakeComponent } from './make/make.component';
import { AddComponent } from './add/add.component';
import { MaterialModule } from '../material/material.module';

@NgModule({
  declarations: [MakeComponent, AddComponent],
  imports: [
    CommonModule,
    MaterialModule,
  ]
})
export class PublicModule { }
