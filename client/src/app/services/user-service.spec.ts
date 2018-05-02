import { TestBed, inject } from '@angular/core/testing';

import { UserService.TsService } from './user-service.ts.service';

describe('UserService.TsService', () => {
  beforeEach(() => {
    TestBed.configureTestingModule({
      providers: [UserService.TsService]
    });
  });

  it('should be created', inject([UserService.TsService], (service: UserService.TsService) => {
    expect(service).toBeTruthy();
  }));
});
