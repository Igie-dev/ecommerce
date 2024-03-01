import { Component } from '@angular/core';
import { LucideAngularModule } from 'lucide-angular';
import { HlmButtonDirective } from '@spartan-ng/ui-button-helm';

@Component({
  selector: 'app-social-media-icons',
  standalone: true,
  imports: [LucideAngularModule, HlmButtonDirective],
  templateUrl: './social-media-icons.component.html',
})
export class SocialMediaIconsComponent {}
