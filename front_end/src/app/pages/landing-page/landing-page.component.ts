import { Component } from '@angular/core';
import { CarouselComponent } from './components/carousel/carousel.component';
@Component({
  selector: 'app-landing-page',
  standalone: true,
  imports: [CarouselComponent],
  templateUrl: './landing-page.component.html',
  styleUrl: './landing-page.component.css',
})
export class LandingPageComponent {}
