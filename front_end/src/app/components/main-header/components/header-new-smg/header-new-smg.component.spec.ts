import { ComponentFixture, TestBed } from '@angular/core/testing';

import { HeaderNewSmgComponent } from './header-new-smg.component';

describe('HeaderNewSmgComponent', () => {
  let component: HeaderNewSmgComponent;
  let fixture: ComponentFixture<HeaderNewSmgComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [HeaderNewSmgComponent]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(HeaderNewSmgComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
