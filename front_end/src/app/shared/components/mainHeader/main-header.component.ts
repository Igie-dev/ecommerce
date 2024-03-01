import { Component } from '@angular/core';
import { NavComponent } from './components/nav/nav.component';
import { NavigationComponent } from './components/navigation/navigation.component';
import { SocialMediaIconsComponent } from '../social-media-icons/social-media-icons.component';
@Component({
  selector: 'app-main-header',
  standalone: true,
  imports: [NavComponent, NavigationComponent, SocialMediaIconsComponent],
  templateUrl: './main-header.component.html',
})
export class MainHeaderComponent {}
