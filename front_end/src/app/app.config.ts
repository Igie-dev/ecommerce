import { ApplicationConfig, importProvidersFrom } from '@angular/core';
import { provideRouter } from '@angular/router';
import { routes } from './app.routes';
import { provideClientHydration } from '@angular/platform-browser';

import {
  LucideAngularModule,
  ShoppingCart,
  Search,
  Facebook,
  Instagram,
  Menu,
  Phone,
  Check,
  ChevronLeft,
  ChevronRight,
} from 'lucide-angular';

export const appConfig: ApplicationConfig = {
  providers: [
    provideRouter(routes),
    provideClientHydration(),
    importProvidersFrom(
      LucideAngularModule.pick({
        ShoppingCart,
        Search,
        Facebook,
        Instagram,
        Menu,
        Phone,
        Check,
        ChevronLeft,
        ChevronRight,
      })
    ),
  ],
};
