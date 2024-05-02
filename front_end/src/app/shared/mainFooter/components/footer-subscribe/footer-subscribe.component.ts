import { Component } from '@angular/core';
import { FormControl, ReactiveFormsModule } from '@angular/forms';
import { HlmButtonDirective } from '@spartan-ng/ui-button-helm';
@Component({
  selector: 'app-footer-subscribe',
  standalone: true,
  imports: [ReactiveFormsModule],
  templateUrl: './footer-subscribe.component.html',
})
export class FooterSubscribeComponent {
  email = new FormControl();

  handleSubmitSubscribe() {
    if (!this.email.value) return;
    console.log(this.email.value);
  }
}
