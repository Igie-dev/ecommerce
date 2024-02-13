import { Component } from '@angular/core';
import { NgIf } from '@angular/common';

@Component({
  selector: 'app-header-new-smg',
  standalone: true,
  imports: [NgIf],
  templateUrl: './header-new-smg.component.html',
})
export class HeaderNewSmgComponent {
  message = '';
}
