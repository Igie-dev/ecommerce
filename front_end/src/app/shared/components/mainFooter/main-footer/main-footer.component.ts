import { Component } from '@angular/core';
import { FooterSubscribeComponent } from '../components/footer-subscribe/footer-subscribe.component';
import { SocialMediaIconsComponent } from '../../social-media-icons/social-media-icons.component';
@Component({
  selector: 'app-main-footer',
  standalone: true,
  imports: [FooterSubscribeComponent, SocialMediaIconsComponent],
  templateUrl: './main-footer.component.html',
})
export class MainFooterComponent {}
