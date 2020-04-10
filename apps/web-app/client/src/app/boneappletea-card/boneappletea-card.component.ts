import { Component, OnInit, Input } from '@angular/core';
import { ApiService } from '../shared/api.service';

@Component({
  selector: 'app-boneappletea-card',
  templateUrl: './boneappletea-card.component.html',
  styleUrls: ['./boneappletea-card.component.css']
})
export class BoneappleteaCardComponent implements OnInit {

    public bats: any;

    constructor(
        private api: ApiService
    ) { }

    ngOnInit() {
        this.api.GetBats()
        .subscribe((data: any) => {
            this.bats = data.boneappleteas;
            console.log(data.boneappleteas); 
        })
    }

    accept(id: string){
        debugger
        this.api.AcceptBat(id);
    }

    deny(id: string){
        debugger
        this.api.DenyBat(id);
    }

}
