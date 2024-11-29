class $Elements {
  private elements: Array<Element>;

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

  // on(event: string, handler: (e) => void, options: any) {
  //   this.elements.forEach((elem) => {});
  // }

  // off(event: string, handler: (e) => void, options: any) {
  //   this.elements.forEach((elem) => {});
  // }
}

export function Query(doc: Document) {
  return function (param: string | (() => void)) {
    // query
    if (typeof param === "string") {
      const proxy = new Proxy(new $Elements(doc.querySelectorAll(param)), {
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
      // load
    } else if (typeof param === "function") {
      window.addEventListener("DOMContentLoaded", () => {
        param();
      });
    }
    return new $Elements([] as unknown as NodeListOf<Element>);
  };
}
