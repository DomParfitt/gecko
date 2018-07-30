import { Component, OnInit, Input } from '@angular/core';

@Component({
  // tslint:disable-next-line:component-selector
  selector: 'svg[app-edge]',
  templateUrl: './edge.component.html',
  styleUrls: ['./edge.component.scss']
})

export class EdgeComponent implements OnInit {

  @Input()
  label = 'edge';
  from: string;
  to: string;

  @Input()
  x = 0;

  @Input()
  y = 0;

  constructor() {
  }

  ngOnInit() {

  }
}
