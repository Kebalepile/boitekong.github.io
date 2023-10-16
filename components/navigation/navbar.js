export  function setupNavigation(element){
    element.innerHTML = `
        <img src="#" alt="website logo" id="logo"/>
        <nav>    
            <ul>
                <li>
                    <a href="#headline"> about </a>
                </li>
                <li>
                    <a href="#contact"> contact info</a>
                </li>
                <li>
                    <h5>Types of Vacanies</h5>
                    <ul>
                        <li>
                            <a href="#pvt-job-board"> private sector jobs </a>
                        </li>
                        <li>
                            <a href="#private-companies"> private sector companies hiring </a>
                        </li>
                        <li>
                            <a href="#public-sector"> public sector departments hiring </a>
                        </li>
                    </ul>
                </li>
                <li>
                    <button id="install-app" class="apply"> install web app</button>
                </li>
            </ul>
        </nav>`;
        
}