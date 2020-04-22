import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { MakeComponent } from './make/make.component';
import { AddComponent } from './add/add.component';
import { MaterialModule } from '../material/material.module';
import { FormsModule } from '@angular/forms';

@NgModule({
  declarations: [MakeComponent, AddComponent],
  imports: [
    CommonModule,
    MaterialModule,
    FormsModule,
  ]
})
export class PublicModule { }
