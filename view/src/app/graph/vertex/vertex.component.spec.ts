import { NO_ERRORS_SCHEMA } from "@angular/core";
import { VertexComponent } from "./vertex.component";
import { ComponentFixture, TestBed } from "@angular/core/testing";

describe("VertexComponent", () => {

  let fixture: ComponentFixture<VertexComponent>;
  let component: VertexComponent;
  beforeEach(() => {
    TestBed.configureTestingModule({
      schemas: [NO_ERRORS_SCHEMA],
      providers: [
      ],
      declarations: [VertexComponent]
    });

    fixture = TestBed.createComponent(VertexComponent);
    component = fixture.componentInstance;

  });

  it("should be able to create component instance", () => {
    expect(component).toBeDefined();
  });
  
});
