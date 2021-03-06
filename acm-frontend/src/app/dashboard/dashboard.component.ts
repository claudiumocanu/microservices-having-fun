import { Component, OnInit } from '@angular/core';
import { SensorService } from '../services/sensor.service';

@Component({
  selector: 'app-dashboard',
  templateUrl: './dashboard.component.html',
  styleUrls: ['./dashboard.component.css']
})
export class DashboardComponent implements OnInit {
  options = {
    height: '500',
    titleTextStyle: {
      color: '3A3',
    },
    animation: {
      startup: true,
      duration: 1000,
      easing: 'inAndOut',
    },
    legend: {
      position: 'none'
    },
    selectionMode: 'multiple',
    is3D: true,
    colors: ['#e0FF0e'],
    backgroundColor: '222',
    curveType: 'function',
    hAxis: {
      title: 'timestamp',
      titleTextStyle: {
        color: '#3A3'
      },
      gridlines: { color: '#3A3' },
      textStyle: { color: '#3A3' }
    },
    vAxis: {
      title: 'reading',
      titleTextStyle: {
        color: '#3A3'
      },
      gridlines: { color: '#3A3' },
      textStyle: { color: '#3A3' }
    }
  };
  sensors: Array<any> = [];
  sensorData: Array<Array<any>> = [
    [1, 37.8, 80.8, 41.8],
    [2, 30.9, 69.5, 32.4],
    [3, 25.4, 57, 25.7],
    [4, 11.7, 18.8, 10.5],
    [5, 11.9, 17.6, 10.4],
    [6, 8.8, 13.6, 7.7],
    [7, 7.6, 12.3, 9.6],
    [8, 12.3, 29.2, 10.6],
    [9, 16.9, 42.9, 14.8],
    [10, 12.8, 30.9, 11.6],
    [11, 5.3, 7.9, 4.7],
    [12, 6.6, 8.4, 5.2],
    [13, 4.8, 6.3, 3.6],
    [14, 4.2, 6.2, 3.4]];

  isLoadingResults = true;

  constructor(private sensorService: SensorService) {

  }

  ngOnInit(): void {
    // Fetch registered sensors from postgres
    this.sensorService.getSensors()
      .subscribe(data => {
        this.sensors = data as [];
        this.isLoadingResults = false;
      })
  }
}
