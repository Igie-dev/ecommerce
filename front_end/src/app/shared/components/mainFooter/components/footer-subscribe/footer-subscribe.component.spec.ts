import { ComponentFixture, TestBed } from '@angular/core/testing';

import { FooterSubscribeComponent } from './footer-subscribe.component';

describe('FooterSubscribeComponent', () => {
  let component: FooterSubscribeComponent;
  let fixture: ComponentFixture<FooterSubscribeComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [FooterSubscribeComponent]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(FooterSubscribeComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
