import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { MakeComponent } from './make/make.component';
import { AddComponent } from './add/add.component';
import { MaterialModule } from '../material/material.module';
import { FormsModule } from '@angular/forms';
import { HomeComponent } from './home/home.component';
import { WordsComponent } from './words/words.component';

@NgModule({
  declarations: [MakeComponent, AddComponent, HomeComponent, WordsComponent],
  imports: [
    CommonModule,
    MaterialModule,
    FormsModule,
  ]
})
export class PublicModule { }
