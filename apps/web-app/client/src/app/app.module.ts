import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { HttpClientModule } from '@angular/common/http';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { BoneappleteaComponent } from './boneappletea/boneappletea.component';
import { BoneappleteaCardComponent } from './boneappletea-card/boneappletea-card.component';

@NgModule({
  declarations: [
    AppComponent,
    BoneappleteaComponent,
    BoneappleteaCardComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    HttpClientModule,
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
