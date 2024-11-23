

class HereWeGO extends HTMLElement {

  constructor() {
    super();
    const text = this.getAttribute("text")
    this.innerHTML = `<h1>${text}</h1>`;
  }

  // Specify which attributes to observe
  static get observedAttributes() {
    return ["test"];
  }

  // React to attribute changes
  attributeChangedCallback(attribute, oldValue, newValue) {
    if (attribute === "test" && oldValue !== newValue) {
      const elem = this.querySelector("h1");
      if (elem) {
        elem.textContent = newValue;
      }
    }
  }

  connectedCallback() {
    console.log("connected...");
    // this.textContent = "implement here wego map"
  }

  disconnectedCallback() {
    console.log("disconnected...");
  }
}

customElements.define('here-wego', HereWeGO);
