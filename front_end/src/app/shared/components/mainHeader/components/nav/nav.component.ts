import { Component } from '@angular/core';
import { ProfileComponent } from './components/profile/profile.component';
import { CartBtnComponent } from './components/cart-btn/cart-btn.component';
import { SearchBtnComponent } from './components/search-btn/search-btn.component';
@Component({
  selector: 'app-nav-btn',
  standalone: true,
  imports: [ProfileComponent, CartBtnComponent, SearchBtnComponent],
  templateUrl: './nav.component.html',
})
export class NavComponent {}
