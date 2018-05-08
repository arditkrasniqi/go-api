import { TestBed, inject } from '@angular/core/testing';

import { UserService.TsService } from './service.ts.service';

describe('Service.TsService', () => {
  beforeEach(() => {
    TestBed.configureTestingModule({
      providers: [UserService.TsService]
    });
  });

  it('should be created', inject([UserService.TsService], (service: UserService.TsService) => {
    expect(service).toBeTruthy();
  }));
});
