import { Component } from '@angular/core';
import { NavComponent } from './components/nav/nav.component';
import { NavigationComponent } from './components/navigation/navigation.component';
import { LucideAngularModule } from 'lucide-angular';
@Component({
  selector: 'app-main-header',
  standalone: true,
  imports: [NavComponent, NavigationComponent, LucideAngularModule],
  templateUrl: './main-header.component.html',
})
export class MainHeaderComponent {}
