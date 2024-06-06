import { Component, Input } from '@angular/core';
import { LucideAngularModule } from 'lucide-angular';
import { SlicePipe, CurrencyPipe } from '@angular/common';

@Component({
  selector: 'product-card',
  standalone: true,
  imports: [LucideAngularModule, SlicePipe, CurrencyPipe],
  templateUrl: './product-card.component.html',
})
export class ProductCardComponent {
  @Input({ required: true })
  product!: TProduct;
  constructor() {}
}
