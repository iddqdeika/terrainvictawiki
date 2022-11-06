import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import {ExtraOptions, RouterModule, Routes} from '@angular/router';

import { AppComponent } from './app.component';
import { HomeComponent } from './home/home.component';
import { TechComponent } from './tech/tech.component';
import { ProjectComponent } from './project/project.component';


const routes: Routes = [
  { path: '', component: HomeComponent },
  { path: 'tech', component: TechComponent},
  { path: 'project', component: ProjectComponent},
  // otherwise redirect to home
  { path: '**', redirectTo: '' }
];

const routerOptions: ExtraOptions = {
  useHash: false,
  anchorScrolling: 'enabled',
  scrollPositionRestoration: 'enabled',
  scrollOffset: [0, 100],
  onSameUrlNavigation: 'reload',
  // ...any other options you'd like to use
};

@NgModule({
  declarations: [
    AppComponent,
    HomeComponent,
    TechComponent,
    ProjectComponent
  ],
  imports: [
    BrowserModule,
    RouterModule,
    RouterModule.forRoot(routes, routerOptions)
  ],
  providers: [],
  bootstrap: [AppComponent]
})

export class AppModule { }
