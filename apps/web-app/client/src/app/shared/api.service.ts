import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';

@Injectable({
  providedIn: 'root'
})

export class ApiService {

    constructor(
        private http: HttpClient
    ) { }

    GetBats() {
        return this.http.get("http://127.0.0.1:8080/get");
    }

    //NOTE: we should have ids on each BAT so we can find them in the DB faster
    AcceptBat(id: string){
        //NOTE: hit router for accept then attach a header of 'id' to it.
        //once in the API, pull the id off the header
        //search for the document in the DB with this id
        //set the flag to true
    }

    //NOTE: we should have ids on each BAT so we can find them in the DB faster
    DenyBat(id: string){
        //NOTE: this might take a little more thinking but I think the delete function should work here
    }
}
