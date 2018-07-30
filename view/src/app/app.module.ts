import { NxGraphComponent } from './nx-graph/nx-graph.component';
import { NgxGraphModule } from '@swimlane/ngx-graph';
import { EdgeComponent } from './graph/edge/edge.component';
import { GraphComponent } from './graph/graph.component';
import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { HttpClientModule } from '@angular/common/http';
import { AppComponent } from './app.component';
import { VertexComponent } from './graph/vertex/vertex.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';

@NgModule({
  declarations: [
    AppComponent,
    GraphComponent,
    VertexComponent,
    EdgeComponent,
    NxGraphComponent
  ],
  imports: [
    BrowserModule,
    HttpClientModule,
    NgxGraphModule,
    BrowserAnimationsModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
