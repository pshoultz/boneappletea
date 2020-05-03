import { Component, OnInit } from '@angular/core';
import { ApiService } from '../../shared/api.service';

@Component({
  selector: 'app-approve',
  templateUrl: './approve.component.html',
  styleUrls: ['./approve.component.scss']
})
export class ApproveComponent implements OnInit {
    public tiles: any[] = [
        {text: 'One', cols: 3, rows: 1, color: 'lightblue'},
        {text: 'Two', cols: 1, rows: 2, color: 'lightgreen'},
        {text: 'Three', cols: 1, rows: 1, color: 'lightpink'},
        {text: 'Four', cols: 2, rows: 1, color: '#DDBDF1'},
    ];

    public boneappleteas: any = {};
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
    accept(root: string, replacement: string){
        this.api.AcceptBats(root, replacement)
        .subscribe((data: any) => {
            console.log(data);
        });
    }

    //NOTE: should delete entry from array in db
    deny(root: string, replacement: string){
        //NOTE: call delete function 
    }

}
