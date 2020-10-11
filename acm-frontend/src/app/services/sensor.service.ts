import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { environment} from '../../environments/environment'

@Injectable({
  providedIn: 'root'
})
export class SensorService {

  url = environment.URL_GO_PG_SENSORS_BACKEND + "/sensor";
  constructor(private http:HttpClient) { 

  }    
  
  getSensors(){
    return this.http.get(this.url);
  }

}
