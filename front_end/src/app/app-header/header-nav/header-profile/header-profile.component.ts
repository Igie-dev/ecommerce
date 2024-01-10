import { Component, OnInit } from '@angular/core';
@Component({
  selector: 'app-header-profile',
  templateUrl: './header-profile.component.html',
  styleUrls: ['./header-profile.component.css'],
})
export class HeaderProfileComponent implements OnInit {
  constructor() {}
  profile_img_src =
    'http://localhost:4200/assets/414103125_1440158746916029_2014991615104848722_n.jpg';
  ngOnInit() {}
}
