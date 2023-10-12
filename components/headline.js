import propersonnelData from "../backend/database/private/Pro-Personnel.json";
import heithaData from "../backend/database/private/heitha-stuffing-group.json";
import privateSectorData from "../backend/database/public/govpage-private-sector.json"
import publicSectorData from "../backend/database/public/govpage-public-sector.json"

export function setupHeadline(element) {
  
  element.innerHTML = `
          <li> ${publicSectorData.blogPosts.length} Government Departments hiring</li>
          <li>${privateSectorData.blogPosts.length} Private Sector Entities/Companies hiring</li>
          <li> ${propersonnelData.blogPosts.length + heithaData.blogPosts.length} Private Sector Job openings</li>`;
}
