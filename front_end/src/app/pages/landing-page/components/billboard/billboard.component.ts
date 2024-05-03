import { Component } from '@angular/core';

import {
  HlmCarouselComponent,
  HlmCarouselContentComponent,
  HlmCarouselItemComponent,
  HlmCarouselNextComponent,
  HlmCarouselPreviousComponent,
} from '@spartan-ng/ui-carousel-helm';

@Component({
  selector: 'landing-page-billboard',
  standalone: true,
  imports: [
    HlmCarouselComponent,
    HlmCarouselContentComponent,
    HlmCarouselItemComponent,
    HlmCarouselNextComponent,
    HlmCarouselPreviousComponent,
  ],
  templateUrl: './billboard.component.html',
})
export class BillboardComponent {
  carouselImages = [
    { id: 1, src: 'assets/hero_carousel_1.png' },
    { id: 2, src: 'assets/hero_carousel_2.jpg' },
    { id: 3, src: 'assets/hero_carousel_3.jpg' },
  ];
}
