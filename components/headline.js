
import{
  stats
} from "./data.js"
export function setupHeadline(element) {
  
  element.innerHTML = `
          <li> ${stats["govDep"]} Government Departments hiring</li>
          <li>${stats["privateComp"]} Private Sector Entities/Companies hiring</li>
          <li> ${stats["privateSectorOpenings"]} Private Sector Job openings</li>`;
}
