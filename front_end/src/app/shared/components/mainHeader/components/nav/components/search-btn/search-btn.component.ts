import { Component } from '@angular/core';
import { LucideAngularModule } from 'lucide-angular';

@Component({
  selector: 'app-header-search-btn',
  standalone: true,
  imports: [LucideAngularModule],
  templateUrl: './search-btn.component.html',
})
export class SearchBtnComponent {}
