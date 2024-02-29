import { Component } from '@angular/core';
import { LucideAngularModule } from 'lucide-angular';

@Component({
  selector: 'app-header-cart-btn',
  standalone: true,
  imports: [LucideAngularModule],
  templateUrl: './cart-btn.component.html',
})
export class CartBtnComponent {}
