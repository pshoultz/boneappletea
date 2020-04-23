import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';

@Injectable({
  providedIn: 'root'
})
export class ApiService {

    private url: string = "http://127.0.0.1:8080/";

    constructor(
        public http: HttpClient
    ) { }

    GetBats(){
        return this.http.get(this.url);
    }

    MakeBat(sentence: string){
        const opts = {
            headers: new HttpHeaders({
                "sentence": sentence
            })
        }

        //return this.http.get(this.url, opts);
        return this.http.get(this.url + "bat", {
            headers: new HttpHeaders({
                "sentence": sentence
            })
        });
    }
}
