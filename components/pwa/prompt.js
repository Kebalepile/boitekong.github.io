export function setupInstallPrormpt(element) {
  let pwaInstallPrompt;

  window.addEventListener("beforeinstallprompt", (e) => {
    try {
      // Prevent Chrome 67 and earlier from automatically showing the prompt
      e.preventDefault();
      // Stash the event so it can be triggered later.
      pwaInstallPrompt = e;
    } catch (err) {
      consol.error(err.message);
    }
  });

  element.addEventListener("click", (e) => {
    try {
      if (pwaInstallPrompt) {
        // Show the prompt
        pwaInstallPrompt.prompt();
        // Wait for the user to respond to the prompt
        pwaInstallPrompt.userChoice.then((choiceResult) => {
          if (choiceResult.outcome === "accepted") {
            console.log("User accepted the A2HS prompt");
          } else {
            console.log("User dismissed the A2HS prompt");
          }
          pwaInstallPrompt = null;
        });
      }
      e.target.blur();
    } catch (err) {
      console.error(err.message);
    }
  });
  /**
   * @description remove install button once web app is installed
   */
  if (window.matchMedia(" (display-mode: standalone)").matches) {
    // PWA is installed

    if (element) element.remove();
  }
}
