import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { NavigationComponent } from './navigation/navigation.component';
import { MaterialModule } from '../material/material.module';
import { AppRoutingModule } from '../app-routing.module';
import { SidenavigationComponent } from './sidenavigation/sidenavigation.component';

@NgModule({
  declarations: [
    NavigationComponent,
    SidenavigationComponent,
  ],
  imports: [
    CommonModule,
    MaterialModule,
    AppRoutingModule,
  ],
  exports:[
      NavigationComponent,
      SidenavigationComponent,
  ]
})
export class SharedModule { }
