

class $Element {
  constructor(private elem: Element) { }

  set(prop: string, value: string) {
    this.elem.setAttribute(prop, value);
  }

  get(prop) {
    if (isNaN(Number(prop))) {
      return this.elem.getAttribute(prop)
    }
    else {
      return this.elem;
    }
  }
}

export function $(query: string) {
  let proxy = new Proxy(new $Element(document.querySelector(query)!), {
    get(obj, prop) {
      if (isNaN(Number(prop)) == true) {
        return obj[prop];
      }
      else {
        return obj.get(prop)
      }
    }
  });
  return proxy;
}
