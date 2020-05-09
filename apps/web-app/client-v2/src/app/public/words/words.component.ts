import { Component, OnInit } from '@angular/core';
import { ApiService } from '../../shared/api.service';

@Component({
  selector: 'app-words',
  templateUrl: './words.component.html',
  styleUrls: ['./words.component.scss']
})
export class WordsComponent implements OnInit {

    private data: any = {};
    private boneappletea: any = {};
    private done: boolean = false;

    constructor(
        private api: ApiService
    ) { }

    ngOnInit() {
    }

    search(form:any){
        this.api.Search(form.value)
            .subscribe((data: any) => {
                if(data === null){
                    this.done = false;
                }else{
                    this.boneappletea = data.boneappletea;
                    console.log(this.boneappletea);
                    this.done = true;
                }
            });
    }

}
