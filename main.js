

import { setupHeadline } from "./components/headline.js";
import {setupPrivateCompanies} from "./components/boards/privateCompanies.js";

document.querySelector("#app").innerHTML = `
  <header></header>
  <div id="headline">
    <h1>Hello Mereko</h1>
    <div id="headline-message" class="card">
       <h3>Looking for a job in the South African Market</h3>
       <br/>
       <h4>Start Here</h4>
       <p>Currently There Are :</p>
        <ul id="headline-job-info"></ul> 
    </div>
  </div>
  <div id="job-board"></div>
  <footer></footer>
`;

setupHeadline(document.querySelector("#headline-job-info"));
setupPrivateCompanies(document.querySelector("#job-board"))
