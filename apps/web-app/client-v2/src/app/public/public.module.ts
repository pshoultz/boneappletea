import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { MakeComponent } from './make/make.component';
import { AddComponent } from './add/add.component';

@NgModule({
  declarations: [MakeComponent, AddComponent],
  imports: [
    CommonModule
  ]
})
export class PublicModule { }
