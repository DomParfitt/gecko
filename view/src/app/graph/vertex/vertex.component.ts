import { Component, OnInit, Input } from '@angular/core';

@Component({
  // tslint:disable-next-line:component-selector
  selector: 'svg[app-vertex]',
  templateUrl: './vertex.component.html',
  styleUrls: ['./vertex.component.scss']
})

export class VertexComponent implements OnInit {

  label = '0';
  isTerminal = false;

  @Input()
  x = 0;

  @Input()
  y = 0;

  constructor() {
  }

  ngOnInit() {

  }
}