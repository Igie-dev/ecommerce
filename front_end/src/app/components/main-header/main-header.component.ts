import { Component } from '@angular/core';
import { NavComponent } from './components/nav/nav.component';
import { NavigationComponent } from './components/navigation/navigation.component';
import { HeaderNewSmgComponent } from './components/header-new-smg/header-new-smg.component';
@Component({
  selector: 'app-main-header',
  standalone: true,
  imports: [NavComponent, NavigationComponent, HeaderNewSmgComponent],
  templateUrl: './main-header.component.html',
})
export class MainHeaderComponent {}
