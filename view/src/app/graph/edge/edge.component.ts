import { Component, OnInit } from '@angular/core';

@Component({
  // tslint:disable-next-line:component-selector
  selector: 'svg[app-edge]',
  templateUrl: './edge.component.html',
  styleUrls: ['./edge.component.scss']
})

export class EdgeComponent implements OnInit {

  label = 'edge';
  from: string;
  to: string;

  constructor() {
  }

  ngOnInit() {

  }
}
