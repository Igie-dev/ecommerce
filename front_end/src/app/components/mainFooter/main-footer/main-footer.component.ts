import { Component } from '@angular/core';
import { LucideAngularModule } from 'lucide-angular';
import { FooterSubscribeComponent } from '../components/footer-subscribe/footer-subscribe.component';
@Component({
  selector: 'app-main-footer',
  standalone: true,
  imports: [LucideAngularModule, FooterSubscribeComponent],
  templateUrl: './main-footer.component.html',
})
export class MainFooterComponent {}
