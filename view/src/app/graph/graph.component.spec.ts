import { NO_ERRORS_SCHEMA } from "@angular/core";
import { GraphComponent } from "./graph.component";
import { ComponentFixture, TestBed } from "@angular/core/testing";

describe("GraphComponent", () => {

  let fixture: ComponentFixture<GraphComponent>;
  let component: GraphComponent;
  beforeEach(() => {
    TestBed.configureTestingModule({
      schemas: [NO_ERRORS_SCHEMA],
      providers: [
      ],
      declarations: [GraphComponent]
    });

    fixture = TestBed.createComponent(GraphComponent);
    component = fixture.componentInstance;

  });

  it("should be able to create component instance", () => {
    expect(component).toBeDefined();
  });
  
});
