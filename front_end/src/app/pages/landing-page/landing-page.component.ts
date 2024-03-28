import { Component } from '@angular/core';
import { CarouselComponent } from './components/carouselContainer/carousel.component';
import { NewProductsComponent } from './components/new-products/new-products.component';
@Component({
  selector: 'app-landing-page',
  standalone: true,
  imports: [CarouselComponent, NewProductsComponent],
  templateUrl: './landing-page.component.html',
})
export class LandingPageComponent {}
