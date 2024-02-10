import { Routes } from '@angular/router';

export const routes: Routes = [
  {
    path: '',
    loadComponent: () =>
      import('./pages/landing-page/landing-page.component').then(
        (c) => c.LandingPageComponent
      ),
  },
  {
    path: '',
    pathMatch: 'full',
    loadComponent: () =>
      import('./pages/landing-page/landing-page.component').then(
        (c) => c.LandingPageComponent
      ),
  },
  {
    path: '**',
    loadComponent: () =>
      import('./pages/notfoundpage/notfoundpage.component').then(
        (mod) => mod.NotfoundpageComponent
      ),
  },
];
