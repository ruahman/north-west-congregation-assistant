export function DefineCustomElement(tagName: string) {
  return function(constructor: CustomElementConstructor) {
    customElements.define(tagName, constructor);
  };
}

