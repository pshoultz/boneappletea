import { Component, OnInit } from '@angular/core';
import { ApiService } from '../../shared/api.service';

@Component({
  selector: 'app-words',
  templateUrl: './words.component.html',
  styleUrls: ['./words.component.scss']
})
export class WordsComponent implements OnInit {

    public data: any = {};
    public boneappleteas: any = [];

    constructor(
        private api: ApiService
    ) { }

    ngOnInit() {
    }

    search(form:any){
        this.api.Search(form.value)
            .subscribe((data: any) => {
                debugger
            });
    }

}
