import { Component, OnInit } from '@angular/core';
import { ApiService } from '../shared/api.service';

@Component({
  selector: 'app-boneappletea',
  templateUrl: './boneappletea.component.html',
  styleUrls: ['./boneappletea.component.css']
})
export class BoneappleteaComponent implements OnInit {

    constructor(
        private api: ApiService
    ) { }

    ngOnInit() {
        this.api.GetBats()
        .subscribe((data: any) => {
           console.log(data); 
        })
    }

}
