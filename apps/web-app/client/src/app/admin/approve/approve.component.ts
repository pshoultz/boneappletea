import { Component, OnInit } from '@angular/core';
import { ApiService } from '../../shared/api.service';

@Component({
  selector: 'app-approve',
  templateUrl: './approve.component.html',
  styleUrls: ['./approve.component.scss']
})
export class ApproveComponent implements OnInit {
    public boneappleteas: any = [];
    public loaded: boolean = false;

    constructor(
        private api: ApiService
    ) { }

    ngOnInit() {
        this.makeCards();
    }

    //NOTE: get all boneappleteas
    makeCards(){
        this.api.GetBats()
        .subscribe((data: any) => {
            this.boneappleteas = data.boneappleteas;
            this.loaded = true;
        });
    }

    //NOTE: accept into db
    accept(root: string, replacement: string, i: number){
        let element = document.getElementById("card-" + i);
        element.remove();
        this.api.AcceptBats(root, replacement)
        .subscribe((data: any) => {
            console.log(data);
            this.makeCards();//NOTE: to refresh the card collection? could be done differently though like just removing the card from dom
        });
    }

    //NOTE: should delete entry from array in db?
    deny(root: string, replacement: string){
        //NOTE: call delete function 
    }

}
