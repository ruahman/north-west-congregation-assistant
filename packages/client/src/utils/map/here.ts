import { DefineCustomElement } from "@utils/web_component";

@DefineCustomElement("here-map")
export class HereMap extends HTMLElement {
  private text: string;
  private template: HTMLTemplateElement;

  constructor() {
    super();
    this.text = this.getAttribute("text")!;
    this.attachShadow({ mode: "open" });
    this.template = document.createElement("template");
    this.template.innerHTML = `<h1 id='here-wego'>${this.text}</h1>`;
  }

  // Specify which attributes to observe
  static get observedAttributes() {
    return ["test"];
  }

  // React to attribute changes
  attributeChangedCallback(_attribute, _oldValue, newValue) {
    const elem = this.shadowRoot?.querySelector("h1");
    if (elem) {
      elem.textContent = newValue;
    }
  }

  set test(val: string) {
    const elem = this.shadowRoot?.querySelector("h1");
    if (elem) {
      elem.textContent = val;
    }
  }

  // now component is in dom
  connectedCallback() {
    this.shadowRoot?.appendChild(this.template.content.cloneNode(true));
  }

  disconnectedCallback() {
    console.log("disconnected...");
  }

  // moved to another documetn
  adoptedCallback() {
    console.log("adopted");
  }
}
