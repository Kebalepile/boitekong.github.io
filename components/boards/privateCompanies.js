import { privateJobs } from "../data.js";
export function setupPrivateCompanies(element) {
  const elems = privateJobs()["blogPosts"].map((p) => {
    const div = document.createElement("div");
    div.classList.add("job-post");
    div.innerHTML = `
        <div class="company-logo">
            <img src=${p["imgSrc"]} alt="company logo" />
        </div>
        <p class="title">${p["title"].replace(/is hiring/gi, "")}</p>
        `;

    return div;
  });

  for (const e of elems) {
    element.appendChild(e);
  }
}
