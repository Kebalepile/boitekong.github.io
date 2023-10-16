import { setupHeadline } from "./components/headline.js";
import { setupPrivateCompanies } from "./components/boards/privateCompanies.js";
import { setupPubliDepartments } from "./components/boards/publicSector.js";
import { setupPrivateSector } from "./components/boards/privateSector.js";
import { setupNavigation } from "./components/navigation/navbar.js";

document.querySelector("#app").innerHTML = `
  <header>
  
  </header>
  <br/>
  <div id="headline">
    <h1>Hello Mereko</h1>
    <div id="headline-message" class="card">
       <h3>Looking for a job in the South African Market</h3>
       <br/>
       <h4>Start Here</h4>
       <h2>Currently There Are :</h2>
       <br/>
        <ul id="headline-job-info"></ul>
        <br/> 
        <button id="share-site" class="apply">share site</button>
    </div>
  </div>
  <div id="job-board"></div>
  <div id="pvt-job-board"></div>

  <dialog id="dialog">
      <article id="info"> 
      </article>
      <button id="close-dialog">Close</button>
  </dialog>

  <footer id="contact">© 2023 K.T MOTSHOANA</footer>
`;
const shareSiteButon = document.getElementById("share-site");

      shareSiteButon.addEventListener("click", async () => {
        const shareData = {
          title:"Boitekong Job Board",
          text: "available job vacancy, might be suitable for you!",
          url: location.origin
        };
        try {
          console.log(shareData)
          await navigator.share(shareData);
        } catch (err) {
          console.error(err);
        }
      });
const closeButton = document.getElementById("close-dialog");

closeButton.addEventListener("click", () => {
  dialog.close();
});


setupHeadline(document.querySelector("#headline-job-info"));
setupPubliDepartments(document.querySelector("#job-board"));
setupPrivateCompanies(document.querySelector("#job-board"));
setupPrivateSector(document.querySelector("#pvt-job-board"));
setupNavigation(document.querySelector("header"));

const installBtn = document.getElementById("install-app");
installBtn.addEventListener('click', () => {
  console.log("install web app")
})
