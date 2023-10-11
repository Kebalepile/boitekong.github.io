import "./style.css";

import { setupHeadline } from "./components/headline";

document.querySelector("#app").innerHTML = `
  <div id="headline">
  
    <h1>Hello Mereko</h1>
    <div id="headline-message" class="card">
       <h3>Looking for a job in the South African Market</h3>
       <br/>
       <h4>Start Here</h4>
       <p>Currently There Are :</p>
        <ul id="headline-job-info">
             
        </ul>
      
    </div>
    
  </div>
`;

setupHeadline(document.querySelector("#headline-job-info"));
