import { EdgeComponent } from './graph/edge/edge.component';
import { GraphComponent } from './graph/graph.component';
import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { HttpClientModule } from '@angular/common/http';
import { AppComponent } from './app.component';
import { VertexComponent } from './graph/vertex/vertex.component';

@NgModule({
  declarations: [
    AppComponent,
    GraphComponent,
    VertexComponent,
    EdgeComponent
  ],
  imports: [
    BrowserModule,
    HttpClientModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
