import { NO_ERRORS_SCHEMA } from '@angular/core';
import { EdgeComponent } from './edge.component';
import { ComponentFixture, TestBed } from '@angular/core/testing';

describe('EdgeComponent', () => {

  let fixture: ComponentFixture<EdgeComponent>;
  let component: EdgeComponent;
  beforeEach(() => {
    TestBed.configureTestingModule({
      schemas: [NO_ERRORS_SCHEMA],
      providers: [
      ],
      declarations: [EdgeComponent]
    });

    fixture = TestBed.createComponent(EdgeComponent);
    component = fixture.componentInstance;

  });

  it('should be able to create component instance', () => {
    expect(component).toBeDefined();
  });

});
