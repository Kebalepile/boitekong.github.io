import { setupHeadline } from "./components/headline.js";
import { setupPrivateCompanies } from "./components/boards/privateCompanies.js";
import { setupPubliDepartments } from "./components/boards/publicSector.js";
import { setupPrivateSector } from "./components/boards/privateSector.js";

document.querySelector("#app").innerHTML = `
  <header></header>
  <div id="headline">
    <h1>Hello Mereko</h1>
    <div id="headline-message" class="card">
       <h3>Looking for a job in the South African Market</h3>
       <br/>
       <h4>Start Here</h4>
       <h2>Currently There Are :</h2>
        <ul id="headline-job-info"></ul> 
    </div>
  </div>
  <div id="job-board"></div>
  <div id="pvt-job-board"></div>
  <footer></footer>
`;

setupHeadline(document.querySelector("#headline-job-info"));
setupPubliDepartments(document.querySelector("#job-board"));
setupPrivateCompanies(document.querySelector("#job-board"));
setupPrivateSector(document.querySelector("#pvt-job-board"));
