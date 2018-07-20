import { TestBed, inject } from '@angular/core/testing';

import { AutomataService } from './automata.service';

describe('AutomataService', () => {
  beforeEach(() => {
    TestBed.configureTestingModule({
      providers: [AutomataService]
    });
  });

  it('should be created', inject([AutomataService], (service: AutomataService) => {
    expect(service).toBeTruthy();
  }));
});
