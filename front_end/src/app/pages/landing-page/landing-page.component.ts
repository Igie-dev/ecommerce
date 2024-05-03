import { Component } from '@angular/core';
import { BillboardComponent } from './components/billboard/billboard.component';
import { NewProductsComponent } from './components/new-products/new-products.component';
@Component({
  selector: 'app-landing-page',
  standalone: true,
  imports: [BillboardComponent, NewProductsComponent],
  templateUrl: './landing-page.component.html',
})
export class LandingPageComponent {}
