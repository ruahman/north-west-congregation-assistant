import H from "@here/maps-api-for-javascript";
import { DefineCustomElement } from "@utils/web-component";
import { _, $Query } from "@utils/dom";

const style = `
#here {
  width: 100%;
  height: 100%;
}
`;

const template = `
<style>${style}</style>
<div id='here'></div>
`;

@DefineCustomElement("here-map")
export class HereMap extends HTMLElement {
  private template: HTMLTemplateElement;
  private query: $Query | undefined;
  private platform: H.service.Platform | undefined;
  private defaultLayers: any;

  constructor() {
    super();

    // create shadow dom
    this.attachShadow({ mode: "open" });

    // setup template
    this.template = document.createElement("template");
    this.template.innerHTML = template;

    // setup query
    if (this.shadowRoot) {
      this.query = _.$(this.shadowRoot);
    }

    // setup Here Map
    const { VITE_HERE_APIKEY: apikey } = import.meta.env;
    if (apikey) {
      this.platform = new H.service.Platform({
        apikey,
      });
    }

    // layers you can apply???
    this.defaultLayers = this.platform?.createDefaultLayers();
  }

  // Specify which attributes to observe
  static get observedAttributes() {
    return ["test"];
  }

  // React to attribute changes
  attributeChangedCallback(attribute, _oldValue, newValue) {
    if (attribute == "test") {
      const elem = this.query?.("#here-map")[0];
      if (elem) {
        elem.textContent = newValue;
      }
    }
  }

  // property you can access from code
  set test(val: string) {
    const elem = this.query?.("#here-map")[0];
    if (elem) {
      elem.textContent = val;
    }
  }

  // now component is in dom
  connectedCallback() {
    // append template
    this.shadowRoot?.appendChild(this.template.content.cloneNode(true));

    // setup map
    const elem = this.query?.("#here")[0];
    new H.Map(elem, this.defaultLayers.vector.normal.map, {
      zoom: 15,
      center: { lat: 18.47383, lng: -66.93851 },
    });
  }
}
