import { Component, ViewChild, ElementRef, HostListener } from '@angular/core';
import { LucideAngularModule } from 'lucide-angular';
@Component({
  selector: 'app-horizontal-scroll-container',
  standalone: true,
  imports: [LucideAngularModule],
  templateUrl: './horizontal-scroll-container.component.html',
  styleUrl: './horizontal-scroll-container.component.css',
})
export class HorizontalScrollContainerComponent {
  @HostListener('scroll', ['$event'])
  @ViewChild('scrollListWrapper')
  listWrapper!: ElementRef<HTMLUListElement>;
  @ViewChild('scrollArrowLeftContainer')
  arrowLeftContainer!: ElementRef<HTMLElement>;
  @ViewChild('scrollArrowRightContainer')
  arrowRightContainer!: ElementRef<HTMLElement>;

  handleButtonsHide() {
    if (this.listWrapper.nativeElement.scrollLeft >= 20) {
      this.arrowLeftContainer.nativeElement.classList.add('show_btn_container');
    } else {
      this.arrowLeftContainer.nativeElement.classList.remove(
        'show_btn_container'
      );
    }

    let maxScrollValue =
      this.listWrapper.nativeElement.scrollWidth -
      this.listWrapper.nativeElement.clientWidth -
      20;

    if (this.listWrapper.nativeElement.scrollLeft >= maxScrollValue) {
      this.arrowRightContainer.nativeElement.classList.remove(
        'show_btn_container'
      );
    } else {
      this.arrowRightContainer.nativeElement.classList.add(
        'show_btn_container'
      );
    }
  }

  handleLeftBtnClick() {
    this.listWrapper.nativeElement.scrollLeft -= 200;
  }

  handleRightBtnClick() {
    this.listWrapper.nativeElement.scrollLeft += 200;
  }

  onScroll(event: Event) {
    this.handleButtonsHide();
  }
}
