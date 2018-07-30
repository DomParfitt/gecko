import { NO_ERRORS_SCHEMA } from "@angular/core";
import { NxGraphComponent } from "./nx-graph.component";
import { ComponentFixture, TestBed } from "@angular/core/testing";

describe("NxGraphComponent", () => {

  let fixture: ComponentFixture<NxGraphComponent>;
  let component: NxGraphComponent;
  beforeEach(() => {
    TestBed.configureTestingModule({
      schemas: [NO_ERRORS_SCHEMA],
      providers: [
      ],
      declarations: [NxGraphComponent]
    });

    fixture = TestBed.createComponent(NxGraphComponent);
    component = fixture.componentInstance;

  });

  it("should be able to create component instance", () => {
    expect(component).toBeDefined();
  });
  
});
