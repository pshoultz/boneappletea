import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-sidenavigation',
  templateUrl: './sidenavigation.component.html',
  styleUrls: ['./sidenavigation.component.scss']
})
export class SidenavigationComponent implements OnInit {

    private show: boolean = false;

    constructor() { }

    ngOnInit() {
    }

    showSideNav(){
        this.show = !this.show;
    }

}
