import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { RouterOutlet } from '@angular/router';
import { MainHeaderComponent } from './components/mainHeader/main-header.component';
import { MainFooterComponent } from './components/mainFooter/main-footer/main-footer.component';
@Component({
  selector: 'app-root',
  standalone: true,
  imports: [
    CommonModule,
    RouterOutlet,
    MainHeaderComponent,
    MainFooterComponent,
  ],
  templateUrl: './app.component.html',
  styleUrl: './app.component.css',
})
export class AppComponent {
  title = 'front_end';
}
