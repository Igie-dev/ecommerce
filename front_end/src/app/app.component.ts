import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { RouterOutlet } from '@angular/router';
import { MainHeaderComponent } from './shared/components/mainHeader/main-header.component';
import { MainFooterComponent } from './shared/components/mainFooter/main-footer/main-footer.component';
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
})
export class AppComponent {
  title = 'front_end';
}
