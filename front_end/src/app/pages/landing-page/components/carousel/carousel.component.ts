import { Component } from '@angular/core';

import {
  HlmCarouselComponent,
  HlmCarouselContentComponent,
  HlmCarouselItemComponent,
  HlmCarouselNextComponent,
  HlmCarouselPreviousComponent,
} from '@spartan-ng/ui-carousel-helm';

@Component({
  selector: 'landing-page-carousel',
  standalone: true,
  imports: [
    HlmCarouselComponent,
    HlmCarouselContentComponent,
    HlmCarouselItemComponent,
    HlmCarouselNextComponent,
    HlmCarouselPreviousComponent,
  ],
  templateUrl: './carousel.component.html',
})
export class CarouselComponent {
  carouselImages = [
    { id: 1, src: 'assets/hero_carousel_1.png' },
    { id: 2, src: 'assets/hero_carousel_2.jpg' },
    { id: 3, src: 'assets/hero_carousel_3.jpg' },
  ];
}
