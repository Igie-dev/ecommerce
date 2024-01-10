import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { LucideAngularModule, ShoppingCart, Search } from 'lucide-angular';
import { AppRoutingModule } from './app-routing.module';
import { NgOptimizedImage } from '@angular/common';
import { AppComponent } from './app.component';
import { AppHeaderComponent } from './app-header/app-header.component';
import { HeaderNavComponent } from './app-header/header-nav/header-nav.component';
import { NavigationComponent } from './app-header/navigation/navigation.component';
import { HeaderCartComponent } from './app-header/header-nav/header-cart/header-cart.component';
import { HeaderSearchComponent } from './app-header/header-nav/header-search/header-search.component';
import { HeaderProfileComponent } from './app-header/header-nav/header-profile/header-profile.component';
@NgModule({
  declarations: [
    AppComponent,
    AppHeaderComponent,
    NavigationComponent,
    HeaderNavComponent,
    HeaderCartComponent,
    HeaderSearchComponent,
    HeaderProfileComponent,
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    NgOptimizedImage,
    LucideAngularModule.pick({ ShoppingCart, Search }),
  ],
  providers: [],
  bootstrap: [AppComponent],
})
export class AppModule {}
