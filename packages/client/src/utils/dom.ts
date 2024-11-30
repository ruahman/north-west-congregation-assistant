export class $Elements {
  private elements: Element[];

  constructor(elems: NodeListOf<Element>) {
    this.elements = Array.from(elems);
  }

  set(prop: string, value: string) {
    this.elements.forEach((elem) => {
      elem.setAttribute(prop, value);
    });
  }

  get(prop) {
    // if prop get attributes
    if (isNaN(Number(prop))) {
      const props = this.elements.map((elem) => {
        return elem.getAttribute(prop);
      });

      if (props.length == 1) {
        return props[0];
      } else {
        return props;
      }
    }
    // if index get element
    else {
      return this.elements[prop];
    }
  }

  on(
    event: string,
    handler: EventListenerOrEventListenerObject,
    options: AddEventListenerOptions,
  ) {
    this.elements.forEach((elem) => {
      elem.addEventListener(event, handler, options);
    });
  }

  off(event: string, handler: EventListenerOrEventListenerObject) {
    this.elements.forEach((elem) => {
      elem.removeEventListener(event, handler);
    });
  }
}

export type $Query = (selector: string) => $Elements;

export const _ = {
  ready(param: () => void) {
    window.addEventListener("DOMContentLoaded", () => {
      param();
    });
  },
  $(doc: Document | ShadowRoot): $Query {
    return function (selector: string): $Elements {
      // query
      const proxy = new Proxy(new $Elements(doc.querySelectorAll(selector)), {
        get(obj, prop) {
          // prop
          if (isNaN(Number(prop)) == true) {
            return obj[prop];
          }
          // index
          else {
            return obj.get(prop);
          }
        },
      });
      return proxy;
    };
  },
};
