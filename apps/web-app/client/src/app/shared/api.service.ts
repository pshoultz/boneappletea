import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';

@Injectable({
  providedIn: 'root'
})
export class ApiService {

    //private url: string = "http://127.0.0.1:8080/";
    private url: string = "http://bat-con-api:8080/";

    constructor(
        public http: HttpClient
    ) { }

    GetBats(){
        return this.http.get(this.url + "get");
    }

    MakeBat(sentence: string){
        return this.http.get(this.url + "bat", {
            headers: new HttpHeaders({
                "sentence": sentence
            })
        });
    }

    AddBat(root:string, replacement: string){
        return this.http.post(this.url + "add",{
            root: root,
            replacement: replacement
        },
        {
            //NOTE: should be json 
            responseType: 'text'
        });
    }

    AcceptBats(root: string, replacement: string){
        return this.http.post(
            this.url + "accept", {
                root: root,
                replacement: replacement
            },
            {
                //NOTE: should be json 
                responseType: 'text'
            });
    }

    SendEmail(address: string, message: string){
        debugger
        return this.http.post(
            this.url + "email", {
                address: address,
                message: message
            },
            {
                //NOTE: should be json 
                responseType: 'text'
            });
    }

    Search(word: string){
        return this.http.get(this.url + "search?value=" + word);
    }
}
